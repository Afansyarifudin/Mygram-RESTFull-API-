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

// CreateComment godoc
//	@Summary		Create Comment
//	@Description	Create Comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
// 	@Param        	photo 	body 	models.CommentReqs true "Create new comment"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/comments [post]
func (h HttpServer) CreateComment(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	contentType := helper.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	Comment := models.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	if err := db.Debug().Find(&models.Comment{}, Comment.PhotoID).Error; err != nil {
		helper.BadRequest(c, "Photo not found")
		return
	}

	Comment.UserID = uint(userId)
	Comment.CreatedAt = time.Now()
	Comment.UpdatedAt = time.Now()

	if err := db.Debug().Create(&Comment).Error; err != nil {
		helper.BadRequest(c, "Failed to create comment")
		return
	}

	helper.Ok(c, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

// GetAllComment godoc
//	@Summary		GetAll Comment
//	@Description	GetAll Comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/comments [get]
func (h HttpServer) GetAllComment(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	Comments := []models.Comment{}

	if err := db.Debug().Preload("User").Preload("Photo").Find(&Comments).Error; err != nil {
		helper.BadRequest(c, "Comment Not found")
		return
	}

	helper.Ok(c, gin.H{
		"Comments": Comments,
	})
}

// GetCommentByID godoc
//	@Summary		GetByID Comment
//	@Description	GetByID Comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
// @Security    Bearer
// 	@Param        	ID 	path 	int true "Id of the comment"
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/comments/{commentId} [get]
func (h HttpServer) GetCommentByID(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}
	GetId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	if err := db.Debug().Where("id = ?", GetId).First(&Comment).Error; err != nil {
		helper.BadRequest(c, "Comment Not Found")
		return
	}

	helper.Ok(c, gin.H{
		"Comment": Comment,
	})
}

// UpdateComment godoc
//	@Summary		Update Comment
//	@Description	Update Comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
// 	@Param        	photo 	body 	models.CommentReqs true "Update comment"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/comments/{commentId} [put]
func (h HttpServer) UpdateComment(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	contentType := helper.GetContentType(c)
	Comment := models.Comment{}
	OldComment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = uint(userId)
	Comment.UpdatedAt = time.Now()

	if err := db.Debug().First(&OldComment, getId).Error; err != nil {
		helper.BadRequest(c, "Comment Not Found")
		return
	}

	if err := db.Debug().Model(&OldComment).Updates(&Comment).Error; err != nil {
		helper.BadRequest(c, "Failed to update comment")
		return
	}

	helper.Ok(c, gin.H{
		"id":         OldComment.Id,
		"message":    OldComment.Message,
		"user_id":    OldComment.UserID,
		"updated_at": OldComment.UpdatedAt,
	})
}

// DeleteComment godoc
//	@Summary		Delete Comment
//	@Description	Delete Comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
// 	@Param        	ID 	path 	int true "Id of the comment"
// @Security    Bearer
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/comments [delete]
func (h HttpServer) DeleteComment(c *gin.Context) {
	db, err := config.InitGorm()
	if err != nil {
		panic(err)
	}

	getId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	if err := db.Debug().First(&Comment, getId).Error; err != nil {
		helper.BadRequest(c, "Comment Not Found")
		return
	}

	if err := db.Debug().Delete(&Comment).Error; err != nil {
		helper.BadRequest(c, "Failed to delete comment")
		return
	}

	helper.Ok(c, "Successfully delete the comment")
}
