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

type SellHandler struct {
	service service.TransactionService
}

func NewSellHandler(s service.TransactionService) *SellHandler {
	return &SellHandler{
		service: s,
	}
}

func (h *SellHandler) SellGold(c *gin.Context) {
	logger.Info("Incoming request to POST /sell")

	var req dto.SellRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Validation error on /sell: %v", err)
		
		formattedErr := validator.FormatValidationError(err)
		response.Error(c, http.StatusBadRequest, formattedErr)
		return
	}

	userID := uint(1)

	transaction, err := h.service.SellGold(c.Request.Context(), userID, req.Gram)
	if err != nil {
		logger.Error("Failed to process Sell transaction: %v", err)
		response.Error(c, http.StatusInternalServerError, "Failed to process transaction")
		return
	}

	logger.LogTransaction(transaction.Type, transaction.Amount, transaction.Gram)

	res := dto.SellResponse{
		Amount: transaction.Amount,
	}
	response.Success(c, http.StatusOK, res)
}
