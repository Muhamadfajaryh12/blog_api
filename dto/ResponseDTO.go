package dto

type ResponseErrorDTO struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"Bad request"`
}

type ResponseSuccessDTO struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"Fetch successfully"`
	Data    any    `json:"data,omitempty"`
}