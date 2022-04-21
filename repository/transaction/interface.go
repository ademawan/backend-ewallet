package transaction

import "backend-ewallet/entities"

type Transaction interface {
	Create(transaction entities.Transaction) (entities.Transaction, error)
	Get(userID string) ([]entities.Transaction, error)
	GetByID(userID, transactionID string) (entities.Transaction, error)

	Update(transactionID string, newTransaction entities.Transaction) (entities.Transaction, error)
	Delete(userID, transactionID string) error
	//jika diperlukan
	Search(q string) ([]entities.Transaction, error)
	// Dummy(length int) bool
}
