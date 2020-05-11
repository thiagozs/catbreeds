package models

import "time"

// TODO: implements all fields
// CatAPI model
type CatAPI struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
