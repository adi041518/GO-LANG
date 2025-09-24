package services

import (
	"example1.com/user-registartion-api/model"
)

type Users struct {
	User []model.User
}

func NewServices() *Users {
	return &Users{
		User: []model.User{},
	}
}

func (u *Users) Uservalidate(us model.User) string {
	if us.Age < 18 {
		return "user is not eligible"
	}
	return "User is eligible"
}
