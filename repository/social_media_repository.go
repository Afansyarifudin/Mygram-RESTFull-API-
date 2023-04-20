package repository

import "mygram/models"

type SocialMediaRepo interface {
	CreateMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	GetAllMedia() (res []models.SocialMedia, err error)
	GetMediaByID(id int64) (res models.SocialMedia, err error)
	UpdateMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	DeleteMedia(id int64) (err error)
}

func (r Repo) CreateMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	return res, nil
}
func (r Repo) GetAllMedia() (res []models.SocialMedia, err error) {
	return res, nil
}
func (r Repo) GetMediaByID(id int64) (res models.SocialMedia, err error) {
	return res, nil
}
func (r Repo) UpdateMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	return res, nil
}
func (r Repo) DeleteMedia(id int64) (err error) {
	return nil
}
