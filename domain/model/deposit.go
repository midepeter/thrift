package model

import "time"

type Deposit struct {
	Id      string    `json:"id"`
	UserID  string    `json:"user_id"`
	Amount  string    `json:"amount"`
	Balance string    `json:"balance"`
	Date    time.Time `json:"date"`
}
