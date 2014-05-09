package bitcoind

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// VERSION represents bicoind package version
	VERSION = 0.1
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	RPCCLIENT_TIMEOUT = 30
)

// A bitpay represents a bitpay client wrapper
type bitcoind struct {
	client *rpcClient
}

// New return a new bitcoind
func New(host string, port int, user, passwd string, useSSL bool) *bitcoind {
	rpcClient := newClient(host, port, user, passwd, useSSL)
	return &bitcoind{rpcClient}
}

// rpcErr2Err return an error from an RPC error
func rpcErr2Err(rpcErr interface{}) error {
	r := rpcErr.(map[string]interface{})
	return errors.New(fmt.Sprintf("(%v) %s", r["code"].(float64), r["message"].(string)))
}

/*type Info struct {
	Balance         Amount
	Blocks          int
	Connections     int
	Difficulty      float64
	Errors          string
	KeyPoolOldest   uint32
	KeyPoolSize     int
	PayTxFee        Amount
	ProtocolVersion int
	Proxy           string
	Testnet         bool
	Version         int
	WalletVersion   int
}*/

// GetInfo return result of "getinfo" command (Amazing !)
func (b *bitcoind) GetInfo() (i info, err error) {
	r, err := b.client.call("getinfo", nil)
	if err != nil {
		return
	}
	if r.Err != nil {
		err = rpcErr2Err(r.Err)
		return
	}
	err = json.Unmarshal(r.Result, &i)
	return
}

// GetNewAddress return a new address for account "account". account parameter is optional
func (b *bitcoind) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetNewAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getnewaddress", account)
	if err != nil {
		return
	}
	if r.Err != nil {
		err = rpcErr2Err(r.Err)
		return
	}
	addr = string(r.Result)
	return
}

// GetAddressesByAccount return addresses associated with account "account"
func (b *bitcoind) GetAddressesByAccount(account string) (addresses []string, err error) {
	params := make([]string, 0)
	params = append(params, account)
	r, err := b.client.call("getaddressesbyaccount", params)
	if err != nil {
		return
	}
	if r.Err != nil {
		err = rpcErr2Err(r.Err)
		return
	}
	err = json.Unmarshal(r.Result, &addresses)
	return
}
