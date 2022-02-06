package entity

import "time"

type Article struct {
	UserId string
	Title string
	Text string `gorm:"unique"`
	Status string
	CreatedAt time.Time
	ModifiedAt time.Time
}


