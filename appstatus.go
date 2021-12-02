package rpi

// AppStatus represents a current app status
type AppStatus struct {
	VPNwithOpenVPN map[string]bool `json:"vpnWithOpenVPN"`
}
