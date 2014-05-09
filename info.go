package bitcoind

type info struct {
	Version         uint32  `json:"version"`
	Protocolversion uint32  `json:"protocolversion"`
	Walletversion   uint32  `json:"walletversion"`
	Balance         float64 `json:"balance"`
	Blocks          uint32  `json:"blocks"`
	Timeoffset      uint32  `json:"timeoffset"`
	Connections     uint32  `json:"connections"`
	Proxy           string  `json:"proxy"`
	Difficulty      float64 `json:"difficulty"`
	Testnet         bool    `json:"testnet"`
	Keypoololdest   uint64  `json:"keypoololdest"`
	Paytxfee        float64 `json:"paytxfee"`
	Relayfee        float64 `json:"relayfee"`
	Errors          string  `json:"errors"`
}
