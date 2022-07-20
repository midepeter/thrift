package model

import "time"

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email_address"`
	PhoneNumber int       `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
