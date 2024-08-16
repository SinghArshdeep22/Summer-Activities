package main

import (
	"coworkingapp/handlers"
	"coworkingapp/middlewares"
	"coworkingapp/models"
	"coworkingapp/utils"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config models.CoworkingConfig

func init() {
	data, err := os.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)

	db, err := gorm.Open(postgres.Open(config.Dsn))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Room{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Booking{})
	seedData(db)

	r := gin.Default()

	r.Use(middlewares.EarlyExitOnPreflightRequests())
	r.Use(middlewares.SetCorsPolicy(config.AllowedOrigin))
	r.Use(func(c *gin.Context) {
		c.Set("DbKey", db)
		c.Set("ConfigKey", config)
		c.Next()
	})

	r.Static("/imgs", "./imgs")
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.GET("/rooms", handlers.GetAllRooms)
	r.GET("/rooms/:id", handlers.GetRoomById)
	r.GET("/rooms/:id/photos", handlers.GetRoomPhotos)
	r.POST("/bookings", middlewares.AuthorizeUser(), handlers.AddBooking)
	r.GET("/bookings", middlewares.AuthorizeUser(), handlers.GetBookingsByUserId)
	r.GET("/bookings/:id", middlewares.AuthorizeUser(), handlers.GetBookingById)
	r.DELETE("/bookings/:id", middlewares.AuthorizeUser(), handlers.DeleteBooking)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func seedData(db *gorm.DB) {
	db.Delete(&models.Booking{}, "1 = 1")
	db.Delete(&models.User{}, "1 = 1")
	db.Delete(&models.Photo{}, "1 = 1")
	db.Delete(&models.Room{}, "1 = 1")

	userId := utils.GetUuid()
	db.Create(&models.User{
		ID:       userId,
		Email:    "hello@gmail.com",
		Username: "Deep",
		Password: "abcd1234!",
	})
	db.Create([]*models.Room{
		{
			ID: utils.GetUuid(), Name: "Green", Cost: 12.50, NumberOfSeats: 4, Category: "Open Space",
			MainPhoto: "green_001.jpg", Photos: []models.Photo{
				{Url: "green_002.jpg"},
				{Url: "green_003.jpg"},
			},
		},
		{
			ID: utils.GetUuid(), Name: "Red", Cost: 100, NumberOfSeats: 50, Category: "Conference Hall",
			MainPhoto: "red_001.jpg", Photos: []models.Photo{
				{Url: "red_002.jpg"},
			},
		},
		{
			ID: utils.GetUuid(), Name: "Yellow", Cost: 4.50, NumberOfSeats: 1, Category: "Shared Desk",
			MainPhoto: "yellow_001.jpg", Photos: []models.Photo{
				{Url: "yellow_002.jpg"},
				{Url: "yellow_003.jpg"},
				{Url: "yellow_004.jpg"},
				{Url: "yellow_005.jpg"},
			},
		},
	})
}
