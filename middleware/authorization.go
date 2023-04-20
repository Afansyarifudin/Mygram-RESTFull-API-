package middleware

import (
	"mygram/config"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := config.InitGorm()
		getId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		Photo := models.Photo{}

		if err := db.Preload("User").Preload("Comments").First(&Photo, getId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Data Not Found",
				"message": err.Error(),
			})
			return
		}

		if uint(UserId) != Photo.UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := config.InitGorm()
		getId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		Comment := models.Comment{}

		if err := db.Preload("User").Preload("Photo").First(&Comment, getId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Data Not Found",
				"message": err.Error(),
			})
		}

		if uint(UserId) != Comment.UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data",
			})
			return
		}
		c.Next()
	}
}

func SosMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := config.InitGorm()
		getId, err := strconv.Atoi(c.Param("mediaId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		Media := models.SocialMedia{}
		if err := db.Preload("User").First(&Media, getId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Data Not Found",
				"message": err.Error(),
			})
			return
		}

		if int(UserId) != int(Media.UserID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data",
			})
			return
		}
		c.Next()
	}
}
