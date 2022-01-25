package rpi

// Software represents the software installed on the RPI
type Software struct {
	IsVNCInstalled              bool `json:"isVNCInstalled"`
	IsOpenVPNInstalled          bool `json:"isOpenVPNInstalled"`
	IsUnzipInstalled            bool `json:"isUnzipInstalled"`
	IsNordVPNInstalled          bool `json:"isNordVPNInstalled"`
	IsSurfSharkVPNInstalled     bool `json:"isSurfSharkVPNInstalled"`
	IsIpVanishVPNInstalled      bool `json:"isIpVanishVPNInstalled"`
	IsVyprVpnVPNInstalled       bool `json:"isVyprVpnVPNInstalled"`
	IsSpecificSoftwareInstalled bool `json:"isSpecificSoftwareInstalled"`
}
