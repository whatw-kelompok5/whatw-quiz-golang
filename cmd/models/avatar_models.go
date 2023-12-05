package models

var AvatarResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    []struct {
		ID    int    `json:"id"`
		Image string `json:"image"`
		Price int    `json:"price"`
	} `json:"data"`
}