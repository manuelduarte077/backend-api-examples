package repository

import (
	"database/sql"
	"gestor_finanzas/domain"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

// Implementación para agregar una transacción
func (r *SQLiteRepository) Add(transaction domain.Transaction) error {
	query := `INSERT INTO transactions(description, amount, type, date) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, transaction.Description, transaction.Amount, transaction.Type, transaction.Date)
	return err
}

// Implementación para listar transacciones
func (r *SQLiteRepository) List() ([]domain.Transaction, error) {
	rows, err := r.db.Query(`SELECT id, description, amount, type, date FROM transactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Description, &transaction.Amount, &transaction.Type, &transaction.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
