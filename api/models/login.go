package models

import "time"

type ReqLogin struct {
	UserName string
	Password string
}

// Login model
type Login struct {
	ID        uint64 `gorm:"primary_key"`
	UserName  string `gorm:"type:varchar(100);index:username;unique;not null;"`
	Password  string `gorm:"type:varchar(200);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (l *Login) TableName() string {
	return "logins"
}
