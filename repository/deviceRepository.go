package repository

import (
	"go_api/database"
	"go_api/models"

	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{
		db: database.RootDatabase.DB,
	}
}

func (r *DeviceRepository) Save(m *models.DeviceModel) error {
	result := r.db.Save(&m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *DeviceRepository) FindAll() ([]models.DeviceModel, error) {
	var modelList []models.DeviceModel
	err := r.db.Preload("User").Find(&modelList).Error
	return modelList, err
}

func (r *DeviceRepository) FindByID(id uint) (models.DeviceModel, error) {
	var model models.DeviceModel
	err := r.db.Preload("User").First(&model, id).Error
	return model, err
}

func (r *DeviceRepository) Exists(m models.DeviceModel) bool {
	var model models.DeviceModel

	result := r.db.First(&model, m.ID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}
		panic(result.Error.Error())
	}

	return true
}

func (r *DeviceRepository) PartialUpdate(m *models.DeviceModel) error {
	return r.db.Model(&m).Updates(m).Error
}

func (r *DeviceRepository) Delete(m models.DeviceModel) error {
	return r.db.Delete(&m).Error
}

func (r *DeviceRepository) FindByIMEI(imei string) (models.DeviceModel, error) {
	var model models.DeviceModel
	err := r.db.Where("device_imei1 = ? OR device_imei2 = ?", imei, imei).First(&model).Error
	return model, err
}

func (r *DeviceRepository) FindByUserID(userID uint) ([]models.DeviceModel, error) {
	var devices []models.DeviceModel
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}

func (r *DeviceRepository) FindByDeviceName(deviceName string) (models.DeviceModel, error) {
	var model models.DeviceModel
	err := r.db.Where("device_name=?", deviceName).First(&model).Error
	return model, err
}

func (r *DeviceRepository) FindByManufacturer(manufacturer string) (models.DeviceModel, error) {
	var model models.DeviceModel
	err := r.db.Where("manufacturer=?", manufacturer).First(&model).Error
	return model, err
}

func (r *DeviceRepository) FindByDeviceModel(deviceModel string) (models.DeviceModel, error) {
	var model models.DeviceModel
	err := r.db.Where("device_model=?", deviceModel).First(&model).Error
	return model, err
}
