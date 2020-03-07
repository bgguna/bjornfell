package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type contactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("benxgoro server started...")

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/send", func(context *gin.Context) {
		message := contactMessage{}
		rawContextData, err := context.GetRawData()
		if err != nil {
			log.Println("Failed to process request.", err)
			context.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		}

		err = json.Unmarshal(rawContextData, &message)
		if err != nil {
			log.Println("Error processing request")
			context.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		}
		log.Printf("Received contact message: %+v", message)
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	router.Run(":3002")
}
