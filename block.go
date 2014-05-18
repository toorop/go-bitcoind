package bitcoind

// Represents a block
type block struct {
	Hash              string   `json:"hash"`
	Confirmations     uint64   `json:"confirmations"`
	Size              uint64   `json:"size"`
	Height            uint64   `json:"height"`
	Version           uint64   `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int64    `json:"time"`
	Nonce             uint64   `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Previousblockhash string   `json:"previousblockhash"`
}
