// Package Navcoind is client library for navcoind JSON RPC API
package navcoind

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	// VERSION represents navcoind package version
	VERSION = 0.1
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	RPCCLIENT_TIMEOUT = 30
)

// A Navcoind represents a Navcoind client
type Navcoind struct {
	client *rpcClient
}

// New return a new navcoind
func New(host string, port int, user, passwd string, useSSL bool, timeoutParam ...int) (*Navcoind, error) {
	var timeout int = RPCCLIENT_TIMEOUT
	// If the timeout is specified in timeoutParam, allow it.
	if len(timeoutParam) != 0 {
		timeout = timeoutParam[0]
	}

	rpcClient, err := newClient(host, port, user, passwd, useSSL, timeout)
	if err != nil {
		return nil, err
	}
	return &Navcoind{rpcClient}, nil
}

// BackupWallet Safely copies wallet.dat to destination,
// which can be a directory or a path with filename on the remote server
func (b *Navcoind) BackupWallet(destination string) error {
	r, err := b.client.call("backupwallet", []string{destination})
	return handleError(err, &r)
}

// DumpPrivKey return private key as string associated to public <address>
func (b *Navcoind) DumpPrivKey(address string) (privKey string, err error) {
	r, err := b.client.call("dumpprivkey", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &privKey)
	return
}

// EncryptWallet encrypts the wallet with <passphrase>.
func (b *Navcoind) EncryptWallet(passphrase string) error {
	r, err := b.client.call("encryptwallet", []string{passphrase})
	return handleError(err, &r)
}

