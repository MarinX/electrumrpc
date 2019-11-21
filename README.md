# Electrum JSON RPC Client

[![Build Status](https://travis-ci.org/MarinX/electrumrpc.svg?branch=master)](https://travis-ci.org/MarinX/electrumrpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarinX/electrumrpc)](https://goreportcard.com/report/github.com/MarinX/electrumrpc)
[![GoDoc](https://godoc.org/github.com/MarinX/electrumrpc?status.svg)](https://godoc.org/github.com/MarinX/electrumrpc)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

**Note**: The library does not have implementations of all Electrum RPC resources[WIP]. 
PRs for new resources and endpoints are welcome, or you can simply implement some yourself as-you-go.

## Preposition
#### 1. Install [Electrum](https://electrum.org/) and create a wallet

#### 2. Set RPC port
By default, it's random port - set it to any port you want
```bash
./run_electrum setconfig rpcport 7777
``` 

#### 3. Set user and password for RPC
```bash
./run_electrum setconfig rpcuser user
```
```bash
./run_electrum setconfig rpcpassword password
```

#### 4. Run Electrum as daemon
```bash
./run_electrum daemon start
```
If you want to start in testnet mode
```bash
./run_electrum --testnet daemon start
```

#### 5. Load daemon wallet
```bash
./run_electrum daemon load_wallet
```
If daemon is running in testnet, you need to specify to load testnet wallet
```bash
./run_electrum --testnet daemon load_wallet
```

Now you have a local Electrum JSON RPC server running - congrats 🥳

If you need to stop it, use
```bash
./run_electrum daemon stop
```
or if running in testnet
```bash
./run_electrum --testnet daemon stop
```

## Install

```bash
go get github.com/MarinX/electrumrpc
```

## Use

```go
import "github.com/MarinX/electrumrpc"
```

## Example
You can find more in [electrumrpc_test.go](https://github.com/MarinX/electrumrpc/blob/master/electrumrpc_test.go)
```go
// httpClient is optional
// if nil, the http.DefaultClient will be used
client := electrumrpc.New("<rpc-user>", "<rpc-password>", "<rpc-endpoint>", nil)

// Call RPC methods
ver, err := client.Version()
if err != nil {
	//handle error
	panic(err)
}
fmt.Println("Electrum version:", ver)
```

Not all endpoints are implemented right now. 
In those case, you can use Call method and point your model
```go
var rpcResponse string
err := client.Call("version", nil, &rpcResponse)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println("Electrum version:", rpcResponse)
```

## Available RPC methods

| RPC Method  | Available |
| --- | --- |
| version  | ✅  |
| getaddressbalance  | ✅  |
| getbalance  | ✅  |
| validateaddress  | ✅  |
| getservers  | ✅  |
| createnewaddress  | ✅  |
| getunusedaddress  | ✅  |
| ismine  | ✅  |
| gettransaction  | ✅  |
| getseed  | ✅  |
| listaddresses  | ✅  |
| addrequest  | ✅  |
| rmrequest  | ✅  |
| clearrequests  | ✅  |
| getrequest  | ✅  |
| getfeerate  | ✅  |

## Contributing
PR's are welcome. Please read [CONTRIBUTING.md](https://github.com/MarinX/electrumrpc/blob/master/CONTRIBUTING.md) for more info

## License
MIT