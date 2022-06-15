package service

import (
	"github.com/jaksonkallio/coinbake/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	EmailAddress string
}

func FindUserByEmailAddress(emailAddress string) *User {
	user := User{}
	database.Handle().Where("email_address = ?", emailAddress).First(&user)
	return &user
}
