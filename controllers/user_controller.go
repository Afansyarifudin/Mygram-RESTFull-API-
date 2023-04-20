package controllers

import (
	"mygram/config"
	"mygram/helper"
	"mygram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
//	@Summary		Register User
//	@Description	Register User
//	@Tags			Users
//	@Accept			json
//	@Produce		json
// 	@Param        	user 	body 	models.UserRegister true "Register New User"
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/users/register [post]
func (h HttpServer) RegisterUser(c *gin.Context) {
	contentType := helper.GetContentType(c)
	_ = contentType

	in := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&in)
	} else {
		c.ShouldBind(&in)
	}

	in.CreatedAt = time.Now()
	in.UpdatedAt = time.Now()

	// call service
	res, err := h.app.RegisterUser(in)
	if err != nil {
		helper.BadRequest(c, "Failed to register user")
		return
	}

	helper.OkWithMessage(c, "Successfully register user", gin.H{
		"id":       res.Id,
		"username": res.Username,
		"email":    res.Email,
	})

}

// LoginUser godoc
//	@Summary		Login User
//	@Description	Login User
//	@Tags			Users
//	@Accept			json
//	@Produce		json
// 	@Param        	user 	body 	models.UserLogin true "Login User"
//	@Success		201	{object} helper.Response
//	@Failure		404	{object} helper.Response
//	@Failure		500	{object} helper.Response
//	@Router			/users/login [post]
func (h HttpServer) LoginUser(c *gin.Context) {
	contentType := helper.GetContentType(c)
	db, _ := config.InitGorm()

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password := User.Password

	if err := db.Debug().Where("email=?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	// fmt.Println((User.Password), (password))
	if comparePass := helper.ComparePassword([]byte(User.Password), []byte(password)); !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	token := helper.GenerateToken(uint(User.GormModel.Id), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

	// fmt.Println((user.Password), (passwordClient))

	// isValid := helper.ComparePassword([]byte(user.Password), []byte(passwordClient))
	// if !isValid {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":   "unauthorized",
	// 		"message": "Invalid email/password",
	// 	})
	// 	return
	// }

}

func (h HttpServer) GetAllUser(c *gin.Context) {

}

func (h HttpServer) UpdateUser(c *gin.Context) {

}

func (h HttpServer) DeleteUser(c *gin.Context) {

}
