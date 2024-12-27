package repository

import (
	"errors"

	"rest.com/pkg/auth"
	"rest.com/pkg/model"
)

var users = []model.User {
	{ID: 1, Name: "user1", Password: "12345", AccessLevel: "User"},
	{ID: 2, Name: "admin1", Password: "12345", AccessLevel: "Admin"},
}


func GetUserByLogin(login string) (model.User, error) {
	for _, user := range users {
		if user.Name == login {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}
func GetAllUsers() []model.User {
	return users
}

func AddUser(login string, password string) {
	users = append(users, model.User{ID: len(users) + 1, Name: login, Password: password, AccessLevel: auth.USER_ROLE})
}