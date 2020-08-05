package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Item struct {
	Id        int    `gorm:"unique;not null"`
	Name      string `gorm:"size:255"`
	Status    int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
