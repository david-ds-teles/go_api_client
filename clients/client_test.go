package clients

import (
	"log"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient(log.Default(), http.DefaultClient, "http://mocked")

	if err != nil {
		t.Fatalf("error creating new client %v", err)
	}

	if client == nil {
		t.Fatalf("client instance invalid. client: %v", client)
	}

	switch client.(type) {
	case defaultClient:
		break
	default:
		t.Fatalf("expecting type: defaultClient, got type: %T", client)
	}
}

func TestNewClientLogDepencyInvalid(t *testing.T) {
	client, err := NewClient(nil, http.DefaultClient, "http://mocked")

	if err == nil {
		t.Fatalf("expecting failure but got error %v", err)
	}

	if client != nil {
		t.Fatalf("expecting failure but got client instance %T", client)
	}

}

func TestNewClientHttpDepencyInvalid(t *testing.T) {
	client, err := NewClient(log.Default(), nil, "http://mocked")

	if err == nil {
		t.Fatalf("expecting failure but got error %v", err)
	}

	if client != nil {
		t.Fatalf("expecting failure but got client instance %T", client)
	}

}

func TestNewClientEndpointDepencyInvalid(t *testing.T) {
	client, err := NewClient(log.Default(), http.DefaultClient, "")

	if err == nil {
		t.Fatalf("expecting failure but got error %v", err)
	}

	if client != nil {
		t.Fatalf("expecting failure but got client instance %T", client)
	}

}

func TestNewFakeClient(t *testing.T) {
	client, err := NewFakeClient(log.Default())

	if err != nil {
		t.Fatalf("error creating fake client %v", err)
	}

	if client == nil {
		t.Fatalf("fake client instance invalid. client: %v", client)
	}

	switch client.(type) {
	case fakeClient:
		break
	default:
		t.Fatalf("expecting type: fakeClient, got type: %T", client)
	}
}

func TestNewFakeClientDepencyInvalid(t *testing.T) {
	client, err := NewFakeClient(nil)

	if err == nil {
		t.Fatalf("expecting failure but got error %v", err)
	}

	if client != nil {
		t.Fatalf("expecting failure but got fake client instance %v", client)
	}
}
