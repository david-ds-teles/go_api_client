package clients

import (
	"log"

	"github.com/david-ds-teles/account_client/models"
	"github.com/google/uuid"
)

// fakeClient created to tests and to development
type fakeClient struct {
	log *log.Logger
}

func newFakeClient(log *log.Logger) fakeClient {
	return fakeClient{log}
}

func (client fakeClient) Create(acc *models.AccountData) (*models.AccountData, error) {
	client.log.Printf("FakeClient.Create: %v", acc)

	acc.ID = uuid.New().String()

	client.log.Printf("FakeClient.Create finished")

	return acc, nil
}

func (client fakeClient) Fetch(id string) (*models.AccountData, error) {
	client.log.Printf("FakeClient.Fetch: %v", id)

	acc := &models.AccountData{}
	acc.ID = id

	client.log.Printf("FakeClient.Fetch finished")
	return acc, nil
}

func (client fakeClient) Delete(id string, version *int64) error {
	client.log.Printf("FakeClient.Delete: %s, %d", id, version)
	client.log.Printf("FakeClient.Delete finished")
	return nil
}
