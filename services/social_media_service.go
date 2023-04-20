package services

import "mygram/models"

type SocialMediaService interface {
	CreateMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	GetAllMedia() (res []models.SocialMedia, err error)
	GetMediaByID(id int64) (res models.SocialMedia, err error)
	UpdateMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	DeleteMedia(id int64) (err error)
}

func (s *Service) CreateMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	return s.repo.CreateMedia(in)
}

func (s *Service) GetAllMedia() (res []models.SocialMedia, err error) {
	res, err = s.repo.GetAllMedia()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetMediaByID(id int64) (res models.SocialMedia, err error) {
	return s.repo.GetMediaByID(id)
}

func (s *Service) UpdateMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	return s.repo.UpdateMedia(in)
}

func (s *Service) DeleteMedia(id int64) (err error) {
	return s.repo.DeleteMedia(id)
}
