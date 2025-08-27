package apirequests

import "go_api/models"

type DeviceCreateRequest struct {
	DeviceName   string `json:"name" form:"name" binding:"required" gorm:"unique"`
	DeviceIMEI1  string `json:"device_imei1" form:"device_imei1" binding:"required" gorm:"unique"`
	DeviceIMEI2  string `json:"device_imei2" form:"device_imei2" gorm:"unique"`
	Manufacturer string `json:"manufacturer" form:"manufacturer" binding:"required"`
	DeviceModel  string `json:"device_model" form:"device_model" binding:"required"`
}

func (r *DeviceCreateRequest) ToModel() models.DeviceModel {
	return models.DeviceModel{
		DeviceName:   r.DeviceName,
		DeviceIMEI1:  r.DeviceIMEI1,
		DeviceIMEI2:  r.DeviceIMEI2,
		Manufacturer: r.Manufacturer,
		DeviceModel:  r.DeviceModel,
	}
}

type DeviceUpdateRequest struct {
	Name         string `json:"name" form:"name" binding:"required" gorm:"unique"`
	DeviceIMEI1  string `json:"device_imei1" form:"device_imei1" binding:"required" gorm:"unique"`
	DeviceIMEI2  string `json:"device_imei2" form:"device_imei2" gorm:"unique"`
	Manufacturer string `json:"manufacturer" form:"manufacturer" binding:"required"`
	DeviceModel  string `json:"device_model" form:"device_model" binding:"required"`
	Status       *bool  `json:"status" form:"status" binding:"required"`
}

func (r *DeviceUpdateRequest) ToModel() models.DeviceModel {
	return models.DeviceModel{
		DeviceName:   r.Name,
		DeviceIMEI1:  r.DeviceIMEI1,
		DeviceIMEI2:  r.DeviceIMEI2,
		Manufacturer: r.Manufacturer,
		DeviceModel:  r.DeviceModel,
		Status:       *r.Status,
	}
}
