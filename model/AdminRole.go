package model

type AdminRole struct {
	AutoID
	Name string `gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Rule string `gorm:"type:text;" json:"rule"`
	Timestamps
}
