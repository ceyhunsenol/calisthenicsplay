package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
	"calisthenics-content-api/integration/calisthenics"
	"calisthenics-content-api/model"
)

type IMediaPlayActionService interface {
	GetPlayAction(token string, cache cache.MediaCache) model.PlayAction
	GetPlayActionByMediaID(token string, mediaID string) model.PlayAction
}

type mediaPlayActionService struct {
	generalInfoCacheService   cache.IGeneralInfoCacheService
	contentAccessCacheService cache.IContentAccessCacheService
	mediaAccessCacheService   cache.IMediaAccessCacheService
	calisthenicsAuthService   calisthenics.ICalisthenicsAuthService

	mediaCacheService cache.IMediaCacheService
}

func NewMediaPlayActionService(generalInfoCacheService cache.IGeneralInfoCacheService,
	contentAccessCacheService cache.IContentAccessCacheService,
	mediaAccessCacheService cache.IMediaAccessCacheService,
	calisthenicsAuthService calisthenics.ICalisthenicsAuthService,
	mediaCacheService cache.IMediaCacheService,
) IMediaPlayActionService {
	return &mediaPlayActionService{
		generalInfoCacheService:   generalInfoCacheService,
		contentAccessCacheService: contentAccessCacheService,
		mediaAccessCacheService:   mediaAccessCacheService,
		calisthenicsAuthService:   calisthenicsAuthService,
		mediaCacheService:         mediaCacheService,
	}
}

func (s *mediaPlayActionService) GetPlayActionByMediaID(token string, mediaID string) model.PlayAction {
	media, err := s.mediaCacheService.GetByID(mediaID)
	if err != nil {
		return model.PlayAction{}
	}
	return s.GetPlayAction(token, media)
}

func (s *mediaPlayActionService) GetPlayAction(token string, cache cache.MediaCache) model.PlayAction {
	accessLevelInfo, accessLevelError := s.generalInfoCacheService.GetByID(string(config.AccessLevel))
	mediaAccess, mediaAccessError := s.mediaAccessCacheService.GetByID(cache.ID)
	contentAccess, contentAccessError := s.contentAccessCacheService.GetByID(cache.ContentID)

	// medianin acess level everyone ise playaction watch donuluyor.
	if mediaAccess.Audience == string(config.Everyone) {
		return model.PlayAction{
			ActionType: string(config.Watch),
			ButtonText: string(config.Watch), // TODO translation
		}
	}

	// mediaya ait bir access verisi yok ise content access verisine bakiliyor
	if mediaAccessError != nil {
		// mediaya sahip olan contentin access level everyone ise playaction watch donuluyor.
		if contentAccess.Audience == string(config.Everyone) {
			return model.PlayAction{
				ActionType: string(config.Watch),
				ButtonText: string(config.Watch), // TODO translation
			}
		}

		// medianin contentine de dair bir access verisi yok ise genel access level infosuna bakiliyor.
		if contentAccessError != nil {
			if accessLevelInfo == string(config.Everyone) {
				return model.PlayAction{
					ActionType: string(config.Watch),
					ButtonText: string(config.Watch), // TODO translation
				}
			}
		}
	}

	// buraya gelirse ya user login ya da subscription istenmektedir.
	_, errorResponse := s.calisthenicsAuthService.GetUser(token)
	if errorResponse != nil {
		return model.PlayAction{
			ActionType: string(config.Register),
			ButtonText: string(config.Register), // TODO translation
		}
	}

	// ileride subscription sistemi yazilinca burada userin sub var mi diye bakilacak

	// medianin access level user ise playaction watch donuluyor.
	if mediaAccess.Audience == string(config.User) {
		return model.PlayAction{
			ActionType: string(config.Watch),
			ButtonText: string(config.Watch), // TODO translation
		}
	}

	// mediaya ait bir access verisi yok ise content access verisine bakiliyor
	if mediaAccessError != nil {
		// mediaya sahip olan contentin access level user ise playaction watch donuluyor.
		if contentAccess.Audience == string(config.User) {
			return model.PlayAction{
				ActionType: string(config.Watch),
				ButtonText: string(config.Watch), // TODO translation
			}
		}

		// medianin contentine de dair bir access verisi yok ise genel access level infosuna bakiliyor.
		if contentAccessError != nil {
			if accessLevelInfo == string(config.User) {
				return model.PlayAction{
					ActionType: string(config.Watch),
					ButtonText: string(config.Watch), // TODO translation
				}
			}
		}
	}

	if accessLevelError != nil {
		return model.PlayAction{
			ActionType: string(config.Watch),
			ButtonText: string(config.Watch), // TODO translation
		}
	}

	return model.PlayAction{
		ActionType: string(config.Register),
		ButtonText: string(config.Register), // TODO translation
	}
}
