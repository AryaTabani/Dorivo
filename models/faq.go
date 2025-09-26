package models

type FAQ struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
