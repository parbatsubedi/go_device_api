package controllers

import (
	apirequests "go_api/apiRequests"
	apiresponses "go_api/apiResponses"
	errorresponse "go_api/apiResponses/errorResponse"
	"go_api/interfaces"
	"go_api/models"
	"go_api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeviceCommandController struct {
}

func MakeDeviceCommandController() *DeviceCommandController {
	return &DeviceCommandController{}
}

var _ interfaces.IDeviceCommandController = (*DeviceCommandController)(nil)

// SendCommand sends a command to a device
func (cx *DeviceCommandController) SendCommand(c *gin.Context) {
	// Get device ID from URL parameter
	deviceIDStr := c.Param("device_id")
	deviceID, err := strconv.ParseUint(deviceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	// Validate request
	var commandRequest apirequests.DeviceCommandRequest
	err = c.ShouldBind(&commandRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	// Verify device exists
	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	_, err = deviceRepo.FindByID(uint(deviceID))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Create command
	command := models.DeviceCommandModel{
		DeviceID:    uint(deviceID),
		CommandType: commandRequest.CommandType,
		CommandData: commandRequest.CommandData,
		Status:      "pending",
	}

	var commandRepo interfaces.IDeviceCommandRepository = repository.NewDeviceCommandRepository()
	if createErr := commandRepo.Create(&command); createErr != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeCreateResourceErrorResponse())
		return
	}

	// Setup Response
	response := apiresponses.NewDeviceCommandResponse()
	response.Data = apiresponses.ToDeviceCommandDto(command)
	c.JSON(http.StatusCreated, response)
}

// GetDeviceCommands gets all commands for a device
func (cx *DeviceCommandController) GetDeviceCommands(c *gin.Context) {
	// Get device ID from URL parameter
	deviceIDStr := c.Param("device_id")
	deviceID, err := strconv.ParseUint(deviceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	// Verify device exists
	var deviceRepo interfaces.IDeviceRepository = repository.NewDeviceRepository()
	_, err = deviceRepo.FindByID(uint(deviceID))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Get commands
	var commandRepo interfaces.IDeviceCommandRepository = repository.NewDeviceCommandRepository()
	commands := commandRepo.FindByDeviceID(uint(deviceID))

	var commandDtos []apiresponses.DeviceCommandDto
	for _, command := range commands {
		commandDtos = append(commandDtos, apiresponses.ToDeviceCommandDto(command))
	}

	response := apiresponses.DeviceCommandListResponse{
		Data: commandDtos,
	}
	c.JSON(http.StatusOK, response)
}

// GetPendingCommands gets pending commands for a device
func (cx *DeviceCommandController) GetPendingCommands(c *gin.Context) {
	deviceIDStr := c.Param("device_id")
	deviceID, err := strconv.ParseUint(deviceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	// Get pending commands
	var commandRepo interfaces.IDeviceCommandRepository = repository.NewDeviceCommandRepository()
	commands := commandRepo.FindPendingCommands(uint(deviceID))

	var commandDtos []apiresponses.DeviceCommandDto
	for _, command := range commands {
		commandDtos = append(commandDtos, apiresponses.ToDeviceCommandDto(command))
	}

	response := apiresponses.DeviceCommandListResponse{
		Data: commandDtos,
	}
	c.JSON(http.StatusOK, response)
}

func (cx *DeviceCommandController) AcknowledgeCommand(c *gin.Context) {
	commandIDStr := c.Param("commandId")
	commandID, err := strconv.ParseUint(commandIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	var commandRepo interfaces.IDeviceCommandRepository = repository.NewDeviceCommandRepository()
	if ackErr := commandRepo.MarkAsAcknowledged(uint(commandID)); ackErr != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeUpdateErrorResponse())
		return
	}

	c.JSON(http.StatusOK, apiresponses.NewSuccessResponse())
}
