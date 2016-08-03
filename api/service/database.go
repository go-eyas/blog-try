package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/go-eyas/blog-try/api/model"
)

type Storage struct {
	db *gorm.DB
}

func Database() *Storage {
	db, err := gorm.Open("sqlite3", "database/blog.db")
	if err != nil {
		panic("database connect fail")
	}

	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Tag{})

	db.Model(&Post{}).Related(&Tag{})

	storage := &Storage{db: db}

	return storage;
}

var DB *gorm.DB;

func InitDatabase() *gorm.DB {
	DB = Database().db
	return DB;
}


