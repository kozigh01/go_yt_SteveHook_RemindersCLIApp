package client

import (
	"net/http"
	"time"
)

type HTTPClient struct {
	BackendUri string
	client     *http.Client
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendUri: uri,
		client:     &http.Client{},
	}
}

func (c HTTPClient) create(title, message string, duration time.Duration) ([]byte, error) {
	return []byte("something created"), nil
}

func (c HTTPClient) edit(id, title, message string, duration time.Duration) ([]byte, error) {
	return []byte("something edited"), nil
}

func (c HTTPClient) fetch(ids []string) ([]byte, error) {
	return []byte("something fetched"), nil
}

func (c HTTPClient) delete(ids []string) error {
	return nil
}

func (c HTTPClient) healthy(host string) bool {
	return true
}
