package usecase

import (
	"gestor_finanzas/domain"
	"time"
)

type TransactionRepository interface {
	Add(transaction domain.Transaction) error
}

type AddTransactionUseCase struct {
	repo TransactionRepository
}

func NewAddTransactionUseCase(repo TransactionRepository) *AddTransactionUseCase {
	return &AddTransactionUseCase{repo: repo}
}

func (uc *AddTransactionUseCase) Execute(description string, amount float64, transactionType domain.TransactionType) error {
	transaction := domain.Transaction{
		Description: description,
		Amount:      amount,
		Type:        transactionType,
		Date:        time.Now(),
	}
	return uc.repo.Add(transaction)
}
