package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	ID      uint `gorm:"primary_key"`
	Title   string
	Tags    []Tag `gorm:"many2many:post_tags;"`
	Content string
}

func (p Post) TableName() string {
	return "b_posts"
}
