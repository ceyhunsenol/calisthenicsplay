package service

import "calisthenics-content-api/job"

type IInitService interface {
	InitCache()
	InitJob()
}

type initService struct {
	// cache
	contentCacheOperations       IContentCacheOperations
	genreCacheOperations         IGenreCacheOperations
	mediaCacheOperations         IMediaCacheOperations
	contentAccessCacheOperations IContentAccessCacheOperations
	generalInfoCacheOperations   IGeneralInfoCacheOperations
	mediaAccessCacheOperations   IMediaAccessCacheOperations

	// job
	jobService job.IJobService
}

func NewInitService(contentCacheOperations IContentCacheOperations,
	genreCacheOperations IGenreCacheOperations,
	mediaCacheOperations IMediaCacheOperations,
	contentAccessCacheOperations IContentAccessCacheOperations,
	generalInfoCacheOperations IGeneralInfoCacheOperations,
	mediaAccessCacheOperations IMediaAccessCacheOperations,
	jobService job.IJobService,

) IInitService {
	return &initService{
		genreCacheOperations:         genreCacheOperations,
		contentCacheOperations:       contentCacheOperations,
		mediaCacheOperations:         mediaCacheOperations,
		contentAccessCacheOperations: contentAccessCacheOperations,
		generalInfoCacheOperations:   generalInfoCacheOperations,
		mediaAccessCacheOperations:   mediaAccessCacheOperations,
		jobService:                   jobService,
	}
}

func (c *initService) InitCache() {
	c.genreCacheOperations.SaveCacheGenres()
	c.contentCacheOperations.SaveCacheContents()
	c.mediaCacheOperations.SaveCacheMedias()

	c.contentAccessCacheOperations.SaveCacheContentAccessList()
	c.generalInfoCacheOperations.SaveCacheGeneralInfos()
	c.mediaAccessCacheOperations.SaveCacheMediaAccessList()
}

func (c *initService) InitJob() {
	c.jobService.LimitedCacheStartJob()
}
