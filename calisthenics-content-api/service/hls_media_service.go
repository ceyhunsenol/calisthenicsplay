package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
	"calisthenics-content-api/model"
	"fmt"
	"github.com/o1egl/paseto"
	"github.com/spf13/viper"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Resolution struct {
	Bandwidth    int
	AvgBandwidth int
	Codecs       string
	Resolution   string
	MediaURL     string
}

type IHLSMediaService interface {
	GetResolutions() []Resolution
	GetMasterPlaylist(t string) (string, *model.ServiceError)
	GetMediaPlaylist(request model.VideoPlaylistRequest) (string, *model.ServiceError)
	GetMediaURL(request model.VideoURLRequest) (model.VideoURLModel, *model.ServiceError)
	GetLicenseKey(token string) (string, *model.ServiceError)
}

type hlSMediaService struct {
	mediaCacheService      cache.IMediaCacheService
	encodingCacheService   cache.IHLSEncodingCacheService
	mediaPlayActionService IMediaPlayActionService
	parameterService       IParameterService
	resolutions            []Resolution
	hlsMediaToken          string
}

func NewHLSMediaService(
	encodingCacheService cache.IHLSEncodingCacheService,
	mediaCacheService cache.IMediaCacheService,
	mediaPlayActionService IMediaPlayActionService,
	parameterService IParameterService,
) IHLSMediaService {
	res1 := Resolution{Bandwidth: 1938572, AvgBandwidth: 1390351, Codecs: "avc1.4d4028,mp4a.40.2", Resolution: "848x480"}
	res2 := Resolution{Bandwidth: 3589944, AvgBandwidth: 2279704, Codecs: "avc1.4d4028,mp4a.40.2", Resolution: "1272x720"}
	res3 := Resolution{Bandwidth: 6974954, AvgBandwidth: 4217619, Codecs: "avc1.4d4028,mp4a.40.2", Resolution: "1906x1080"}
	resolutions := make([]Resolution, 0)
	resolutions = append(resolutions, res1, res2, res3)
	return &hlSMediaService{
		encodingCacheService:   encodingCacheService,
		mediaCacheService:      mediaCacheService,
		parameterService:       parameterService,
		resolutions:            resolutions,
		mediaPlayActionService: mediaPlayActionService,
		hlsMediaToken:          viper.GetString("media.hls.paseto-token"),
	}
}

func (s *hlSMediaService) createHLSToken(key []byte, footer interface{}) (string, *model.ServiceError) {
	v2 := paseto.NewV2()
	payload := paseto.JSONToken{
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(40 * time.Minute),
		NotBefore:  time.Now(),
	}
	token, err := v2.Encrypt(key, payload, footer)
	if err != nil {
		return "", &model.ServiceError{Code: http.StatusUnauthorized, Message: "Token not valid"}
	}
	return token, nil
}

func (s *hlSMediaService) getHLSMediaTokenModel(key []byte, token string) (model.HLSMediaTokenModel, *model.ServiceError) {
	v2 := paseto.NewV2()
	var payload paseto.JSONToken
	var hlsToken model.HLSMediaTokenModel
	err := v2.Decrypt(token, key, &payload, &hlsToken)
	if err != nil {
		return model.HLSMediaTokenModel{}, &model.ServiceError{Code: http.StatusUnauthorized, Message: "Token not valid"}
	}
	if payload.Expiration.Before(time.Now()) {
		return model.HLSMediaTokenModel{}, &model.ServiceError{Code: http.StatusUnauthorized, Message: "Token expired"}
	}
	return hlsToken, nil
}

func (s *hlSMediaService) GetMasterPlaylist(t string) (string, *model.ServiceError) {
	key := []byte(s.hlsMediaToken)
	_, err := s.getHLSMediaTokenModel(key, t)
	if err != nil {
		return "", err
	}
	playlistContent := "#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-INDEPENDENT-SEGMENTS\n"
	playlistPath, playlistError := s.parameterService.GetHLSPlaylistPath()
	if playlistError != nil {
		return "", &model.ServiceError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	re := regexp.MustCompile(`x(\d+)`)
	for _, resolution := range s.resolutions {
		matches := re.FindStringSubmatch(resolution.Resolution)
		playlistContent += fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,AVERAGE-BANDWIDTH=%d,CODECS=\"%s\",RESOLUTION=%s\n%s\n",
			resolution.Bandwidth, resolution.AvgBandwidth, resolution.Codecs, resolution.Resolution, playlistPath+"/"+matches[1]+".m3u8?t="+t)
	}
	return playlistContent, nil
}

