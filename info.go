package navcoind

type Info struct {
	Version         uint32  `json:"version"`
	Protocolversion uint32  `json:"protocolversion"`
	Walletversion   uint32  `json:"walletversion"`
	Balance         float64 `json:"balance"`
	Blocks          uint32  `json:"blocks"`
	Timeoffset      int32   `json:"timeoffset"`
	Connections     uint32  `json:"connections"`
	Proxy           string  `json:"proxy,omitempty"`
	Difficulty      float64 `json:"difficulty"`
	Testnet         bool    `json:"testnet"`
	Keypoololdest   uint64  `json:"keypoololdest"`
	KeypoolSize     uint32  `json:"keypoolsize,omitempty"`
	UnlockedUntil   int64   `json:"unlocked_until,omitempty"`
	Paytxfee        float64 `json:"paytxfee"`
	Relayfee        float64 `json:"relayfee"`
	Errors          string  `json:"errors"`
}

type BlockchainInfo struct {
	Chain                string                  `json:"chain"`
	Blocks               uint64                  `json:"blocks"`
	Headers              uint64                  `json:"headers"`
	BestBlockHash        string                  `json:"bestblockhash"`
	Difficulty           float64                 `json:"difficulty"`
	MedianTime           uint64                  `json:"mediantime"`
	VerificationProgress float64                 `json:"verificationprocess"`
	ChainWork            string                  `json:"chainwork"`
	Pruned               bool                    `json:"pruned"`
	SoftForks            []SoftFork              `json:"softforks"`
	Bip9SoftForks        map[string]Bip9SoftFork `json:"bip9_softforks"`
}

type SoftFork struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
	Enforce struct {
		Status   bool `json:"status"`
		Found    uint `json:"found"`
		Required uint `json:"required"`
		Window   uint `json:"window"`
	}
	Reject struct {
		Status   bool `json:"status"`
		Found    uint `json:"found"`
		Required uint `json:"required"`
		Window   uint `json:"window"`
	}
}

type Bip9SoftFork struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	Bit       uint   `json:"bit"`
	StartTime uint64 `json:"starttime"`
	Timeout   uint64 `json:"timeout"`
}
