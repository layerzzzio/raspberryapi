package rpi

// SoftwareConfig represents software configs installed on the RPI
type SoftwareConfig struct {
	VPNCountries map[string][]string `json:"VPNCountries"`
}
