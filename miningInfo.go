package bitcoind

type MiningInfo struct {
	// The current block
	Blocks uint64 `json:"block"`
	// The last block size
	CurrentBlocksize uint64  `json:"currentblocksize"`
	CurrentBlockTx   uint64  `json:"currentblocktx"` // The last block transaction
	Difficulty       float64 `json:"difficulty"`
	Errors           string  `json:"errors"`
	GenProcLimit     int32   `json:"genproclimit"`
	NetworkHashps    uint64  `json:"networkhashps"`
	PooledtTx        uint64  `json:"pooledtx"`
	Testnet          bool    `json:"testnet"`
	Generate         bool    `json:"generate"`
	HashesPersec     uint64  `json:"hashespersec"`
}
