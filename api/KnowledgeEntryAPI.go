package api

import (
	"github.com/gin-gonic/gin"
	"kdb/httputil"
	"kdb/model"
	"kdb/repository"
	"net/http"
)

var knowledgeEntryRepository repository.KnowledgeEntryRepository

func init() {
	knowledgeEntryRepository = knowledgeEntryRepository.InitKnowledgeEntryRepository()
}

// GetAllEntries godoc
// @Summary      GetAllEntries
// @Description  GetAllEntries
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.KnowledgeEntry
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry [get]
func GetAllEntries(c *gin.Context) {
	e, err := knowledgeEntryRepository.GetAllKnowledgeEntries()

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, e)
}

// GetEntryById godoc
// @Summary      GetEntryById
// @Description  GetEntryById
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Param        id path int  true  "Id of knowledgeentry"  Format(int)
// @Success      200  {object}   model.KnowledgeEntry
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry/{id} [get]
func GetEntryById(c *gin.Context) {
	id := c.Param("id")

	e, err := knowledgeEntryRepository.GetKnowledgeEntryById(id)

	if e.KnowledgeEntryId == "0" || e.KnowledgeEntryId == "" {
		c.AbortWithStatus(404)
		return
	}

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, e)
}

// CreateEntry godoc
// @Summary      CreateEntry
// @Description  CreateEntry
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Param        knowledgeentry body model.KnowledgeEntry true "Create knowledgeentry"
// @Success      200  {object}  model.KnowledgeEntry
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry [post]
func CreateEntry(c *gin.Context) {
	var e model.KnowledgeEntry

	c.ShouldBindJSON(&e)

	communities, err := knowledgeEntryRepository.AddKnowledgeEntry(&e)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, communities)
}

// UpdateEntry godoc
// @Summary      UpdateEntry
// @Description  UpdateEntry
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Param id path int true "ID of the KnowledgeEntryVersion to be updated"
// @Param KnowledgeEntryVersion body model.KnowledgeEntryVersion true "Updated knowledgeentry"
// @Success      200  {object}   model.KnowledgeEntryVersion
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry/{id} [put]
func UpdateEntry(c *gin.Context) {
	var e model.KnowledgeEntryVersion
	id := c.Param("id")

	c.ShouldBindJSON(&e)

	r, err := knowledgeEntryRepository.ModifyKnowledgeEntry(id, &e)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, r)
}

// DeleteEntry godoc
// @Summary      DeleteEntry
// @Description  DeleteEntry
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Param        id path string true "ID of the KnowledgeEntryVersion to be deleted"
// @Success 	 204 "No Content"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry [delete]
func DeleteEntry(c *gin.Context) {
	id := c.Param("id")

	err := knowledgeEntryRepository.DeleteKnowledgeEntry(id)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, id)
}

// GetAllEntryVersionsByEntry godoc
// @Summary      GetAllEntryVersionsByEntry
// @Description  GetAllEntryVersionsByEntry
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Param        id path int  true  "Id of knowledgeentry"  Format(int)
// @Success      200  {array}   model.KnowledgeEntry
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry/versions/{id} [get]
func GetAllEntryVersionsByEntry(c *gin.Context) {
	id := c.Param("id")

	e, err := knowledgeEntryRepository.GetKnowledgeEntryVersionsById(id)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, e)
}

// GetEntryVersionById godoc
// @Summary      GetEntryVersionById
// @Description  GetEntryVersionById
// @Tags         KnowledgeEntry
// @Accept       json
// @Produce      json
// @Param        id path int  true  "Id of knowledgeentry"  Format(int)
// @Success      200  {object}   model.KnowledgeEntry
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /knowledgeentry/version/{id} [get]
func GetEntryVersionById(c *gin.Context) {
	id := c.Param("id")

	e, err := knowledgeEntryRepository.GetEntryVersionById(id)

	if e.KnowledgeEntryVersionId == "0" || e.KnowledgeEntryVersionId == "" {
		c.AbortWithStatus(404)
		return
	}

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, e)
}
