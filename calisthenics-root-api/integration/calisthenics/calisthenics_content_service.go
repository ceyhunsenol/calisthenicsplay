package calisthenics

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

type ICalisthenicsContentService interface {
	Refresh(request RefreshRequest) *ErrorResponse
	RefreshWithMedias(contentID string) *ErrorResponse
}

type calisthenicsContentService struct {
	baseURL url.URL
}

func NewCalisthenicsContentService() ICalisthenicsContentService {
	baseURLString := viper.GetString("integration.calisthenics.content")
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

func (c *calisthenicsContentService) RefreshWithMedias(contentID string) *ErrorResponse {
	requestURL := c.baseURL.JoinPath("v1/cache/refresh/content-with-medias", contentID)
	response, err := http.Get(requestURL.String())
	if err != nil {
		return &ErrorResponse{Message: "Request error"}
	}
	defer response.Body.Close()

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
