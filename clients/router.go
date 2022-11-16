package clients

import "context"

type CreateTransactionResult struct {
	TransactionID string
	VendorName    string
	VendorID      string
	Status        Status
}

type TransactionRouter interface {
	CreateTransaction(ctx context.Context, transactionID string, recipient Recipient, amount Money) (CreateTransactionResult, Status, error)
}

type Router struct {
}

func NewRouter(clients []Client) *Router {
	return &Router{}
}

func (r *Router) CreateTransaction(ctx context.Context, transactionID string, recipient Recipient, amount Money) (CreateTransactionResult, error) {
	return CreateTransactionResult{
		TransactionID: transactionID,
		Status:        StatusFailed,
	}, nil
}
