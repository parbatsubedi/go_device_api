package repository

import (
	"errors"
	"go_api/database"
	"go_api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository initialization from tenant
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.RootDatabase.DB,
	}
}

func (r *UserRepository) Save(m *models.UserModel) error {
	result := r.db.Save(&m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindAll() []models.UserModel {
	var modelList []models.UserModel
	r.db.Find(&modelList)
	return modelList
}

func (r *UserRepository) Exists(m models.UserModel) bool {
	var model models.UserModel

	result := r.db.First(&model, m.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}

		panic(result.Error.Error())
	}

	return true
}

func (r *UserRepository) FindByID(id uint) (models.UserModel, bool) {
	var model models.UserModel
	result := r.db.First(&model, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model, false
		}

		panic(result.Error.Error())
	}

	return model, true

}

func (r *UserRepository) FindByEmail(email string) (models.UserModel, bool) {
	var model models.UserModel
	result := r.db.Where("email=?", email).First(&model)
	if result.Error != nil {
		return model, false
	}
	return model, true
}

func (r *UserRepository) FindByMobileNo(mobileNo string) (models.UserModel, bool) {
	var model models.UserModel
	result := r.db.Where("mobile_no=?", mobileNo).First(&model)
	if result.Error != nil {
		return model, false
	}
	return model, true
}

func (r *UserRepository) PartialUpdate(m models.UserModel) bool {
	if updateErr := r.db.Model(&m).Updates(m).Error; updateErr != nil {
		return false
	}
	return true
}

func (r *UserRepository) Delete(m models.UserModel) bool {
	if deleteErr := r.db.Delete(&m).Error; deleteErr != nil {
		return false
	}
	return true
}
