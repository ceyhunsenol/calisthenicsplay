package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
)

type IParameterService interface {
	GetAccessLevel() (string, error)
	GetMasterM3u8Path() (string, error)
	GetHLSCdnURL() (string, error)
	GetHLSLicensePath() (string, error)
	GetHLSPlaylistPath() (string, error)
}

type parameterService struct {
	generalInfoCacheService cache.IGeneralInfoCacheService
}

func NewParameterService(generalInfoCacheService cache.IGeneralInfoCacheService) IParameterService {
	return &parameterService{
		generalInfoCacheService: generalInfoCacheService,
	}
}

func (s *parameterService) GetAccessLevel() (string, error) {
	id, err := s.generalInfoCacheService.GetByID(string(config.AccessLevel))
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *parameterService) GetMasterM3u8Path() (string, error) {
	id, err := s.generalInfoCacheService.GetByID("MASTER_HLS_MASTER_M3U8_PATH")
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *parameterService) GetHLSCdnURL() (string, error) {
	id, err := s.generalInfoCacheService.GetByID("HLS_CDN_URL")
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *parameterService) GetHLSLicensePath() (string, error) {
	id, err := s.generalInfoCacheService.GetByID("HLS_LICENSE_PATH")
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *parameterService) GetHLSPlaylistPath() (string, error) {
	id, err := s.generalInfoCacheService.GetByID("HLS_PLAYLIST_PATH")
	if err != nil {
		return "", err
	}
	return id, nil
}
