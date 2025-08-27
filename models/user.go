package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name                 string    `json:"name" binding:"required"`
	Password             string    `json:"password" binding:"required"`
	Email                string    `json:"email"`
	UserTypeId           int       `json:"user_type_id"`
	Role                 string    `json:"role"`
	MobileNo             string    `json:"mobile_no" binding:"required"`
	Address              string    `json:"address"`
	ProfilePhoto         string    `json:"profile_photo"`
	EmailVerifiedAt      time.Time `json:"email_verified_at"`
	RememberToken        string    `json:"remember_token"`
	PhoneVerifiedAt      time.Time `json:"phone_verified_at"`
	Gender               string    `json:"gender" binding:"required"`
	ShouldForgetPassword int       `json:"should_forget_password"`
	Status               bool      `gorm:"default:false" json:"status"`
	
	Devices              []DeviceModel `gorm:"foreignKey:UserID"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserModelDto struct {
	gorm.Model
	Name                 string  `json:"name" binding:"required"`
	Email                string  `json:"email"`
	UserTypeId           int     `json:"user_type_id"`
	Role                 string  `json:"role"`
	MobileNo             string  `json:"mobile_no" binding:"required"`
	Address              string  `json:"address"`
	ProfilePhoto         string  `json:"profile_photo"`
	EmailVerifiedAt      string  `json:"email_verified_at"`
	RememberToken        string  `json:"remember_token"`
	PhoneVerifiedAt      string  `json:"phone_verified_at"`
	Gender               string  `json:"gender" binding:"required"`
	ShouldForgetPassword int     `json:"should_forget_password"`
	Status               bool    `gorm:"default:false" json:"status"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	DeletedAt            *string `json:"deleted_at"`
}

func (m *UserModel) ToDto() UserModelDto {
	var modelDto UserModelDto
	// Custom copying logic for non-directly mappable fields
	copier.CopyWithOption(&modelDto, &m, copier.Option{IgnoreEmpty: true})
	// Custom conversion for HumanReadableTime
	modelDto.CreatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	modelDto.UpdatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	if !m.DeletedAt.Time.IsZero() {
		d := m.DeletedAt.Time.Format("2006-01-02 15:04:05")
		modelDto.DeletedAt = &d
	}
	return modelDto
}
