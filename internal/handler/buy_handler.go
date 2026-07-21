package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udindwy/halogold-backend/internal/dto"
	"github.com/udindwy/halogold-backend/internal/service"
	"github.com/udindwy/halogold-backend/internal/validator"
	"github.com/udindwy/halogold-backend/pkg/logger"
	"github.com/udindwy/halogold-backend/pkg/response"
)

type BuyHandler struct {
	service service.TransactionService
}

func NewBuyHandler(s service.TransactionService) *BuyHandler {
	return &BuyHandler{
		service: s,
	}
}

func (h *BuyHandler) BuyGold(c *gin.Context) {
	logger.Info("Incoming request to POST /buy")

	var req dto.BuyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Validation error on /buy: %v", err)
		
		formattedErr := validator.FormatValidationError(err)
		response.Error(c, http.StatusBadRequest, formattedErr)
		return
	}

	userID := uint(1)

	transaction, err := h.service.BuyGold(c.Request.Context(), userID, req.Amount)
	if err != nil {
		logger.Error("Failed to process Buy transaction: %v", err)
		response.Error(c, http.StatusInternalServerError, "Failed to process transaction")
		return
	}

	logger.LogTransaction(transaction.Type, transaction.Amount, transaction.Gram)

	res := dto.BuyResponse{
		Gram:  transaction.Gram,
		Price: service.GoldPrice, 
	}
	response.Success(c, http.StatusOK, res)
}
