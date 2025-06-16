package config

import (
	"firebase-playground/internal/domain/model/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("firebase.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if !db.Migrator().HasTable(&user.User{}) {
		db.Migrator().CreateTable(&user.User{})
	}

	return db, nil
}
