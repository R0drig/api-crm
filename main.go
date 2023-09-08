package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

)

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&User{})

	app := gin.Default()
	handlers := newHandlers(db)
	app.POST("/register", handlers.registerUser)
	app.POST("/login", handlers.loginUser)

	auth := app.Group("/auth")
	auth.Use(authenticateUser)
	{
		auth.GET("/user", handlers.getUserInfo)
		auth.POST("/register-lead", handlers.createLead)
		auth.GET("/leads", handlers.showAllUsersLeads)
		auth.PATCH("/update-lead/:id", handlers.updateUserLead)
		auth.DELETE("/update-lead/:id", handlers.deleteUserLead)
	}

	app.Run(":8080")
}
