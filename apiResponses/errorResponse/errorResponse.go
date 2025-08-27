package errorresponse

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Errors  any    `json:"errors"`
}

func MakeDuplicateResourceErrorResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Resource already exists",
		Code:    http.StatusBadRequest,
	}
}

func MakeResourceNotFoundErrorResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Resource not found",
		Code:    http.StatusBadRequest,
	}
}

func MakeUpdateErrorResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Resource could not be updated",
		Code:    http.StatusBadRequest,
	}
}

func MakeDeleteErrorResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Resource could not be deleted",
		Code:    http.StatusBadRequest,
	}
}

func MakeInvalidRequestResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Invalid Request",
		Code:    http.StatusBadRequest,
	}
}

func MakeInvalidResourceIdResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Invalid Request",
		Code:    http.StatusBadRequest,
	}
}

func MakeCreateResourceErrorResponse() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Resource could not be created",
		Code:    http.StatusBadRequest,
	}
}

func MakeUnAuthorizedErrorResponse(msg string) ErrorResponse {
	if msg == "" {
		msg = "Unauthorized"
	}
	return ErrorResponse{
		Status:  "error",
		Message: msg,
		Code:    http.StatusBadRequest,
	}
}

func MakeInternalServerError() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Something went wrong. Please try again.",
		Code:    http.StatusInternalServerError,
	}
}

func MakeCustomErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func MakeValidationErrorsResponse(err error) ErrorResponse {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			// Use a more user-friendly error message that doesn't expose struct field names
			fieldName := fe.Field()
			out[i] = ErrorMsg{fieldName, getErrorMsg(fe)}
		}
		return ErrorResponse{
			Status:  "error",
			Message: "Validation failed. Please check your input.",
			Errors:  out,
			Code:    http.StatusBadRequest,
		}
	}
	return ErrorResponse{
		Status:  "error",
		Message: "An unexpected error occurred.",
		Code:    http.StatusBadRequest,
	}
}
