package main

import (
	"calisthenics-content-api/cache"
	"fmt"
)

func main() {
	//Run()
	cacheManager := cache.NewCacheManager()
	mediaCacheService := cache.NewMediaCacheService(cacheManager)

	//mediaCache2 := cache.MediaCache{
	//	ID:              "Media2",
	//	DescriptionCode: "Desc2",
	//	URL:             "URL2",
	//	Type:            "VIDEO",
	//	ContentID:       "Content1",
	//}
	//mediaCache := cache.MediaCache{
	//	ID:              "Media1",
	//	DescriptionCode: "Desc1",
	//	URL:             "URL1",
	//	Type:            "VIDEO",
	//	ContentID:       "Content1",
	//}
	//mediaCache3 := cache.MediaCache{
	//	ID:              "Media3",
	//	DescriptionCode: "Desc3",
	//	URL:             "URL3",
	//	Type:            "VIDEO",
	//	ContentID:       "Content1",
	//}
	//mediaCache4 := cache.MediaCache{
	//	ID:              "Media4",
	//	DescriptionCode: "Desc4",
	//	URL:             "URL4",
	//	Type:            "VIDEO",
	//	ContentID:       "Content1",
	//}
	//mediaCache5 := cache.MediaCache{
	//	ID:              "Media5",
	//	DescriptionCode: "Desc5",
	//	URL:             "URL5",
	//	Type:            "VIDEO",
	//	ContentID:       "Content5",
	//}
	//mediaCacheService.Save(mediaCache)
	//mediaCacheService.Save(mediaCache2)
	//mediaCacheService.Save(mediaCache3)
	//mediaCacheService.Save(mediaCache4)
	//mediaCacheService.Save(mediaCache5)
	//id, err := mediaCacheService.GetByID("Media1")
	//if err != nil {
	//	return
	//}
	//
	//all := mediaCacheService.GetAll()
	//for _, i := range all {
	//	fmt.Println(i)
	//}
	//fmt.Println(id)
	//
	//mediaIDs := mediaCacheService.GetAllByContentID("Content1", "VIDEO")
	//for _, d := range mediaIDs {
	//	fmt.Println(d)
	//}
	//
	//fmt.Println("-----------------------------")
	//mediaCacheService.Remove("Media1")
	//
	//mediaIDs = mediaCacheService.GetAllByContentID("Content1", "VIDEO")
	//for _, d := range mediaIDs {
	//	fmt.Println(d)
	//}

	fmt.Println("-----------------------------")
	cachesx := []cache.MediaCache{
		{ID: "1", ContentID: "A", Type: "VIDEO", Active: true},
		{ID: "2", ContentID: "B", Type: "VIDEO", Active: true},
		{ID: "3", ContentID: "A", Type: "VIDEO", Active: false},
		{ID: "4", ContentID: "C", Type: "VIDEO", Active: true},
		{ID: "5", ContentID: "A", Type: "GIF", Active: true},
	}
	//
	mediaCacheService.SaveAllSlice(cachesx)
	//
	//all := mediaCacheService.GetAll()
	//for _, mediaCache := range all {
	//	fmt.Println(mediaCache)
	//}
	//
	//fmt.Println("----------------------")
	//
	//ids := mediaCacheService.GetAllByContentID("A", "VIDEO")
	//fmt.Println(ids)
	//fmt.Println("----------------------")
	//ids = mediaCacheService.GetAllByContentID("A", "GIF")
	//fmt.Println(ids)
	//
	//fmt.Println("----------------------")
	//
	//ids = mediaCacheService.GetAllByContentID("B", "VIDEO")
	//fmt.Println(ids)
	//
	//fmt.Println("----------------------")
	//
	//ids = mediaCacheService.GetAllByContentID("C", "VIDEO")
	//fmt.Println(ids)
	//caches := cache.MediaCache{
	//	ID:        "MediaCache1",
	//	Type:      "VIDEO",
	//	Active:    true,
	//	ContentID: "content1",
	//}
	//
	//langCache := cache.NewMultiLangCache("Media1_cache_Code")
	//langCache.SetByLang("base", "This is media cache1 description")
	//langCache.SetByLang("en", "This is media cache1 description")
	//langCache.SetByLang("tr", "Bu media cache1 açıklamasıdır")
	//
	//caches.DescriptionMultiLang = langCache
	//lang := caches.DescriptionMultiLang.GetByLang("tr")
	//fmt.Println(lang)

	genreCache1 := cache.GenreCache{
		ID:     "GenreCache1",
		Type:   "CATEGORY",
		Active: true,
	}

	genreCache2 := cache.GenreCache{
		ID:     "GenreCache2",
		Type:   "CATEGORY",
		Active: true,
	}

	genreCache3 := cache.GenreCache{
		ID:     "GenreCache3",
		Type:   "BREAKDOWN",
		Active: true,
	}

	genreCache4 := cache.GenreCache{
		ID:     "GenreCache4",
		Type:   "CATEGORY",
		Active: true,
	}

	genreCacheService := cache.NewGenreCacheService(cacheManager)
	genreCacheService.SaveAll(genreCache1, genreCache2, genreCache3, genreCache4)

	all := genreCacheService.GetAll()
	fmt.Println(all)

	byType := genreCacheService.GetAllByType("CATEGORY")
	fmt.Println(byType)

}
