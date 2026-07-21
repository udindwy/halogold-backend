package service

import (
	"context"
	"errors"

	"github.com/udindwy/halogold-backend/internal/model"
	"github.com/udindwy/halogold-backend/internal/repository"
)

const GoldPrice = 1945200.0

type TransactionService interface {
	GetPrice(ctx context.Context) (float64, error)
	BuyGold(ctx context.Context, userID uint, amount float64) (*model.Transaction, error)
	SellGold(ctx context.Context, userID uint, gram float64) (*model.Transaction, error)
	GetTransactions(ctx context.Context) ([]model.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

func (s *transactionService) GetPrice(ctx context.Context) (float64, error) {
	return GoldPrice, nil
}

func (s *transactionService) BuyGold(ctx context.Context, userID uint, amount float64) (*model.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	gram := amount / GoldPrice

	transaction := &model.Transaction{
		UserID: userID,
		Type:   "BUY",
		Amount: amount,
		Gram:   gram,
	}

	err := s.repo.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) SellGold(ctx context.Context, userID uint, gram float64) (*model.Transaction, error) {
	if gram <= 0 {
		return nil, errors.New("gram must be greater than zero")
	}

	amount := gram * GoldPrice

	transaction := &model.Transaction{
		UserID: userID,
		Type:   "SELL",
		Amount: amount,
		Gram:   gram,
	}

	err := s.repo.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	return s.repo.FindAll(ctx)
}
