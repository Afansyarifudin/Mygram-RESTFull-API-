package services

import (
	"mygram/models"
)

type CommentService interface {
	CreateComment(in models.Comment) (res models.Comment, err error)
	GetAllComment() (res []models.Comment, err error)
	GetCommentByID(id int64) (res models.Comment, err error)
	UpdateComment(in models.Comment) (res models.Comment, err error)
	DeleteComment(id int64) (err error)
}

func (s *Service) CreateComment(in models.Comment) (res models.Comment, err error) {
	return s.repo.CreateComment(in)
}

func (s *Service) GetAllComment() (res []models.Comment, err error) {
	res, err = s.repo.GetAllComment()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetCommentByID(id int64) (res models.Comment, err error) {
	return s.repo.GetCommentByID(id)
}

func (s *Service) UpdateComment(in models.Comment) (res models.Comment, err error) {
	return s.repo.UpdateComment(in)
}

func (s *Service) DeleteComment(id int64) (err error) {
	return s.repo.DeleteComment(id)
}
