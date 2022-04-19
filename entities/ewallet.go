package entities

type Ewallet struct {
	EwalletID string `gorm:"index;unique;type:varchar(22)" json:"ewallet_id"`
	UserID    string
	Saldo     int
}
