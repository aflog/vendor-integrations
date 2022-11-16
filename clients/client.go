package clients

import (
	"context"
	"time"
)

type ListTransactionParams struct {
	// If present, only return transaction after this time
	From *time.Time

	// If present, only return transaction before this time
	To *time.Time

	// Use for pagination
	Page     int
	PageSize int
}

type Client interface {
	// Return the balance available on the account
	GetBalance(ctx context.Context) (Money, error)

	// List transaction filtered by the parameters in params
	ListTransactions(ctx context.Context, params ListTransactionParams) ([]Transaction, error)

	// Fetch a transaction by its vendor ID or an error if not found
	GetTransactionByID(ctx context.Context, vendorID string) (Transaction, error)

	// Create a new trasaction with the given ID, recipient and money, the ID should be set on the vendor side if
	// possible. The function must return the vendor ID for this transaction and the Status of the transaction
	CreateTransaction(ctx context.Context, transactionID string, recipient Recipient, amount Money) (string, Status, error)
}
