package service

import (
	"calisthenics-auth-api/data"
	"calisthenics-auth-api/model"
	"calisthenics-auth-api/pkg"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type IAuthService interface {
	Register(request model.RegisterRequest) (*model.TokenModel, *model.ServiceError)
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

func (u *authService) Register(request model.RegisterRequest) (*model.TokenModel, *model.ServiceError) {
	existsEmail, err := u.userService.EmailExists(request.Email)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "An error occurred while checking the email."}
	}
	if existsEmail {
		return nil, &model.ServiceError{Code: http.StatusUnauthorized, Message: "A user with this email already exists."}
	}

	existsUsername, err := u.userService.UsernameExists(request.Username)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "An error occurred while checking the username."}
	}
	if existsUsername {
		return nil, &model.ServiceError{Code: http.StatusUnauthorized, Message: "A user with this username already exists."}
	}

	hash, err := pkg.HashPassword(request.Password)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to process register."}
	}

	user := data.User{
		Username: request.Username,
		Password: hash,
		Email:    request.Email,
	}
	savedUser, err := u.userService.Save(user)
	if err != nil {
		return nil, &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to process register."}
	}
	return u.getTokenModel(*savedUser)
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
