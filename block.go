package bitcoind

// Represents a block
type Block struct {
	// The block hash
	Hash string `json:"hash"`

	// The number of confirmations
	Confirmations uint64 `json:"confirmations"`

	// The block size
	Size uint64 `json:"size"`

	// The block height or index
	Height uint64 `json:"height"`

	// The block version
	Version uint32 `json:"version"`

	// The merkle root
	Merkleroot string `json:"merkleroot"`

	// Slice on transaction ids
	Tx []string `json:"tx"`

	// The block time in seconds since epoch (Jan 1 1970 GMT)
	Time int64 `json:"time"`

	// The nonce
	Nonce uint64 `json:"nonce"`

	// The bits
	Bits string `json:"bits"`

	// The difficulty
	Difficulty float64 `json:"difficulty"`

	// Total amount of work in active chain, in hexadecimal
	Chainwork string `json:"chainwork,omitempty"`

	// The hash of the previous block
	Previousblockhash string `json:"previousblockhash"`

	// The hash of the next block
	Nextblockhash string `json:"nextblockhash"`
}
