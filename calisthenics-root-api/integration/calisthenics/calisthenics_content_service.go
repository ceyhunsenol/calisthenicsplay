package calisthenics

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type ICalisthenicsContentService interface {
	Refresh(request RefreshRequest) *ErrorResponse
}

type calisthenicsContentService struct {
	baseURL url.URL
}

func NewCalisthenicsContentService() ICalisthenicsContentService {
	//baseURLString := viper.GetString("integration.calisthenics.content")
	baseURLString := "http://localhost:1322"
	baseURL, _ := url.Parse(baseURLString)
	return &calisthenicsContentService{
		baseURL: *baseURL,
	}
}

func (c *calisthenicsContentService) Refresh(request RefreshRequest) *ErrorResponse {
	requestURL := c.baseURL.JoinPath("v1/cache/refresh", request.CacheType, request.ID)
	response, err := http.Get(requestURL.String())
	if err != nil {
		return &ErrorResponse{Message: "Request error"}
	}
	defer response.Body.Close()

	// Yanıtın içeriğini oku
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &ErrorResponse{Message: "IO error"}
	}

	var errorResponse ErrorResponse
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		return &ErrorResponse{Message: "Deserialization error"}
	}

	if response.StatusCode != http.StatusOK {
		return &ErrorResponse{Message: errorResponse.Message}
	}
	return nil
}
