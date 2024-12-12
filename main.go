package main

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/di"
	"github.com/GuilhermeOCamargo/go-wallet-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	walletController := di.Initialize()

	api := r.Group("/api")
	routes.WalletRoutes(api, walletController)

	// Inicie o servidor
	port := ":8080"
	log.Printf("Iniciando servidor na porta %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %s", err)
	}
}
