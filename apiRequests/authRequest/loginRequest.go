package authrequest

type LoginRequest struct {
	MobileNo string `json:"mobile_no" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgetPasswordRequest struct {
	MobileNo string `json:"mobile_no" binding:"required" minLength:"10" maxLength:"10"`
}

type ResetPasswordRequest struct {
	MobileNo string `json:"mobile_no" binding:"required"`
	Password string `json:"password" binding:"required"`
	// ConfirmPassword string `json:"confirm_password" binding:"required"`
	Otp string `form:"otp" binding:"required" minLength:"6" maxLength:"6"`
}
