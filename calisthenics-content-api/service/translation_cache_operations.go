package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type ITranslationCacheOperations interface {
	SaveCacheTranslationList() *model.ServiceError
	SaveCacheTranslation(code string) interface{}
}

type translationCacheOperations struct {
	translationService      ITranslationService
	translationCacheService cache.ITranslationCacheService
}

func NewTranslationCacheOperations(translationCacheService cache.ITranslationCacheService,
	translationService ITranslationService,
) ITranslationCacheOperations {
	return &translationCacheOperations{
		translationCacheService: translationCacheService,
		translationService:      translationService,
	}
}

func (o *translationCacheOperations) SaveCacheTranslationList() *model.ServiceError {
	o.translationCacheService.RemoveAll()
	contentCodes, err := o.translationService.GetAllDistinctCodesByDomain("API_CONTENT")
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	langCaches := make([]cache.MultiLangCache, 0)
	for _, code := range contentCodes {
		translation, codeError := o.translationService.GetAllByCode(code)
		if codeError != nil {
			return &model.ServiceError{Code: http.StatusNotFound, Message: "Translation not found"}
		}

		langCache := cache.NewMultiLangCache(code)
		for _, d := range translation {
			if d.Active {
				langCache.SetByLang(d.LangCode, d.Translate)
			}
		}
		langCache.SetByLang("base", langCache.GetByLang("EN"))
		langCaches = append(langCaches, *langCache)
	}
	o.translationCacheService.SaveAllSlice(langCaches)
	return nil
}

func (o *translationCacheOperations) SaveCacheTranslation(code string) interface{} {
	o.translationCacheService.Remove(code)
	translation, err := o.translationService.GetAllByCode(code)
	if err != nil {
		return nil
	}
	langCache := cache.NewMultiLangCache(code)
	for _, d := range translation {
		if d.Active {
			langCache.SetByLang(d.LangCode, d.Translate)
		}
	}
	langCache.SetByLang("base", langCache.GetByLang("en"))
	o.translationCacheService.Save(*langCache)
	return langCache
}
