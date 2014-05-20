package bitcoind

type Peer struct {
	// The ip address and port of the peer
	Addr string `json:"addr"`

	// Local address
	Addrlocal string `json:"addrlocal"`

	// The services
	Services string `json:"services"`

	// The time in seconds since epoch (Jan 1 1970 GMT) of the last send
	Lastsend uint64 `json:"lastsend"`

	// The time in seconds since epoch (Jan 1 1970 GMT) of the last receive
	Lastrecv uint64 `json:"lastrecv"`

	// The total bytes sent
	Bytessent uint64 `json:"bytessent"`

	// The total bytes received
	Bytesrecv uint64 `json:"bytesrecv"`

	// The connection time in seconds since epoch (Jan 1 1970 GMT)
	Conntime uint64 `json:"conntime"`

	// Ping time
	Pingtime float64 `json:"pingtime"`

	// Ping Wait
	Pingwait float64 `json:"pingwait"`

	// The peer version, such as 7001
	Version uint32 `json:"version"`

	// The string version
	Subver string `json:"subver"`

	// Inbound (true) or Outbound (false)
	Inbound bool `json:"inbound"`

	//  The starting height (block) of the peer
	Startingheight int32 `json:"startingheight"`

	// The ban score (stats.nMisbehavior)
	Banscore int32 `json:"banscore"`

	// If sync node
	Syncnode bool `json:"syncnode"`
}
