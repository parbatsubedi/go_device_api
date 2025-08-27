package apiresponses

type SuccessResponse struct {
	State   bool   `json:"state"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewSuccessResponse() SuccessResponse {
	return SuccessResponse{
		State:   true,
		Status:  "SUCCESS",
		Message: "Success",
	}
}

// GenericSuccessResponse is a generic success response with dynamic data
type GenericSuccessResponse struct {
	SuccessResponse
	Data interface{} `json:"data"`
}

func NewGenericSuccessResponse(data any) GenericSuccessResponse {
	return GenericSuccessResponse{
		SuccessResponse: NewSuccessResponse(),
		Data:            data,
	}
}
