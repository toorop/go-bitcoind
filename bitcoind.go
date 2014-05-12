package bitcoind

import (
	"encoding/json"
	"errors"
	//"fmt"
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
func New(host string, port int, user, passwd string, useSSL bool) (*bitcoind, error) {
	rpcClient, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}
	return &bitcoind{rpcClient}, nil
}

// BackupWallet Safely copies wallet.dat to destination,
// which can be a directory or a path with filename on the remote server
func (b *bitcoind) BackupWallet(destination string) error {
	r, err := b.client.call("backupwallet", []string{destination})
	return handleError(err, &r)
}

// DumpPrivKey return private key as string associated to public <address>
func (b *bitcoind) DumpPrivKey(address string) (privKey string, err error) {
	r, err := b.client.call("dumpprivkey", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &privKey)
	return
}

// EncryptWallet encrypts the wallet with <passphrase>.
func (b *bitcoind) EncryptWallet(passphrase string) error {
	r, err := b.client.call("encryptwallet", []string{passphrase})
	return handleError(err, &r)
}

// GetAccount returns the account associated with the given address.
func (b *bitcoind) GetAccount(address string) (account string, err error) {
	r, err := b.client.call("getaccount", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	account = string(r.Result)
	return
}

// GetAccountAddress Returns the current bitcoin address for receiving
// payments to this account.
// If account does not exist, it will be created along with an
// associated new address that will be returned.
func (b *bitcoind) GetAccountAddress(account string) (address string, err error) {
	r, err := b.client.call("getaccountaddress", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	address = string(r.Result)
	return

}

// GetInfo return result of "getinfo" command (Amazing !)
func (b *bitcoind) GetInfo() (i info, err error) {
	r, err := b.client.call("getinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &i)
	return
}

// GetNewAddress return a new address for account [account].
func (b *bitcoind) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetNewAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getnewaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	addr = string(r.Result)
	return
}

// GetAddressesByAccount return addresses associated with account <account>
func (b *bitcoind) GetAddressesByAccount(account string) (addresses []string, err error) {
	r, err := b.client.call("getaddressesbyaccount", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addresses)
	return
}
