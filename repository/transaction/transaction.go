package transaction

import (
	"backend-ewallet/entities"
	"errors"
	"fmt"

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
		err := ur.database.Transaction(func(tx *gorm.DB) error {
			var user entities.User
			result := tx.Model(entities.User{}).Where("user_id=?", transaction.Sender).First(&user)
			if err := result.Error; err != nil {
				return err
			}
			if user.Saldo < transaction.Amount {
				return errors.New("saldo tidak cukup")
			} else {
				if err := tx.Debug().Model(entities.User{}).Where("user_id =?", transaction.Sender).UpdateColumn("saldo", gorm.Expr("saldo - ?", transaction.Amount)).Error; err != nil {
					return err
				}

				if err := tx.Debug().Model(entities.User{}).Where("user_id =?", transaction.Recipient).UpdateColumn("saldo", gorm.Expr("saldo + ?", transaction.Amount)).Error; err != nil {
					return errors.New("transfer gagal")
				}
				sql := fmt.Sprintf("INSERT INTO transactions ('transaction_id','sender','recipient','transaction_type','amount','created_at','deleted_at') VALUES ('%s','%s','%s','transfer',%d,'%s',NULL)", transaction.TransactionID, transaction.Sender, transaction.Recipient, transaction.Amount, transaction.CreatedAt)

				if err := tx.Raw(sql).Scan(&transaction); err != nil {
					return errors.New("transfer failed")
				}
			}
			return nil
		})
		if err != nil {
			return entities.Transaction{}, err
		}
	} else if transaction.TransactionType == "topup" {
		err := ur.database.Transaction(func(tx *gorm.DB) error {
			var user entities.User
			result := tx.Debug().Model(entities.User{}).Where("user_id=?", transaction.Recipient).First(&user)
			if err := result.Error; err != nil {
				return err
			}

			if err := tx.Debug().Model(entities.User{}).Where("user_id =?", transaction.Recipient).UpdateColumn("saldo", gorm.Expr("saldo + ?", transaction.Amount)).Error; err != nil {
				return errors.New("transfer failed")
			}

			return nil
		})
		if err != nil {
			return entities.Transaction{}, err
		}

	} else {
		return entities.Transaction{}, errors.New("payment type not allowed")

	}

	return transaction, nil
}

func (ur *TransactionRepository) TransferAmount(RecipientID string, amount uint) error {

	if err := ur.database.Model(entities.User{}).Where("user_id =?", RecipientID).UpdateColumn("saldo", gorm.Expr("saldo + ?", amount)).Error; err != nil {
		return errors.New("transfer failed")
	}

	return nil
}

func (ur *TransactionRepository) Get(userID string) ([]entities.Transaction, error) {
	arrTransaction := []entities.Transaction{}

	result := ur.database.Where("sender_id =? OR recipient_id =?", userID, userID).Find(&arrTransaction)
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}
	if err := result.Error; err != nil {
		return arrTransaction, err
	}

	return arrTransaction, nil
}
func (ur *TransactionRepository) GetTransactionSend(userID string) ([]entities.Transaction, error) {
	arrTransaction := []entities.Transaction{}

	result := ur.database.Where("sender_id =?", userID).Find(&arrTransaction)
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}
	if err := result.Error; err != nil {
		return arrTransaction, err
	}

	return arrTransaction, nil
}
func (ur *TransactionRepository) GetTransactionReceived(userID string) ([]entities.Transaction, error) {
	arrTransaction := []entities.Transaction{}

	result := ur.database.Where("recipient_id =?", userID).Find(&arrTransaction)
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}
	if err := result.Error; err != nil {
		return arrTransaction, err
	}

	return arrTransaction, nil
}

func (ur *TransactionRepository) GetByID(senderID, transactionID string) (entities.Transaction, error) {
	arrTransaction := entities.Transaction{}

	result := ur.database.Where("transaction_id =? AND sender_id=?", transactionID, senderID).First(&arrTransaction)
	if result.RowsAffected == 0 {
		return arrTransaction, errors.New("record not found")
	}
	if err := result.Error; err != nil {
		return arrTransaction, err
	}

	return arrTransaction, nil
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
