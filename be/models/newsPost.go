package models

import "time"

type NewsPost struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProfileId    int
	NewsCategory []*NewsCategory `gorm:"many2many:categories;"`
}
