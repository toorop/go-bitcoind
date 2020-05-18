package navcoind

type Proposal struct {
	Version             uint32           `json:"version"`
	Hash                string           `json:"hash"`
	BlockHash           string           `json:"blockHash"`
	Description         string           `json:"description"`
	RequestedAmount     string           `json:"requestedAmount"`
	NotPaidYet          string           `json:"notPaidYet"`
	NotRequestedYet     string           `json:"notRequestedYet"`
	UserPaidFee         string           `json:"userPaidFee"`
	PaymentAddress      string           `json:"paymentAddress"`
	ProposalDuration    uint64           `json:"proposalDuration"`
	ExpiresOn           uint64           `json:"expiresOn"`
	VotesYes            uint             `json:"votesYes"`
	VotesAbs            uint             `json:"votesAbs"`
	VotesNo             uint             `json:"votesNo"`
	VotingCycle         uint             `json:"votingCycle"`
	Status              string           `json:"status"`
	State               uint             `json:"state"`
	StateChangedOnBlock string           `json:"stateChangedOnBlock,omitempty"`
	PaymentRequests     []PaymentRequest `json:"paymentRequests"`
}

type PaymentRequest struct {
	Version             uint32 `json:"version"`
	Hash                string `json:"hash"`
	BlockHash           string `json:"blockHash"`
	ProposalHash        string `json:"proposalHash,omitempty"`
	Description         string `json:"description"`
	RequestedAmount     string `json:"requestedAmount"`
	VotesYes            uint   `json:"votesYes"`
	VotesAbs            uint   `json:"votesAbs"`
	VotesNo             uint   `json:"votesNo"`
	VotingCycle         uint   `json:"votingCycle"`
	Status              string `json:"status"`
	State               uint   `json:"state"`
	StateChangedOnBlock string `json:"stateChangedOnBlock,omitempty"`
	PaidOnBlock         string `json:"paidOnBlock,omitempty"`
}

type CFundStats struct {
	Consensus Consensus `json:"consensus"`
}

type Consensus struct {
	BlocksPerVotingCycle                uint    `json:"blocksPerVotingCycle"`
	MinSumVotesPerVotingCycle           uint    `json:"minSumVotesPerVotingCycle"`
	MaxCountVotingCycleProposals        uint    `json:"maxCountVotingCycleProposals"`
	MaxCountVotingCyclePaymentRequests  uint    `json:"maxCountVotingCyclePaymentRequests"`
	VotesAcceptProposalPercentage       uint    `json:"votesAcceptProposalPercentage"`
	VotesRejectProposalPercentage       uint    `json:"votesRejectProposalPercentage"`
	VotesAcceptPaymentRequestPercentage uint    `json:"votesAcceptPaymentRequestPercentage"`
	VotesRejectPaymentRequestPercentage uint    `json:"votesRejectPaymentRequestPercentage"`
	ProposalMinimalFee                  float64 `json:"proposalMinimalFee"`
}

type ConsensusParameter struct {
	Id          int    `json:"id"`
	Description string `json:"desc"`
	Type        int    `json:"type"`
	Value       int    `json:"value"`
}

type Consultation struct {
	Version                  uint32         `json:"version"`
	Hash                     string         `json:"hash"`
	BlockHash                string         `json:"blockhash"`
	Question                 string         `json:"question"`
	Support                  int            `json:"support"`
	Abstain                  int            `json:"abstain"`
	Answers                  []*Answer      `json:"answers"`
	Min                      int            `json:"min"`
	Max                      int            `json:"max"`
	VotingCyclesFromCreation int            `json:"votingCyclesFromCreation"`
	VotingCycleForState      Cycle          `json:"votingCycleForState"`
	Status                   string         `json:"status"`
	State                    int            `json:"state"`
	StateChangedOnBlock      string         `json:"stateChangedOnBlock"`
	MapState                 map[int]string `json:"mapState"`
}

type Cycle struct {
	Current int `json:"current"`
	Max     int `json:"max"`
}

type Answer struct {
	Version             uint32         `json:"version"`
	Answer              string         `json:"answer"`
	Support             int            `json:"support"`
	Votes               int            `json:"votes"`
	Status              string         `json:"status"`
	State               int            `json:"state"`
	StateChangedOnBlock string         `json:"stateChangedOnBlock"`
	TxBlockHash         string         `json:"txblockhash"`
	Parent              string         `json:"parent"`
	Hash                string         `json:"hash"`
	MapState            map[int]string `json:"mapState"`
}
