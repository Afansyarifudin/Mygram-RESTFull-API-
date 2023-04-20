package controllers

import (
	"mygram/config"
	"mygram/helper"
	"mygram/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateSocialMedia godoc
//	@Summary		Create Social Media
//	@Description	Create Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
// 	@Param        	photo 	body 	models.SocialMediaReqs true "Create new social media"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/socialmedias [post]
func (h HttpServer) CreateMedia(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	contentType := helper.GetContentType(c)
	Media := models.SocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.UserID = uint(userId)
	Media.CreatedAt = time.Now()
	Media.UpdatedAt = time.Now()

	if err := db.Debug().Create(&Media).Error; err != nil {
		helper.BadRequest(c, "Failed to create social media")
		return
	}

	helper.Ok(c, gin.H{
		"id":               Media.Id,
		"name":             Media.Name,
		"social_media_url": Media.SocialMediaUrl,
		"user_id":          Media.UserID,
		"created_at":       Media.CreatedAt,
	})
}

// GetAllSocialMedia godoc
//	@Summary		GetAll Social Media
//	@Description	GetAll Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/socialmedias [get]
func (h HttpServer) GetAllMedia(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	Medias := []models.SocialMedia{}

	if err := db.Debug().Preload("User").Find(&Medias).Error; err != nil {
		helper.BadRequest(c, "Social Media Not Found")
		return
	}

	helper.Ok(c, gin.H{
		"media": Medias,
	})
}

// GetByIDSocialMedia godoc
//	@Summary		GetByID Social Media
//	@Description	GetByID Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
// 	@Param        	Id 	path 	int true "Id of the social media"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/socialmedias/{mediaId} [get]
func (h HttpServer) GetMediaByID(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	getId, _ := strconv.Atoi(c.Param("mediaId"))
	Media := models.SocialMedia{}

	if err := db.Debug().Where("id = ?", getId).First(&Media).Error; err != nil {
		helper.BadRequest(c, "Social Media Not Found")
		return
	}
	helper.Ok(c, gin.H{
		"id":               Media.Id,
		"name":             Media.Name,
		"social_media_url": Media.SocialMediaUrl,
		"user_id":          Media.UserID,
		"created_at":       Media.CreatedAt,
	})
}

// UpdateSocialMedia godoc
//	@Summary		Update Social Media
//	@Description	Update Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
// 	@Param        	photo 	body 	models.SocialMediaReqs true "Create new social media"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/socialmedias [put]
func (h HttpServer) UpdateMedia(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	contentType := helper.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("mediaId"))

	Media := models.SocialMedia{}
	OldMedia := models.SocialMedia{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.UserID = uint(userId)
	Media.UpdatedAt = time.Now()

	if err := db.Debug().First(&OldMedia, getId).Error; err != nil {
		helper.BadRequest(c, "Social Media Not Found")
		return
	}

	if err := db.Debug().Model(&OldMedia).Updates(&Media).Error; err != nil {
		helper.BadRequest(c, "Failed to update social media")
		return
	}

	helper.Ok(c, gin.H{
		"id":               OldMedia.Id,
		"name":             OldMedia.Name,
		"social_media_url": OldMedia.SocialMediaUrl,
		"user_id":          OldMedia.UserID,
		"updated_at":       OldMedia.UpdatedAt,
	})
}

// DeleteSocialMedia godoc
//	@Summary		Delete Social Media
//	@Description	Delete Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
// 	@Param        	Id 	path 	int true "Id of the social media"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/socialmedias/{mediaId} [delete]
func (h HttpServer) DeleteMedia(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	getId, _ := strconv.Atoi(c.Param("mediaId"))
	Media := models.SocialMedia{}

	if err := db.Debug().First(&Media, getId).Error; err != nil {
		helper.BadRequest(c, "Media Not Found")
		return
	}

	if err := db.Debug().Delete(&Media).Error; err != nil {
		helper.BadRequest(c, "Failed to delete data")
		return
	}

	helper.Ok(c, "Successfully delete the social media")
}
