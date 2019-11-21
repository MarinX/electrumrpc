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

// Transaction is bitcoin transaction model
type Transaction struct {
	Hex      string `json:"hex"`
	Complete bool   `json:"complete"`
	Final    bool   `json:"final"`
}

// PaymentRequest is electrum payment request model
type PaymentRequest struct {
	ID         string `json:"id"`
	Amount     uint64 `json:"amount"`
	Expiration uint64 `json:"exp"`
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	URI        string `json:"URI"`
	Status     string `json:"status"`
	AmountBTC  string `json:"amount (BTC)"`
	Time       uint64 `json:"time"`
}
