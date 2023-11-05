package main

func main() {
	Run()
	//mediaCache1 := &cache.MediaCache{
	//	ID:        "media1",
	//	Active:    true,
	//	ContentID: "content1",
	//}
	//mediaCache2 := &cache.MediaCache{
	//	ID:        "media1",
	//	Active:    true,
	//	ContentID: "content1",
	//}
	//mediaCache3 := &cache.MediaCache{
	//	ID:        "media1",
	//	Active:    true,
	//	ContentID: "content1",
	//}
	//mediaCache4 := &cache.MediaCache{
	//	ID:        "media1",
	//	Active:    true,
	//	ContentID: "content2",
	//}
	//
	//medias := []*cache.MediaCache{mediaCache1, mediaCache2, mediaCache3, mediaCache4}
	//
	//contentCache := &cache.ContentCache{
	//	ID:     "content1",
	//	Active: false,
	//}
	//
	//contentCache2 := &cache.ContentCache{
	//	ID:     "content2",
	//	Active: false,
	//}
	//
	//contents := map[string]*cache.ContentCache{
	//	"content1": contentCache,
	//	"content2": contentCache2,
	//}
	//
	//grouped := pkg.GroupByField(medias, func(media *cache.MediaCache) string {
	//	return media.ContentID
	//})
	//
	//for key, value := range grouped {
	//	content, err := contents[key]
	//	if !err || content.Active {
	//		continue
	//	}
	//
	//	for _, media := range value {
	//		media.Active = false
	//	}
	//}
	//
	//for _, media := range medias {
	//	fmt.Print(media)
	//}
}
