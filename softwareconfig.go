package rpi

// SoftwareConfig represents software configs installed on the RPI
type SoftwareConfig struct {
	NordVPN NordVPN `json:"nordVpn"`
}

type NordVPN struct {
	TCPFiles []string `json:"tcpFiles"`
	UDPFiles []string `json:"udpFiles"`
}
