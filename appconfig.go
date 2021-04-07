package rpi

// AppConfig represents configs of apps installed on the RPI
type AppConfig struct {
	VPNCountries map[string][]string `json:"VPNCountries"`
}
