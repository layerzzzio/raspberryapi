package sys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// Process represents an empty process entity on the current system.
type Process struct{}

// List returns a list of disk stats
func (p Process) List(pinfo []metrics.PInfo) ([]rpi.ProcessSummary, error) {
	var result []rpi.ProcessSummary

	for i := range pinfo {
		result = append(
			result,
			rpi.ProcessSummary{
				ID:         pinfo[i].ID,
				Name:       pinfo[i].Name,
				CPUPercent: pinfo[i].CPUPercent,
				MemPercent: pinfo[i].MemPercent,
			},
		)
	}
	return result, nil
}

// View returns a disk stats
func (p Process) View(id int32, pinfo []metrics.PInfo) (rpi.ProcessDetails, error) {
	var result rpi.ProcessDetails

	for i := range pinfo {
		if id == pinfo[i].ID {
			result = rpi.ProcessDetails{
				ID:           pinfo[i].ID,
				Name:         pinfo[i].Name,
				Username:     pinfo[i].Username,
				CommandLine:  pinfo[i].CommandLine,
				Status:       pinfo[i].Status,
				CreationTime: pinfo[i].CreationTime,
				Foreground:   pinfo[i].Foreground,
				Background:   pinfo[i].Background,
				IsRunning:    pinfo[i].IsRunning,
				CPUPercent:   pinfo[i].CPUPercent,
				MemPercent:   pinfo[i].MemPercent,
				ParentP:      pinfo[i].ParentP,
			}
			break
		}
	}
	return result, nil
}
