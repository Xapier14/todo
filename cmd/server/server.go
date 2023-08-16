package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/routes"
)

func setupRouter(config string) *gin.Engine {
	if config != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	fmt.Println("Config: " + config)
	if config == "DEBUG" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	}
	routes.InitializeRoutes(router, "")
	return router
}

func main() {
	fmt.Println("Loading configuration...")
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	config := os.Getenv("CONFIG")
	if config == "" {
		config = "DEBUG"
	}
	
	fmt.Println("Initializing database connection...")
	err := db.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Auto migrating database...")
	err = db.AutoMigrate()
	if err != nil {
		panic(err)
	}

	fmt.Println("Initializing router...")
	router := setupRouter(config)

	fmt.Printf("Listening on port %s...", port)
	router.Run(":" + port)
}