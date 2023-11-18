package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
	"calisthenics-content-api/integration/calisthenics"
	"calisthenics-content-api/model"
)

type IMediaPlayActionService interface {
	GetPlayAction(request model.PlayActionRequest, cache cache.MediaCache) model.PlayAction
	GetPlayActionByMediaID(request model.PlayActionRequest, mediaID string) model.PlayAction
}

type mediaPlayActionService struct {
	generalInfoCacheService   cache.IGeneralInfoCacheService
	contentAccessCacheService cache.IContentAccessCacheService
	mediaAccessCacheService   cache.IMediaAccessCacheService
	calisthenicsAuthService   calisthenics.ICalisthenicsAuthService
	translationCacheService   cache.ITranslationCacheService
	mediaCacheService         cache.IMediaCacheService
}

func NewMediaPlayActionService(generalInfoCacheService cache.IGeneralInfoCacheService,
	contentAccessCacheService cache.IContentAccessCacheService,
	mediaAccessCacheService cache.IMediaAccessCacheService,
	calisthenicsAuthService calisthenics.ICalisthenicsAuthService,
	mediaCacheService cache.IMediaCacheService,
	translationCacheService cache.ITranslationCacheService,
) IMediaPlayActionService {
	return &mediaPlayActionService{
		generalInfoCacheService:   generalInfoCacheService,
		contentAccessCacheService: contentAccessCacheService,
		mediaAccessCacheService:   mediaAccessCacheService,
		calisthenicsAuthService:   calisthenicsAuthService,
		mediaCacheService:         mediaCacheService,
		translationCacheService:   translationCacheService,
	}
}

func (s *mediaPlayActionService) GetPlayActionByMediaID(request model.PlayActionRequest, mediaID string) model.PlayAction {
	media, err := s.mediaCacheService.GetByID(mediaID)
	if err != nil {
		registerButtonTranslate := s.translationCacheService.GetByLang("PLAY_ACTION_BUTTON_REGISTER", request.LangCode)
		return model.PlayAction{
			ActionType: config.Register,
			ButtonText: registerButtonTranslate,
		}
	}
	return s.GetPlayAction(request, media)
}

func (s *mediaPlayActionService) GetPlayAction(request model.PlayActionRequest, cache cache.MediaCache) model.PlayAction {
	accessLevelInfo, accessLevelError := s.generalInfoCacheService.GetByID(string(config.AccessLevel))
	mediaAccess, mediaAccessError := s.mediaAccessCacheService.GetByID(cache.ID)
	contentAccess, contentAccessError := s.contentAccessCacheService.GetByID(cache.ContentID)

	watchTranslate := s.translationCacheService.GetByLang("PLAY_ACTION_BUTTON_WATCH", request.LangCode)
	// medianin acess level everyone ise playaction watch donuluyor.
	if mediaAccess.Audience == string(config.Everyone) {
		return model.PlayAction{
			ActionType: config.Watch,
			ButtonText: watchTranslate,
		}
	}

	// mediaya ait bir access verisi yok ise content access verisine bakiliyor
	if mediaAccessError != nil {
		// mediaya sahip olan contentin access level everyone ise playaction watch donuluyor.
		if contentAccess.Audience == string(config.Everyone) {
			return model.PlayAction{
				ActionType: config.Watch,
				ButtonText: watchTranslate,
			}
		}

		// medianin contentine de dair bir access verisi yok ise genel access level infosuna bakiliyor.
		if contentAccessError != nil {
			if accessLevelInfo == string(config.Everyone) {
				return model.PlayAction{
					ActionType: config.Watch,
					ButtonText: watchTranslate,
				}
			}
		}
	}

	registerButtonTranslate := s.translationCacheService.GetByLang("PLAY_ACTION_BUTTON_REGISTER", request.LangCode)
	// buraya gelirse ya user login ya da subscription istenmektedir.
	_, errorResponse := s.calisthenicsAuthService.GetUser(request.Token)
	if errorResponse != nil {
		return model.PlayAction{
			ActionType: config.Register,
			ButtonText: registerButtonTranslate,
		}
	}

	// ileride subscription sistemi yazilinca burada userin sub var mi diye bakilacak

	// medianin access level user ise playaction watch donuluyor.
	if mediaAccess.Audience == string(config.User) {
		return model.PlayAction{
			ActionType: config.Watch,
			ButtonText: watchTranslate,
		}
	}

	// mediaya ait bir access verisi yok ise content access verisine bakiliyor
	if mediaAccessError != nil {
		// mediaya sahip olan contentin access level user ise playaction watch donuluyor.
		if contentAccess.Audience == string(config.User) {
			return model.PlayAction{
				ActionType: config.Watch,
				ButtonText: watchTranslate,
			}
		}

		// medianin contentine de dair bir access verisi yok ise genel access level infosuna bakiliyor.
		if contentAccessError != nil {
			if accessLevelInfo == string(config.User) {
				return model.PlayAction{
					ActionType: config.Watch,
					ButtonText: watchTranslate,
				}
			}
		}
	}

	if accessLevelError != nil {
		return model.PlayAction{
			ActionType: config.Watch,
			ButtonText: watchTranslate,
		}
	}

	return model.PlayAction{
		ActionType: config.Register,
		ButtonText: registerButtonTranslate,
	}
}
