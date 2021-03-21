package rpi

// Software represents the software installed on the RPI
type Software struct {
	IsVNC     bool `json:"isVNC"`
	IsOpenVPN bool `json:"isOpenVPN"`
}
