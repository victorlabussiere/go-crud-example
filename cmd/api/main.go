package main

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/database"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Falha ao iniciar o banco de dados: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
