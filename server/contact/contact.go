package contact

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// contactMessage represents the structure of an incoming Contact message.
type contactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// HandleNewMsg handles incoming messages from the Contact tab.
func HandleNewMsg() func(context *gin.Context) {
	return func(context *gin.Context) {
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
	}
}
