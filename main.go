package main

import (
	"log"
	"net/http"
	"os"

	"github.com/david-ds-teles/account_client/clients"
	"github.com/david-ds-teles/account_client/models"
)

const localApiEndpoint = "http://localhost:8080"

func main() {
	log := log.New(os.Stdout, "client: ", log.Default().Flags())

	client, err := clients.NewClient(log, http.DefaultClient, localApiEndpoint)

	if err != nil {
		log.Fatalf("error trying initialize a new client. %v", err)
	}

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

	createdAcc, err := client.Create(account)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fetchedAcc, err := client.Fetch(createdAcc.ID)

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = client.Delete(fetchedAcc.ID, fetchedAcc.Version)

	if err != nil {
		log.Fatalf(err.Error())
	}
}
