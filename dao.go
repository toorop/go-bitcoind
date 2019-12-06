package navcoind

type Proposal struct {
	Version             uint32           `json:"version"`
	Hash                string           `json:"hash"`
	BlockHash           string           `json:"blockHash"`
	Description         string           `json:"description"`
	RequestedAmount     string           `json:"requestedAmount"`
	NotPaidYet          string           `json:"notPaidYet"`
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
	ProposalHash        string `json:"proposalHash"`
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
