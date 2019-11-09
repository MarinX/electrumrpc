package electrumrpc

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

type RequestData struct {
	request *http.Request
	body    string
}

var httpServer *httptest.Server
var client *Client
var requestChan = make(chan *RequestData, 1)
var responseBody = ""

func TestMain(t *testing.M) {

	httpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		requestChan <- &RequestData{r, string(data)}
		fmt.Fprintf(w, responseBody)
	}))
	defer httpServer.Close()
	client = New("user", "password", httpServer.URL, nil)
	os.Exit(t.Run())
}

func TestClientHeaders(t *testing.T) {
	RegisterTestingT(t)

	client.Call("method", nil, nil)

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", "user", "password")))

	req := (<-requestChan).request
	Expect(req.Method).To(Equal("POST"))
	Expect(req.Header.Get("Content-Type")).To(Equal("application/json"))
	Expect(req.Header.Get("Authorization")).To(Equal(fmt.Sprintf("Basic %s", auth)))
}

func TestVersion(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": "3.3.8", "id": 5577006791947779410, "error": null}`
	res, err := client.Version()
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res).To(Equal("3.3.8"))
}

func TestGetAddressBalance(t *testing.T) {
	RegisterTestingT(t)
	responseBody = `{"result": {"confirmed": "0.04027708", "unconfirmed": "0"}, "id": 5577006791947779410, "error": null}`
	res, err := client.GetAddressBalance("1BTCAddressABC")
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res.Confirmed).To(Equal("0.04027708"))
	Expect(res.Unconfirmed).To(Equal("0"))
}

func TestValidateBalance(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": false, "id": 5577006791947779410, "error": null}`
	res, err := client.ValidateAddress("UnvalidBTCAddress")
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res).To(BeFalse())
}

func TestGetServers(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": {"testnet.qtornado.com": {"pruning": "-", "s": "51002", "t": "51001", "version": "1.4"}}, "error": null}`
	res, err := client.GetServers()
	<-requestChan

	t.Log(res["testnet.qtornadeo.com"].S)
	Expect(err).To(BeNil())

	for k, v := range res {
		Expect(k).NotTo(BeEmpty())
		Expect(v.S).To(Equal("51002"))
	}
}
