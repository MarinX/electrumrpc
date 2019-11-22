package electrumrpc

// Version returns the version of Electrum.
func (c *Client) Version() (res string, err error) {
	err = c.Call("version", nil, &res)
	return
}

// GetAddressBalance returns balance for given address
func (c *Client) GetAddressBalance(address string) (res Balance, err error) {
	err = c.Call("getaddressbalance", address, &res)
	return
}

// GetBalance return the balance of your wallet.
func (c *Client) GetBalance() (res Balance, err error) {
	err = c.Call("getbalance", nil, &res)
	return
}

// ValidateAddress check that an address is valid
func (c *Client) ValidateAddress(address string) (res bool, err error) {
	err = c.Call("validateaddress", address, &res)
	return
}

// GetServers returns the list of available servers
func (c *Client) GetServers() (res map[string]Server, err error) {
	err = c.Call("getservers", nil, &res)
	return
}

// CreateNewAddress creates a new receiving address, beyond the gap limit of the wallet
func (c *Client) CreateNewAddress() (res string, err error) {
	err = c.Call("createnewaddress", nil, &res)
	return
}

// GetUnusedAddress returns the first unused address of the wallet, or None if all addresses are used
// An address is considered as used if it has received a transaction, or if it is used in a payment request.
func (c *Client) GetUnusedAddress() (res string, err error) {
	err = c.Call("getunusedaddress", nil, &res)
	return
}

// IsMine checks if address is in wallet.
// Return true if and only address is in wallet
func (c *Client) IsMine(address string) (res bool, err error) {
	err = c.Call("ismine", address, &res)
	return
}

// GetTransaction retrieve a transaction by id
func (c *Client) GetTransaction(txid string) (res Transaction, err error) {
	err = c.Call("gettransaction", txid, &res)
	return
}

// GetSeed returns the generation seed of your wallet
func (c *Client) GetSeed(password string) (res string, err error) {
	var pass *string
	if len(password) > 0 {
		pass = &password
	}
	// cannot pass empty string, we need to pass nil for none password wallet
	err = c.Call("getseed", pass, &res)
	return
}

// ListAdddresses returns the list of all addresses in your wallet
// @TODO, support optional arguments to filter the results
func (c *Client) ListAdddresses() (res []string, err error) {
	err = c.Call("listaddresses", nil, &res)
	return
}

// AddRequest creates a payment request, using the first unused address of the wallet
func (c *Client) AddRequest(amount float64, memo string, expiration uint64) (res PaymentRequest, err error) {
	err = c.Call("addrequest", []interface{}{amount, memo, expiration}, &res)
	return
}

// ListRequest lists the payment requests you made
func (c *Client) ListRequest(pending, expired, paid bool) (res []PaymentRequest, err error) {
	err = c.Call("listrequests", []bool{pending, expired, paid}, &res)
	return
}

// RemoveRequest removes a payment request
// Returns true if removal was successful
func (c *Client) RemoveRequest(btcAddress string) (res bool, err error) {
	err = c.Call("rmrequest", btcAddress, &res)
	return
}

// GetRequest returns a payment request
func (c *Client) GetRequest(btcAddress string) (res PaymentRequest, err error) {
	err = c.Call("getrequest", btcAddress, &res)
	return
}

// ClearRequests removes all payment requests
func (c *Client) ClearRequests() (err error) {
	err = c.Call("clearrequests", nil, nil)
	// invalid JSON RPC response
	// check if result is null which is valid response
	if err != nil && err.Error() == "result is null" {
		err = nil
	}
	return
}

// GetFeeRate returns current suggested fee rate (in sat/kvByte), according to config
// settings or supplied parameters
func (c *Client) GetFeeRate(feeMethod FeeMethod) (res uint64, err error) {
	if feeMethod == "" {
		err = c.Call("getfeerate", nil, &res)
		return
	}
	err = c.Call("getfeerate", feeMethod, &res)
	return
}

// SignMessage signs a message with a key
func (c *Client) SignMessage(btcAddress, message string) (res string, err error) {
	req := &SignMessageRequest{
		Address: btcAddress,
		Message: message,
	}
	err = c.Call("signmessage", req, &res)
	return
}

// VerifyMessage verifies a signature
func (c *Client) VerifyMessage(btcAddress, signature, message string) (res bool, err error) {
	req := &VerifyMessageRequest{
		Address:   btcAddress,
		Signature: signature,
		Message:   message,
	}
	err = c.Call("verifymessage", req, &res)
	return
}
