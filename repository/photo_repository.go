package repository

import "mygram/models"

type PhotoRepo interface {
	CreatePhoto(in models.Photo) (res models.Photo, err error)
	GetAllPhoto() (res []models.Photo, err error)
	GetPhotoByID(id int64) (res models.Photo, err error)
	UpdatePhoto(in models.Photo) (res models.Photo, err error)
	DeletePhoto(id int64) (err error)
}

func (r Repo) CreatePhoto(in models.Photo) (res models.Photo, err error) {

	if err := r.gorm.Debug().Create(&in).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllPhoto() (res []models.Photo, err error) {

	in := []models.Photo{}

	if err := r.gorm.Debug().Preload("Comments").Preload("User").Find(&in).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetPhotoByID(id int64) (res models.Photo, err error) {

	if err := r.gorm.Debug().First(&res, id).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdatePhoto(in models.Photo) (res models.Photo, err error) {

	if err := r.gorm.Debug().First(&in, in.Id).Error; err != nil {
		return res, err
	}

	if err := r.gorm.Debug().Model(&in).Updates(models.Photo{
		// Title:    in.Title,
		// Caption:  in.Caption,
		// PhotoUrl: in.PhotoUrl,
		// UserID:   in.UserID,
	}).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeletePhoto(id int64) (err error) {

	in := models.Photo{}

	if err := r.gorm.Debug().First(&in, in.Id).Error; err != nil {
		return err
	}

	if err := r.gorm.Debug().Delete(&in).Error; err != nil {
		return err
	}

	return nil
}
