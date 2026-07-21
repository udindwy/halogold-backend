package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/udindwy/halogold-backend/internal/handler"
)

func SetupRoutes(r *gin.Engine, priceHandler *handler.PriceHandler, buyHandler *handler.BuyHandler, sellHandler *handler.SellHandler, transactionHandler *handler.TransactionHandler) {
	r.GET("/price", priceHandler.GetPrice)
	r.POST("/buy", buyHandler.BuyGold)
	r.POST("/sell", sellHandler.SellGold)
	r.GET("/transactions", transactionHandler.GetTransactions)
}
