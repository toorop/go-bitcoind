package main

import (
	"fmt"
	"github.com/toorop/go-bitcoind"
)

const (
	SERVER_HOST = ""
	SERVER_PORT = 8334
	USER        = "bitcoinrpc"
	PASSWD      = ""
	USESSL      = false
)

func main() {
	bc := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	// getinfo
	/*
		info, err := bc.GetInfo()
		fmt.Println(err, info.Difficulty)
	*/

	// getnewaddress
	/*
		newAddress, err := bc.GetNewAddress("tests", "toto")
		fmt.Println(err, newAddress)
	*/

	// GetAddressesByAccount
	addresses, err := bc.GetAddressesByAccount("tests")
	fmt.Println(err, addresses)
}
