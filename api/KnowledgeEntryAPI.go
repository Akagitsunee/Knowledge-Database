package api

import (
	"github.com/gin-gonic/gin"
	"kdb/model"
	"net/http"
)

func GetAllEntries(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func GetEntryById(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func CreateEntry(c *gin.Context) {
	var entry model.KnowledgeEntry

	if err := c.BindJSON(&entry); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, entry)
}

func UpdateEntry(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func DeleteEntry(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}
