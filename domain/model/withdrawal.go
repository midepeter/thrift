package model

import "time"

type Withdrawal struct {
	Id          string    `json:"id"`
	UserId      int       `json:"user_id"`
	Amount      float32   `json:"amount"`
	DateCreated time.Time `json:date_created"`
}
