package navcoind

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vin struct {
	Coinbase  string    `json:"coinbase,omitempty"`
	Txid      string    `json:"txid,omitempty"`
	Vout      int       `json:"vout,omitempty"`
	ScriptSig ScriptSig `json:"scriptSig,omitempty"`
	Value     float64   `json:"value,omitempty"`
	ValueSat  uint64    `json:"valuesat,omitempty"`
	Address   string    `json:"address,omitempty"`
	Sequence  uint32    `json:"sequence"`
}

type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs,omitempty"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses,omitempty"`
	Hash      string   `json:"hash,omitempty"`
}

type Vout struct {
	Value        float64      `json:"value"`
	ValueSat     uint64       `json:"valuesat"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type RawTransaction struct {
	Hex             string `json:"hex"`
	Txid            string `json:"txid"`
	Hash            string `json:"hash"`
	Size            uint64 `json:"size"`
	VSize           uint64 `json:"vsize"`
	Version         uint32 `json:"version"`
	LockTime        uint32 `json:"locktime"`
	AnonDestination string `json:"anon-destination"`
	Vin             []Vin  `json:"vin"`
	Vout            []Vout `json:"vout"`
	BlockHash       string `json:"blockhash,omitempty"`
	Height          uint64 `json:"height"`
	Confirmations   uint64 `json:"confirmations,omitempty"`
	Time            int64  `json:"time,omitempty"`
	BlockTime       int64  `json:"blocktime,omitempty"`
}

type TransactionDetails struct {
	Account  string  `json:"account"`
	Address  string  `json:"address,omitempty"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Fee      float64 `json:"fee,omitempty"`
}

type Transaction struct {
	Amount          float64              `json:"amount"`
	Account         string               `json:"account,omitempty"`
	Address         string               `json:"address,omitempty"`
	Category        string               `json:"category,omitempty"`
	Fee             float64              `json:"fee,omitempty"`
	Confirmations   int64                `json:"confirmations"`
	BlockHash       string               `json:"blockhash"`
	BlockIndex      int64                `json:"blockindex"`
	BlockTime       int64                `json:"blocktime"`
	TxID            string               `json:"txid"`
	WalletConflicts []string             `json:"walletconflicts"`
	Time            int64                `json:"time"`
	TimeReceived    int64                `json:"timereceived"`
	Details         []TransactionDetails `json:"details,omitempty"`
	Hex             string               `json:"hex,omitempty"`
}

type UTransactionOut struct {
	Bestblock     string       `json:"bestblock"`
	Confirmations uint32       `json:"confirmations"`
	Value         float64      `json:"value"`
	ScriptPubKey  ScriptPubKey `json:"scriptPubKey"`
	Version       uint32       `json:"version"`
	Coinbase      bool         `json:"coinbase"`
}

type TransactionOutSet struct {
	Height          uint32  `json:"height"`
	Bestblock       string  `json:"bestblock"`
	Transactions    float64 `json:"transactions"`
	TxOuts          float64 `json:"txouts"`
	BytesSerialized float64 `json:"bytes_serialized"`
	HashSerialized  string  `json:"hash_serialized"`
	TotalAmount     float64 `json:"total_amount"`
}
