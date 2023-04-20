package repository

import "mygram/models"

type CommentRepo interface {
	CreateComment(in models.Comment) (res models.Comment, err error)
	GetAllComment() (res []models.Comment, err error)
	GetCommentByID(id int64) (res models.Comment, err error)
	UpdateComment(in models.Comment) (res models.Comment, err error)
	DeleteComment(id int64) (err error)
}

func (r Repo) CreateComment(in models.Comment) (res models.Comment, err error) {
	return res, nil
}
func (r Repo) GetAllComment() (res []models.Comment, err error) {
	return res, nil
}
func (r Repo) GetCommentByID(id int64) (res models.Comment, err error) {
	return res, nil
}
func (r Repo) UpdateComment(in models.Comment) (res models.Comment, err error) {
	return res, nil
}
func (r Repo) DeleteComment(id int64) (err error) {
	return nil
}
