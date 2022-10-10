package main

import (
	"admin/config"
	"admin/models/database"
	"admin/routers"
	"log"
	"net/http"

	"admin/models"
)

func main() {
	// f, err := os.OpenFile("server.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("Error")
	// }
	// defer f.Close()
	// log.SetOutput(f)

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
