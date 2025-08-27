package authrequest

import (
	"go_api/models"

	"github.com/jinzhu/copier"
)

type RegisterRequest struct {
	Name                 string `json:"name" binding:"required"`
	Password             string `json:"password" binding:"required"`
	Email                string `json:"email" binding:"required"`
	MobileNo             string `json:"mobile_no" binding:"required"`
	UserTypeId           int    `json:"user_type_id" binding:"required"`
	Role                 string `json:"role" binding:"required"`
	Address              string `json:"address" binding:"required"`
	ProfilePhoto         string `json:"profile_photo"`
	Gender               string `json:"gender" binding:"required"`
	ShouldForgetPassword int    `json:"should_forget_password"`
}

func (r *RegisterRequest) ToModel() models.UserModel {
	var model models.UserModel
	copier.CopyWithOption(&model, &r, copier.Option{IgnoreEmpty: true})
	return model
}
