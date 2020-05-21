package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/net"
)

// Net mock
type Net struct {
	ListFn func([]net.InterfaceStat) ([]rpi.Net, error)
	ViewFn func(int, []net.InterfaceStat, []net.IOCountersStat) (rpi.Net, error)
}

// List mock
func (n Net) List(netInfo []net.InterfaceStat) ([]rpi.Net, error) {
	return n.ListFn(netInfo)
}

// View mock
func (n Net) View(id int, netInfo []net.InterfaceStat, netStats []net.IOCountersStat) (rpi.Net, error) {
	return n.ViewFn(id, netInfo, netStats)
}
