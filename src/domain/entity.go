package domain

import "time"

type UserEntity struct {
	Id           int64
	NationalCode string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AdminEntity struct {
	Id        int64
	Name      string
	Family    string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
