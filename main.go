package main

import (
	"fmt"
	"io"
	"os"
	"time"
	// "github.com/joho/godotenv"
	"github.com/MandaKausthubh/GraphLibrary/utils"
	// "github.com/MandaKausthubh/GraphLibrary/GraphLib"
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
}
