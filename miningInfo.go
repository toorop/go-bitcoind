package bitcoind

type MiningInfo struct {
	Block            uint64  `json:"block"`
	CurrentBlocksize uint64  `json:"currentblocksize"`
	CurrentBlockTx   uint64  `json:"currentblocktx"`
	Difficulty       float64 `json:"difficulty"`
	Errors           string  `json:"errors"`
	GenProcLimit     int32   `json:"genproclimit"`
	NetworkHashps    uint64  `json:"networkhashps"`
	PooledtTx        uint64  `json:"pooledtx"`
	Testnet          bool    `json:"testnet"`
	Generate         bool    `json:"generate"`
	HashesPersec     uint64  `json:"hashespersec"`
}
