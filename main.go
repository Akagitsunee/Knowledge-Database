package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kdb/api"
	docs "kdb/docs"
)

// @title           Knowledge Database
// @version         1.0
// @description     Knowledge Database
// @termsOfService

// @contact.name   API Support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      164.128.186.124/
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	docs.SwaggerInfo.Title = "Knowledge Database"

	router.GET("/knowledgeentry", api.GetAllEntries)
	router.GET("/knowledgeentry/:id", api.GetEntryById)
	router.POST("/knowledgeentry", api.CreateEntry)
	router.PUT("/knowledgeentry/:id", api.UpdateEntry)
	router.DELETE("/knowledgeentry/delete/:id", api.DeleteEntry)
	router.GET("/knowledgeentry/versions/:id", api.GetAllEntryVersionsByEntry)
	router.GET("/knowledgeentry/version/:id", api.GetEntryVersionById)
	router.GET("/knowledgeentry/community/:id", api.GetEntriesByCommunity)

	router.GET("/community", api.GetAllCommunities)
	router.GET("/community/:id", api.GetCommunityById)
	router.POST("/community", api.CreateCommunity)
	router.PUT("/community/:id", api.UpdateCommunity)
	router.DELETE("/community/:id", api.DeleteCommunity)
	router.POST("/cty/:communityId/employee/:employeeId", api.AddEmployee)
	router.DELETE("/cty/:communityId/employee/:employeeId", api.RemoveEmployee)
	router.GET("/community/employees/:id", api.GetAllCommunityMemebers)

	router.GET("/employee", api.GetAllEmployees)
	router.GET("/employee/:id", api.GetEmployeeById)
	router.POST("/employee", api.CreateEmployee)
	router.PUT("/employee/:id", api.UpdateEmployee)
	router.DELETE("/employee/:id", api.DeleteEmployee)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
