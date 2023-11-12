package calisthenics

import (
	"calisthenics-content-api/cache"
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

type ICalisthenicsAuthService interface {
	GetUser(token string) (UserResponse, *ErrorResponse)
}

type calisthenicsContentService struct {
	baseURL             url.URL
	client              http.Client
	limitedCacheService cache.ILimitedCacheService
}

func NewCalisthenicsAuthService(limitedCacheService cache.ILimitedCacheService) ICalisthenicsAuthService {
	baseURLString := viper.GetString("integration.calisthenics.auth")
	baseURL, _ := url.Parse(baseURLString)
	return &calisthenicsContentService{
		baseURL:             *baseURL,
		client:              http.Client{},
		limitedCacheService: limitedCacheService,
	}
}

func (c *calisthenicsContentService) GetUser(token string) (UserResponse, *ErrorResponse) {
	obj, err := c.limitedCacheService.GetByID(token)
	if err == nil {
		return obj.(UserResponse), nil
	}
	requestURL := c.baseURL.JoinPath("v1/user/user-info")

	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return UserResponse{}, &ErrorResponse{Message: "Request error"}
	}

	req.Header.Add("X-Authorization", token)
	response, err := c.client.Do(req)
	if err != nil {
		return UserResponse{}, &ErrorResponse{Message: "Request error"}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return UserResponse{}, &ErrorResponse{Message: "IO error"}
	}

	if response.StatusCode != http.StatusOK {
		var errorResponse *ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return UserResponse{}, &ErrorResponse{Message: "Deserialization error"}
		}
		return UserResponse{}, errorResponse
	}

	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return UserResponse{}, &ErrorResponse{Message: "Deserialization error"}
	}
	limitedCache := cache.NewLimitedCache(token, userResponse)
	c.limitedCacheService.Save(limitedCache)
	return userResponse, nil
}
