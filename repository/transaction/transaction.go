package transaction

import (
	"backend-ewallet/entities"
	"errors"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		database: db,
	}
}

func (ur *TransactionRepository) Create(transaction entities.Transaction) (entities.Transaction, error) {

	uid := shortuuid.New()
	transaction.TransactionID = uid

	if transaction.TransactionType == "transfer" {
		var user entities.User
		result := ur.database.Model(entities.User{}).Where("user_id=?", transaction.SenderID).First(&user)
		if err := result.Error; err != nil {
			return entities.Transaction{}, err
		}
		if user.Saldo < transaction.Amount {
			return entities.Transaction{}, errors.New("saldo tidak cukup")
		} else {
			ur.database.Model(entities.User{}).Where("user_id =?", transaction.SenderID).UpdateColumn("saldo", gorm.Expr("saldo - ?", transaction.Amount))

			res := ur.TransferAmount(transaction.RecipientID, transaction.Amount)
			if res != nil {
				return entities.Transaction{}, res
			}

			if err := ur.database.Create(&transaction).Error; err != nil {
				return transaction, errors.New("transaction failed")
			}
			return transaction, nil
		}
	}

	return entities.Transaction{}, errors.New("tidak dapat di prosses")
}

func (ur *TransactionRepository) TransferAmount(RecipientID string, amount uint) error {

	if err := ur.database.Model(entities.User{}).Where("user_id =?", RecipientID).UpdateColumn("saldo", gorm.Expr("saldo + ?", amount)).Error; err != nil {
		return errors.New("failed")
	}

	return nil
}

func (ur *TransactionRepository) Get(userID string) ([]entities.Transaction, error) {
	arrTransaction := []entities.Transaction{}

	result := ur.database.Where("sender_id =?", userID).First(&arrTransaction)
	if err := result.Error; err != nil {
		return arrTransaction, err
	}
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}

	return arrTransaction, nil
}

func (ur *TransactionRepository) GetByID(senderID, transactionID string) (entities.Transaction, error) {
	arrTransaction := entities.Transaction{}

	result := ur.database.Where("transaction_id =? AND sender_id=?", transactionID, senderID).First(&arrTransaction)
	if err := result.Error; err != nil {
		return arrTransaction, err
	}
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}

	return arrTransaction, nil
}

func (ur *TransactionRepository) Update(transactionID string, newTransaction entities.Transaction) (entities.Transaction, error) {
	userID := newTransaction.SenderID
	newTransaction.SenderID = ""
	var transaction entities.Transaction
	result := ur.database.Where("transaction_id =? AND sender_id =?", transactionID, userID).First(&transaction)

	if result.Error != nil {
		return entities.Transaction{}, errors.New("failed to update transaction")
	}
	if result.RowsAffected == 0 {
		return entities.Transaction{}, errors.New("transaction not found")
	}

	if err := ur.database.Model(&transaction).Where("transaction_id =? AND sender_id =?", transactionID, userID).Updates(&newTransaction).Error; err != nil {
		return entities.Transaction{}, err
	}

	return transaction, nil
}

func (ur *TransactionRepository) Delete(userID, transactionID string) error {

	result := ur.database.Where("transaction_id =? AND sender_id =?", transactionID, userID).Delete(&entities.Transaction{})
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (ur *TransactionRepository) Search(q string) ([]entities.Transaction, error) {
	arrTransaction := []entities.Transaction{}

	if len(q) < 3 {
		if len(q) == 1 {
			ur.database.Debug().Where("a =?", q).Find(&arrTransaction)
			return arrTransaction, nil
		}
		if len(q) == 2 {
			ur.database.Debug().Where("b =?", q).Find(&arrTransaction)
			return arrTransaction, nil

		}
		if len(q) == 3 {
			ur.database.Debug().Where("c =?", q).Find(&arrTransaction)
			return arrTransaction, nil

		}
	}
	sql := "%" + q + "%"

	result := ur.database.Debug().Where("name like ?", sql).Find(&arrTransaction)
	if err := result.Error; err != nil {
		return arrTransaction, err
	}
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}

	return arrTransaction, nil
}
