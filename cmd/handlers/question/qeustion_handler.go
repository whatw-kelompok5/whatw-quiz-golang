package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"whatw-golang/cmd/models"
	pb_laravel_question "whatw-golang/pb/laravel/question"
)

type ServerQuestion struct {
	pb_laravel_question.UnimplementedQuestionServiceServer
}
func (*ServerQuestion) GetQuestion(ctx context.Context, req *pb_laravel_question.QuestionRequest) (*pb_laravel_question.QuestionResponse, error) {
	apiURL := os.Getenv("API_URL_LARAVEL_QUESTION")
	client := http.Client{
		Timeout: 10 * time.Second, // Set a timeout for the HTTP request
	}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response body: %v", err)
	}

	// Check if the response is HTML
	if strings.HasPrefix(string(data), "<") {
		return nil, fmt.Errorf("API returned HTML error page: %s", string(data))
	}

	if err := json.Unmarshal(data, &models.ResponseQuestion); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	if !models.ResponseQuestion.Success || len(models.ResponseQuestion.Data) == 0 {
		return nil, fmt.Errorf("API response indicates failure or empty data")
	}

	var dataRespose []*pb_laravel_question.Question
	for _, q := range models.ResponseQuestion.Data {
		question := &pb_laravel_question.Question{
			Id:          int32(q.ID),
			Difficulty:  q.Difficulty,
			Question:    q.Question,
			Options:     q.Options,
			Answer:      q.Answer,
		}
		dataRespose = append(dataRespose, question)
	}

	return &pb_laravel_question.QuestionResponse{
		Data: dataRespose,
	}, nil
}
