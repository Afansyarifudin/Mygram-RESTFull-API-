package repository

import "gorm.io/gorm"

type Repo struct {
	gorm *gorm.DB
}

type RepoInterface interface {
	UserRepo
	CommentRepo
	PhotoRepo
	SocialMediaRepo
}

func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}
