package main

import (
	"fmt"
	"log"

	"database/internal/db"
	"database/internal/service"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Println("DB Connected")

	userService := service.NewUserService(database)

	userService.CreateTable()

	userService.CreateUser("John", "john@example.com")
	userService.CreateUser("Alice", "alice@example.com")

	user := userService.GetUser(1)
	fmt.Println("User:", user)

	users := userService.GetAllUsers()
	fmt.Println("Users:", users)

	userService.UpdateUser(1, "John Updated")
	userService.DeleteUser(2)
}
