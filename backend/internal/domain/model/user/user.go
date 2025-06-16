package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:string;size:255;unique;not null;comment:メールアドレス"`
	Password string `gorm:"type:string;size:255;not null;comment:パスワード"`
}
