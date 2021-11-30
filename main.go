package main

import (
	"github.com/gin-gonic/gin"
	"kdb/api"
)

func main() {
	router := gin.Default()

	router.GET("/knowledgeentry", api.GetAll)
	router.GET("/knowledgeentry/:id", api.GetById)
	router.POST("/knowledgeentry", api.Create)
	router.PUT("/knowledgeentry/:id", api.Update)
	router.DELETE("/knowledgeentry/:id", api.Delete)

	router.GET("/knowledgeentryversion", api.GetAll)
	router.GET("/knowledgeentryversion/:id", api.GetById)
	router.POST("/knowledgeentryversion", api.Create)
	router.PUT("/knowledgeentryversion/:id", api.Update)
	router.DELETE("/knowledgeentryversion/:id", api.Delete)

	router.GET("/space", api.GetAll)
	router.GET("/space/:id", api.GetById)
	router.POST("/space", api.Create)
	router.PUT("/space/:id", api.Update)
	router.DELETE("/space/:id", api.Delete)

	router.GET("/user", api.GetAll)
	router.GET("/user/:id", api.GetById)
	router.POST("/user", api.Create)
	router.PUT("/user/:id", api.Update)
	router.DELETE("/user/:id", api.Delete)

	router.Run(":8080")
}
