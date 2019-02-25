package model

import (
	"time"
)

type AutoID struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
}

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

