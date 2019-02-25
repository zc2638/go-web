package model

type Admin struct {
	AutoID
	Name     string `gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Password string `gorm:"type:varchar(255);not null;" json:"password"`
	Role     uint	`gorm:"type:int;default:0;not null;" json:"role"`
	Timestamps
}
