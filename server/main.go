package main

import (
	"bjornfell/contact"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("BjornFell Server started...")

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/send", contact.HandleNewMsg())

	router.Run(":3002")
}
