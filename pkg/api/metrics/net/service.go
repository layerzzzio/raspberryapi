package net

import (
	"github.com/raspibuddy/rpi"
	"github.com/shirou/gopsutil/net"
)

// Service represents all Net application services.
type Service interface {
	List() ([]rpi.Net, error)
	View(int) (rpi.Net, error)
}

// Net represents a Net application service.
type Net struct {
	nsys NSYS
	mt   Metrics
}

// NSYS represents a Net repository service.
type NSYS interface {
	List([]net.InterfaceStat) ([]rpi.Net, error)
	View(int, []net.InterfaceStat, []net.IOCountersStat) (rpi.Net, error)
}

// Metrics represents the system metrics interface
type Metrics interface {
	NetInfo() ([]net.InterfaceStat, error)
	NetStats() ([]net.IOCountersStat, error)
}

// New creates a Net application service instance.
func New(nsys NSYS, mt Metrics) *Net {
	return &Net{nsys: nsys, mt: mt}
}