func (s *hlSMediaService) GetMediaPlaylist(request model.VideoPlaylistRequest) (string, *model.ServiceError) {
	key := []byte(s.hlsMediaToken)
	tokenModel, serviceError := s.getHLSMediaTokenModel(key, request.Token)
	if serviceError != nil {
		return "", serviceError
	}

	re := regexp.MustCompile(`(\d+)\.m3u8`)
	resolutionMatch := re.FindStringSubmatch(request.Resolution)
	if len(resolutionMatch) < 1 {
		return "", &model.ServiceError{Code: http.StatusBadRequest, Message: "Invalid request format"}
	}

	found := false
	for _, resolution := range s.resolutions {
		if strings.Contains(resolution.Resolution, resolutionMatch[1]) {
			found = true
			break
		}
	}

	if !found {
		return "", &model.ServiceError{Code: http.StatusBadRequest, Message: "Invalid resolution"}
	}

	var builder strings.Builder

	builder.WriteString("#EXTM3U\n")
	builder.WriteString(fmt.Sprintf("#EXT-X-VERSION:%d\n", 3))
	builder.WriteString(fmt.Sprintf("#EXT-X-TARGETDURATION:%d\n", 12))
	builder.WriteString(fmt.Sprintf("#EXT-X-MEDIA-SEQUENCE:%d\n", 0))
	builder.WriteString(fmt.Sprintf("#EXT-X-PLAYLIST-TYPE:%s\n", "VOD"))
	encodingCache, err := s.encodingCacheService.GetByID(tokenModel.EncodingID)
	if err != nil {
		return "", &model.ServiceError{Code: http.StatusNotFound, Message: "Encoding not found"}
	}

	url, err := s.parameterService.GetHLSCdnURL()
	if err != nil {
		return "", &model.ServiceError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	licenseURL, err := s.parameterService.GetHLSLicensePath()
	if err != nil {
		return "", nil
	}

	r := regexp.MustCompile(`_(\d+)_`)
	resolutionFiles := make([]cache.HLSEncodingFileCache, 0)
	for _, file := range encodingCache.Files {
		matches := r.FindStringSubmatch(file.FileName)
		if resolutionMatch[1] == matches[1] {
			resolutionFiles = append(resolutionFiles, file)
		}
	}

	for i, file := range resolutionFiles {
		if i == 0 {
			builder.WriteString("#EXT-X-DISCONTINUITY\n")
		}
		builder.WriteString(fmt.Sprintf("#EXTINF:%.2f,\n", 10.0))
		builder.WriteString(fmt.Sprintf("#EXT-X-KEY:METHOD=%s,URI=\"%s\",IV=%s\n", "AES-128", tokenModel.Host+licenseURL+"?t="+request.Token, file.IV))
		builder.WriteString(url + "/" + file.FileName + "\n")
	}

	builder.WriteString("#EXT-X-ENDLIST\n")

	return builder.String(), nil
}

func (s *hlSMediaService) GetMediaURL(request model.VideoURLRequest) (model.VideoURLModel, *model.ServiceError) {
	mediaCache, err := s.mediaCacheService.GetByID(request.MediaID)
	if err != nil {
		return model.VideoURLModel{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Media not found"}
	}
	action := s.mediaPlayActionService.GetPlayAction(request.Token, mediaCache)
	if action.ActionType != config.Watch {
		return model.VideoURLModel{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Media cannot be accessed"}
	}
	encodingCache, err := s.encodingCacheService.GetByID(mediaCache.EncodingID)
	if err != nil {
		return model.VideoURLModel{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Encoding not found"}
	}

	key := []byte(s.hlsMediaToken)
	hlSToken := model.HLSMediaTokenModel{
		MediaID:      request.MediaID,
		EncodingID:   encodingCache.ID,
		UserAgent:    request.UserAgent,
		CallerIP:     request.CallerIP,
		PlatformType: request.PlatformType,
		Host:         request.Host,
	}
	token, serviceError := s.createHLSToken(key, hlSToken)
	if serviceError != nil {
		return model.VideoURLModel{}, &model.ServiceError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	path, err := s.parameterService.GetMasterM3u8Path()
	if err != nil {
		return model.VideoURLModel{}, &model.ServiceError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	resolutions := make([]model.VideoURLResolutionModel, 0)
	re := regexp.MustCompile(`x(\d+)`)
	for _, resolution := range s.resolutions {
		matches := re.FindStringSubmatch(resolution.Resolution)
		height, _ := strconv.Atoi(matches[1])
		resolutionModel := model.VideoURLResolutionModel{
			Height:       height,
			Bandwidth:    resolution.Bandwidth,
			AvgBandwidth: resolution.AvgBandwidth,
		}
		resolutions = append(resolutions, resolutionModel)
	}
	urlModel := model.VideoURLModel{
		VideoURL:    request.Host + path + "?t=" + token,
		Resolutions: resolutions,
	}
	return urlModel, nil
}

func (s *hlSMediaService) GetResolutions() []Resolution {
	return s.resolutions
}

func (s *hlSMediaService) GetLicenseKey(token string) (string, *model.ServiceError) {
	key := []byte(s.hlsMediaToken)
	tokenModel, serviceError := s.getHLSMediaTokenModel(key, token)
	if serviceError != nil {
		return "", serviceError
	}

	encodingCache, err := s.encodingCacheService.GetByID(tokenModel.EncodingID)
	if err != nil {
		return "", &model.ServiceError{Code: http.StatusNotFound, Message: "Encoding not found"}
	}

	return encodingCache.LicenseKey, nil
}
