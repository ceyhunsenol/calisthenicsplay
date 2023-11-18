package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IGeneralInfoCacheOperations interface {
	SaveCacheGeneralInfos() *model.ServiceError
	SaveCacheGeneralInfo(ID string) interface{}
}

type generalInfoCacheOperations struct {
	generalInfoCacheService cache.IGeneralInfoCacheService
	generalInfoService      IGeneralInfoService
}

func NewGeneralInfoCacheOperations(generalInfoCacheService cache.IGeneralInfoCacheService,
	generalInfoService IGeneralInfoService,
) IGeneralInfoCacheOperations {
	return &generalInfoCacheOperations{
		generalInfoCacheService: generalInfoCacheService,
		generalInfoService:      generalInfoService,
	}
}

func (o *generalInfoCacheOperations) SaveCacheGeneralInfos() *model.ServiceError {
	o.generalInfoCacheService.RemoveAll()
	infos, err := o.generalInfoService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	infoCaches := make([]cache.GeneralInfoCache, 0)
	for _, value := range infos {
		contentCache := cache.GeneralInfoCache{
			Key:   value.InfoKey,
			Value: value.InfoValue,
		}
		infoCaches = append(infoCaches, contentCache)
	}
	o.generalInfoCacheService.SaveAllSlice(infoCaches)
	return nil
}

func (o *generalInfoCacheOperations) SaveCacheGeneralInfo(ID string) interface{} {
	o.generalInfoCacheService.Remove(ID)
	info, err := o.generalInfoService.GetByID(ID)
	if err != nil {
		return nil
	}
	infoCache := cache.GeneralInfoCache{
		Key:   info.InfoKey,
		Value: info.InfoValue,
	}
	o.generalInfoCacheService.Save(infoCache)
	return &infoCache
}
