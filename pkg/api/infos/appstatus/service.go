package appstatus

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all AppStatus application services.
type Service interface {
	List() (rpi.AppStatus, error)
}

// AppStatus represents an AppStatus application service.
type AppStatus struct {
	apsfsys APSFSYS
	i       Infos
}

// APSFSYS represents an AppStatus repository service.
type APSFSYS interface {
	List(
		map[string]bool,
	) (rpi.AppStatus, error)
}

// Infos represents the infos interface
type Infos interface {
	StatusVPNWithOpenVPN(string, string) map[string]bool
}

// New creates a AppStatus application service instance.
func New(apsfsys APSFSYS, i Infos) *AppStatus {
	return &AppStatus{apsfsys: apsfsys, i: i}
}
