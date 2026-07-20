package domain

import "time"

type Notes struct {
	ID 			uint			`gorm:"primaryKey" json:"id"`
	Title 		string			`gorm:"type:varchar(255);not null" json:"title" validate:"required,min=3,max=255"`
	Body 		string			`gorm:"type:text" json:"body"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
}