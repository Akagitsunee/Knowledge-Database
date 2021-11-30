package api

import (
	"github.com/gin-gonic/gin"
	"kdb/model"
	"net/http"
)

func GetAllSpaces(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func GetSpaceById(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func CreateSpace(c *gin.Context) {
	var space model.Space

	if err := c.BindJSON(&space); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, space)
}

func UpdateSpace(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func DeleteSpace(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}
