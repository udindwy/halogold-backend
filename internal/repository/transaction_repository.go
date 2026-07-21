package repository

import (
	"context"

	"github.com/udindwy/halogold-backend/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *model.Transaction) error
	FindAll(ctx context.Context) ([]model.Transaction, error)
	FindByID(ctx context.Context, id uint) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *model.Transaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *transactionRepository) FindAll(ctx context.Context) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindByID(ctx context.Context, id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&transaction).Error
	
	if err != nil {
		return nil, err
	}
	
	return &transaction, nil
}
