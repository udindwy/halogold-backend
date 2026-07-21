package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udindwy/halogold-backend/internal/dto"
	"github.com/udindwy/halogold-backend/internal/service"
	"github.com/udindwy/halogold-backend/pkg/logger"
	"github.com/udindwy/halogold-backend/pkg/response"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: s,
	}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	logger.Info("Incoming request to GET /transactions")

	transactions, err := h.service.GetTransactions(c.Request.Context())
	if err != nil {
		logger.Error("Failed to fetch transactions: %v", err)
		response.Error(c, http.StatusInternalServerError, "Failed to fetch transactions")
		return
	}

	var res []dto.TransactionResponse
	for _, t := range transactions {
		res = append(res, dto.TransactionResponse{
			ID:        t.ID,
			UserID:    t.UserID,
			Type:      t.Type,
			Amount:    t.Amount,
			Gram:      t.Gram,
			CreatedAt: t.CreatedAt,
		})
	}

	if res == nil {
		res = []dto.TransactionResponse{}
	}

	response.Success(c, http.StatusOK, res)
}
