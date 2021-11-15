package clients

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/david-ds-teles/account_client/models"
	"github.com/google/uuid"
)

// a defaultCliente to interact with accounts API using an httpClient
type defaultClient struct {
	log         *log.Logger
	http        *http.Client
	apiEndpoint string
}

//account endpoint to concatenate it with apiEndpoint
const accountEndpoint = "/v1/organisation/accounts"

// default content type used
const contentType = "application/json"

// constructor for the default client
func newDefaultClient(log *log.Logger, http *http.Client, apiEndpoint string) defaultClient {
	return defaultClient{log, http, apiEndpoint}
}

// responsible to check every http response comparing the expected statusCode with received statusCode
// if status is not the same, an error is returned
func (client defaultClient) checkHttpResponse(rsp *http.Response, expectedStatusCode int) error {
	client.log.Printf("checking http response for expected statucCode: %d", expectedStatusCode)

	// success
	if rsp.StatusCode == expectedStatusCode {
		return nil
	}

	client.log.Printf("http response failed. StatusCode expected: %d, received: %d ", expectedStatusCode, rsp.StatusCode)

	// error
	rspMessage, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return fmt.Errorf("error calling http: %v", err)
	}

	return fmt.Errorf("operation failed with httpStatus: %d, Message: %s", rsp.StatusCode, string(rspMessage))
}

// models.AccountData need be provided. The respective IDs are created using github.com/google/uuid
// if created, an AccountData is returned with all provided fields
func (client defaultClient) Create(acc *models.AccountData) (*models.AccountData, error) {
	client.log.Println("starting defaultClient.Create")

	if acc == nil {
		return nil, errors.New("invalid account, nil")
	}

	acc.ID = uuid.New().String()
	acc.OrganisationID = uuid.New().String()
	acc.Type = "accounts"
	version := int64(0)
	acc.Version = &version

	buffer, err := acc.ToJSON()

	if err != nil {
		return nil, err
	}

	rsp, err := client.http.Post(fmt.Sprintf("%s%s", client.apiEndpoint, accountEndpoint), contentType, buffer)

	if err != nil {
		return nil, fmt.Errorf("error trying call API. Error %v", err)
	}

	defer rsp.Body.Close()
	err = client.checkHttpResponse(rsp, http.StatusCreated)

	if err != nil {
		return nil, err
	}

	client.log.Printf("account successfully created ID: %s", acc.ID)

	result := &models.AccountData{}
	err = result.FromJSON(rsp.Body)

	if err != nil {
		return nil, err
	}

	client.log.Println("ending defaultClient.Create")
	return result, nil
}

// Fetch an AccountData using id provided. If success, an AccountData is returned
func (client defaultClient) Fetch(id string) (*models.AccountData, error) {
	client.log.Printf("starting defaultClient.Fetch id: %s", id)

	if id == "" {
		return nil, errors.New("invalid account id")
	}

	url := fmt.Sprintf("%s%s/%s", client.apiEndpoint, accountEndpoint, id)
	rsp, err := client.http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error trying call API. Error %v", err)
	}

	defer rsp.Body.Close()
	err = client.checkHttpResponse(rsp, http.StatusOK)

	if err != nil {
		return nil, err
	}

	result := &models.AccountData{}
	err = result.FromJSON(rsp.Body)

	if err != nil {
		return nil, err
	}

	client.log.Println("account successfully fetched")
	client.log.Println("ending defaultClient.Fetch")
	return result, nil
}

// to delete an AccountData its id and version need to be provided. If deleted, nil is returned
func (client defaultClient) Delete(id string, version *int64) error {
	client.log.Printf("starting defaultClient.Delete id: %s\n", id)

	if id == "" || version == nil {
		return fmt.Errorf("invalid params, id: %s , version: %d", id, *version)
	}

	url := fmt.Sprintf("%s%s/%s?version=%d", client.apiEndpoint, accountEndpoint, id, *version)
	delReq, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return fmt.Errorf("error trying create http.DELETE. %v", err)
	}

	rsp, err := client.http.Do(delReq)

	if err != nil {
		return fmt.Errorf("error trying call API. Http error: %v", err)
	}

	err = client.checkHttpResponse(rsp, http.StatusNoContent)

	if err != nil {
		return err
	}

	client.log.Println("account successfully deleted")
	client.log.Println("ending defaultClient.Delete")

	return nil
}
