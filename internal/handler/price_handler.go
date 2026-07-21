package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udindwy/halogold-backend/internal/dto"
	"github.com/udindwy/halogold-backend/internal/service"
	"github.com/udindwy/halogold-backend/pkg/logger"
	"github.com/udindwy/halogold-backend/pkg/response"
)

type PriceHandler struct {
	service service.TransactionService
}

func NewPriceHandler(s service.TransactionService) *PriceHandler {
	return &PriceHandler{
		service: s,
	}
}

func (h *PriceHandler) GetPrice(c *gin.Context) {
	logger.Info("Incoming request to GET /price")

	price, err := h.service.GetPrice(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get gold price: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res := dto.PriceResponse{
		Price: price,
	}

	logger.Info("Successfully returned gold price")
	response.Success(c, http.StatusOK, res)
}
