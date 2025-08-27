package auth

import (
	apiresponses "go_api/apiResponses"
	"go_api/models"
)

type LoginSuccessResponse struct {
	apiresponses.SuccessResponse
	Token string              `json:"access_token"`
	User  models.UserModelDto `json:"user"`
}

func NewLoginSuccessResponse() LoginSuccessResponse {
	return LoginSuccessResponse{
		apiresponses.NewSuccessResponse(),
		"",
		models.UserModelDto{},
		// []string{},
	}
}

type CurrentUserResponse struct {
	apiresponses.SuccessResponse
	UserModelDto models.UserModelDto `json:"data"`
}

func NewCurrentUserResponse() CurrentUserResponse {
	return CurrentUserResponse{
		apiresponses.NewSuccessResponse(),
		models.UserModelDto{},
	}
}
