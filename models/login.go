package models

import "time"

// Login model
type Login struct {
	ID        uint `gorm:"primary_key"`
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (l *Login) TableName() string {
	return "logins"
}
