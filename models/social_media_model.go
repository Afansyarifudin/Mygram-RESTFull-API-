package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name_sosmed" form:"name_sosmed" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media url is required"`
	UserID         uint   `gorm:"user_id"`
	User           *User
}

func (m *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (m *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

type SocialMediaReqs struct {
	Name           string `json:"name_sosmed"`
	SocialMediaUrl string `json:"social_media_url"`
}
