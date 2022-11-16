package clients

import (
	"github.com/shopspring/decimal"
)

type Status string

const (
	StatusSuccess = "success"
	StatusFailed  = "failed"
	StatusPending = "pending"
	StatusUnknown = "unknown"
)

type Recipient struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Country     string
}

type Money struct {
	Amount   decimal.Decimal
	Currency string
}

type Transaction struct {
	ID        string
	Recipient Recipient
	Money     Money
	Status    Status
}
