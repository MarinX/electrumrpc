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

func TestCreateNewAddress(t *testing.T) {
	RegisterTestingT(t)
	responseBody = `{"result": "tb1q7k0d4yyx253t9te92nrlkzvy88l72f38dwhu72", "id": 5577006791947779410, "error": null}`
	res, err := client.CreateNewAddress()
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res).To(Equal("tb1q7k0d4yyx253t9te92nrlkzvy88l72f38dwhu72"))
}

func TestGetUnusedAddress(t *testing.T) {
	RegisterTestingT(t)
	responseBody = `{"result": "tb1qnf5dx9d3swffc08qkrhfjxqyrc6yq8qrcx6d4m", "id": 5577006791947779410, "error": null}`
	res, err := client.GetUnusedAddress()
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res).To(Equal("tb1qnf5dx9d3swffc08qkrhfjxqyrc6yq8qrcx6d4m"))
}

func TestIsMine(t *testing.T) {
	RegisterTestingT(t)
	responseBody = `{"result": true, "id": 5577006791947779410, "error": null}`

	res, err := client.IsMine("tb1qnf5dx9d3swffc08qkrhfjxqyrc6yq8qrcx6d4m")
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res).To(BeTrue())
}

func TestGetTransaction(t *testing.T) {
	RegisterTestingT(t)
	responseBody = `{"result": {"hex": "0200000003d1a4362c08dad0e1e1d9750ce672c50cfd5d898e3556b51bf7600ea57c245e14000000006a4730440220251763207d6fd06c8d846a3cc0ae4b285bc6f846de755936b2608ab25155cb90022007d85c30ba12d858001546472f430c35dc5f6e7cfa8b0531069dbed139779bfc0121036086223dc93e6dbac9a68869c653cc91c6df07025ff808f3748f6bcea16ac694feffffff7bd2d753f042705d826f7e3bcb94eafcd829cd0679dc0912107fc566834973b1010000006a473044022040003f31c20685b39b26749146348c2983ef0a13449e7563c569cd998895296c022064405083da407a2d1a95257633c072765e715aa450b3b1cd425376606727427d012103ff11526fc8dc65ffebd7e549dee3d946b1846a2ea36884c38f691e2c06eb63b0feffffff32e4389427afc0d376c648d9a0fe727150cda66b185f302a29cea33ad26c589f000000006a47304402206d74b803148c538039a149869d2917bd8975885e90cca551456b2b8cb6dc857f02200b474ecb3d8ef181022135ab45f536b8f5fbc1685033d6e927d0927d867858ae012103ff11526fc8dc65ffebd7e549dee3d946b1846a2ea36884c38f691e2c06eb63b0feffffff02d0ad1300000000001976a914dd0f8fc6874a1768ee774664fe69c3cc78c6115888ace0930400000000001600143dfcc8bfed24c5bf7ee8a7c8139d95a84ee20c2427891800", "complete": true, "final": true}, "id": 5577006791947779410, "error": null}`

	res, err := client.GetTransaction("063aaf441c45e95c8924f18157011ad240b2337c263e575b1bb0a3ce0eabf94a")
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res.Complete).To(BeTrue())
	Expect(res.Final).To(BeTrue())
	Expect(res.Hex).NotTo(BeEmpty())
}

func TestGetSeed(t *testing.T) {
	RegisterTestingT(t)
	responseBody = `{"result": "negative miracle small debris crime employ crash confirm inform unique pride hello", "id": 5577006791947779410, "error": null}`

	res, err := client.GetSeed("")
	<-requestChan

	Expect(err).To(BeNil())
	Expect(res).NotTo(BeEmpty())

}
