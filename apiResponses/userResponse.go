package apiresponses

import "go_api/models"

type UserResponse struct {
	SuccessResponse

	Data models.UserModelDto `json:"data"`
}

func NewUserResponse() UserResponse {
	return UserResponse{
		NewSuccessResponse(),
		models.UserModelDto{},
	}
}

type UserListResponse struct {
	SuccessResponse
	// Data models.Pagination `json:"data"`
	Data []models.UserModelDto `json:"data"`
}

func NewUserListResponse() UserListResponse {
	return UserListResponse{
		NewSuccessResponse(),
		// models.Pagination{},
		[]models.UserModelDto{},
	}
}
