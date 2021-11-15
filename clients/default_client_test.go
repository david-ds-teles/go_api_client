package clients

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/david-ds-teles/account_client/models"
	"github.com/google/uuid"
)

const apiEndpointEnvVar = "API_ENDPOINT"

func getTestClient(t *testing.T) Client {
	apiEndpoint := os.Getenv(apiEndpointEnvVar)

	if apiEndpoint == "" {
		panic("API_ENDPOINT env not configured")
	}

	client, err := NewClient(log.Default(), http.DefaultClient, apiEndpoint)

	if err != nil {
		panic(fmt.Errorf("client instance expected, but got %v", err))
	}

	return client
}

func accountTestData() *models.AccountData {
	country := "GB"
	account := &models.AccountData{
		Attributes: &models.AccountAttributes{
			Name:          []string{"David Teles"},
			BankID:        "400300",
			BankIDCode:    "GBDSC",
			Bic:           "NWBKGB22",
			Country:       &country,
			BaseCurrency:  "GBP",
			Iban:          "GB11NWBK40030041426819",
			AccountNumber: "41426819",
		},
	}

	return account
}

func TestCreate(t *testing.T) {
	client := getTestClient(t)

	account := accountTestData()

	result, err := client.Create(account)

	if err != nil {
		t.Fatalf("create failed, got error: %v", err)
	}

	if result.ID == "" {
		t.Fatalf("expect valid ID, got: %s", result.ID)
	}
}

func TestCreateInvalidParams(t *testing.T) {
	client := getTestClient(t)

	result, err := client.Create(nil)

	if err == nil {
		t.Fatalf("expected error but got: %v", err)
	}

	if result != nil {
		t.Fatalf("expected result nil but got: %v", result)
	}

	expect := "invalid account, nil"

	if err.Error() != expect {
		t.Fatalf("expected %s but got: %v", expect, err)
	}
}

func TestCreateInvalidAccount(t *testing.T) {
	client := getTestClient(t)

	result, err := client.Create(&models.AccountData{})

	if err == nil {
		t.Fatalf("expected error but got: %v", err)
	}

	if result != nil {
		t.Fatalf("expected result nil but got: %v", result)
	}

	expected := "httpStatus: 400"
	want := regexp.MustCompile(`\b` + expected + `\b`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`expected: %s, got: %v`, expected, err)
	}
}

func TestFetch(t *testing.T) {
	client := getTestClient(t)

	created, err := client.Create(accountTestData())

	if err != nil {
		t.Fatalf("failed created before fetch: %v", err)
	}

	result, err := client.Fetch(created.ID)

	if err != nil {
		t.Fatalf("fetch failed, got error: %v", err)
	}

	if result.ID != created.ID {
		t.Fatalf("invalid ID, expect %s  got: %s", created.ID, result.ID)
	}
}

func TestFetchInvalidParam(t *testing.T) {
	client := getTestClient(t)

	result, err := client.Fetch("")

	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}

	if result != nil {
		t.Fatalf("expected nil, got: %v", result)
	}

	expect := "invalid account id"

	if err.Error() != expect {
		t.Fatalf("expected: %s, got: %s", expect, err.Error())
	}
}

func TestFetchNotFound(t *testing.T) {
	client := getTestClient(t)

	result, err := client.Fetch(uuid.New().String())

	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}

	if result != nil {
		t.Fatalf("expected nil, got: %v", result)
	}

	expected := "httpStatus: 404"
	want := regexp.MustCompile(`\b` + expected + `\b`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`expected: %s, got: %v`, expected, err)
	}
}

func TestDelete(t *testing.T) {
	client := getTestClient(t)

	created, err := client.Create(accountTestData())

	if err != nil {
		t.Fatalf("failed created before fetch: %v", err)
	}

	result, err := client.Fetch(created.ID)

	if err != nil {
		t.Fatalf("fetch failed, got error: %v", err)
	}

	err = client.Delete(result.ID, result.Version)

	if err != nil {
		t.Fatalf("failed to delete, expected error nil, got: %v", err)
	}

}

func TestDeleteInvalidParams(t *testing.T) {
	client := getTestClient(t)

	version := int64(0)
	err := client.Delete("", &version)

	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}

	expect := "invalid params, id:  , version: 0"

	if err.Error() != expect {
		t.Fatalf("expected %s , got: %s", expect, err.Error())
	}

}

func TestDeleteNotFound(t *testing.T) {
	client := getTestClient(t)

	version := int64(0)
	err := client.Delete(uuid.New().String(), &version)

	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}

	expected := "httpStatus: 404"
	want := regexp.MustCompile(`\b` + expected + `\b`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`expected: %s, got: %v`, expected, err)
	}

}
