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
