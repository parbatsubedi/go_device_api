package controllers

import (
	apirequests "go_api/apiRequests"
	apiresponses "go_api/apiResponses"
	errorresponse "go_api/apiResponses/errorResponse"
	"go_api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceLocationController struct {
}

func MakeDeviceLocationController() *DeviceLocationController {
	return &DeviceLocationController{}
}

func (cx *DeviceLocationController) Create(c *gin.Context) {
	//validate request
	var deviceLocationCreateRequest apirequests.CreateDeviceLocationRequest

	err := c.ShouldBind(&deviceLocationCreateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	deviceLocationModel := deviceLocationCreateRequest.ToModel()

	deviceLocationRepo := repository.NewDeviceLocationRepository()

	// Create Resource
	if createErr := deviceLocationRepo.Create(&deviceLocationModel); createErr != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeCreateResourceErrorResponse())
		return // unreachable code
	}
	// Setup Response
	c.JSON(http.StatusCreated, apiresponses.NewSuccessResponse())
	// return // unreachable code so commented
}
