package main

import (
	"github.com/toorop/go-bitcoind"
	"log"
)

const (
	SERVER_HOST        = ""
	SERVER_PORT        = 8334
	USER               = "bitcoinrpc"
	PASSWD             = "sss"
	USESSL             = false
	WALLET_PASSPHRASE  = "p1"
	WALLET_PASSPHRASE2 = "p2"
)

func main() {
	bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}

	//walletpassphrase
	err = bc.WalletPassphrase(WALLET_PASSPHRASE, 3600)
	log.Println(err)

	// backupwallet
	/*
		err = bc.BackupWallet("/tmp/wallet.dat")
		log.Println(err)
	*/

	// dumpprivkey
	/*
		privKey, err := bc.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
		log.Println(err, privKey)
	*/

	// encryptwallet
	/*
		err = bc.EncryptWallet(WALLET_PASSPHRASE)
		log.Println(err)
	*/

	// getaccount
	/*
		account, err := bc.GetAccount("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
		log.Println(err, account)
	*/

	//getaccountaddress
	/*
		address, err := bc.GetAccountAddress("tests")
		log.Println(err, address)
	*/
	// GetAddressesByAccount
	/*
		addresses, err := bc.GetAddressesByAccount("tests")
		log.Println(err, addresses)
	*/

	// getaddressbyaccount
	/*
		address, err := bc.GetAddressesByAccount("tests")
		log.Println(err, address)
	*/

	// getbalance
	/*
		balance, err := bc.GetBalance("tests", 0)
		log.Println(err, balance)
	*/

	// getbestblockhash
	/*
		bestblockhash, err := bc.GetBestBlockhash()
		log.Println(err, bestblockhash)
	*/

	// getblock
	/*
		block, err := bc.GetBlock("00000000000000003f8d1861d035e44d4297c49bd2517dc0a44ad73c7091926c")
		log.Println(err, block)
	*/

	// getblockcount
	/*
		count, err := bc.GetBlockCount()
		log.Println(err, count)
	*/

	// getblockhash
	/*
		hash, err := bc.GetBlockHash(0)
		log.Println(err, hash)
	*/
	// TODO a finir
	// getBlockTemplate
	/*
		template, err := bc.GetBlockTemplate([]string{"longpoll"}, "template")
		log.Println(err, template)
	*/

	// getconnectioncount
	/*
		count, err := bc.GetConnectionCount()
		log.Println(err, count)
	*/
	// getdifficulty
	/*
		difficulty, err := bc.GetDifficulty()
		log.Println(err, difficulty)
	*/

	// getgenerate
	/*
		generate, err := bc.GetGenerate()
		log.Println(err, generate)
	*/

	// gethashespersec
	/*
		hashpersec, err := bc.GetHashesPerSec()
		log.Println(hashpersec, err)
	*/

	// getinfo
	/*
		info, err := bc.GetInfo()

		log.Println(err, info)
	*/

	// getmininginfo
	/*
		mininginfo, err := bc.GetMiningInfo()
		log.Println(err, mininginfo)
	*/

	// getnewaddress
	/*
		newAddress, err := bc.GetNewAddress("test2")
		log.Println(err, newAddress)
	*/

	// getpeerinfo
	/*
		peerInfo, err := bc.GetPeerInfo()
		log.Println(err, peerInfo)
	*/

	// GetRawChangeAddress
	/*
		rawAddress, err := bc.GetRawChangeAddress("tests")
		log.Println(err, rawAddress)
	*/

	// GetRawMempool
	/*
		txIds, err := bc.GetRawMempool()
		log.Println(err, txIds)
	*/

	// getrawtransaction
	/*
		 rawTransaction, err := bc.GetRawTransaction("00010589f7c108a4fd546df03a17bf485ede3baf52b35ddd5b83e974ec360abf", 1)
		log.Println(err, rawTransaction)
	*/

	// GetReceivedByAccount
	/*
		amount, err := bc.GetReceivedByAccount("all", 1)
		log.Println(err, amount)
	*/

	// GetReceivedByAddress
	/*
		amount, err := bc.GetReceivedByAddress("1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy", 0)
		log.Println(err, amount)
	*/

	// GetTransaction
	/*
		transaction, err := bc.GetTransaction("a1b7093d041bc1b763ba1ad894d2bd5376b38e6c7369613684e7140e8d9f7515")
		log.Println(err, transaction)
	*/

	// GetTxOut
	/*
		transaction, err := bc.GetTxOut("a1b7093d041bc1b763ba1ad894d2bd5376b38e6c7369613684e7140e8d9f7515", 1, false)
		log.Println(err, transaction)
	*/

	// GetTxOutsetInfo
	/*
		tx, err := bc.GetTxOutsetInfo()
		log.Println(err, tx)
	*/

	// GetWork
	/*
		work, err := bc.GetWork("00000002439aa25207d0b0d91162ecaad3e306d7df6a514d0f1c372c0000000000000000ea834758ea6ed14b00f9951befe6a6c819dcf9f32c20b37ba1308e66305fca4e537cf6b4187c305300000000000000800000000000000000000000000000000000000000000000000000000000000000000000000000000080020000")
		//work, err := bc.GetWork()
		log.Println(err, work)
	*/

	// ImportPrivKey
	/*
		err = bc.ImportPrivKey("KFAKEZmymFAKECuxFAKECmFAKEcFAKEjxFAKEUoFAKE85FAKEi9", "imported from space", false)
		log.Println(err)
	*/

	// KeyPoolRefill
	/*
		err = bc.KeyPoolRefill()
		log.Println(err)
	*/

	// ListAccount
	/*
		accounts, err := bc.ListAccounts(1)
		log.Println(err, accounts)
	*/

	// ListAddressGroupings
	/*
		list, err := bc.ListAddressGroupings()
		log.Println(err, list)
	*/

	// ListReceivedByAccount
	/*
		list, err := bc.ListReceivedByAccount(1, true)
		log.Println(err, list)
	*/
	// ListReceivedByAddress
	/*
		list, err := bc.ListReceivedByAddress(1, true)
		log.Println(err, list)
	*/

	// ListSinceBlock
	/*
		 transactions, err := bc.ListSinceBlock("00000000000000003f8d1861d035e44d4297c49bd2517dc0a44ad73c7091926c", 1)
		log.Println(err, transactions)
	*/

	// ListTransactions
	/*
		transactions, err := bc.ListTransactions("tests", 10, 0)
		log.Println(err, transactions)
	*/

	// ListUnspent
	/*
		transactions, err := bc.ListUnspent(1, 9999)
		log.Println(err, transactions)
	*/

	// LockUnspent
	/*
		success, err := bc.LockUnspent(false, []bitcoind.UnspendableOutput{{"61195c9a04eb4bb6ef7c1d360e472b1620c4befed611ddcab46a6b2711344cd5", 0}, {"a1b7093d041bc1b763ba1ad894d2bd5376b38e6c7369613684e7140e8d9f7515", 0}})
		log.Println(err, success)
	*/

	// ListLockUnspent
	/*
		unspentLocked, err := bc.ListLockUnspent()
		log.Println(err, unspentLocked)
	*/

	// Move
	/*
		success, err := bc.Move("tests1", "test2", 0.0001, 1, "Move test")
		log.Println(err, success)
	*/

	// senfrom
	/*
		txID, err := bc.SendFrom("test2", "1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB", 0.0001, 1, "Comment: Envoit depuis test2", "commentTo: envoit depuis test2")
		log.Println(err, txID)
	*/

	// sendmany
	/*
		amounts := make(map[string]float64)
		amounts["1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB"] = 0.0001
		amounts["1Ldfez73eanxUZhudrS62BXqk8BrLxYQFj"] = 0.0001
		txID, err := bc.SendMany("tests", amounts, 1, "test sendMany")
		log.Println(err, txID)
	*/

	// sendtoaddress
	/*
		txID, err := bc.SendToAddress("1Ldfez73eanxUZhudrS62BXqk8BrLxYQFj", 0.0001, "send to address test", "send to cx")
		log.Println(err, txID)
	*/

	// setaccount
	/*
		err = bc.SetAccount("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3", "tests")
		log.Println(err)
	*/

	// setgenerate
	/*
		err = bc.SetGenerate(false, 1)
		log.Println(err)
	*/

	// settxfee
	/*
		err = bc.SetTxFee(0.00001)
		log.Println(err)
	*/

	// SignMessage
	/*
		sig, err := bc.SignMessage("1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy", "test message")
		log.Println(err, sig)
	*/

	// stop
	/*
		err = bc.Stop()
		log.Println(err)
	*/

	// ValidateAddress
	/*
		va, err := bc.ValidateAddress("1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB")
		log.Println(err, va)
	*/

	// verifymessage
	/*
		err = bc.VerifyMessage("1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy", "IDIM8ITWfjyf/e9jaY3+P1ErTDRVdpD3YKmTlVlhuBhwZ/qgnLf+M0GTSXgFWCv4dL3rV+x5ymZJagqy19iRFg4=", "test message")
		log.Println(err)
	*/

	// walletlock
	/*
		err = bc.WalletLock()
		log.Println(err)
	*/

	// WalletPassphraseChange
	/*
		err = bc.WalletPassphraseChange(WALLET_PASSPHRASE2, WALLET_PASSPHRASE)
		log.Println(err)
	*/
}
