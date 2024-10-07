package main

import (
	"GolangProject/controllers"
	"GolangProject/db"
	"GolangProject/middlewares"
	"GolangProject/migrations"
	"GolangProject/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitDB()
	migrations.Migrate(db.DB)

	// Start background process
	go services.StartTransferBackground()

	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.POST("/topup", controllers.TopUp)
	auth.POST("/transfer", controllers.Transfer)

	r.Run()
}
