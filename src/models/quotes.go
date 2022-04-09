package models

import "time"

type Quotes struct {
	ID        int64 `gorm:"primary_key"`
	Content   string
	Author    string
	Like      int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type QuotesOfTheDay struct {
	ID        int64 `gorm:"primary_key"`
	QuotesID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
