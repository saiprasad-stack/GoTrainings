package main

import (
	"gorm-gin-postgres/config"
	"gorm-gin-postgres/database"
	"gorm-gin-postgres/routes"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg)
	router := routes.SetupRouter()
	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
