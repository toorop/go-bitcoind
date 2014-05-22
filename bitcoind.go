// Package Bitcoind is client librari for bitcoind JSON RPC API
package bitcoind

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	// VERSION represents bicoind package version
	VERSION = 0.1
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	RPCCLIENT_TIMEOUT = 30
)

// A Bitcoind represents a Bitcoind client
type Bitcoind struct {
	client *rpcClient
}

// New return a new bitcoind
func New(host string, port int, user, passwd string, useSSL bool) (*Bitcoind, error) {
	rpcClient, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}
	return &Bitcoind{rpcClient}, nil
}

// BackupWallet Safely copies wallet.dat to destination,
// which can be a directory or a path with filename on the remote server
func (b *Bitcoind) BackupWallet(destination string) error {
	r, err := b.client.call("backupwallet", []string{destination})
	return handleError(err, &r)
}

// DumpPrivKey return private key as string associated to public <address>
func (b *Bitcoind) DumpPrivKey(address string) (privKey string, err error) {
	r, err := b.client.call("dumpprivkey", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &privKey)
	return
}

// EncryptWallet encrypts the wallet with <passphrase>.
func (b *Bitcoind) EncryptWallet(passphrase string) error {
	r, err := b.client.call("encryptwallet", []string{passphrase})
	return handleError(err, &r)
}

// GetAccount returns the account associated with the given address.
func (b *Bitcoind) GetAccount(address string) (account string, err error) {
	r, err := b.client.call("getaccount", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &account)
	return
}

// GetAccountAddress Returns the current bitcoin address for receiving
// payments to this account.
// If account does not exist, it will be created along with an
// associated new address that will be returned.
func (b *Bitcoind) GetAccountAddress(account string) (address string, err error) {
	r, err := b.client.call("getaccountaddress", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &address)
	return
}

// GetAddressesByAccount return addresses associated with account <account>
func (b *Bitcoind) GetAddressesByAccount(account string) (addresses []string, err error) {
	r, err := b.client.call("getaddressesbyaccount", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addresses)
	return
}

// GetBalance return the balance of the server or of a specific account
//If [account] is "", returns the server's total available balance.
//If [account] is specified, returns the balance in the account
func (b *Bitcoind) GetBalance(account string, minconf uint64) (balance float64, err error) {
	r, err := b.client.call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// GetBestBlockhash returns the hash of the best (tip) block in the longest block chain.
func (b *Bitcoind) GetBestBlockhash() (bestBlockHash string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &bestBlockHash)
	return
}

