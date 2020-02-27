package navcoind

type Block struct {
	Hash                       string   `json:"hash"`
	Confirmations              uint64   `json:"confirmations"`
	StrippedSize               uint64   `json:"strippedsize"`
	Size                       uint64   `json:"size"`
	Weight                     uint64   `json:"weight"`
	Height                     uint64   `json:"height"`
	Version                    uint32   `json:"version"`
	VersionHex                 string   `json:"versionHex"`
	MerkleRoot                 string   `json:"merkleroot"`
	Tx                         []string `json:"tx"`
	TxCount                    uint     `json:"tx_count"`
	ProposalCount              uint     `json:"proposal_count"`
	PaymentRequestCount        uint     `json:"payment_request_count"`
	ProposalVotesCount         uint     `json:"proposal_votes_count"`
	PaymentRequestVotesCount   uint     `json:"payment_request_votes_count"`
	PaymentRequestPayoutsCount uint     `json:"payment_request_payouts_count"`
	Time                       int64    `json:"time"`
	MedianTime                 int64    `json:"mediantime"`
	Nonce                      uint64   `json:"nonce"`
	Bits                       string   `json:"bits"`
	Difficulty                 float64  `json:"difficulty"`
	ChainWork                  string   `json:"chainwork,omitempty"`
	PreviousBlockHash          string   `json:"previousblockhash"`
	NextBlockHash              string   `json:"nextblockhash"`
}

type BlockHeader struct {
	Hash              string      `json:"hash"`
	Confirmations     uint64      `json:"confirmations"`
	Height            uint64      `json:"height"`
	Version           uint32      `json:"version"`
	VersionHex        string      `json:"versionHex"`
	MerkleRoot        string      `json:"merkleroot"`
	Time              int64       `json:"time"`
	MedianTime        int64       `json:"mediantime"`
	Mint              float64     `json:"mint"`
	Nonce             uint64      `json:"nonce"`
	Bits              string      `json:"bits"`
	Difficulty        float64     `json:"difficulty"`
	ChainWork         string      `json:"chainwork"`
	NcfSupply         string      `json:"ncfsupply"`
	NcfLocked         string      `json:"ncflocked"`
	Flags             string      `json:"flags"`
	ProofHash         string      `json:"proofhash"`
	EntropyBit        int64       `json:"entropybit"`
	Modifier          string      `json:"modifier"`
	CfundVotes        []CfundVote `json:"cfund_votes"`
	PreviousBlockHash string      `json:"previousblockhash"`
	NextBlockHash     string      `json:"nextblockhash"`
}

type CfundVote struct {
	Hash string `json:"hash"`
	Vote int    `json:"vote"`
}
