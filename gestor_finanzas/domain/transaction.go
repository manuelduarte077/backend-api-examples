package domain

import "time"

// Tipo de transacción: ingreso o gasto
type TransactionType string

const (
	Income  TransactionType = "ingreso"
	Expense TransactionType = "gasto"
)

// Transaction representa una transacción de gasto o ingreso
type Transaction struct {
	ID          int64
	Description string
	Amount      float64
	Type        TransactionType
	Date        time.Time
}