// GetBlock returns information about the block with the given hash.
func (b *Bitcoind) GetBlock(blockHash string) (block Block, err error) {
	r, err := b.client.call("getblock", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

// GetBlockCount returns the number of blocks in the longest block chain.
func (b *Bitcoind) GetBlockCount() (count uint64, err error) {
	r, err := b.client.call("getblockcount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetBlockHash returns hash of block in best-block-chain at <index>
func (b *Bitcoind) GetBlockHash(index uint64) (hash string, err error) {
	r, err := b.client.call("getblockhash", []uint64{index})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hash)
	return
}

// getBlockTemplateParams reperesent parameters for GetBlockTemplate
type getBlockTemplateParams struct {
	Mode         string   `json:"mode,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}

// TODO a finir
// GetBlockTemplate Returns data needed to construct a block to work on.
// See BIP_0022 for more info on params.
func (b *Bitcoind) GetBlockTemplate(capabilities []string, mode string) (template string, err error) {
	params := getBlockTemplateParams{
		Mode:         mode,
		Capabilities: capabilities,
	}
	// TODO []interface{}{mode, capa}
	r, err := b.client.call("getblocktemplate", []getBlockTemplateParams{params})
	if err = handleError(err, &r); err != nil {
		return
	}
	fmt.Println(json.Unmarshal(r.Result, &template))
	return
}

// GetConnectionCount returns the number of connections to other nodes.
func (b *Bitcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.client.call("getconnectioncount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetDifficulty returns the proof-of-work difficulty as a multiple of
// the minimum difficulty.
func (b *Bitcoind) GetDifficulty() (difficulty float64, err error) {
	r, err := b.client.call("getdifficulty", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	difficulty, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// GetGenerate returns true or false whether bitcoind is currently generating hashes
func (b *Bitcoind) GetGenerate() (generate bool, err error) {
	r, err := b.client.call("getgenerate", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &generate)
	return
}

// GetHashesPerSec returns a recent hashes per second performance measurement while generating.
func (b *Bitcoind) GetHashesPerSec() (hashpersec float64, err error) {
	r, err := b.client.call("gethashespersec", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hashpersec)
	return
}

// GetInfo return result of "getinfo" command (Amazing !)
func (b *Bitcoind) GetInfo() (i Info, err error) {
	r, err := b.client.call("getinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &i)
	return
}

// GetMiningInfo returns an object containing mining-related information
func (b *Bitcoind) GetMiningInfo() (miningInfo MiningInfo, err error) {
	r, err := b.client.call("getmininginfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &miningInfo)
	return
}

// GetNewAddress return a new address for account [account].
func (b *Bitcoind) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetNewAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getnewaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addr)
	return
}

// GetPeerInfo returns data about each connected node
func (b *Bitcoind) GetPeerInfo() (peerInfo []Peer, err error) {
	r, err := b.client.call("getpeerinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &peerInfo)
	return
}

// GetRawChangeAddress Returns a new Bitcoin address, for receiving change.
// This is for use with raw transactions, NOT normal use.
func (b *Bitcoind) GetRawChangeAddress(account ...string) (rawAddress string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetRawChangeAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getrawchangeaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &rawAddress)
	return
}

// GetRawMempool returns all transaction ids in memory pool
func (b *Bitcoind) GetRawMempool() (txId []string, err error) {
	r, err := b.client.call("getrawmempool", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txId)
	return
}

// GetRawTransaction returns raw transaction representation for given transaction id.
func (b *Bitcoind) GetRawTransaction(txId string, verbose bool) (rawTx interface{}, err error) {
	r, err := b.client.call("getrawtransaction", []interface{}{txId, verbose})
	if err = handleError(err, &r); err != nil {
		return
	}
	if !verbose {
		err = json.Unmarshal(r.Result, &rawTx)
	} else {
		var t RawTransaction
		err = json.Unmarshal(r.Result, &t)
		rawTx = t
	}
	return
}

// GetReceivedByAccount Returns the total amount received by addresses with [account] in
// transactions with at least [minconf] confirmations. If [account] is set to all return
// will include all transactions to all accounts
func (b *Bitcoind) GetReceivedByAccount(account string, minconf uint32) (amount float64, err error) {
	if account == "all" {
		account = ""
	}
	r, err := b.client.call("getreceivedbyaccount", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &amount)
	return
}

// Returns the amount received by <address> in transactions with at least [minconf] confirmations.
// It correctly handles the case where someone has sent to the address in multiple transactions.
// Keep in mind that addresses are only ever used for receiving transactions. Works only for addresses
// in the local wallet, external addresses will always show 0.
func (b *Bitcoind) GetReceivedByAddress(address string, minconf uint32) (amount float64, err error) {
	r, err := b.client.call("getreceivedbyaddress", []interface{}{address, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &amount)
	return
}

// GetTransaction returns a Bitcoind.Transation struct about the given transaction
func (b *Bitcoind) GetTransaction(txid string) (transaction Transaction, err error) {
	r, err := b.client.call("gettransaction", []interface{}{txid})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

// GetTxOut returns details about an unspent transaction output (UTXO)
func (b *Bitcoind) GetTxOut(txid string, n uint32, includeMempool bool) (transactionOut UTransactionOut, err error) {
	r, err := b.client.call("gettxout", []interface{}{txid, n, includeMempool})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transactionOut)
	return
}

// GetTxOutsetInfo returns statistics about the unspent transaction output (UTXO) set
func (b *Bitcoind) GetTxOutsetInfo() (txOutSet TransactionOutSet, err error) {
	r, err := b.client.call("gettxoutsetinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txOutSet)
	return
}

// GetWork
// If [data] is not specified, returns formatted hash data to work on
// If [data] is specified, tries to solve the block and returns true if it was successful.
func (b *Bitcoind) GetWork(data ...string) (response interface{}, err error) {
	if len(data) > 1 {
		err = errors.New("Bad parameters for GetWork: you can set 0 or 1 parameter data")
		return
	}
	var r rpcResponse

	if len(data) == 0 {
		r, err = b.client.call("getwork", nil)
		if err = handleError(err, &r); err != nil {
			return
		}
		var work Work
		err = json.Unmarshal(r.Result, &work)
		response = work
	} else {
		r, err = b.client.call("getwork", data)
		if err = handleError(err, &r); err != nil {
			return
		}
		var t bool
		err = json.Unmarshal(r.Result, &t)
		response = t
	}
	return
}

// ImportPrivKey Adds a private key (as returned by dumpprivkey) to your wallet.
// This may take a while, as a rescan is done, looking for existing transactions.
// Optional [rescan] parameter added in 0.8.0.
// Note: There's no need to import public key, as in ECDSA (unlike RSA) this
// can be computed from private key.
func (b *Bitcoind) ImportPrivKey(privKey, label string, rescan bool) error {
	r, err := b.client.call("importprivkey", []interface{}{privKey, label, rescan})
	return handleError(err, &r)
}

// walletPassphrase stores the wallet decryption key in memory for <timeout> seconds.
func (b *Bitcoind) WalletPassphrase(passPhrase string, timeout uint64) error {
	r, err := b.client.call("walletpassphrase", []interface{}{passPhrase, timeout})
	return handleError(err, &r)
}
