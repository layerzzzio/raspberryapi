package port

import (
	"github.com/raspibuddy/rpi"
)

// View returns a Port model.
func (p *Port) View(port int32) (rpi.Port, error) {
	isPortListening := p.i.IsPortListening(port)
	return p.psys.View(port, isPortListening)
}
