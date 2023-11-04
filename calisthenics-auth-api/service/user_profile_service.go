package service

import (
	"calisthenics-auth-api/data"
	"calisthenics-auth-api/model"
	"net/http"
	"time"
)

const layout = "2006-01-02"

type IUserProfileService interface {
	UserProfile(request model.UserProfileRequest) *model.ServiceError
}

type userProfileService struct {
	profileService IProfileService
}

func NewUserProfileService(profileService IProfileService) IUserProfileService {
	return &userProfileService{profileService: profileService}
}

func (u *userProfileService) UserProfile(request model.UserProfileRequest) *model.ServiceError {
	date, err := time.Parse(layout, request.DateOfBirth)
	if err != nil {
		return nil
	}
	profile := data.Profile{
		UserID:      request.UserID,
		DateOfBirth: date,
		AvatarURL:   request.AvatarURL,
		Bio:         request.Bio,
	}
	_, err = u.profileService.SaveOrUpdate(profile)
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "An error occurred while saving the profile."}
	}
	return nil
}
