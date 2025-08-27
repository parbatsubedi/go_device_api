package auth

import (
	authrequest "go_api/apiRequests/authRequest"
	apiresponses "go_api/apiResponses"
	"go_api/apiResponses/auth"
	errorresponse "go_api/apiResponses/errorResponse"
	"go_api/repository"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
}

func NewAuthController() AuthController {
	return AuthController{}
}

func (cx *AuthController) Login(c *gin.Context) {

	var loginRequest authrequest.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		slog.Info("Login Error", "Error", err)
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	slog.Info("Login Request", "request", loginRequest)

	userRepo := repository.NewUserRepository()
	user, err := userRepo.FindByMobileNo(loginRequest.MobileNo)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeUnAuthorizedErrorResponse("User does not exists"))
		return
	}

	checkHashError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if checkHashError != nil {
		c.JSON(http.StatusBadRequest, errorresponse.MakeUnAuthorizedErrorResponse("Invalid Password"))
		return
	}

	//token generate and with expiry 1 day
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//get complete encoded token in response
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeInternalServerError())
		return
	}

	loginSuccessResponse := auth.NewLoginSuccessResponse()
	loginSuccessResponse.Token = tokenString
	loginSuccessResponse.User = user.ToDto()

	c.JSON(http.StatusOK, loginSuccessResponse)
	// return //unreachable code so commented
}

func (cx *AuthController) Register(c *gin.Context) {
	var registerRequest authrequest.RegisterRequest

	err := c.ShouldBind(&registerRequest)
	if err != nil {
		slog.Info("Register Error", "Error", err)
		c.JSON(http.StatusBadRequest, errorresponse.MakeValidationErrorsResponse(err))
		return
	}

	slog.Info("Register Request", "request", registerRequest)

	userRepo := repository.NewUserRepository()

	userByMobile, err := userRepo.FindByMobileNo(registerRequest.MobileNo)
	if err == nil && userByMobile.ID != 0 {
		c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("User Already Exists"))
		return
	}

	userByEmail, err := userRepo.FindByEmail(registerRequest.Email)
	if err == nil && userByEmail.ID != 0 {
		c.JSON(http.StatusBadRequest, errorresponse.MakeCustomErrorResponse("Email Already Exists"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.MakeCustomErrorResponse("Failed to hash password"))
		return
	}
	registerRequest.Password = string(hash)

	user := registerRequest.ToModel()
	userRepo.Save(&user)

	c.JSON(http.StatusOK, apiresponses.NewSuccessResponse())
}
