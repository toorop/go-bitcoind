package bitcoind

// A Work represents a formatted hash data to work on
type Work struct {
	Midstate string `json:"midstate"`
	Data     string `json:"data"`
	Hash1    string `json:"hash1"`
	Target   string `json:"target"`
}
