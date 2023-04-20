package services

import "mygram/models"

type UserService interface {
	RegisterUser(in models.User) (res models.User, err error)
	LoginUser(in models.User) (res bool, err error)
	GetAllUser() (res []models.User, err error)
	UpdateUser(in models.User) (res models.User, err error)
	DeleteUser(id int64) (err error)
}

func (s *Service) RegisterUser(in models.User) (res models.User, err error) {
	return s.repo.RegisterUser(in)
}

func (s *Service) LoginUser(in models.User) (res bool, err error) {
	return s.repo.LoginUser(in)
}

func (s *Service) GetAllUser() (res []models.User, err error) {
	res, err = s.repo.GetAllUser()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) UpdateUser(in models.User) (res models.User, err error) {
	return s.repo.UpdateUser(in)
}

func (s *Service) DeleteUser(id int64) (err error) {
	return s.repo.DeleteUser(id)
}
