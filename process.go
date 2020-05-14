package rpi

import (
	"time"
)

// ProcessSummary is
type ProcessSummary struct {
	ID         int32   `json:"id"`
	Name       string  `json:"name"`
	CPUPercent float64 `json:"cpuPercent"`
	MemPercent float32 `json:"memPercent"`
}

// ProcessDetails is
type ProcessDetails struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	CommandLine  string    `json:"commandLine"`
	Status       string    `json:"status"`
	CreationTime time.Time `json:"creationTime"`
	Foreground   bool      `json:"foreground"`
	Background   bool      `json:"background"`
	IsRunning    bool      `json:"isRunning"`
	CPUPercent   float64   `json:"cpuPercent"`
	MemPercent   float32   `json:"memPercent"`
	ParentP      int32     `json:"parentPID"`
}
