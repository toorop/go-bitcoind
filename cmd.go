// @Time    : 2017/10/17 下午11:27
// @Author  :
// @Site    :
// @File    : cmd.go
// @Software: Gogland
package bitcoind

// NewCreateMultisigCmd returns a new instance which can be used to issue a
// createmultisig JSON-RPC command.
func NewCreateMultisigCmd(nRequired int, keys []string) *CreateMultisigCmd {
	return &CreateMultisigCmd{
		NRequired: nRequired,
		Keys:      keys,
	}
}

// NewAddMultisigAddressCmd returns a new instance which can be used to issue a
// addmultisigaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewAddMultisigAddressCmd(nRequired int, keys []string, account *string) *AddMultisigAddressCmd {
	return &AddMultisigAddressCmd{
		NRequired: nRequired,
		Keys:      keys,
		Account:   account,
	}
}

// NewSignRawTransactionCmd returns a new instance which can be used to issue a
// signrawtransaction JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewSignRawTransactionCmd(hexEncodedTx string, inputs *[]RawTxInput, privKeys *[]string, flags *string) *SignRawTransactionCmd {
	return &SignRawTransactionCmd{
		RawTx:    hexEncodedTx,
		Inputs:   inputs,
		PrivKeys: privKeys,
		Flags:    flags,
	}
}

// NewSendRawTransactionCmd returns a new instance which can be used to issue a
// sendrawtransaction JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewSendRawTransactionCmd(hexTx string, allowHighFees *bool) *SendRawTransactionCmd {
	return &SendRawTransactionCmd{
		HexTx:         hexTx,
		AllowHighFees: allowHighFees,
	}
}
