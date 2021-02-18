package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// HumanUser mock
type HumanUser struct {
	ListFn func(lines []string) ([]rpi.HumanUser, error)
}

// List mock
func (hu HumanUser) List(lines []string) ([]rpi.HumanUser, error) {
	return hu.ListFn(lines)
}
