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
	Version             uint32    `json:"version"`
	Hash                string    `json:"hash"`
	BlockHash           string    `json:"blockhash"`
	Question            string    `json:"question"`
	Support             int       `json:"support"`
	Abstain             int       `json:"abstain"`
	Answers             []*Answer `json:"answers"`
	Min                 int       `json:"min"`
	Max                 int       `json:"max"`
	VotingCycle         int       `json:"votingCycle"`
	Status              string    `json:"status"`
	State               int       `json:"state"`
	StateChangedOnBlock string    `json:"stateChangedOnBlock"`
}

type Answer struct {
	Version uint32 `json:"version"`
	String  string `json:"string"`
	Support int    `json:"support,omitempty"`
	Votes   int    `json:"votes,omitempty"`
	Status  string `json:"status,omitempty"`
	State   int    `json:"state,omitempty"`
	Parent  string `json:"parent,omitempty"`
	Hash    string `json:"hash,omitempty"`
}
