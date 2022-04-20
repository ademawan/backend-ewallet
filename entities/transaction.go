package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionID   string `gorm:"index;unique;type:varchar(22)" json:"transaction_id"`
	SenderID        string
	RecipientID     string
	RecievedAmount  int    `gorm:"type:int(30)" json:"recieved_amount"`
	SentAmount      int    `gorm:"type:int(30)" json:"sent_amount"`
	TransactionType string `gorm:"type:varchar(30)" json:"transaction_type"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `gorm:"-" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
