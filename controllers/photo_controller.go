package controllers

import (
	"mygram/config"
	"mygram/helper"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
//	@Summary		Create Photo
//	@Description	Create Photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
// 	@Param        	photo 	body 	models.PhotoReqs true "Create new photo"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/photos [post]
func (h HttpServer) CreatePhoto(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	Photo := models.Photo{}
	contextType := helper.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint(userId)
	Photo.CreatedAt = time.Now()
	Photo.UpdatedAt = time.Now()

	if err := db.Debug().Create(&Photo).Error; err != nil {
		helper.BadRequest(c, "Failed to create photo")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

// GetAllPhoto godoc
//	@Summary		GetAll Photo
//	@Description	GetAll Photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/photos [get]
func (h HttpServer) GetAllPhoto(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}
	Photos := []models.Photo{}

	if err := db.Debug().Preload("Comments").Preload("User").Find(&Photos).Error; err != nil {
		helper.BadRequest(c, "Photo Not Found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": Photos,
	})
}

// GetPhotoByID godoc
//	@Summary		GetByID Photo
//	@Description	GetByID Photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Param			ID	path		int	true	"ID of the photo"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/photos/{photoId} [get]
func (h HttpServer) GetPhotoByID(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}
	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err := db.Debug().Where("id = ?", GetId).First(&Photo).Error; err != nil {
		helper.BadRequest(c, "Photo Not Found")
		return
	}
	helper.Ok(c, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})

}

// UpdatePhoto godoc
//	@Summary		Update Photo
//	@Description	Update Photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Param			ID	path		int	true	"ID of the photo"
// 	@Param        	photo 	body 	models.PhotoReqs true "Create new photo"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/photos/{photoId} [put]
func (h HttpServer) UpdatePhoto(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}
	GetId, _ := strconv.Atoi(c.Param("photoId"))

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"]

	contextType := helper.GetContentType(c)
	Photo := models.Photo{}
	OldPhoto := models.Photo{}

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UpdatedAt = time.Now()
	Photo.UserID = uint(userId.(float64))

	if err := db.Debug().First(&OldPhoto, GetId).Error; err != nil {
		helper.BadRequest(c, "Photo not found")
		return
	}

	if err := db.Debug().Model(&OldPhoto).Updates(&Photo).Error; err != nil {
		helper.BadRequest(c, "Failed to update photo")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldPhoto.Id,
		"title":      OldPhoto.Title,
		"caption":    OldPhoto.Caption,
		"photo_url":  OldPhoto.PhotoUrl,
		"user_id":    OldPhoto.UserID,
		"updated_at": OldPhoto.UpdatedAt,
	})
}

// DeletePhoto godoc
//	@Summary		Delete Photo
//	@Description	Delete Photo
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Param			ID	path		int	true	"ID of the photo"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/photos/{photoId} [delete]
func (h HttpServer) DeletePhoto(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err := db.Debug().First(&Photo, GetId).Error; err != nil {
		helper.BadRequest(c, "Photo not found")
		return
	}

	if err := db.Debug().Delete(&Photo).Error; err != nil {
		helper.BadRequest(c, "Failed to delete photo")
		return
	}

	helper.Ok(c, "Successfully delete the photo")

}
