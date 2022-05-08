package model

import "time"

type Todo struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`
	Priority  uint      `json:"priority"`
	Done      bool      `json:"is_done"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
