package api

import (
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type HLSMediaTestController struct {
}

func NewHLSMediaTestController() *HLSMediaTestController {
	return &HLSMediaTestController{}
}

func (u *HLSMediaTestController) InitHLSMediaTestRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/hls.m3u8", u.Test)
	v1.GET("/:id.m3u8", u.Test2)
	v1.GET("/license.key", u.Test3)
	v1.GET("/key", u.Key)
}

type A struct {
	Uri      string
	ExtInf   float64
	Discount bool
	Name     string
	IV       string
}

type License struct {
	Name string
	Iv   string
	Key  string
}

func (u *HLSMediaTestController) Key(c echo.Context) error {
	licenseKeyBytes, err := hex.DecodeString("7a48b3a83054f3541a4a5b03592efa53")
	if err != nil {
		fmt.Println("Hata:", err)
		return nil
	}
	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", licenseKeyBytes)
}

func (u *HLSMediaTestController) Test3(c echo.Context) error {
	id := c.QueryParam("id")
	l1 := License{
		Name: "output_480p0.ts",
		Key:  "7a48b3a83054f3541a4a5b03592efa53",
	}

	l2 := License{
		Name: "output_480p1.ts",
		Key:  "7a48b3a83054f3541a4a5b03592efa53",
	}

	l3 := License{
		Name: "output_480p2.ts",
		Key:  "7a48b3a83054f3541a4a5b03592efa53",
	}

	lMap := make(map[string]License)
	lMap["output_480p0"] = l1
	lMap["output_480p1"] = l2
	lMap["output_480p2"] = l3

	license := lMap[id]

	licenseKeyBytes, err := hex.DecodeString(license.Key)
	if err != nil {
		fmt.Println("Hata:", err)
		return nil
	}

	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", licenseKeyBytes)
}

func (u *HLSMediaTestController) Test2(c echo.Context) error {

	var builder strings.Builder

	a1 := A{
		Uri:      "http://localhost:1322/downloads/output_480p0.ts",
		ExtInf:   10.300000,
		Discount: true,
		Name:     "output_480p0",
		IV:       "0x00000000000000000000000000000000",
	}

	a2 := A{
		Uri:      "http://localhost:1322/downloads/output_480p1.ts",
		ExtInf:   9.966667,
		Discount: false,
		Name:     "output_480p1",
		IV:       "0x00000000000000000000000000000001",
	}

	a3 := A{
		Uri:      "http://localhost:1322/downloads/output_480p2.ts",
		ExtInf:   6.900000,
		Discount: false,
		Name:     "output_480p2",
		IV:       "0x00000000000000000000000000000002",
	}

	ax := make([]A, 0)
	ax = append(ax, a1, a2, a3)

	builder.WriteString("#EXTM3U\n")
	builder.WriteString(fmt.Sprintf("#EXT-X-VERSION:%d\n", 3))
	builder.WriteString(fmt.Sprintf("#EXT-X-TARGETDURATION:%d\n", 12))
	builder.WriteString(fmt.Sprintf("#EXT-X-MEDIA-SEQUENCE:%d\n", 0))
	builder.WriteString(fmt.Sprintf("#EXT-X-PLAYLIST-TYPE:%s\n", "VOD"))

	for _, seg := range ax {
		if seg.Discount {
			builder.WriteString("#EXT-X-DISCONTINUITY\n")
		}
		builder.WriteString(fmt.Sprintf("#EXTINF:%.2f,\n", seg.ExtInf))
		builder.WriteString(fmt.Sprintf("#EXT-X-KEY:METHOD=%s,URI=\"%s\",IV=%s\n", "AES-128", "http://localhost:1322/v1/license.key?id="+seg.Name, seg.IV))
		builder.WriteString(seg.Uri + "\n")
	}

	builder.WriteString("#EXT-X-ENDLIST\n")

	result := builder.String()

	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", []byte(result))
}

func (u *HLSMediaTestController) Test(c echo.Context) error {
	resolutions := []struct {
		Bandwidth    int
		AvgBandwidth int
		Codecs       string
		Resolution   string
		MediaURL     string
	}{
		{Bandwidth: 1938572, AvgBandwidth: 1390351, Codecs: "avc1.4d4028,mp4a.40.2", Resolution: "848x480", MediaURL: "/v1/480.m3u8"},
		{Bandwidth: 3589944, AvgBandwidth: 2279704, Codecs: "avc1.4d4028,mp4a.40.2", Resolution: "1272x720", MediaURL: "/v1/720.m3u8"},
		{Bandwidth: 6974954, AvgBandwidth: 4217619, Codecs: "avc1.4d4028,mp4a.40.2", Resolution: "1906x1080", MediaURL: "/v1/1080.m3u8"},
	}

	playlistContent := "#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-INDEPENDENT-SEGMENTS\n"

	for _, resolution := range resolutions {
		playlistContent += fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,AVERAGE-BANDWIDTH=%d,CODECS=\"%s\",RESOLUTION=%s\n%s\n",
			resolution.Bandwidth, resolution.AvgBandwidth, resolution.Codecs, resolution.Resolution, resolution.MediaURL)
	}

	return c.Blob(http.StatusOK, "application/vnd.apple.mpegurl", []byte(playlistContent))
}
