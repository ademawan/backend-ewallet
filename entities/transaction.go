package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionID   string `gorm:"index;unique;type:varchar(22)" json:"transaction_id"`
	SenderID        string
	RecipientID     string
	TransactionType string `gorm:"type:varchar(30)" json:"transaction_type"`
	Amount          uint   `gorm:"type:int(30)" json:"amount"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `gorm:"-" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

//RecievedAmount  uint   `gorm:"type:int(30)" json:"recieved_amount"`
//	SentAmount      uint   `gorm:"type:int(30)" json:"sent_amount"`
