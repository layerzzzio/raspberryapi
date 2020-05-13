package rpi

import (
	"time"
)

// Process is
type Process struct {
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
	CPUTimes     CPUStats  `json:"cpuTimes"`
	Threads      int32     `json:"threads"`
	MemPercent   float32   `json:"memPercent"`
	MemInfo      Mem       `json:"memInfo"`
	ParentP      int32     `json:"parentPID"`
	ChildrenP    []int32   `json:"childrenPID"`
}
