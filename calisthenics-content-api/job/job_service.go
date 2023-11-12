package job

import (
	"calisthenics-content-api/cache"
	"time"
)

type IJobService interface {
	LimitedCacheStartJob()
}

type jobService struct {
	limitedCacheService cache.ILimitedCacheService
}

func NewJobService(limitedCacheService cache.ILimitedCacheService) IJobService {
	return &jobService{
		limitedCacheService: limitedCacheService,
	}
}

func (j *jobService) LimitedCacheStartJob() {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			all := j.limitedCacheService.GetAll()
			for _, limitedCache := range all {
				if limitedCache.GetLimitedEndDate().Before(time.Now()) {
					j.limitedCacheService.Remove(limitedCache.Key)
				}
			}
		}
	}
}
