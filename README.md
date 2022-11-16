# Clients


We work with multiple vendors who can pay transactions in multiple countries and currencies.
We want to write an abstraction that allows the rest of our system to create transactions without
having to worry about the actual execution of transactions and lets us easily add new vendors
without having to change the rest of the system.

There is a test server emulating multiple vendors with different behaviours that we will use 
the exercise. To simplify things they are all on the same IP address and port with different path
prefixes, but the clients shouldn't rely on this.

## Requirements

* Implement the `Client` interface for all the vendors
* All times must be in UTC
* Clients must to handle common types of failures consistently
* Implement the basic router to send transactions to the right clients


## Vendors

All vendors do transactions asynchronously 

You can use the following test data to test your clients (otherwise they all succeed by default)
* recipient name = "Test Failed" => status = failed
* recipient name = "Test Success" => status = success

### Vendor 1

This vendor is available for payments in `Tanzania`.

#### Get balance

Path: `/vendor1/api/balance`
Query args: None

```
$ curl -i http://127.0.0.1:8080/vendor1/api/balance
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 31 Oct 2022 08:05:16 GMT
Content-Length: 7

990000
```

#### Get transactions

Path: `/vendor1/api/transactions`
Query args: 
* `page`: int for the page number (optional)

The API returns up to 10 results per page.


```
$ curl -i http://127.0.0.1:8080/vendor1/api/transactions?page=3
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 31 Oct 2022 08:39:36 GMT
Content-Length: 1866

{
  "transactions": [
    {
      "amount": 100,
      "recipient_name": "ABC DEF",
      "recipient_phone_number": "1234567890",
      "status": "success",
      "created": "2022-08-12 11:39:29"
    },
    ...
    {
      "amount": 100,
      "recipient_name": "ABC DEF",
      "recipient_phone_number": "1234567890",
      "status": "success",
      "created": "2022-08-21 11:39:29"
    }
  ],
  "total_transactions": 101
}
```

#### Create transaction

Path: `/vendor1/api/transactions`
Query args: None
Body: 
- `amount`: int amount in TZS
- `recipient_name`: string
- `recipient_phone_number`: string in local format
- `external_reference`: string unique identifier for the transaction

```
$ curl -i -XPOST 127.0.0.1:8080/vendor1/api/transactions -d '{"amount": 3, "Recipient_name": "test name", "recipient_phone_number": "0766610550"}'
HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 31 Oct 2022 08:16:35 GMT
Content-Length: 16

{
  "id": 101
}
```


### Vendor 2

This vendor is available for payments in `Kenya`.

#### Get balance

Path: `/vendor2/api/balance`
Query args: None

```
$ curl -i http://127.0.0.1:8080/vendor2/api/balance
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 31 Oct 2022 10:30:46 GMT
Content-Length: 24

{
  "balance": 997900
}
```

#### List transactions

Path :`/vendor2/api/transfers`
Query args:
- `page`: int for the page number (optional)
- `from`: string datetime in RFC3339 format to exclude any transactions created before that time (optional)
- `to`: string datetime in RFC3339 format to exclude any transactions created after that time (optional)


```
$ curl -i "http://127.0.0.1:8080/vendor2/api/transfers?from=2022-10-23T01:02:33Z&to=2022-10-25T00:00:00Z&page=1"
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 31 Oct 2022 11:42:48 GMT
Content-Length: 554

{
  "transactions": [
    {
      "id": 14,
      "amount": 100,
      "recipient_name": "ABC DEF",
      "recipient_phone_number": "1234567890",
      "status": "success",
      "external_reference": "ref-13",
      "created": "2022-10-23T13:41:06.18439365+02:00"
    },
    {
      "id": 15,
      "amount": 100,
      "recipient_name": "ABC DEF",
      "recipient_phone_number": "1234567890",
      "status": "success",
      "external_reference": "ref-14",
      "created": "2022-10-24T13:41:06.18439388+02:00"
    }
  ],
  "total_transactions": 2
}
```

#### Create transaction

Path: `/vendor2/api/transfers`
Query args: None
Body:
- `amount`: float amount in KES
- `recipient_name`: string
- `recipient_phone_number`: string in international format
- `external_reference`: string unique identifier for the transaction

```
$ curl -i -XPOST 127.0.0.1:8080/vendor2/api/transfers -d '{"amount": 3, "Recipient_name": "test name", "recipient_phone_number": "+254766610550", "external_reference": "ddj"}'
HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 31 Oct 2022 11:48:07 GMT
Content-Length: 15

{
  "id": 22
}
```