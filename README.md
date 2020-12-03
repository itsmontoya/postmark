# Postmark
Postmark is a Client SDK for the Postmark API

## Usage
### New
```go
func ExampleNew() {
	testClient = New(testAPIKey)
}
```

### Client.Email
```go
func ExampleClient_Email() {
	var err error
	e := MakeEmail("from@email.com", "to@email.com", "test email", "<p>Hello world</p>")

	var resp EmailResponse
	if resp, err = testClient.Email(e); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response!", resp)
}
```