// GetAccount returns the account associated with the given address.
func (b *Navcoind) GetAccount(address string) (account string, err error) {
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
func (b *Navcoind) GetAccountAddress(account string) (address string, err error) {
	r, err := b.client.call("getaccountaddress", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &address)
	return
}

// GetAddressesByAccount return addresses associated with account <account>
func (b *Navcoind) GetAddressesByAccount(account string) (addresses []string, err error) {
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
func (b *Navcoind) GetBalance(account string, minconf uint64) (balance float64, err error) {
	r, err := b.client.call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

type BlockHeader struct {
	Hash              string
	Confirmations     int
	Height            int
	Version           uint32
	VersionHex        string
	Merkleroot        string
	Time              int64
	Mediantime        int64
	Nonce             uint32
	Bits              uint32
	Difficulty        float64
	Chainwork         string
	Txes              int    `json:"nTx"`
	Previousblockhash string `json:"omitempty"`
	Nextblockhash     string `json:"omitempty"`
}

func (b *Navcoind) GetBlockheader(blockHash string) (*BlockHeader, error) {
	r, err := b.client.call("getblockheader", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return nil, err
	}

	var blockHeader BlockHeader
	err = json.Unmarshal(r.Result, &blockHeader)

	return &blockHeader, nil
}

// GetBestBlockhash returns the hash of the best (tip) block in the longest block chain.
func (b *Navcoind) GetBestBlockhash() (bestBlockHash string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &bestBlockHash)
	return
}

// GetBlock returns information about the block with the given hash.
func (b *Navcoind) GetBlock(blockHash string) (block Block, err error) {
	r, err := b.client.call("getblock", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

// GetRawBlock returns information about the block with the given hash.
func (b *Navcoind) GetRawBlock(blockHash string) (str string, err error) {
	r, err := b.client.call("getblock", []interface{}{blockHash, false})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &str)
	return
}

// GetBlockCount returns the number of blocks in the longest block chain.
func (b *Navcoind) GetBlockCount() (count uint64, err error) {
	r, err := b.client.call("getblockcount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetBlockHash returns hash of block in best-block-chain at <index>
func (b *Navcoind) GetBlockHash(index uint64) (hash string, err error) {
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
func (b *Navcoind) GetBlockTemplate(capabilities []string, mode string) (template string, err error) {
	params := getBlockTemplateParams{
		Mode:         mode,
		Capabilities: capabilities,
	}
	// TODO []interface{}{mode, capa}
	r, err := b.client.call("getblocktemplate", []getBlockTemplateParams{params})
	if err = handleError(err, &r); err != nil {
		return
	}
	return
}

type ChainTip struct {
	// The height of the current tip
	Height int
	// The hash of the highest tip
	Hash string
	// The length of the associated blockchain branch
	BranchLen int
	// The status of the current tip.
	Status string
}

func (b *Navcoind) GetChainTips() (tips []ChainTip, err error) {
	r, err := b.client.call("getchaintips", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &tips)
	return
}

// GetConnectionCount returns the number of connections to other nodes.
func (b *Navcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.client.call("getconnectioncount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetDifficulty returns the proof-of-work difficulty as a multiple of
// the minimum difficulty.
func (b *Navcoind) GetDifficulty() (difficulty float64, err error) {
	r, err := b.client.call("getdifficulty", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	difficulty, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// GetGenerate returns true or false whether navcoind is currently generating hashes
func (b *Navcoind) GetGenerate() (generate bool, err error) {
	r, err := b.client.call("getgenerate", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &generate)
	return
}

// GetHashesPerSec returns a recent hashes per second performance measurement while generating.
func (b *Navcoind) GetHashesPerSec() (hashpersec float64, err error) {
	r, err := b.client.call("gethashespersec", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hashpersec)
	return
}

// GetInfo return result of "getinfo" command
func (b *Navcoind) GetInfo() (i Info, err error) {
	r, err := b.client.call("getinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &i)
	return
}

// GetBlockchainInfo return result of "getblockchaininfo" command
func (b *Navcoind) GetBlockchainInfo() (i BlockchainInfo, err error) {
	r, err := b.client.call("getblockchaininfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &i)
	return
}

// GetMiningInfo returns an object containing mining-related information
func (b *Navcoind) GetMiningInfo() (miningInfo MiningInfo, err error) {
	r, err := b.client.call("getmininginfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &miningInfo)
	return
}

// GetNewAddress return a new address for account [account].
func (b *Navcoind) GetNewAddress(account ...string) (addr string, err error) {
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
func (b *Navcoind) GetPeerInfo() (peerInfo []Peer, err error) {
	r, err := b.client.call("getpeerinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &peerInfo)
	return
}

// GetRawChangeAddress Returns a new Bitcoin address, for receiving change.
// This is for use with raw transactions, NOT normal use.
func (b *Navcoind) GetRawChangeAddress(account ...string) (rawAddress string, err error) {
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
func (b *Navcoind) GetRawMempool() (txId []string, err error) {
	r, err := b.client.call("getrawmempool", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txId)
	return
}

type VerboseTx struct {
	// Virtual transaction size as defined in BIP 141
	Size uint32
	// Transaction fee in BTC
	Fee float64
	// Transaction fee with fee deltas used for mining priority
	ModifiedFee float64
	// Local time when tx entered pool
	Time uint32
	// Block height when tx entered pool
	Height uint32
	// Number of inpool descendents (including this one)
	DescendantCount uint32
	// Virtual transaction size of in-mempool descendants (including this one)
	DescendantSize uint32
	// Modified fees (see above) of in-mempool descendants (including this one)
	DescendantFees float64
	// Number of in-mempool ancestor transactions (including this one)
	AncestorCount uint32
	// Virtual transaction size of in-mempool ancestors (including this one)
	AncestorSize uint32
	// Modified fees (see above) of in-mempool ancestors (including this one)
	AncestorFees uint32
	// Hash of serialized transaction, including witness data
	WTxId string
	// Unconfirmed transactions used as inputs for this transaction
	Depends []string
	// Used by Bitcoin Unlimited RPC
	SpentBy []string
}

// GetRawMempoolVerbose returns a verbose set of transactions
// map [TxId] => VerboseTx
func (b *Navcoind) GetRawMempoolVerbose() (txs map[string]VerboseTx, err error) {
	r, err := b.client.call("getrawmempool", []bool{true})
	if err = handleError(err, &r); err != nil {
		return
	}
	// marshall the data into the string txs map
	err = json.Unmarshal(r.Result, &txs)
	return
}

// GetRawTransaction returns raw transaction representation for given transaction id.
func (b *Navcoind) GetRawTransaction(txId string, verbose bool) (rawTx interface{}, err error) {
	intVerbose := 0
	if verbose {
		intVerbose = 1
	}
	r, err := b.client.call("getrawtransaction", []interface{}{txId, intVerbose})
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
func (b *Navcoind) GetReceivedByAccount(account string, minconf uint32) (amount float64, err error) {
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
func (b *Navcoind) GetReceivedByAddress(address string, minconf uint32) (amount float64, err error) {
	r, err := b.client.call("getreceivedbyaddress", []interface{}{address, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &amount)
	return
}

// GetTransaction returns a Navcoind.Transation struct about the given transaction
func (b *Navcoind) GetTransaction(txid string) (transaction Transaction, err error) {
	r, err := b.client.call("gettransaction", []interface{}{txid})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

// GetTxOut returns details about an unspent transaction output (UTXO)
func (b *Navcoind) GetTxOut(txid string, n uint32, includeMempool bool) (transactionOut UTransactionOut, err error) {
	r, err := b.client.call("gettxout", []interface{}{txid, n, includeMempool})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transactionOut)
	return
}

// GetTxOutsetInfo returns statistics about the unspent transaction output (UTXO) set
func (b *Navcoind) GetTxOutsetInfo() (txOutSet TransactionOutSet, err error) {
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
func (b *Navcoind) GetWork(data ...string) (response interface{}, err error) {
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
func (b *Navcoind) ImportPrivKey(privKey, label string, rescan bool) error {
	r, err := b.client.call("importprivkey", []interface{}{privKey, label, rescan})
	return handleError(err, &r)
}

// KeyPoolRefill fills the keypool, requires wallet passphrase to be set.
func (b *Navcoind) KeyPoolRefill() error {
	r, err := b.client.call("keypoolrefill", nil)
	return handleError(err, &r)
}

// ListAccounts returns Object that has account names as keys, account balances as values.
func (b *Navcoind) ListAccounts(minconf int32) (accounts map[string]float64, err error) {
	r, err := b.client.call("listaccounts", []int32{minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &accounts)
	return
}

// ListAddressResult represents a result composing ListAddressGroupings slice reply
type ListAddressResult struct {
	Address string
	Amount  float64
	Account string
}

// ListAddressGroupings returns all addresses in the wallet and info used for coincontrol.
func (b *Navcoind) ListAddressGroupings() (list []ListAddressResult, err error) {
	r, err := b.client.call("listaddressgroupings", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	// hum.....
	var t [][][]interface{}
	err = json.Unmarshal(r.Result, &t)
	for _, tt := range t {
		for _, ttt := range tt {
			list = append(list, ListAddressResult{ttt[0].(string), ttt[1].(float64), ttt[2].(string)})
		}
	}
	return
}

// ReceivedByAccount represents how much coin a account have recieved
type ReceivedByAccount struct {
	// the account of the receiving addresses
	Account string
	// total amount received by addresses with this account
	Amount float64
	// number of confirmations of the most recent transaction included
	Confirmations uint32
}

// ListReceivedByAccount Returns an slice of AccountRecieved:
func (b *Navcoind) ListReceivedByAccount(minConf uint32, includeEmpty bool) (list []ReceivedByAccount, err error) {
	r, err := b.client.call("listreceivedbyaccount", []interface{}{minConf, includeEmpty})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &list)
	return
}

// ReceivedByAddress represents how much coin a account have recieved
type ReceivedByAddress struct {
	//  receiving address
	Address string
	// The corresponding account
	Account string
	// total amount received by addresses with this account
	Amount float64
	// number of confirmations of the most recent transaction included
	Confirmations uint32
	// Tansactions ID
	TxIds []string
}

// ListReceivedByAccount Returns an slice of AccountRecieved:
func (b *Navcoind) ListReceivedByAddress(minConf uint32, includeEmpty bool) (list []ReceivedByAddress, err error) {
	r, err := b.client.call("listreceivedbyaddress", []interface{}{minConf, includeEmpty})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &list)
	return
}

// ListSinceBlock
func (b *Navcoind) ListSinceBlock(blockHash string, targetConfirmations uint32) (transaction []Transaction, err error) {
	r, err := b.client.call("listsinceblock", []interface{}{blockHash, targetConfirmations})
	if err = handleError(err, &r); err != nil {
		return
	}
	type ts struct {
		Transactions []Transaction
	}
	var result ts
	if err = json.Unmarshal(r.Result, &result); err != nil {
		return
	}
	transaction = result.Transactions
	return
}

// ListTransactions returns up to [count] most recent transactions skipping the first
// [from] transactions for account [account]. If [account] not provided it'll return
// recent transactions from all accounts.
func (b *Navcoind) ListTransactions(account string, count, from uint32) (transaction []Transaction, err error) {
	r, err := b.client.call("listtransactions", []interface{}{account, count, from})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

// ListUnspent returns array of unspent transaction inputs in the wallet.
func (b *Navcoind) ListUnspent(minconf, maxconf uint32) (transactions []Transaction, err error) {
	if maxconf > 999999 {
		maxconf = 999999
	}
	r, err := b.client.call("listunspent", []interface{}{minconf, maxconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transactions)
	return
}

// UnspendableOutput represents a unspendable (locked) output
type UnspendableOutput struct {
	TxId string `json:"txid"`
	Vout uint64 `json:"vout"`
}

// ListLockUnspent returns list of temporarily unspendable outputs
func (b *Navcoind) ListLockUnspent() (unspendableOutputs []UnspendableOutput, err error) {
	r, err := b.client.call("listlockunspent", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &unspendableOutputs)
	return
}

// LockUnspent updates(lock/unlock) list of temporarily unspendable outputs
func (b *Navcoind) LockUnspent(lock bool, outputs []UnspendableOutput) (success bool, err error) {
	r, err := b.client.call("lockunspent", []interface{}{lock, outputs})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &success)
	return
}

// Move from one account in your wallet to another
func (b *Navcoind) Move(formAccount, toAccount string, amount float64, minconf uint32, comment string) (success bool, err error) {
	r, err := b.client.call("move", []interface{}{formAccount, toAccount, amount, minconf, comment})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &success)
	return

}

// SendFrom send amount from fromAccount to toAddress
//  amount is a real and is rounded to 8 decimal places.
//  Will send the given amount to the given address, ensuring the account has a valid balance using [minconf] confirmations.
func (b *Navcoind) SendFrom(fromAccount, toAddress string, amount float64, minconf uint32, comment, commentTo string) (txID string, err error) {
	r, err := b.client.call("sendfrom", []interface{}{fromAccount, toAddress, amount, minconf, comment, commentTo})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}

// SenMany send multiple times
func (b *Navcoind) SendMany(fromAccount string, amounts map[string]float64, minconf uint32, comment string) (txID string, err error) {
	r, err := b.client.call("sendmany", []interface{}{fromAccount, amounts, minconf, comment})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}

// SendToAddress send an amount to a given address
func (b *Navcoind) SendToAddress(toAddress string, amount float64, comment, commentTo string) (txID string, err error) {
	r, err := b.client.call("sendtoaddress", []interface{}{toAddress, amount, comment, commentTo})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}

// SetAccount sets the account associated with the given address
func (b *Navcoind) SetAccount(address, account string) error {
	r, err := b.client.call("setaccount", []interface{}{address, account})
	return handleError(err, &r)
}

// SetGenerate turns generation on or off.
// Generation is limited to [genproclimit] processors, -1 is unlimited.
func (b *Navcoind) SetGenerate(generate bool, genProcLimit int32) error {
	r, err := b.client.call("setgenerate", []interface{}{generate, genProcLimit})
	return handleError(err, &r)
}

// SetTxFee set the transaction fee per kB
func (b *Navcoind) SetTxFee(amount float64) error {
	r, err := b.client.call("settxfee", []interface{}{amount})
	return handleError(err, &r)
}

// Stop stop bitcoin server.
func (b *Navcoind) Stop() error {
	r, err := b.client.call("stop", nil)
	return handleError(err, &r)
}

// SignMessage sign a message with the private key of an address
func (b *Navcoind) SignMessage(address, message string) (sig string, err error) {
	r, err := b.client.call("signmessage", []interface{}{address, message})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &sig)
	return
}

// Verifymessage Verify a signed message.
func (b *Navcoind) VerifyMessage(address, sign, message string) (success bool, err error) {
	r, err := b.client.call("verifymessage", []interface{}{address, sign, message})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &success)
	return
}

// ValidateAddressResponse represents a response to "validateaddress" call
type ValidateAddressResponse struct {
	IsValid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	IsMine       bool   `json:"ismine"`
	IsScript     bool   `json:"isscript"`
	PubKey       string `json:"pubkey"`
	IsCompressed bool   `json:"iscompressed"`
	Account      string `json:"account"`
}

// ValidateAddress return information about <bitcoinaddress>.
func (b *Navcoind) ValidateAddress(address string) (va ValidateAddressResponse, err error) {
	r, err := b.client.call("validateaddress", []interface{}{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &va)
	return
}

// WalletLock Removes the wallet encryption key from memory, locking the wallet.
// After calling this method, you will need to call walletpassphrase again before being
// able to call any methods which require the wallet to be unlocked.
func (b *Navcoind) WalletLock() error {
	r, err := b.client.call("walletlock", nil)
	return handleError(err, &r)
}

// walletPassphrase stores the wallet decryption key in memory for <timeout> seconds.
func (b *Navcoind) WalletPassphrase(passPhrase string, timeout uint64) error {
	r, err := b.client.call("walletpassphrase", []interface{}{passPhrase, timeout})
	return handleError(err, &r)
}

func (b *Navcoind) WalletPassphraseChange(oldPassphrase, newPassprhase string) error {
	r, err := b.client.call("walletpassphrasechange", []interface{}{oldPassphrase, newPassprhase})
	return handleError(err, &r)
}

// GetProposal returns information about the proposal with the given hash.
func (b *Navcoind) GetProposal(hash string) (proposal Proposal, err error) {
	r, err := b.client.call("getproposal", []string{hash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &proposal)
	return
}

// CfundStats Returns statistics about the community fund.
func (b *Navcoind) CfundStats(hash string) (cfundStats CFundStats, err error) {
	r, err := b.client.call("cfundstats", []string{hash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &cfundStats)
	return
}

// GetPaymentRequest returns information about the proposal with the given hash.
func (b *Navcoind) GetPaymentRequest(hash string) (paymentRequest PaymentRequest, err error) {
	r, err := b.client.call("getpaymentrequest", []string{hash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &paymentRequest)
	return
}
