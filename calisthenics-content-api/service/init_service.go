package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/job"
	"calisthenics-content-api/model"
	"net/http"
)

type IInitService interface {
	InitCache()
	InitJob()
	CallFromCacheFuncAllByCacheType(cacheType string) *model.ServiceError
	CallFromCacheFuncByCacheType(cacheType, id string) interface{}
}

type CacheAllFunc func() *model.ServiceError

type CacheFunc func(string) interface{}

type initService struct {
	// cache
	contentCacheOperations       IContentCacheOperations
	genreCacheOperations         IGenreCacheOperations
	mediaCacheOperations         IMediaCacheOperations
	contentAccessCacheOperations IContentAccessCacheOperations
	generalInfoCacheOperations   IGeneralInfoCacheOperations
	mediaAccessCacheOperations   IMediaAccessCacheOperations
	hlSEncodingCacheOperations   IHLSEncodingCacheOperations
	translationCacheOperations   ITranslationCacheOperations

	funcMapAll map[string]CacheAllFunc
	funcMap    map[string]CacheFunc

	// job
	jobService job.IJobService
}

func NewInitService(contentCacheOperations IContentCacheOperations,
	genreCacheOperations IGenreCacheOperations,
	mediaCacheOperations IMediaCacheOperations,
	contentAccessCacheOperations IContentAccessCacheOperations,
	generalInfoCacheOperations IGeneralInfoCacheOperations,
	mediaAccessCacheOperations IMediaAccessCacheOperations,
	hlSEncodingCacheOperations IHLSEncodingCacheOperations,
	translationCacheOperations ITranslationCacheOperations,
	jobService job.IJobService,

) IInitService {
	funcMapAll := map[string]CacheAllFunc{
		string(cache.Genre):         genreCacheOperations.SaveCacheGenres,
		string(cache.Content):       contentCacheOperations.SaveCacheContents,
		string(cache.Media):         mediaCacheOperations.SaveCacheMedias,
		string(cache.ContentAccess): contentAccessCacheOperations.SaveCacheContentAccessList,
		string(cache.GeneralInfo):   generalInfoCacheOperations.SaveCacheGeneralInfos,
		string(cache.MediaAccess):   mediaAccessCacheOperations.SaveCacheMediaAccessList,
		string(cache.HLS):           hlSEncodingCacheOperations.SaveCacheHLSEncodingList,
		string(cache.Translation):   translationCacheOperations.SaveCacheTranslationList,
	}

	funcMap := map[string]CacheFunc{
		string(cache.Genre): genreCacheOperations.SaveCacheGenre,
		string(cache.Content): func(id string) interface{} {
			return contentCacheOperations.SaveCacheContent(id)
		},
		string(cache.Media):         mediaCacheOperations.SaveCacheMedia,
		string(cache.ContentAccess): contentAccessCacheOperations.SaveCacheContentAccess,
		string(cache.GeneralInfo):   generalInfoCacheOperations.SaveCacheGeneralInfo,
		string(cache.MediaAccess):   mediaAccessCacheOperations.SaveCacheMediaAccess,
		string(cache.HLS):           hlSEncodingCacheOperations.SaveCacheHLSEncoding,
		string(cache.Translation):   translationCacheOperations.SaveCacheTranslation,
	}

	return &initService{
		genreCacheOperations:         genreCacheOperations,
		contentCacheOperations:       contentCacheOperations,
		mediaCacheOperations:         mediaCacheOperations,
		contentAccessCacheOperations: contentAccessCacheOperations,
		generalInfoCacheOperations:   generalInfoCacheOperations,
		mediaAccessCacheOperations:   mediaAccessCacheOperations,
		hlSEncodingCacheOperations:   hlSEncodingCacheOperations,
		translationCacheOperations:   translationCacheOperations,
		jobService:                   jobService,
		funcMapAll:                   funcMapAll,
		funcMap:                      funcMap,
	}
}

func (c *initService) InitCache() {
	_ = c.genreCacheOperations.SaveCacheGenres()
	_ = c.contentCacheOperations.SaveCacheContents()
	_ = c.mediaCacheOperations.SaveCacheMedias()

	_ = c.contentAccessCacheOperations.SaveCacheContentAccessList()
	_ = c.generalInfoCacheOperations.SaveCacheGeneralInfos()
	_ = c.mediaAccessCacheOperations.SaveCacheMediaAccessList()
	_ = c.hlSEncodingCacheOperations.SaveCacheHLSEncodingList()

	_ = c.translationCacheOperations.SaveCacheTranslationList()
}

func (c *initService) InitJob() {
	c.jobService.LimitedCacheJob()
}

func (c *initService) CallFromCacheFuncAllByCacheType(cacheType string) *model.ServiceError {
	callFunc := c.funcMapAll[cacheType]
	if callFunc == nil {
		return &model.ServiceError{Code: http.StatusBadRequest, Message: "Not implemented"}
	}
	return callFunc()
}

func (c *initService) CallFromCacheFuncByCacheType(cacheType, id string) interface{} {
	callFunc := c.funcMap[cacheType]
	if callFunc == nil {
		return &model.ServiceError{Code: http.StatusBadRequest, Message: "Not implemented"}
	}
	return callFunc(id)
}
