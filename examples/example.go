package main

import (
	"github.com/toorop/go-bitcoind"
	"log"
)

const (
	SERVER_HOST       = ""
	SERVER_PORT       = 8334
	USER              = "bitcoinrpc"
	PASSWD            = ""
	USESSL            = false
	WALLET_PASSPHRASE = ""
)

func main() {
	bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}

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

	// getbalance
	/*
		balance, err := bc.GetBalance("tests", 10)
		log.Println(err, balance)
	*/

	// getbestblockhash
	/*
		bestblockhash, err := bc.GetBestBlockhash()
		log.Println(err, bestblockhash)
	*/

	// getbestblockhash

	/*
		 bestBlockHash, err := bc.GetBestBlockhash()
		log.Println(err, bestBlockHash)
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
	difficulty, err := bc.GetDifficulty()
	log.Println(err, difficulty)

	// getinfo
	/*
		info, err := bc.GetInfo()
		log.Println(err, info)
	*/

	// getnewaddress
	/*
		newAddress, err := bc.GetNewAddress("tests", "toto")
		log.Println(err, newAddress)
	*/

}
