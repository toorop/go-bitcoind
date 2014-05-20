package bitcoind

// A MiningInfo represents a mininginfo response
type MiningInfo struct {
	// The current block
	Blocks uint64 `json:"blocks"`

	// The last block size
	CurrentBlocksize uint64 `json:"currentblocksize"`

	// The last block transaction
	CurrentBlockTx uint64 `json:"currentblocktx"`

	// The current difficulty
	Difficulty float64 `json:"difficulty"`

	// Current errors
	Errors string `json:"errors"`

	// The processor limit for generation. -1 if no generation. (see getgenerate or setgenerate calls)
	GenProcLimit int32 `json:"genproclimit"`

	// The size of the mem pool
	PooledtTx uint64 `json:"pooledtx"`

	// If using testnet or not
	Testnet bool `json:"testnet"`

	// If the generation is on or off (see getgenerate or setgenerate calls)
	Generate bool `json:"generate"`

	// The network hashrate
	NetworkHashps uint64 `json:"networkhashps"`

	// Node hashrate
	HashesPersec uint64 `json:"hashespersec"`
}
