package models

import "time"

type Profile struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Fullname  string `json:"full_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
