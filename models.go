package electrumrpc

// Balance model
type Balance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
	Unmatured   string `json:"unmatured"`
	Lightning   string `json:"lightning"`
}

// Server is electrum server information
type Server struct {
	S       string `json:"s"`
	Pruning string `json:"pruning"`
	Version string `json:"version"`
}
