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
		c.JSON(http.StatusNotFound, APIResponse{
			Status: "error",
			Error:  "Node not found",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Status: "success",
		Data:   node,
	})
}

func GetSubgraphHandler(c *gin.Context) {
	nodeID := c.Param("id")

	subgraph, err := graph.GenerateSubgraph(nodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Status: "error",
			Error:  "Could not generate subgraph",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Status: "success",
		Data:   subgraph,
	})
}
