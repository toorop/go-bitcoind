package navcoind

type AddressBalance struct {
	Balance  uint64 `json:"balance"`
	Received uint64 `json:"received"`
}
