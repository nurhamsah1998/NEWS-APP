package models

import "time"

type NewsCategory struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
	NewsPost  []*NewsPost `gorm:"many2many:news_post"`
}
