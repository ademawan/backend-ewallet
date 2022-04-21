package transaction

type TransactionCreateResponse struct {
	TransactionID   string `json:"transaction_id"`
	SenderID        string `json:"sender_id"`
	RecipientID     string `json:"recipient_id"`
	RecievedAmount  int    `json:"recieved_amound"`
	SentAmount      int    `json:"sent_amount"`
	TransactionType string `json:"transaction_type"`
}

type TransactionUpdateResponse struct {
	TransactionID   string `json:"transaction_id"`
	SenderID        string `json:"sender_id"`
	RecipientID     string `json:"recipient_id"`
	RecievedAmount  int    `json:"recieved_amound"`
	SentAmount      int    `json:"sent_amount"`
	TransactionType string `json:"transaction_type"`
}
type TransactionGetByIdResponse struct {
	TransactionID   string `json:"transaction_id"`
	SenderID        string `json:"sender_id"`
	RecipientID     string `json:"recipient_id"`
	RecievedAmount  int    `json:"recieved_amound"`
	SentAmount      int    `json:"sent_amount"`
	TransactionType string `json:"transaction_type"`
}

//=========================================================

// =================== Create User Request =======================
type CreateTransactionRequestFormat struct {
	SenderID        string
	RecipientID     string `json:"recipient_id" form:"recipient_id"`
	RecievedAmount  int    `json:"recieved_amound" form:"received_amount"`
	SentAmount      int    `json:"sent_amount" form:"sent_amount"`
	TransactionType string `json:"transaction_type"`
}

// =================== Update User Request =======================
type UpdateTransactionRequestFormat struct {
	SenderID        string
	RecipientID     string `json:"recipient_id" form:"recipient_id"`
	RecievedAmount  int    `json:"recieved_amound" form:"received_amount"`
	SentAmount      int    `json:"sent_amount" form:"sent_amount"`
	TransactionType string `json:"transaction_type"`
}
