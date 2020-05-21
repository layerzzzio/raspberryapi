package rpi

// User represents the current host active virtual users.
type User struct {
	User     string `json:"user"`
	Terminal string `json:"terminal"`
	Started  int    `json:"started"`
}
