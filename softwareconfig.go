package rpi

// SoftwareConfig represents software configs installed on the RPI
type SoftwareConfig struct {
	NordVPN NordVPN `json:"nordVpn"`
}

type NordVPN struct {
	TCPCountries []string `json:"tcpCountries "`
	UDPCountries []string `json:"udpCountries "`
}
