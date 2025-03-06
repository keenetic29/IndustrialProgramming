package repository

import (
	"rest.com/pkg/db"
	"rest.com/pkg/model"
)

func GetUserByLogin(login string) (model.User, error) {
	var user model.User
	err := db.DB.Where("username = ?", login).First(&user).Error
	return user, err
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := db.DB.Find(&users).Error
	return users, err
}

func AddUser(login string, password string) error {
	return db.DB.Create(&model.User{Username: login, Password: password}).Error
}