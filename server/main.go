package main

import (
	"log"
	"os"

	"github.com/bgguna/benxgoro/server/contact"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("benxgoro server started...")

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/send", contact.HandleNewMsg())

	router.Run(":3002")
}
