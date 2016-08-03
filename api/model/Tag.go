package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	ID uint `gorm:"primary_key"`
	Name string
}

func (t Tag) TableName() string {
	return "b_tags"
}