package apirequests

import "go_api/models"

type CreateUserRequest struct {
	Name                 string `json:"name" form:"name" binding:"required"`
	Password             string `json:"password" form:"password" binding:"required"`
	Email                string `json:"email" form:"email" binding:"required"`
	MobileNo             string `json:"mobile_no" form:"mobile_no" binding:"required"`
	UserTypeId           int    `json:"user_type_id" form:"user_type_id" binding:"required"`
	Role                 string `json:"role" form:"role" binding:"required"`
	Address              string `json:"address" form:"address" binding:"required"`
	ProfilePhoto         string `json:"profile_photo" form:"profile_photo"`
	EmailVerifiedAt      string `json:"email_verified_at" form:"email_verified_at"`
	RememberToken        string `json:"remember_token" form:"remember_token"`
	PhoneVerifiedAt      string `json:"phone_verified_at" form:"phone_verified_at"`
	Gender               string `json:"gender" form:"gender" binding:"required"`
	ShouldForgetPassword int    `json:"should_forget_password" form:"should_forget_password"`
	Status               bool   `json:"status" form:"status"`
}

func (r *CreateUserRequest) ToModel() models.UserModel {
	return models.UserModel{
		Name:         r.Name,
		Email:        r.Email,
		MobileNo:     r.MobileNo,
		Password:     r.Password,
		UserTypeId:   r.UserTypeId,
		Gender:       r.Gender,
		ProfilePhoto: r.ProfilePhoto,
		Address:      r.Address,
	}
}

type UpdateUserRequest struct {
	Name                 string `json:"name" binding:"required"`
	Password             string `json:"password" binding:"required"`
	Email                string `json:"email" binding:"required"`
	MobileNo             string `json:"mobile_no" binding:"required"`
	UserTypeId           int    `json:"user_type_id" binding:"required"`
	Role                 string `json:"role" binding:"required"`
	Address              string `json:"address" binding:"required"`
	ProfilePhoto         string `json:"profile_photo"`
	EmailVerifiedAt      string `json:"email_verified_at"`
	RememberToken        string `json:"remember_token"`
	PhoneVerifiedAt      string `json:"phone_verified_at"`
	Gender               string `json:"gender" binding:"required"`
	ShouldForgetPassword int    `json:"should_forget_password"`
	Status               bool   `json:"status"`
}

func (r *UpdateUserRequest) ToModel() models.UserModel {
	return models.UserModel{
		Name:         r.Name,
		Email:        r.Email,
		MobileNo:     r.MobileNo,
		Password:     r.Password,
		UserTypeId:   r.UserTypeId,
		Gender:       r.Gender,
		ProfilePhoto: r.ProfilePhoto,
		Address:      r.Address,
		Status:       r.Status,
		// EmailVerifiedAt: r.EmailVerifiedAt,
		// PhoneVerifiedAt: r.PhoneVerifiedAt,
		RememberToken:        r.RememberToken,
		ShouldForgetPassword: r.ShouldForgetPassword,
		Role:                 r.Role,
	}
}
