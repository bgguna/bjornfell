package contact

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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
		db, _ := sql.Open("sqlite3", "./storage/benxgoro.db")
		message := contactMessage{}
		rawContextData, err := context.GetRawData()
		if err != nil {
			log.Println("Failed to process request.", err)
			context.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		}

		err = json.Unmarshal(rawContextData, &message)
		if err != nil {
			log.Println("Error processing request.", err)
			context.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		}
		log.Printf("Received contact message: %+v", message)
		statement, err := db.Prepare("INSERT INTO contact (name, email, message) VALUES (?, ?, ?)")
		if err != nil {
			log.Println("Error preparing to store contact message.", err)
			context.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		}

		_, err = statement.Exec(message.Name, message.Email, message.Message)
		if err != nil {
			log.Println("Error storing contact message.", err)
			context.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		}

		log.Println("Contact message stored.")
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
