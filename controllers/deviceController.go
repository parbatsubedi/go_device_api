package controllers

import (
	apirequests "go_api/apiRequests"
	apiresponses "go_api/apiResponses"
	errorresponse "go_api/apiResponses/errorResponse"
	"go_api/helpers"
	"go_api/interfaces"
	"go_api/models"
	"go_api/repository"
	"go_api/services"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	auditService *services.AuditService
}

func MakeDeviceController() *DeviceController {
	return &DeviceController{
		auditService: services.NewAuditService(),
	}
}

func (cx *DeviceController) Create(c *gin.Context) {
	//validate request
	var deviceCreateRequest apirequests.DeviceCreateRequest

	err := c.ShouldBind(&deviceCreateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	deviceModel := deviceCreateRequest.ToModel()
	// Set the user ID from the authenticated user context
	deviceModel.UserID = helpers.GetCurrentUserID(c)

	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()

	// Logical Validations
	_, err = deviceRepo.FindByDeviceName(deviceModel.DeviceName)
	if err == nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Device Name Already Exists"))
		return
	}

	_, err = deviceRepo.FindByIMEI(deviceModel.DeviceIMEI1)
	if err == nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Device IMEI1 Already Exists"))
		return
	}

	if deviceModel.DeviceIMEI2 != "" {
		_, err = deviceRepo.FindByIMEI(deviceModel.DeviceIMEI2)
		if err == nil {
			c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Device IMEI2 Already Exists"))
			return
		}
	}

	// Create Resource
	if createErr := deviceRepo.Save(&deviceModel); createErr != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeCreateResourceErrorResponse())
		return // unreachable code
	}

	// Log the create action to audit trail
	ipAddress := c.ClientIP()
	userID := helpers.GetCurrentUserID(c) // Get the user ID from the context
	if err := cx.auditService.LogCreate("devices", deviceModel.ID, &userID, ipAddress, ""); err != nil {
		slog.Error("Failed to log device create action", slog.Any("error", err))
	}

	// Setup Response
	c.JSON(http.StatusCreated, apiresponses.NewSuccessResponse())
	// return // unreachable code so commented
}

func (cx *DeviceController) GetAll(c *gin.Context) {
	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	devices, err := deviceRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeInternalServerError())
		return
	}

	var deviceDtos []models.DeviceModelDto
	for _, device := range devices {
		deviceDtos = append(deviceDtos, device.ToDto())
	}

	response := apiresponses.NewDeviceListResponse()
	response.Data = deviceDtos
	c.JSON(http.StatusOK, response)
}

func (cx *DeviceController) GetByID(c *gin.Context) {
	resourceIDStr := c.Param("id")
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()

	// Find from DB
	var model models.DeviceModel
	model, err = deviceRepo.FindByID(uint(resourceID))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Setup Response
	response := apiresponses.NewDeviceResponse()
	response.Data = model.ToDto()
	c.JSON(http.StatusOK, response)
	// return //unreachable code so commented
}

func (cx *DeviceController) Update(c *gin.Context) {
	// Validate and Bind the incoming request
	var updateRequest apirequests.DeviceUpdateRequest
	err := c.ShouldBind(&updateRequest)
	if err != nil {
		slog.Error("Validation Error: ", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}
	// Request Struct to the respective Model
	model := updateRequest.ToModel()
	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	// Logical Validations
	resourceIDStr := c.Param("id")
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}
	// Cast ResourceID to uint
	model.ID = uint(resourceID)
	model.UserID = helpers.GetCurrentUserID(c)
	// Check if resource exists
	existingDevice, err := deviceRepo.FindByID(model.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Logical Validations - Check for duplicate device name (if name is being changed)
	if model.DeviceName != existingDevice.DeviceName {
		_, err := deviceRepo.FindByDeviceName(model.DeviceName)
		if err == nil {
			c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Device Name Already Exists"))
			return
		}
	}

	// Logical Validations - Check for duplicate IMEI1 (if IMEI1 is being changed)
	if model.DeviceIMEI1 != existingDevice.DeviceIMEI1 {
		_, err := deviceRepo.FindByIMEI(model.DeviceIMEI1)
		if err == nil {
			c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Device IMEI1 Already Exists"))
			return
		}
	}

	// Logical Validations - Check for duplicate IMEI2 (if IMEI2 is being changed and not empty)
	if model.DeviceIMEI2 != "" && model.DeviceIMEI2 != existingDevice.DeviceIMEI2 {
		_, err := deviceRepo.FindByIMEI(model.DeviceIMEI2)
		if err == nil {
			c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Device IMEI2 Already Exists"))
			return
		}
	}

	// Update Resource
	if updateErr := deviceRepo.PartialUpdate(&model); updateErr != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeUpdateErrorResponse())
		return // unreachable code
	}

	// Log the update action to audit trail
	ipAddress := c.ClientIP()
	userID := helpers.GetCurrentUserID(c)
	if err := cx.auditService.LogUpdate("devices", model.ID, &userID, ipAddress, "", ""); err != nil {
		slog.Error("Failed to log device update action", slog.Any("error", err))
	}

	// Setup Response
	c.JSON(http.StatusOK, apiresponses.NewSuccessResponse())
	// return //unreachable code so commented
}

