package controllers

import (
	"fmt"
	apirequests "go_api/apiRequests"
	apiresponses "go_api/apiResponses"
	errorresponse "go_api/apiResponses/errorResponse"
	"go_api/models"
	"go_api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
}

func MakeUserController() UserController {
	return UserController{}
}

func (cx *UserController) Create(c *gin.Context) {
	// Validate and Bind the incoming request
	var createRequest apirequests.CreateUserRequest
	err := c.ShouldBind(&createRequest) // use ShouldBind to support multiple content types and ShouldBindJSON for only JSON(raw request)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	// Request Struct to the respective Model
	model := createRequest.ToModel()

	userRepo := repository.NewUserRepository()

	// Logical Validations
	_, emailExists := userRepo.FindByEmail(model.Email)
	if emailExists {
		c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Email Already Exists"))
		return
	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(model.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeCustomErrorResponse("Failed to hash password"))
		return
	}
	model.Password = string(hash)

	// Create Resource
	if err := userRepo.Save(&model); err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeCreateResourceErrorResponse())
		return
	}

	// Setup Response
	c.JSON(http.StatusCreated, apiresponses.NewSuccessResponse())
	// return //unreachable code so commented
}

func (cx *UserController) GetById(c *gin.Context) {
	resourceIDStr := c.Param("id")
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	userRepo := repository.NewUserRepository()

	// Find from DB
	var model models.UserModel
	model, exists := userRepo.FindByID(uint(resourceID))
	if !exists {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Setup Response
	response := apiresponses.NewUserResponse()
	response.Data = model.ToDto()
	c.JSON(http.StatusOK, response)
	// return //unreachable code so commented
}

func (cx *UserController) Update(c *gin.Context) {
	// Validate and Bind the incoming request
	var updateRequest apirequests.UpdateUserRequest
	err := c.ShouldBind(&updateRequest) // use ShouldBind to support multiple content types and ShouldBindJSON for only JSON(raw request)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	// Request Struct to the respective Model
	model := updateRequest.ToModel()

	userRepo := repository.NewUserRepository()

	// Logical Validations
	resourceIDStr := c.Param("id")
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}
	// Cast ResourceID to uint
	model.ID = uint(resourceID)

	// Check if resource exists
	_, exists := userRepo.FindByID(model.ID)
	if !exists {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	// Update Resource
	updateError := userRepo.PartialUpdate(model)
	if !updateError {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeUpdateErrorResponse())
		return
	}

	// Reload the updated resource to reflect the changes
	updatedUser, _ := userRepo.FindByID(model.ID)

	// Setup Response
	userResponse := apiresponses.NewUserResponse()
	userResponse.Data = updatedUser.ToDto()

	c.JSON(http.StatusOK, userResponse)
	// return //unreachable code so commented
}

func (cx *UserController) GetAll(c *gin.Context) {

	userRepo := repository.NewUserRepository()

	// Fetch all users from the database
	modelList := userRepo.FindAll()
	fmt.Printf("Found %d users\n", len(modelList))

	// Convert models to DTOs
	modelDtoList := []models.UserModelDto{}
	for _, model := range modelList {
		modelDtoList = append(modelDtoList, model.ToDto())
	}

	// Setup Response
	response := apiresponses.NewUserListResponse()
	response.Data = modelDtoList

	c.JSON(http.StatusOK, response)
	// return //unreachable code so commented
}

func (cx *UserController) Delete(c *gin.Context) {
	resourceIDStr := c.Param("id")
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeInvalidResourceIdResponse())
		return
	}

	userRepo := repository.NewUserRepository()

	var model models.UserModel
	model, exists := userRepo.FindByID(uint(resourceID))
	if !exists {
		c.JSON(http.StatusBadRequest, errorresponse.MakeResourceNotFoundErrorResponse())
		return
	}

	deleteSuccess := userRepo.Delete(model)
	if !deleteSuccess {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeInternalServerError())
		return
	}

	// c.JSON(http.StatusNoContent, gin.H{
	// 	"success": true,
	// 	"message": "User deleted successfully",
	// })
	c.JSON(http.StatusNoContent, apiresponses.NewSuccessResponse())
}
