package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"calisthenics-root-api/pkg"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type IAuthService interface {
	Login(loginRequest model.LoginRequest) (*model.TokenModel, *model.ServiceError)
	RefreshToken(refreshTokenModel model.RefreshTokenRequest) (*model.TokenModel, *model.ServiceError)
}

type authService struct {
	userService           IUserService
	accessTokenSecretKey  string
	refreshTokenSecretKey string
}

func NewAuthService(userService IUserService) IAuthService {
	return &authService{
		userService:           userService,
		accessTokenSecretKey:  viper.GetString("security.jwt.access-token-secret-key"),
		refreshTokenSecretKey: viper.GetString("security.jwt.refresh-token-secret-key"),
	}
}

func (u *authService) getTokenModel(user data.User) (*model.TokenModel, *model.ServiceError) {
	accessToken, err := pkg.CreateToken(user.ID, time.Hour*1, u.accessTokenSecretKey)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to process."}
	}

	refreshToken, err := pkg.CreateToken(user.ID, time.Hour*720, u.refreshTokenSecretKey)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to process."}
	}
	return model.NewTokenModel(user.Username, accessToken, refreshToken), nil
}

func (u *authService) Login(loginRequest model.LoginRequest) (*model.TokenModel, *model.ServiceError) {
	var user data.User
	var err error
	if loginRequest.Email != "" {
		user, err = u.userService.GetByEmail(loginRequest.Email)
		if err != nil {
			return nil, &model.ServiceError{Code: http.StatusUnauthorized, Message: "User not found."}
		}
	} else if loginRequest.Username != "" {
		user, err = u.userService.GetByUsername(loginRequest.Username)
		if err != nil {
			return nil, &model.ServiceError{Code: http.StatusUnauthorized, Message: "User not found."}
		}
	}
	hashControlled := pkg.CheckPasswordHash(loginRequest.Password, user.Password)
	if !hashControlled {
		return nil, &model.ServiceError{Code: http.StatusUnauthorized, Message: "Incorrect password."}
	}
	return u.getTokenModel(user)
}

func (u *authService) RefreshToken(refreshTokenModel model.RefreshTokenRequest) (*model.TokenModel, *model.ServiceError) {
	userID, err := pkg.GetUserIDFromToken(refreshTokenModel.RefreshToken, u.refreshTokenSecretKey)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to process login."}
	}
	user, err := u.userService.GetById(userID)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusUnauthorized, Message: "User not found."}
	}
	return u.getTokenModel(*user)
}
