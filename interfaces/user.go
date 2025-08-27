package interfaces

import "go_api/models"

// IUserRepository defines the interface for user data operations.
type IUserRepository interface {
	Save(m *models.UserModel) error
	FindAll() ([]models.UserModel, error)
	Exists(m models.UserModel) (bool, error)
	FindByID(id uint) (models.UserModel, error)
	FindByEmail(email string) (models.UserModel, error)
	FindByMobileNo(mobileNo string) (models.UserModel, error)
	PartialUpdate(m models.UserModel) error
	Delete(m models.UserModel) error
}
