package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IHLSEncodingCacheOperations interface {
	SaveCacheHLSEncodingList() *model.ServiceError
	SaveCacheHLSEncoding(ID string) interface{}
}

type hlSEncodingCacheOperations struct {
	hlSEncodingCacheService cache.IHLSEncodingCacheService
	hlSEncodingService      IEncodingService
}

func NewHLSEncodingCacheOperations(
	hlSEncodingCacheService cache.IHLSEncodingCacheService,
	hlSEncodingService IEncodingService,
) IHLSEncodingCacheOperations {
	return &hlSEncodingCacheOperations{
		hlSEncodingCacheService: hlSEncodingCacheService,
		hlSEncodingService:      hlSEncodingService,
	}
}

func (o *hlSEncodingCacheOperations) SaveCacheHLSEncodingList() *model.ServiceError {
	o.hlSEncodingCacheService.RemoveAll()
	hLSMediaList, err := o.hlSEncodingService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	hlSEncodingCache := make([]cache.HLSEncodingCache, 0)
	for _, value := range hLSMediaList {
		fileCaches := make([]cache.HLSEncodingFileCache, 0)
		for _, val := range value.EncodingFiles {
			fileCache := cache.HLSEncodingFileCache{
				FileName:   val.FileName,
				EncodingID: val.EncodingID,
				IV:         val.IV,
				Ext:        val.Ext,
			}
			fileCaches = append(fileCaches, fileCache)
		}
		cac := cache.HLSEncodingCache{
			ID:         value.ID,
			LicenseKey: value.LicenseKey,
			MediaID:    value.MediaID,
			Files:      fileCaches,
		}
		hlSEncodingCache = append(hlSEncodingCache, cac)
	}
	o.hlSEncodingCacheService.SaveAllSlice(hlSEncodingCache)
	return nil
}

func (o *hlSEncodingCacheOperations) SaveCacheHLSEncoding(ID string) interface{} {
	o.hlSEncodingCacheService.Remove(ID)
	hlsEncoding, err := o.hlSEncodingService.GetByID(ID)
	if err != nil {
		return nil
	}
	fileCaches := make([]cache.HLSEncodingFileCache, 0)
	for _, val := range hlsEncoding.EncodingFiles {
		fileCache := cache.HLSEncodingFileCache{
			FileName:   val.FileName,
			EncodingID: val.EncodingID,
			IV:         val.IV,
			Ext:        val.Ext,
		}
		fileCaches = append(fileCaches, fileCache)
	}
	hlsCache := cache.HLSEncodingCache{
		ID:         hlsEncoding.ID,
		LicenseKey: hlsEncoding.LicenseKey,
		MediaID:    hlsEncoding.MediaID,
		Files:      fileCaches,
	}
	o.hlSEncodingCacheService.Save(hlsCache)
	return &hlsCache
}
