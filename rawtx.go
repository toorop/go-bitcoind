package bitcoind

import (
	"fmt"
	"strings"
)

type SignRawTransactionResponse struct {
	Hex      string                    `json:"hex"`
	Complete bool                      `json:"complete"`
	Errors   []SignRawTransactionError `json:"errors,omitempty"`
}

func (r *SignRawTransactionResponse) HasError() bool {
	return len(r.Errors) > 0
}

func (r *SignRawTransactionResponse) Error() string {
	errors := make([]string, 0, 0)
	for _, e := range r.Errors {
		errors = append(errors, e.Message)
	}
	return strings.Join(errors, ", ")
}

type SignRawTransactionError struct {
	TxID      string `json:"txid"`
	Vout      int    `json:"vout"`
	ScriptSig string `json:"scriptSig"`
	Sequence  int64  `json:"sequence"`
	Message   string `json:"error"`
}

func (e *SignRawTransactionError) Error() string {
	return fmt.Sprintf("%s:%d - %s", e.TxID, e.Vout, e.Message)
}

type RawTransactionInput struct {
	TxID string `json:"txid"`
	Vout int    `json:"vout"`
}

type RawTransactionOutput map[string]float64
