package services

import "mygram/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	UserService
	CommentService
	PhotoService
	SocialMediaService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
