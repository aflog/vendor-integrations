package vendor1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aflog/vendor-integrations/clients"
	"github.com/aflog/vendor-integrations/clients/vendors"
	"github.com/shopspring/decimal"
)

type Vendor1 struct {
	url string
}

func New(url string) *Vendor1 {
	return &Vendor1{url: url}
}

type TransactionsPage struct {
	Transactions      []Transaction
	TotalTransactions int
}

type Transaction struct {
	Amount               float64 `json:"amount"`
	RecipientName        string  `json:"recipient_name"`
	RecipientPhoneNumber string  `json:"recipient_phone_number"`
	Status               string  `json:"status"`
	Created              string  `json:"created"`
	ID                   int     `json:"id"`
}

func (v *Vendor1) GetBalance(ctx context.Context) (clients.Money, error) {
	return clients.Money{}, &vendors.Error{ErrType: vendors.ErrTypeUnimplemented}
}

// List transaction filtered by the parameters in params
func (v *Vendor1) ListTransactions(ctx context.Context, params clients.ListTransactionParams) ([]clients.Transaction, error) {
	var transactions []clients.Transaction

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(fmt.Sprintf("%s/vendor1/api/transactions", v.url))
	if err != nil {
		return transactions, &vendors.Error{ErrType: vendors.ErrorTypeConnection, Msg: err.Error()}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return transactions, err
	}

	page := &TransactionsPage{}
	if err = json.Unmarshal(body, page); err != nil {
		return transactions, &vendors.Error{ErrType: vendors.ErrTypeResponseFormat, Msg: err.Error()}
	}

	for _, t := range page.Transactions {
		transactions = append(transactions, parseToGeneric(t))
	}

	return transactions, nil
}

func (v *Vendor1) GetTransactionByID(ctx context.Context, vendorID string) (clients.Transaction, error) {
	return clients.Transaction{}, &vendors.Error{ErrType: vendors.ErrTypeUnimplemented}
}

func (v *Vendor1) CreateTransaction(ctx context.Context, transactionID string, recipient clients.Recipient, amount clients.Money) (string, clients.Status, error) {
	return "", clients.StatusSuccess, &vendors.Error{ErrType: vendors.ErrTypeUnimplemented}
}

func parseToGeneric(t Transaction) clients.Transaction {
	nameParts := strings.SplitN(t.RecipientName, " ", 2)

	recipient := clients.Recipient{
		FirstName:   nameParts[0],
		PhoneNumber: t.RecipientPhoneNumber,
	}

	if len(nameParts) > 1 {
		recipient.LastName = nameParts[1]
	}
	return clients.Transaction{ID: fmt.Sprintf("%d", t.ID), Recipient: recipient, Money: clients.Money{Amount: decimal.NewFromFloat(t.Amount)}, Status: parseStatus(t.Status)}
}

func parseStatus(s string) clients.Status {
	switch s {
	case "success":
		return clients.StatusSuccess
	default:
		return clients.StatusUnknown

	}
}