func (cx *DeviceController) Delete(c *gin.Context) {
	resourceIDStr := c.Param("id")
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()

	// Find from DB
	var model models.DeviceModel
	model, err = deviceRepo.FindByID(uint(resourceID))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Delete Resource
	if deleteErr := deviceRepo.Delete(model); deleteErr != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeDeleteErrorResponse())
		return // unreachable code
	}

	// Log the delete action to audit trail
	ipAddress := c.ClientIP()
	userID := helpers.GetCurrentUserID(c)
	if err := cx.auditService.LogDelete("devices", model.ID, &userID, ipAddress, ""); err != nil {
		slog.Error("Failed to log device delete action", slog.Any("error", err))
	}

	// Setup Response
	c.JSON(http.StatusOK, apiresponses.NewSuccessResponse())
}

func (cx *DeviceController) GetUserDevices(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	devices, err := deviceRepo.FindByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeInternalServerError())
		return
	}

	var deviceDtos []models.DeviceModelDto
	for _, device := range devices {
		deviceDtos = append(deviceDtos, device.ToDto())
	}

	response := apiresponses.NewDeviceListResponse()
	response.Data = deviceDtos
	c.JSON(http.StatusOK, response)
}

func (cx *DeviceController) GetDashboardData(c *gin.Context) {
	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	deviceLocationRepo := repository.NewDeviceLocationRepository()

	devices, err := deviceRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeInternalServerError())
		return
	}

	// Add logging to capture the response
	slog.Info("Fetching dashboard data", slog.Any("devices", devices))

	activeDevices := 0
	offlineDevices := 0
	totalBattery := 0
	alerts := 0

	// Get device details with their latest location data
	deviceDetails := make([]gin.H, 0)

	for _, device := range devices {
		// Get latest location for battery level
		latestLocation, err := deviceLocationRepo.FindLatestByDeviceID(device.ID)
		batteryLevel := 0

		if err == nil && latestLocation.ID != 0 {
			batteryLevel = latestLocation.BatteryLevel
		}

		deviceStatus := device.Status
		if deviceStatus {
			activeDevices++
			totalBattery += batteryLevel
		} else {
			offlineDevices++
		}

		// Check for low battery alert
		if batteryLevel < 20 {
			alerts++
		}

		deviceDetails = append(deviceDetails, gin.H{
			"id":            device.ID,
			"name":          device.DeviceName,
			"status":        device.Status,
			"battery_level": batteryLevel,
			"last_seen":     device.LastSeenAt,
			"model":         device.DeviceModel,
			"manufacturer":  device.Manufacturer,
		})
	}

	averageBattery := 0
	if activeDevices > 0 {
		averageBattery = totalBattery / activeDevices
	}

	response := gin.H{
		"active_devices":  activeDevices,
		"offline_devices": offlineDevices,
		"average_battery": averageBattery,
		"alerts":          alerts,
		"devices":         deviceDetails,
	}

	c.JSON(http.StatusOK, response)
}

func (cx *DeviceController) GetDashboardDataForTemplate(c *gin.Context) (gin.H, error) {
	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	deviceLocationRepo := repository.NewDeviceLocationRepository()

	devices, err := deviceRepo.FindAll()
	if err != nil {
		return gin.H{
			"title":           "Admin Dashboard",
			"active_devices":  0,
			"average_battery": 0,
			"alerts":          0,
			"offline_devices": 0,
			"devices":         []interface{}{},
		}, err
	}

	activeDevices := 0
	offlineDevices := 0
	totalBattery := 0
	alerts := 0

	// Get device details with their latest location data
	deviceDetails := make([]gin.H, 0)

	for _, device := range devices {
		// Get latest location for battery level
		latestLocation, err := deviceLocationRepo.FindLatestByDeviceID(device.ID)
		batteryLevel := 0

		if err == nil && latestLocation.ID != 0 {
			batteryLevel = latestLocation.BatteryLevel
		}

		deviceStatus := device.Status
		if deviceStatus {
			activeDevices++
			totalBattery += batteryLevel
		} else {
			offlineDevices++
		}

		// Check for low battery alert
		if batteryLevel < 20 {
			alerts++
		}

		deviceDetails = append(deviceDetails, gin.H{
			"id":            device.ID,
			"name":          device.DeviceName,
			"status":        device.Status,
			"battery_level": batteryLevel,
			"last_seen":     device.LastSeenAt,
			"model":         device.DeviceModel,
			"manufacturer":  device.Manufacturer,
		})
	}

	averageBattery := 0
	if activeDevices > 0 {
		averageBattery = totalBattery / activeDevices
	}

	return gin.H{
		"title":           "Admin Dashboard",
		"active_devices":  activeDevices,
		"offline_devices": offlineDevices,
		"average_battery": averageBattery,
		"alerts":          alerts,
		"devices":         deviceDetails,
	}, nil
}
