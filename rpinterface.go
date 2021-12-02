package rpi

// RpInterface represents an RPI interface
type RpInterface struct {
	IsStartXElf        bool              `json:"isStartXElf"`
	IsCamera           bool              `json:"isCamera"`
	IsSSH              bool              `json:"isSSH"`
	IsSSHKeyGenerating bool              `json:"isSSHKeyGenerating"`
	IsVNC              bool              `json:"isVNC"`
	IsVNCInstalled     bool              `json:"isVNCInstalled"`
	IsSPI              bool              `json:"isSPI"`
	IsI2C              bool              `json:"isI2C"`
	IsOneWire          bool              `json:"isOneWire"`
	IsRemoteGpio       bool              `json:"isRemoteGpio"`
	IsWifiInterfaces   bool              `json:"isWifiInterfaces"`
	IsWpaSupCom        map[string]bool   `json:"isWpaSupCom"`
	ZoneInfo           map[string]string `json:"zoneInfo"`
}
