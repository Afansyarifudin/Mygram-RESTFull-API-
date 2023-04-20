package services

import "mygram/models"

type PhotoService interface {
	CreatePhoto(in models.Photo) (res models.Photo, err error)
	GetAllPhoto() (res []models.Photo, err error)
	GetPhotoByID(id int64) (res models.Photo, err error)
	UpdatePhoto(in models.Photo) (res models.Photo, err error)
	DeletePhoto(id int64) (err error)
}

func (s *Service) CreatePhoto(in models.Photo) (res models.Photo, err error) {
	return s.repo.CreatePhoto(in)
}
func (s *Service) GetAllPhoto() (res []models.Photo, err error) {
	res, err = s.repo.GetAllPhoto()
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) GetPhotoByID(id int64) (res models.Photo, err error) {
	return s.repo.GetPhotoByID(id)
}
func (s *Service) UpdatePhoto(in models.Photo) (res models.Photo, err error) {
	return s.repo.UpdatePhoto(in)
}
func (s *Service) DeletePhoto(id int64) (err error) {
	return s.repo.DeletePhoto(id)
}
