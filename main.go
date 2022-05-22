package main

import (
	"admin/config"
	"admin/routers"
	"log"
	"net/http"

	"admin/models"
)

func main() {

	var user models.User
	user.Username = "admin"
	user.Password = "admin"
	user.Name = "administrador"
	user.Role = "admin"

	// database.InsertUser(&user)

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

	var counter models.Counter
	counter.Code = "MyCounter"
	counter.Counter = 0
	// database.InsertCounter(&counter)

	port := config.ServerPort()
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println("..." + port)
	log.Fatal(server.ListenAndServe())

}

// git push  https://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@github.com/ltru0495/salesadmin.git
