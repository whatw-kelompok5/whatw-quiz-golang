package question

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	pb_laravel_question "whatw-golang/pb/laravel/question"
)

type ServerQuestion struct {
	pb_laravel_question.UnimplementedQuestionServiceServer
}

var apiResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	Data []struct {
		ID int `json:"id"`
		Difficulty         string `json:"difficulty"`
		Question           string `json:"question"`
		CorrectAnswer      string `json:"correct_answer"`
		IncoorrectAnswers1 string `json:"incorrect_answer1"`
		IncoorrectAnswers2 string `json:"incorrect_answer2"`
		IncoorrectAnswers3 string `json:"incorrect_answer3"`
	} `json:"data"`
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

	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	if !apiResponse.Success || len(apiResponse.Data) == 0 {
		return nil, fmt.Errorf("API response indicates failure or empty data")
	}

	var dataRespose []*pb_laravel_question.Question
	for _, q := range apiResponse.Data {
		dataRespose = append(dataRespose, &pb_laravel_question.Question{
			Id:                 int32(q.ID),
			Difficulty:         q.Difficulty,
			Question:           q.Question,
			CorrectAnswer:      q.CorrectAnswer,
			IncorrectAnswer1: q.IncoorrectAnswers1,
			IncorrectAnswer2: q.IncoorrectAnswers2,
			IncorrectAnswer3: q.IncoorrectAnswers3,
		})
	}

	return &pb_laravel_question.QuestionResponse{
		Data: dataRespose,
	}, nil
}
