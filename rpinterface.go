package rpi

// RpInterface represents an RPI interface
type RpInterface struct {
	IsStartXElf bool `json:"isStartXElf"`
	IsCamera    bool `json:"isCamera"`
	IsSSH       bool `json:"isSSH"`
}
