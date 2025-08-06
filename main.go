package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/MandaKausthubh/GraphLibrary/internal/api"
	"github.com/MandaKausthubh/GraphLibrary/internal/db"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
	"github.com/MandaKausthubh/GraphLibrary/internal/redis_cache"
)

func main() {
	// Initialize Redis Cache
	cacheClient := cache.NewRedisClient("localhost:6379", "", 0)

	// Initialize PostgreSQL DB
	dbClient, err := db.NewPostgresDB("postgres://user:password@localhost:5432/yourdb?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	// Create repository and services
	nodeRepo := db.NewNodeRepository(dbClient)
	edgeRepo := db.NewEdgeRepository(dbClient)

	graphService := graph.NewGraphService(edgeRepo, nodeRepo, cacheClient)

	handler := api.NewHandler(graphService)

	// Setup Gin Router
	r := gin.Default()

	// Register routes
	r.GET("/edge", handler.GetEdgeHandler)
	r.POST("/edge", handler.CreateEdgeHandler)

	r.GET("/node/:id", handler.GetNodeHandler)
	r.POST("/node", handler.CreateNodeHandler)

	// Start the server
	log.Println("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
