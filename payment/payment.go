package payment


import  "github.com/rpip/paystack-go"

  apiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"

// second param is an optional http client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
client := paystack.NewClient(apiKey)

recipient := &TransferRecipient{
Type:          "Nuban",
Name:          "Customer 1",
Description:   "Demo customer",
AccountNumber: "0100000010",
BankCode:      "044",
Currency:      "NGN",
Metadata:      map[string]interface{}{"job": "tour operator"},
}

recipient1, err := client.Transfer.CreateRecipient(recipient)

req := &TransferRequest{
Source:    "balance",
Reason:    "to visit torist center",
Amount:    30,
Recipient: recipient1.RecipientCode,
}

transfer, err := client.Transfer.Initiate(req)
if err != nil {
// do something with error
}

// retrieve list of plans
plans, err := client.Plan.List()

for i, plan := range plans.Values {
fmt.Printf("%+v", plan)
}

cust := &Customer{
FirstName: "User123",
LastName:  "AdminUser",
Email:     "user123@gmail.com",
Phone:     "+23400000000000000",
}
// create the customer
customer, err := client.Customer.Create(cust)
if err != nil {
// do something with error
}

// Get customer by ID
customer, err := client.Customers.Get(customer.ID)