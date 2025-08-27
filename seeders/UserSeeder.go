package seeders

import (
	"go_api/database"
	"go_api/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserSeeder struct {
}

func (r UserSeeder) Seed() error {
	var count int64
	db := database.RootDatabase.DB
	db.Model(&models.UserModel{}).Count(&count)
	if count == 0 {
		password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

		user := models.UserModel{
			Name:            "Parbat Subedi",
			Email:           "parbatsubedi000@gmail.com",
			MobileNo:        "9843723270",
			Password:        string(password),
			Status:          true,
			Gender:          "Male",
			EmailVerifiedAt: time.Now(),
			PhoneVerifiedAt: time.Now(),
			// ShouldChangePassword: 0,
			UserTypeId: 1,
			// UserCategory:         "Admin",
		}
		db.Create(&user)

	}
	return nil
}
