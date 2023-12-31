package models

import(
	"go-jwt/database"

 	"golang.org/x/crypto/bcrypt"
 	"gorm.io/gorm"
)

// user define in db

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

// create user record in db

func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.error != nil {
		return result.Error
	}
	return nil
}

// encrypt user password

func (user *User) HasPassword(password string) error  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// check user password

func (user *User) CheckPassword(providePassword string) error{
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providePassword))
	if err != nil {
		return err
	}
	return nil
}