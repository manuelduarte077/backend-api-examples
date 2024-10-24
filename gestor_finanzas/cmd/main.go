package main

import (
	"fmt"
	"gestor_finanzas/domain"
	"gestor_finanzas/infrastructure/db"
	"gestor_finanzas/interface/repository"
	"gestor_finanzas/usecase"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Inicializar la base de datos
	database, err := db.InitDB("finanzas.db")
	if err != nil {
		fmt.Println("Error al conectar con la base de datos:", err)
		os.Exit(1)
	}
	repo := repository.NewSQLiteRepository(database)

	// Comando 'add'
	var description string
	var amount float64
	var transactionType string

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Agregar una nueva transacción",
		Run: func(cmd *cobra.Command, args []string) {
			useCase := usecase.NewAddTransactionUseCase(repo)
			err := useCase.Execute(description, amount, domain.TransactionType(transactionType))
			if err != nil {
				fmt.Println("Error al agregar la transacción:", err)
			} else {
				fmt.Println("Transacción agregada exitosamente.")
			}
		},
	}
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Descripción de la transacción")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Monto de la transacción")
	addCmd.Flags().StringVarP(&transactionType, "type", "t", "", "Tipo de transacción (ingreso o gasto)")

	// Comando 'list'
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Listar todas las transacciones",
		Run: func(cmd *cobra.Command, args []string) {
			transactions, err := repo.List()
			if err != nil {
				fmt.Println("Error al listar las transacciones:", err)
			} else {
				for _, t := range transactions {
					fmt.Printf("%d: %s - %f - %s - %s\n", t.ID, t.Description, t.Amount, t.Type, t.Date)
				}
			}
		},
	}

	// Crear raíz de la CLI
	var rootCmd = &cobra.Command{Use: "finanzas"}
	rootCmd.AddCommand(addCmd, listCmd)
	rootCmd.Execute()
}
