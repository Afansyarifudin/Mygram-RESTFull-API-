package repository

import (
	"fmt"
	"mygram/helper"
	"mygram/models"
)

type UserRepo interface {
	RegisterUser(in models.User) (res models.User, err error)
	LoginUser(in models.User) (res bool, err error)
	GetAllUser() (res []models.User, err error)
	UpdateUser(in models.User) (res models.User, err error)
	DeleteUser(id int64) (err error)
}

func (r Repo) RegisterUser(in models.User) (res models.User, err error) {
	if err := r.gorm.Debug().Create(&in).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) LoginUser(in models.User) (res bool, err error) {
	password := in.Password

	if err := r.gorm.Debug().Where("email = ?", in.Email).Take(&in).Error; err != nil {
		return res, err
	}
	// fmt.Println((in.Password), (password))

	isValid := helper.ComparePassword([]byte(in.Password), []byte(password))
	if !isValid {
		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"error":   "unauthorized",
		// 	"message": "Invalid email/password",
		// })
		fmt.Println("invalid username or password")
		return
	}

	return true, nil
}

func (r Repo) GetAllUser() (res []models.User, err error) {
	return res, nil
}

func (r Repo) UpdateUser(in models.User) (res models.User, err error) {
	return res, nil
}

func (r Repo) DeleteUser(id int64) (err error) {
	return nil
}
