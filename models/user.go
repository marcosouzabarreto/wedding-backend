package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	BaseModel
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"-"`
}

type UserInput struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
