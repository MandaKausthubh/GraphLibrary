package api

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// Node routes
	router.GET("/nodes/:id", GetNodeHandler)
	router.GET("/subgraph/:id", GetSubgraphHandler)

	// Edge routes
	router.POST("/edges", CreateEdgeHandler)
	router.GET("/edges/:fromID/:toID", GetEdgeHandler)

	// Cache routes
	router.GET("/cache/edge/:fromID/:toID", GetEdgeCacheHandler)
	router.POST("/cache/edge", SetEdgeCacheHandler)

	return router
}
