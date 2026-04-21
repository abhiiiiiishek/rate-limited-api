package model

type Request struct {
	UserID  string      `json:"user_id"`
	Payload interface{} `json:"payload"`
}
