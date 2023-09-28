package dtos

type AuthDto struct {
	PhoneNumber *string `json:"phoneNumber"`
	Otp         *string `json:"otp"`
}
