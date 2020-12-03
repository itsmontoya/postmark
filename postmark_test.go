package postmark

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	testAPIKey    = os.Getenv("POSTMARK_API_KEY")
	testToEmail   = os.Getenv("POSTMARK_TEST_TO_EMAIL")
	testFromEmail = os.Getenv("POSTMARK_TEST_FROM_EMAIL")
	testClient    *Client
)

func TestNew(t *testing.T) {
	c := New(testAPIKey)
	if c.apiKey != testAPIKey {
		t.Fatalf("invalid api key, expected <%s> and received <%s>", testAPIKey, c.apiKey)
	}
}

func TestClient_Email(t *testing.T) {
	var err error
	c := New(testAPIKey)
	e := MakeEmail(testToEmail, testFromEmail, "test email", "<p>Hello world</p>")

	var resp EmailResponse
	if resp, err = c.Email(e); err != nil {
		t.Fatal(err)
	}

	if resp.To != e.To {
		t.Fatalf("invalid \"to\" address, expected <%s> and received <%s>", e.To, resp.To)
	}
}

func ExampleNew() {
	testClient = New(testAPIKey)
}

func ExampleClient_Email() {
	var err error
	e := MakeEmail("from@email.com", "to@email.com", "test email", "<p>Hello world</p>")

	var resp EmailResponse
	if resp, err = testClient.Email(e); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response!", resp)
}
