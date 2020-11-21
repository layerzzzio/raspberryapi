package sys

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Destroy represents an empty Destroy entity on the current system.
type Destroy struct{}

// ExecuteDF returns an action response
func (des Destroy) ExecuteDF(execs map[int]rpi.Exec) (rpi.Action, error) {
	// redefine the steps for this actions
	steps := map[int]string{
		1: actions.DeleteFile,
	}

	// and make sure no steps are forgoten
	if len(steps) != len(execs) {
		return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "length steps & execs are different")
	}

	// check if the first key equals 1
	for key := range execs {
		if key != 1 {
			return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "first key in execs map is not equal 1")
		}
		break
	}

	// the exit status is the status of the last exec
	return rpi.Action{
		Name:          actions.DeleteFile,
		Steps:         steps,
		NumberOfSteps: uint16(len(steps)),
		Executions:    execs,
		ExitStatus:    execs[len(execs)-1].ExitStatus,
		StartTime:     execs[1].StartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteDU returns an action response
func (des Destroy) ExecuteDU(execs map[int]rpi.Exec) (rpi.Action, error) {
	fmt.Println("to disconnect based on name and tty")
	return rpi.Action{}, nil
}

// ExecuteKP returns an action response
func (des Destroy) ExecuteKP(execs map[int]rpi.Exec) (rpi.Action, error) {
	return rpi.Action{}, nil
}
