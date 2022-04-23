package transaction

type TransactionCreateResponse struct {
	TransactionID   string `json:"transaction_id"`
	SenderID        string `json:"sender_id"`
	RecipientID     string `json:"recipient_id"`
	Amount          uint   `json:"amount"`
	TransactionType string `json:"transaction_type"`
}

type TransactionGetByIdResponse struct {
	TransactionID   string `json:"transaction_id"`
	SenderID        string `json:"sender_id"`
	RecipientID     string `json:"recipient_id"`
	Amount          uint   `json:"amount"`
	TransactionType string `json:"transaction_type"`
}

//=========================================================

// =================== Create Transaction Request =======================
type CreateTransactionRequestFormat struct {
	SenderID        string
	RecipientID     string `json:"recipient_id" form:"recipient_id"`
	Amount          uint   `json:"amount" form:"amount"`
	TransactionType string `json:"transaction_type"`
}
