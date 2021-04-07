package rpi

// Software represents the software installed on the RPI
type Software struct {
	IsVNC          bool `json:"isVNC"`
	IsOpenVPN      bool `json:"isOpenVPN"`
	IsUnzip        bool `json:"isUnzip"`
	IsNordVPN      bool `json:"isNordVpn"`
	IsSurfSharkVPN bool `json:"isSurfSharkVPN"`
	IsIpVanishVPN  bool `json:"isIpVanishVPN"`
	IsVyprVpnVPN   bool `json:"isVyprVpnVPN"`
}
