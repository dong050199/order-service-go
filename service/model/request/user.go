package request

import (
	"order-service/service/model/entity"
)

type UserRegister struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u UserRegister) ToEntityModel() entity.User {
	return entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		UserName:  u.UserName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedBy: u.UserName,
	}
}

type UserLogin struct {
	UserNameOrEmail string `json:"user_name_or_email"`
	Password        string `json:"password"`
}
