package event

import "time"

type Event struct {
	ID        string     `json:"id" gorm:"primary_key;unique"`
	Name      string     `json:"name" binding:"required"`
	Location  string     `json:"location" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
