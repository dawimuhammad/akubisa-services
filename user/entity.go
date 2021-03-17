package user

import "time"

// User struct to accomodate table users
type User struct {
	ID             int
	Name           string
	Occuppation    string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
