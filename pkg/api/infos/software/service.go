package software

import (
	"github.com/raspibuddy/rpi"
)

// Service represents all Software application services.
type Service interface {
	List() (rpi.Software, error)
	View(string) (rpi.Software, error)
}

// Software represents an Software application service.
type Software struct {
	sofsys SOFSYS
	i      Infos
}

// INTSYS represents an Software repository service.
type SOFSYS interface {
	List(
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
	) (rpi.Software, error)

	View(bool) (rpi.Software, error)
}

// Infos represents the infos interface
type Infos interface {
	IsDPKGInstalled(string) bool
	HasDirectoryAtLeastOneFile(string, bool) bool
}

// New creates a Software application service instance.
func New(sofsys SOFSYS, i Infos) *Software {
	return &Software{sofsys: sofsys, i: i}
}
