package rpi

// User represents the virtual users connected to the system
type User struct {
	User     string `json:"user"`
	Terminal string `json:"terminal"`
	Started  int    `json:"started"`
}
