package avatar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	pb_laravel_avatar "whatw-golang/pb/laravel/avatar"

)

type ServerAvatar struct {
	pb_laravel_avatar.AvatarServiceServer
}

var apiResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	Data []struct {
		ID int `json:"id"`
		Image string `json:"image"`
		Price int `json:"price"`
	} `json:"data"`
}
func (s *ServerAvatar) GetAvatar(ctx context.Context, req *pb_laravel_avatar.AvatarRequest) (*pb_laravel_avatar.AvatarResponse, error) {
	apiURL := os.Getenv("API_URL_LARAVEL_AVATAR")
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response body: %v", err)
	}


	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}
	
	if !apiResponse.Success || len(apiResponse.Data) == 0 {
		return nil, fmt.Errorf("API response indicates failure or empty data")
	}
	

	var dataRespose []*pb_laravel_avatar.Avatar
	for _, q := range apiResponse.Data {
		dataRespose = append(dataRespose, &pb_laravel_avatar.Avatar{
			Id: int32(q.ID),
			Image: q.Image,
			Price: int32(q.Price),
		})
	}

	return &pb_laravel_avatar.AvatarResponse{
		Data: dataRespose,
	}, nil
}