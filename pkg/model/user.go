package model

import (
	"time"
)

// User represents someone with access to our system.
type User struct {
	ID           string    `json:"id"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
}
