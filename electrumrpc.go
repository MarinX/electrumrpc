package electrumrpc

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/MarinX/rpc/json"
)

// Client holds the JSON RPC HTTP client with Basic Auth
type Client struct {
	auth       string
	addr       string
	httpClient *http.Client
}

// New creates new client with username and password
// Optional, you can set your own http.Client
// If http.Client is not specified, the http.DefaultClient will be used
func New(username, password, addr string, httpClient *http.Client) *Client {
	auth := []byte(fmt.Sprintf("%s:%s", username, password))
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		auth:       base64.StdEncoding.EncodeToString(auth),
		httpClient: httpClient,
		addr:       addr,
	}
}

// Call is a method for calling JSON RPC endpoint
// Method is public - so if some func is missing, you can call here with your own model
func (c *Client) Call(method string, params interface{}, result interface{}) error {
	msg, err := json.EncodeClientRequest(method, params)
	if err != nil {
		return err
	}

	req, err := c.newRequest(msg)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.DecodeClientResponse(resp.Body, result)
}

// newRequest creates new HTTP POST request with Basic Auth headers
func (c *Client) newRequest(body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", c.addr, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.auth))
	return req, nil
}
