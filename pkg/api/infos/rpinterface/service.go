package rpinterface

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all RpInterface application services.
type Service interface {
	List() (rpi.RpInterface, error)
}

// RpInterface represents an RpInterface application service.
type RpInterface struct {
	intsys INTSYS
	i      Infos
}

// INTSYS represents an RpInterface repository service.
type INTSYS interface {
	List(
		[]string,
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
		[]string,
		map[string]bool,
		map[string]string,
	) (rpi.RpInterface, error)
}

// Infos represents the infos interface
type Infos interface {
	ReadFile(string) ([]string, error)
	IsFileExists(string) bool
	GetConfigFiles() map[string]rpi.ConfigFileDetails
	IsQuietGrep(string, string, string) bool
	IsSSHKeyGenerating(string) bool
	IsDPKGInstalled(string) bool
	IsSPI(string) bool
	IsI2C(string) bool
	IsVariableSet([]string, string, string) bool
	ListWifiInterfaces(string) []string
	IsWpaSupCom() map[string]bool
	ZoneInfo(string) map[string]string
}

// New creates a RpInterface application service instance.
func New(intsys INTSYS, i Infos) *RpInterface {
	return &RpInterface{intsys: intsys, i: i}
}
