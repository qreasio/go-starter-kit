package model

import (
	"time"
)

// User represents someone with access to our system.
type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Firstname   string    `json:"first_name" db:"first_name"`
	Lastname    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email"`
	Password    []byte    `json:"-"`
	DateJoined  time.Time `json:"date_joined" db:"date_joined"`
	LastLogin   time.Time `json:"last_login" db:"last_login"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	IsStaff     bool      `json:"is_staff" db:"is_staff"`
	IsSuperuser bool      `json:"is_superuser" db:"is_superuser"`
}
