package main

import (
	"context"
	"log"

	"github.com/di3gor35/ApiRest/config"
	"github.com/di3gor35/ApiRest/pkg/database"
	//"github.com/labstack/echo/v4"
)

func main(){
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando config %v", err)
		return
	}

	client, err := database.NewMongoDBConnection(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Error al conectar con MongoDB %v", err)
		return
	}
	defer client.Disconnect(context.Background())
	log.Println("Conexión exitosa con MongoDB")
	client.Database(cfg.MongoDBName)
}