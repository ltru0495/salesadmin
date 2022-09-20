package main

import (
	"admin/models/database"

	"admin/models"
)

func main() {

	var user models.User
	user.Username = "admin2"
	user.Password = "admin"
	user.Name = "administrador"
	user.Role = "admin"

	database.InsertUser(&user)

	user.Username = "cajamarca"
	user.Name = "CAJAMARCA"
	user.Password = "123456"
	user.Role = "user"
	// database.InsertUser(&user)

	user.Username = "loreto"
	user.Name = "LORETO"
	user.Password = "123456"
	user.Role = "user"
	// database.InsertUser(&user)
}

