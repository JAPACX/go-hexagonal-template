package entities

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	UpdatedAt time.Time
	CreatedAt time.Time
}
