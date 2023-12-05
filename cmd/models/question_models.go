package models

var ResponseQuestion struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	Data []struct {
		ID int `json:"id"`
		Difficulty         string `json:"difficulty"`
		Question           string `json:"question"`
		Options            []string `json:"options"`
		Answer             string `json:"answer"`
	} `json:"data"`
}