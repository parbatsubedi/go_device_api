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

func (r *UserRepository) FindAll() ([]models.UserModel, error) {
	var modelList []models.UserModel
	result := r.db.Find(&modelList)
	if result.Error != nil {
		return nil, result.Error
	}
	return modelList, nil
}

func (r *UserRepository) Exists(m models.UserModel) (bool, error) {
	var model models.UserModel

	result := r.db.First(&model, m.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (r *UserRepository) FindByID(id uint) (models.UserModel, error) {
	var model models.UserModel
	result := r.db.First(&model, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model, result.Error
		}
		return model, result.Error
	}

	return model, nil
}

func (r *UserRepository) FindByEmail(email string) (models.UserModel, error) {
	var model models.UserModel
	result := r.db.Where("email=?", email).First(&model)
	if result.Error != nil {
		return model, result.Error
	}
	return model, nil
}

func (r *UserRepository) FindByMobileNo(mobileNo string) (models.UserModel, error) {
	var model models.UserModel
	result := r.db.Where("mobile_no=?", mobileNo).First(&model)
	if result.Error != nil {
		return model, result.Error
	}
	return model, nil
}

func (r *UserRepository) PartialUpdate(m models.UserModel) error {
	if updateErr := r.db.Model(&m).Updates(m).Error; updateErr != nil {
		return updateErr
	}
	return nil
}

func (r *UserRepository) Delete(m models.UserModel) error {
	if deleteErr := r.db.Delete(&m).Error; deleteErr != nil {
		return deleteErr
	}
	return nil
}
