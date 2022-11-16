package vendor1

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aflog/vendor-integrations/clients"
	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
)

var fakeResponse = `{
	"transactions": [
	  {
		"id":1,
		"amount": 100,
		"recipient_name": "ABC DEF",
		"recipient_phone_number": "1234567890",
		"status": "success",
		"created": "2022-08-12 11:39:29"
	  },
	  {
		"id": 2,
		"amount": 100,
		"recipient_name": "ABC DEF",
		"recipient_phone_number": "1234567890",
		"status": "success",
		"created": "2022-08-21 11:39:29"
	  }
	],
	"total_transactions": 101
  }`

func TestListTransaction(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fakeResponse)
	}))

	defer testServer.Close()

	vendor := New(testServer.URL)

	// vendor := New("https://f3a1-91-126-210-10.ngrok.io") // not available ATM
	results, err := vendor.ListTransactions(context.Background(), clients.ListTransactionParams{})
	if err != nil {
		t.Errorf("Received error: %s", err)
	}

	expected := []clients.Transaction{
		{ID: "1", Recipient: clients.Recipient{FirstName: "ABC", LastName: "DEF", PhoneNumber: "1234567890"}, Money: clients.Money{Amount: decimal.NewFromInt(int64(100))}, Status: clients.StatusSuccess},
		{ID: "2", Recipient: clients.Recipient{FirstName: "ABC", LastName: "DEF", PhoneNumber: "1234567890"}, Money: clients.Money{Amount: decimal.NewFromInt(int64(100))}, Status: clients.StatusSuccess},
	}

	if diff := cmp.Diff(expected, results); diff != "" {
		t.Errorf("Result not as expected (- missing from result) (+ added to result):\n %s", diff)
	}

}
