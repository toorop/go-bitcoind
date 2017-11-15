// @Time    : 2017/10/14 上午1:27
// @Author  :
// @Site    :
// @File    : address.go
// @Software: Gogland
package bitcoind

// CreateMultisigCmd defines the createmultisig JSON-RPC command.
type CreateMultisigCmd struct {
	NRequired int
	Keys      []string
}

// AddMultisigAddressCmd defines the addmutisigaddress JSON-RPC command.
type AddMultisigAddressCmd struct {
	NRequired int
	Keys      []string
	Account   *string
}
