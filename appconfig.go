package rpi

// AppConfigVPNWithOvpn represents configs of VPN apps installed on the RPI
type AppConfigVPNWithOvpn struct {
	VPNCountries map[string]map[string]string `json:"VPNCountries"`
}
