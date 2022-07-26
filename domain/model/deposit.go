package model

import "time"

type Deposit struct {
	Id      string    `json:"id"`
	UserID  string    `json:"user_id"`
	Amount  float32    `json:"amount"`
	Balance float32    `json:"balance"`
	Date    time.Time `json:"date"`
}
