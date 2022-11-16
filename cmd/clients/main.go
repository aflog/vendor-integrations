package main

import (
	"context"
	"fmt"

	"github.com/aflog/vendor-integrations/clients"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	router := clients.NewRouter([]clients.Client{})

	transactions := []clients.Transaction{
		{
			ID: "1",
			Recipient: clients.Recipient{
				FirstName:   "John",
				LastName:    "Doe",
				PhoneNumber: "+255766610550",
				Country:     "TZ",
			},
			Money: clients.Money{
				Amount:   decimal.NewFromFloat(3500),
				Currency: "TZS",
			},
		},
		{
			ID: "2",
			Recipient: clients.Recipient{
				FirstName:   "Jane",
				LastName:    "Doe",
				PhoneNumber: "+254723123456",
				Country:     "KE",
			},
			Money: clients.Money{
				Amount:   decimal.NewFromFloat(230),
				Currency: "KES",
			},
		},
	}

	for _, transaction := range transactions {
		response, err := router.CreateTransaction(ctx, transaction.ID, transaction.Recipient, transaction.Money)
		if err != nil {
			logrus.WithField("transaction_id", transaction.ID).WithError(err).Error("failed to create transaction")
			continue
		}
		fmt.Println(response)
	}
}
