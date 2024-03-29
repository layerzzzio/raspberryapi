package rpi

// Let represents a current host net interface.
type Net struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Flags       []string `json:"flags"`
	IPv4        string   `json:"ipv4"`
	BytesSent   uint64   `json:"bytesSent,omitempty"`
	BytesRecv   uint64   `json:"bytesRecv,omitempty"`
	PacketsSent uint64   `json:"packetsSent,omitempty"`
	PacketsRecv uint64   `json:"packetsRecv,omitempty"`
}
