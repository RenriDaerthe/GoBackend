package main

import (
	"fmt"
	"gemini_backend/controllers"
	"gemini_backend/db"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🔥 Gemini backend starting up...")

	db.Connect()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "✅ Gemini Backend is running!",
		})
	})

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Update this line to listen on 0.0.0.0 to allow external connections
	r.Run("0.0.0.0:3000") // This will allow external devices (like the emulator) to connect
}
