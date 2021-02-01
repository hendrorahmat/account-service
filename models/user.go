package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string `gorm:"PrimaryKey" json:"id"`
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
