package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/MandaKausthubh/GraphLibrary/internal/db"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
)

func GetNodeHandler(c *gin.Context) {
	nodeID := c.Param("id")

	node, err := db.GetNodeByID(nodeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  "Node not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "success",
		"data":   node,
	})
}

func GetSubgraphHandler(c *gin.Context) {
	nodeID := c.Param("id")

	subgraph, err := graph.GetSubGraph(nodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Could not generate subgraph",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   subgraph,
	})
}
