package middlewares

import (
	"fmt"
	errorresponse "go_api/apiResponses/errorResponse"
	"go_api/repository"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(userRepo repository.IUserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		if reqToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorresponse.MakeUnAuthorizedErrorResponse("Authorization header is required"))
			return
		}

		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorresponse.MakeUnAuthorizedErrorResponse("Invalid token format"))
			return
		}
		reqToken = splitToken[1]
		slog.Debug("Auth token received")

		// Decode/Validate it
		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorresponse.MakeUnAuthorizedErrorResponse("Invalid token"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorresponse.MakeUnAuthorizedErrorResponse("Token has expired"))
				return
			}

			// Find the user with token sub
			userId := uint(claims["sub"].(float64))
			user, err := userRepo.FindByID(userId)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorresponse.MakeUnAuthorizedErrorResponse("User associated with token not found"))
				return
			}

			// Attach to Req
			c.Set("user", user.ToDto())
			c.Set("user_id", userId)
			c.Set("user_model", user) // original user object

			//Continue
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, errorresponse.MakeUnAuthorizedErrorResponse("Invalid token claims"))
	}
}
