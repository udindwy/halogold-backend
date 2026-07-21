package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/udindwy/halogold-backend/config"
	"github.com/udindwy/halogold-backend/internal/handler"
	"github.com/udindwy/halogold-backend/internal/repository"
	"github.com/udindwy/halogold-backend/internal/service"
	"github.com/udindwy/halogold-backend/pkg/logger"
	"github.com/udindwy/halogold-backend/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}

	config.ConnectDatabase(cfg)
	defer config.CloseDatabase()

	logger.DBConnected()

	txRepo := repository.NewTransactionRepository(config.DB)
	txService := service.NewTransactionService(txRepo)
	
	priceHandler := handler.NewPriceHandler(txService)
	buyHandler := handler.NewBuyHandler(txService)
	sellHandler := handler.NewSellHandler(txService)
	transactionHandler := handler.NewTransactionHandler(txService)

	r := gin.Default()
	routes.SetupRoutes(r, priceHandler, buyHandler, sellHandler, transactionHandler)

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	logger.ServerStart(port)
	if err := r.Run(":" + port); err != nil {
		logger.Error("Server gagal berjalan: %v", err)
	}
}
