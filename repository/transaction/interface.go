package transaction

import "backend-ewallet/entities"

type Transaction interface {
	Create(transaction entities.Transaction) (entities.Transaction, error)
	// CreateTopup(transaction entities.Transaction) (entities.Transaction, error)
	Get(userID string) ([]entities.Transaction, error)
	GetTransactionSend(userID string) ([]entities.Transaction, error)
	GetTransactionReceived(userID string) ([]entities.Transaction, error)
	GetByID(userID, transactionID string) (entities.Transaction, error)

	//jika diperlukan
	Search(q string) ([]entities.Transaction, error)
	// Dummy(length int) bool
}
