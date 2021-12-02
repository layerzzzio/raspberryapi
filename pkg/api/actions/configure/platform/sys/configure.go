package sys

import (
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
)

// Configure represents an empty Configure entity on the current system.
type Configure struct{}

// ExecuteCH returns an action response after changing hostname
func (con Configure) ExecuteCH(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.ChangeHostname,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteCP returns an action response after changing password
func (con Configure) ExecuteCP(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.ChangePassword,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteWNB returns an action response after enabling or disable wait for network at boot
func (con Configure) ExecuteWNB(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.WaitForNetworkAtBoot,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteOV returns an action response after enabling or disable overscan
func (con Configure) ExecuteOV(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.Overscan,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteBL returns an action response after enabling or disable blanking
func (con Configure) ExecuteBL(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.Blanking,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteAUS returns an action response after adding a user
func (con Configure) ExecuteAUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.AddUser,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteDUS returns an action response after deleting a user
func (con Configure) ExecuteDUS(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.DeleteUser,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteCA returns an action response after enabling or disable camera
func (con Configure) ExecuteCA(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.CameraInterface,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteSSH returns an action response after enabling or disable ssh
func (con Configure) ExecuteSSH(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.SSH,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteVNC returns an action response after enabling or disable vnc
func (con Configure) ExecuteVNC(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.VNC,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteSPI returns an action response after enabling or disable spi
func (con Configure) ExecuteSPI(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.SPI,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteI2C returns an action response after enabling or disable i2c
func (con Configure) ExecuteI2C(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.I2C,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteONW returns an action response after enabling or disable one-wire
func (con Configure) ExecuteONW(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.OneWire,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteRG returns an action response after enabling or disable remote gpio
func (con Configure) ExecuteRG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.RGPIO,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteUPD returns an action response after updating the system
func (con Configure) ExecuteUPD(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.Update,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteUPG returns an action response after upgrading the system
func (con Configure) ExecuteUPG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.Upgrade,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteUPDG returns an action response after updating & upgrading the system
func (con Configure) ExecuteUPDG(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.UpDateGrade,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}

// ExecuteWC returns an action response after changing system the wifi country
func (con Configure) ExecuteWC(plan map[int](map[int]actions.Func)) (rpi.Action, error) {
	actionStartTime := uint64(time.Now().Unix())
	progressInit := actions.FlattenPlan(plan)
	progress, exitStatus := actions.ExecutePlan(plan, progressInit)

	return rpi.Action{
		Name:          actions.WifiCountry,
		NumberOfSteps: uint16(len(progressInit)),
		Progress:      progress,
		ExitStatus:    exitStatus,
		StartTime:     actionStartTime,
		EndTime:       uint64(time.Now().Unix()),
	}, nil
}
