package main

import (
	"fmt"
	"go-notes-service/internal/bootstrap"
	"go-notes-service/internal/db"
	"log"

	"github.com/sirupsen/logrus"
)

const (
	serverAddr = ":8080"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// database := mustConnectDB()
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close(database)
	db.RunMigrations(database)

	fmt.Println("Server is listening on:localhost:", serverAddr)
	if err := bootstrap.NewApp(database).Listen(serverAddr); err != nil {
		log.Printf("server stopped with error: %v", err)
	}
}
