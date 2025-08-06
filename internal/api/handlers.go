package api

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/MandaKausthubh/GraphLibrary/internal/db"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
	"github.com/MandaKausthubh/GraphLibrary/internal/redis_cache"
	"github.com/MandaKausthubh/GraphLibrary/internal/router"
)


type Handler struct {
	GHApiKey string
	EdgeDB db.EdgeRepository
	NodeDB db.NodeRepository
	RedisClient  *cache.RedisClient
}

func (h *Handler) GetNodeHandler(c *gin.Context) {
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


func (h *Handler) CreateNodeHandler(c *gin.Context) {
	var node graph.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request data",
		})
		return
	}

	if err := h.NodeDB.CreateNode(&node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Could not create node",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
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

type CreateEdgeRequest struct {
    FromNodeID    string  `json:"from_node_id" binding:"required"`
    ToNodeID      string  `json:"to_node_id" binding:"required"`
    DistanceKm    float64 `json:"distance_km" binding:"required"`
    TravelTimeSec int     `json:"travel_time_sec" binding:"required"`
    Metadata      string  `json:"metadata"` // optional
}

// TODO: Implement the following handlers
func (h *Handler) CreateEdgeHandler(c *gin.Context) {
	// This function will handle the creation of an edge
	// It will parse the request body, validate it, and then create the edge in the database
	// If successful, it will return the created edge as a JSON response
	var req CreateEdgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	edge := &graph.Edge{
		FromNodeID:    req.FromNodeID,
		ToNodeID:      req.ToNodeID,
		DistanceKm:    req.DistanceKm,
		TravelTimeSec: req.TravelTimeSec,
		Metadata:      req.Metadata,
		CreatedAt:     time.Now(),
	}

	if err := h.EdgeDB.CreateEdge(edge); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create edge"})
		return
	}

	// Cache the edge in RedisClient
	if err := h.RedisClient.CacheEdge(edge, time.Hour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache edge"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   edge,
	})
}

func (h *Handler) GetEdgeHandler(c *gin.Context) {
	fromID := c.Param("from")
	toID := c.Param("to")

	// Step 1: Check Redis
	edge, err := h.RedisClient.GetCachedEdge(fromID, toID)
	if err == nil && edge != nil {
		c.JSON(http.StatusOK, edge)
		return
	}

	// Step 2: Check DB
	edge, err = h.EdgeDB.GetEdge(fromID, toID)
	if err == nil && edge != nil {
		// Update Redis
		_ = h.RedisClient.CacheEdge(edge, time.Hour)
		c.JSON(http.StatusOK, edge)
		return
	}

	// Step 3: Get node details to call GraphHopper
	fromNode, err1 := h.NodeDB.GetNodeByID(fromID)
	toNode, err2 := h.NodeDB.GetNodeByID(toID)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node IDs"})
		return
	}

	// Step 4: Query GraphHopper
	// gh_response := &router.GHResponse{}
	start := router.GHPoint{ fromNode.Latitude, fromNode.Latitude }
	end := router.GHPoint{ toNode.Latitude, toNode.Latitude }
	gh_response, err := router.CallGraphHopper(start, end, h.GHApiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from GraphHopper"})
		return
	}
	edge, err = router.ConvertGHResponseToEdge(gh_response, fromNode.ID, toNode.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GraphHopper Response conversion failed"})
		return
	}

	// Step 5: Save to DB and Redis
	// if err := h.DB.CreateEdge(edge); err != nil {
	// 	// Optional: log, don’t fail
	// }
	_ = h.RedisClient.CacheEdge(edge, time.Hour)

	c.JSON(http.StatusOK, edge)
}

