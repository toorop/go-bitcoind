package navcoind

import "time"

type AddressBalance struct {
	Balance  uint64 `json:"balance"`
	Received uint64 `json:"received"`
}

type AddressHistoryRequest struct {
	Addresses []string `json:"addresses"`
	Start     uint64   `json:"start"`
	End       uint64   `json:"end"`
}

type AddressHistory struct {
	Block   string    `json:"block"`
	TxIndex uint      `json:"txindex"`
	Time    time.Time `json:"time"`
	TxId    string    `json:"txid"`
	Address string    `json:"address"`
	Changes struct {
		Balance      int64 `json:"balance"`
		Stakable     int64 `json:"stakable"`
		VotingWeight int64 `json:"voting_weight"`
		Flags        uint  `json:"flags"`
	}
	Result struct {
		Balance      int64 `json:"balance"`
		Stakable     int64 `json:"stakable"`
		VotingWeight int64 `json:"voting_weight"`
	}
}
