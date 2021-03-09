package sys

import (
	"testing"

	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCH(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.ChangeHostname,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "change_hostname",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			changeHostname, err := s.ExecuteCH(tc.plan)
			assert.Equal(t, tc.wantedDataName, changeHostname.Name)
			assert.Equal(t, tc.wantedDataNumSteps, changeHostname.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, changeHostname.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, changeHostname.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteCP(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.ChangePassword,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "change_password",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			changePassword, err := s.ExecuteCP(tc.plan)
			assert.Equal(t, tc.wantedDataName, changePassword.Name)
			assert.Equal(t, tc.wantedDataNumSteps, changePassword.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, changePassword.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, changePassword.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteWNB(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.WaitForNetworkAtBoot,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "wait_for_network_at_boot",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			changePassword, err := s.ExecuteWNB(tc.plan)
			assert.Equal(t, tc.wantedDataName, changePassword.Name)
			assert.Equal(t, tc.wantedDataNumSteps, changePassword.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, changePassword.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, changePassword.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteOV(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.Overscan,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "overscan",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteOV(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteBL(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.Blanking,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "blanking",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteBL(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteAUS(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.AddUser,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "add_user",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteAUS(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteDUS(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DeleteUser,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "delete_user",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteDUS(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteCA(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.CameraInterface,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "camera_interface",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteCA(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSSH(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.SSH,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "ssh",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteSSH(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteVNC(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.VNC,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "vnc",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteVNC(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSPI(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.SPI,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "spi",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteSPI(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteI2C(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.I2C,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "i2c",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteI2C(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteONW(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.OneWire,
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			wantedDataName:        "one_wire",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			overscan, err := s.ExecuteONW(tc.plan)
			assert.Equal(t, tc.wantedDataName, overscan.Name)
			assert.Equal(t, tc.wantedDataNumSteps, overscan.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, overscan.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, overscan.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
