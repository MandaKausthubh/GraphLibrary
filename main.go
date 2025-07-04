package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/MandaKausthubh/GraphLibrary/utils"
	"github.com/MandaKausthubh/GraphLibrary/GraphLib"
	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"%s - %s %s %d %s %s\n",
			params.ClientIP,
			params.Method,
			params.Path,
			params.StatusCode,
			params.TimeStamp.Format(time.RFC822Z),
			params.Latency,
		)
	})
}

func setupLogFile() {
	logFile, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}



func main() {
	fmt.Println("Hello, World!");
	fmt.Println("The sum of 3 and 5 is:", utils.Add(3, 5))


	apiKey := os.Getenv("GRAPH_HOPPER_API_KEY")
	if apiKey == "" {
        fmt.Println("API key not found in environment")
        return
    }

	graphlib.CallGraphHopper(
		graphlib.GHPoint{11.539421, 48.118477},
		graphlib.GHPoint{11.559023,48.12228},
		apiKey,
	)

	server := gin.New()
	server.Use(CustomLogger(), gin.Recovery())
	setupLogFile()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" : "OK!!",
		})
	})

	server.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" : "JUST OK!!",
		})
	})

	server.Run(":8080")
}
