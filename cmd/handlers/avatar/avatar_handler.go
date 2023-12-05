package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"whatw-golang/cmd/models"
	pb_laravel_avatar "whatw-golang/pb/laravel/avatar"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ServerAvatar struct {
	pb_laravel_avatar.UnimplementedAvatarServiceServer
}

func (*ServerAvatar) GetAvatar(ctx context.Context, req *pb_laravel_avatar.AvatarRequest) (*pb_laravel_avatar.AvatarResponse, error) {
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

	if err := json.Unmarshal(data, &models.AvatarResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	if !models.AvatarResponse.Success || len(models.AvatarResponse.Data) == 0 {
		return nil, fmt.Errorf("API response indicates failure or empty data")
	}

	var dataRespose []*pb_laravel_avatar.Avatar
	for _, q := range models.AvatarResponse.Data {
		price, err := anypb.New(&wrapperspb.Int32Value{Value: int32(q.Price)})
		if err != nil {
			return nil, fmt.Errorf("failed to marshal price: %v", err)
		}
		avatar := &pb_laravel_avatar.Avatar{
			Id:    int32(q.ID),
			Image: q.Image,
			Price: price,
		}
		dataRespose = append(dataRespose, avatar)
	}

	return &pb_laravel_avatar.AvatarResponse{
		Data: dataRespose,
	}, nil
}
