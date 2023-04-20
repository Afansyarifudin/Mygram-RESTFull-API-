package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `data:"data"`
}

var (
	ErrorNotFound = "record Not Found"
)

// status Ok

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func OkWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

// Status Bad

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	obj := gin.H{"status": http.StatusBadRequest, "message": message}
	if len(data) > 0 {
		obj["data"] = data[0]
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: message,
	})
}

// Unautenticated
// func Unauthenticated(c *gin.Context, message string) {
// 	c.JSON(http.StatusUnauthorized, Response{
// 		Status:  http.StatusUnauthorized,
// 		Message: "Error Unauthenticated",
// 	})
// }
