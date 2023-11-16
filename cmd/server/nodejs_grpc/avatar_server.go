package nodejs_grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	pb_nodejs "whatw-golang/pb/nodejs"

	"github.com/joho/godotenv"
)

type ServerNodejs struct {
	pb_nodejs.AvatarServiceServer
}

var apiResponseNodejs []struct {
	ID	int    `json:"id"`
	Image string `json:"image"`
	Price int    `json:"price"`
}

func (s *ServerNodejs) GetAvatars(ctx context.Context, req *pb_nodejs.AvatarsRequest) (*pb_nodejs.AvatarResponse, error) {
	godotenv.Load(".env")
	apiURL := os.Getenv("API_URL_NODEJS_AVATAR")
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response body: %v", err)
	}

	if err := json.Unmarshal(data, &apiResponseNodejs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}
	
	if len(apiResponseNodejs) == 0 {
		return nil, fmt.Errorf("API response indicates failure or empty data")
	}
	

	var dataRespose []*pb_nodejs.Avatars
	for _, q := range apiResponseNodejs {

		dataRespose = append(dataRespose, &pb_nodejs.Avatars{
			Id:    int32(q.ID),
			Image: q.Image,
			Price: int32(q.Price),
		})
	}

	return &pb_nodejs.AvatarResponse{
		Data: dataRespose,
	}, nil
}
