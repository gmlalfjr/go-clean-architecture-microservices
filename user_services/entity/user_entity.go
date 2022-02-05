package entity

import "time"

type User struct {
	Id int
	FullName string
	Email string `gorm:"unique"`
	Password string
	CreatedAt time.Time
	ModifiedAt time.Time
}


