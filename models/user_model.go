package models

import (
	"mygram/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string        `gorm:"not null;uniqueIndex;type:varchar(50)" json:"username" form:"username" valid:"required~Username is required"`
	Email    string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email formal"`
	Password string        `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)~Password mas have minimum length 6 characters"`
	Age      uint          `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,range(8|70)~Minimum age is 8 year old"`
	Photos   []Photo       `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments []Comment     `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Medias   []SocialMedia `json:"medias" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helper.HashPassword(u.Password)

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helper.HashPassword(u.Password)
	err = nil
	return
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
