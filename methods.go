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
