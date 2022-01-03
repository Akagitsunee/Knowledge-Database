package api

import (
	"github.com/gin-gonic/gin"
	"kdb/httputil"
	"kdb/model"
	"kdb/repository"
	"net/http"
)

var communityRepository repository.CommunityRepository

func init() {
	communityRepository = communityRepository.InitCommunityRepository()
}

// GetAllCommunities godoc
// @Summary      List communities
// @Description  get communities
// @Tags         Community
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Community
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /community [get]
func GetAllCommunities(c *gin.Context) {
	communities, err := communityRepository.GetAllCommunitys()

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, communities)
}

// GetCommunityById godoc
// @Summary      GetCommunityById
// @Description  GetCommunityById
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        id path int  true  "Id of community"  Format(int)
// @Success      200  {object}   model.Community
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /community/{id} [get]
func GetCommunityById(c *gin.Context) {
	id := c.Param("id")

	communities, err := communityRepository.GetCommunityById(id)

	if communities.CommunityId == "0" || communities.CommunityId == "" {
		c.AbortWithStatus(404)
		return
	}

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, communities)
}

// CreateCommunity godoc
// @Summary      CreateCommunity
// @Description  CreateCommunity
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        community body model.Community true "Create community"
// @Success      200  {object}  model.Community
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /community [post]
func CreateCommunity(c *gin.Context) {
	var e model.Community

	c.ShouldBindJSON(&e)

	communities, err := communityRepository.AddCommunity(&e)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, communities)
}

// UpdateCommunity godoc
// @Summary      UpdateCommunity
// @Description  UpdateCommunity
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param id path int true "ID of the community to be updated"
// @Param community body model.Community true "Updated community"
// @Success      200  {object}   model.Community
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /community/{id} [put]
func UpdateCommunity(c *gin.Context) {
	var e model.Community
	id := c.Param("id")

	c.ShouldBindJSON(&e)

	communities, err := communityRepository.ModifyCommunity(id, &e)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, communities)
}

// DeleteCommunity godoc
// @Summary      DeleteCommunity
// @Description  DeleteCommunity
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        id path string true "ID of the community to be deleted"
// @Success 	 204 "No Content"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /community/{id} [delete]
func DeleteCommunity(c *gin.Context) {
	id := c.Param("id")

	err := communityRepository.DeleteCommunity(id)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, id)
}

// AddEmployee godoc
// @Summary      AddEmployee
// @Description  AddEmployee
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        communityId   path      int  true  "Community ID"
// @Param        employeeId   path      int  true  "employee ID"
// @Success      200   {string}  string    "ok"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /cty/{communityId}/employee/{employeeId} [post]
func AddEmployee(c *gin.Context) {
	cId := c.Param("communityId")
	eId := c.Param("employeeId")

	err := communityRepository.AddEmployee(cId, eId)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// RemoveEmployee godoc
// @Summary      RemoveEmployee
// @Description  RemoveEmployee
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        communityId   path      int  true  "Community ID"
// @Param        employeeId   path      int  true  "employee ID"
// @Success      200   {string}  string    "ok"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /cty/{communityId}/employee/{employeeId} [delete]
func RemoveEmployee(c *gin.Context) {
	cId := c.Param("communityId")
	eId := c.Param("employeeId")

	err := communityRepository.RemoveEmployee(cId, eId)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

// GetAllCommunityMemebers godoc
// @Summary      GetAllCommunityMemebers
// @Description  GetAllCommunityMemebers
// @Tags         Community
// @Accept       json
// @Produce      json
// @Param        id path int  true  "Id of community"  Format(int)
// @Success      200  {array}   model.Employee
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /community/employees/{id} [get]
func GetAllCommunityMemebers(c *gin.Context) {
	id := c.Param("id")

	e, err := communityRepository.GetAllCommunityMembersById(id)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, e)
}
