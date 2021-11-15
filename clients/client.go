package clients

import (
	"fmt"
	"log"
	"net/http"

	"github.com/david-ds-teles/account_client/models"
)

// port responsible to provide communication externaly of this package
type Client interface {
	Create(acc *models.AccountData) (*models.AccountData, error)
	Fetch(id string) (*models.AccountData, error)
	Delete(id string, version *int64) error
}

// create a fakeClient constructor
func NewFakeClient(log *log.Logger) (Client, error) {
	if log == nil {
		return nil, fmt.Errorf("the dependency of NewFakeClient was not provided %v", log)
	}
	return newFakeClient(log), nil
}

// default client constructor
func NewClient(log *log.Logger, http *http.Client, apiEndpoint string) (Client, error) {
	if log == nil || http == nil || apiEndpoint == "" {
		return nil, fmt.Errorf("one or more dependencies of NewClient was not provided. Log: %v, http: %v, apiEndpoint: %v", log, http, apiEndpoint)
	}
	return newDefaultClient(log, http, apiEndpoint), nil
}
