package main

import (
	"log"
	"studi-kasus/handlers"
	"studi-kasus/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Koneksi Database
	dsn := "host=localhost user=postgres password=123 dbname=user_manag port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Connection Failed")
	}

	// Migrate
	db.AutoMigrate(&models.User{})

	userRepository := models.NewRepository(db)
	userService := models.NewService(userRepository)

	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/user", userHandler.UserHandlers)
	r.GET("/user/:id", userHandler.UserByIdHandlers)
	r.PUT("/user/:id", userHandler.UpdateUserHandler)
	r.GET("/user/:id/:action", userHandler.UserByLimitOffset)
	r.POST("/user", userHandler.AddUserHandlers)
	r.DELETE("/user/:id", userHandler.DeleteUserHandlers)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Hash Password

func hashAndSalt(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
