package api

import (
	"github.com/gin-gonic/gin"
	"kdb/model"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}
