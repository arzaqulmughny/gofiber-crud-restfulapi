package domain

import "time"

type Users struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:text" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
