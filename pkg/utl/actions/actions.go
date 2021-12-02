package actions

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/constants"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
	"github.com/shirou/gopsutil/host"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

const (
	// DefaultFilePerm is the default file permission
	DefaultFilePerm = uint32(0644)

	// Enable is a flag to enable a configuration
	Enable = "enable"

	// Disable is a flag to enable a configuration
	Disable = "disable"

	// Add is a flag to add a configuration or a user
	Add = "add"

	// Delete is a flag to delete a configuration or a user
	Delete = "delete"

	// Comment is a flag to comment lines
	Comment = "comment"

	// Uncomment is a flag to uncomment lines
	Uncomment = "uncomment"

	// IpRegex is the regex used to detect ip addresses in strings
	HostnameChangeInHostsRegex = `^\s*127.0.1.1.*`
	// IpRegex = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"

	// OverscanRegex is the regex used to detect disable_overscan variable in /boot/config.txt
	DisableOrEnableOverscanRegex = `^\s*#?\s*disable_overscan\s*=.*`

	// CommentOverscanRegex is the regex used to detect overscan variable in /boot/config.txt
	CommentOverscanRegex = `^\s*overscan_(left|right|top|bottom)\s*=.*`

	// UncommentOverscanRegex is the regex used to detect overscan variable in /boot/config.txt
	UncommentOverscanRegex = `^\s*#?\s*overscan_(left|right|bottom|top)\s*=.*`

	// DisableOrEnableCameraRegex is the regex used to detect start_x variable in /boot/config.txt
	DisableOrEnableCameraRegex = `^\s*#?\s*start_x\s*=.*`

	// DisableOrEnableSPIRegex is the regex used to detect dtparam=spi variable in /boot/config.txt
	DisableOrEnableSPIRegex = `^\s*#?\s*dtparam\s*=\s*spi\s*=.*`

	// DisableOrEnableI2CRegex is the regex used to detect dtparam=i2c_arm variable in /boot/config.txt
	DisableOrEnableI2CRegex = `^\s*#?\s*dtparam\s*=\s*i2c_arm\s*=.*`

	// StartxCameraRegex is the regex used to detect startx variable in /boot/config.txt
	StartxCameraRegex = `^\s*startx.*`

	// OneWireCommentRegex is the regex used to comment one-wire variable in /boot/config.txt
	OneWireCommentRegex = `^\s*#?\s*dtoverlay\s*=\s*w1-gpio.*`

	// FixupFileCameraRegex is the regex used to detect fixup_file variable in /boot/config.txt
	FixupFileCameraRegex = `^\s*fixup_file.*`

	// StartFileCameraRegex is the regex used to detect start_file variable in /boot/config.txt
	StartFileCameraRegex = `^\s*start_file.*`

	// GpuMemRegex regex
	GpuMemRegex = `^\s*gpu_mem\s*=.*`

	// GpuMemRegex regex
	// GpuMemCameraRegex = `^\s*gpu_mem\s*=\s*([0-1]\s*[0-2]\s*[0-7]\s*.*|\s*)$`

	// Separator separates parent and child execution
	Separator = "<|>"

	// DeleteFile is the name of the delete file exec
	DeleteFile = "delete_file"

	// Reboot is the name of the reboot exec
	Reboot = "reboot"

	// Shutdown is the name of the shutdown exec
	Shutdown = "shutdown"

	// RebootShutdown is the name of the reboot or shutdown exec
	RebootShutdown = "RebootShutdown"

	// KillProcess is the name of the kill process exec
	KillProcess = "kill_process"

	// KillProcessByName is the name of the kill process by name exec
	KillProcessByName = "kill_process_by_name"

	// StopUserSession is the name of the disconnect user action
	StopUserSession = "stop_user_sessions"

	// ChangeHostname is the name of the change username action
	ChangeHostname = "change_hostname"

	// ChangeHostnameInHostnameFile is the name of the change username action
	ChangeHostnameInHostnameFile = "change_hostname_in_hostname_file"

	// ChangeHostnameInHostFile is the name of the change username action
	ChangeHostnameInHostsFile = "change_hostname_in_hosts_file"

	// ChangePassword is the name of the change password action
	ChangePassword = "change_password"

	// WaitForNetworkAtBoot is the name of the wait for network at boot
	WaitForNetworkAtBoot = "wait_for_network_at_boot"

	// DisableOrEnableRemoteGpio is the name of the disable or enable rgpio
	DisableOrEnableRemoteGpio = "disable_or_enable_rgpio"

	// Overscan is the name of the overscan actions
	Overscan = "overscan"

	// DisableOrEnableCameraInterface is the name of the disable or enable camera interface actions
	DisableOrEnableCameraInterface = "disable_or_enable_camera_interface"

	// DisableOrEnableSPIInterface is the name of the disable or enable spi interface actions
	DisableOrEnableSPIInterface = "disable_or_enable_spi_interface"

	// DisableOrEnableI2CInterface is the name of the disable or enable i2c interface actions
	DisableOrEnableI2CInterface = "disable_or_enable_i2c_interface"

	// DisableOrEnableONWInterface is the name of the disable or enable one-wore interface actions
	DisableOrEnableONWInterface = "disable_or_enable_one_wire_interface"

	// CameraInterface is the name of the camera interface actions
	CameraInterface = "camera_interface"

	// DisableOrEnableOverscan is the name of the disable or enable overscan method
	DisableOrEnableOverscan = "disable_or_enable_overscan"

	// CommentOverscan is the name of the comment overscan method
	CommentOverscan = "comment_overscan"

	// DisableOrEnableBlanking is the name of the disable or enable blanking function
	DisableOrEnableBlanking = "disable_or_enable_blanking"

	// Blanking is the name of the disable or enable blanking method
	Blanking = "blanking"

	// AddUser is the name of the adding user method
	AddUser = "add_user"

	// DeleteUser is the name of the deleting user method
	DeleteUser = "delete_user"

	// SSH is the name of the ssh user method
	SSH = "ssh"

	// VNC is the name of the vnc user method
	VNC = "vnc"

	// SPI is the name of the spi user method
	SPI = "spi"

	// I2C is the name of the i2c user method
	I2C = "i2c"

	// OneWire is the name of the one-wire user method
	OneWire = "one_wire"

	// RGPIO is the name of the remote gpio user method
	RGPIO = "remote_gpio"

	// Update is the name of the update method
	Update = "update"

	// Upgrade is the name of the upgrade method
	Upgrade = "upgrade"

	// UpDateGrade is the name of the update & upgrade method
	UpDateGrade = "update_upgrade"

	// WifiCountry is the name of the wifi country method
	WifiCountry = "wifi_country"

	// InstallAptGet is the name of the install with apt-get method
	InstallAptGet = "install_apt_get"

	// DeployVersion is the name of deploy a specific version on the device
	DeployVersion = "deploy_api_version"

	// InstallVPNWithOVPN is the name of the install install vpn with opvn method
	InstallVPNWithOVPN = "install_vpn_with_ovpn"

	// ActionVPNWithOVPN is the name of the actions vpn with opvn method
	ActionVPNWithOVPN = "action_vpn_with_ovpn"

	// ConfirmVPNAuthentication is the name of the action to confirm if VPN authentication workded or not
	ConfirmVPNAuthentication = "confirm_vpn_auth"
)

var (
	// RepTypeAllOccurrences is a flag meaning all occurrences of a word should be replaced
	RepTypeAllOccurrences = "all_occurrences"

	// RepTypeEntireLine is a flag meaning all occurrences of an entire file line should be replaced
	RepTypeEntireLine = "entire_line"
)

// Service represents several system scripts.
type Service struct{}

// Actions represents multiple system related action scripts.
type Actions interface{}

// New creates a service instance.
func New() *Service {
	return &Service{}
}

// Params holds the Func dependencies values
type OtherParams struct {
	Value map[string]string
}

// Func represents a function to be called by function Call
// attribute Arguments content should be ordered
type Func struct {
	Name      string
	Reference interface{}
	Argument  []interface{}
	// Example: "1" + action.Separator + "2" = "1<|>2"
	// Why not another function name ?
	// Reason : ensure uniqueness of the dependency
	Dependency OtherParams
}

// Error is returned by Actions when the argument evaluation fails
type Error struct {
	// Name is the file name for which the error occurred.
	Arguments []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("at least one argument is empty: %v", e.Arguments)
}

type KP struct {
	Pid string
}

// KillProcess kill a given process
func (s Service) KillProcess(arg interface{}) (rpi.Exec, error) {
	var pid string

	switch v := arg.(type) {
	case KP:
		pid = v.Pid
	case OtherParams:
		pid = arg.(OtherParams).Value["pid"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"pid"}}
	}

	var stdErr string
	startTime := uint64(time.Now().Unix())
	pidNum, err := strconv.Atoi(pid)
	if err != nil {
		return rpi.Exec{
			Name:       KillProcess,
			StartTime:  startTime,
			EndTime:    uint64(time.Now().Unix()),
			ExitStatus: uint8(1),
			Stderr:     "pid is not an int",
		}, nil
	} else {
		exitStatus := 0
		ps, _ := os.FindProcess(pidNum)
		e := ps.Kill()

		if e != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(e)
		}
		// execution end time
		endTime := uint64(time.Now().Unix())
		return rpi.Exec{
			Name:       KillProcess,
			StartTime:  startTime,
			EndTime:    endTime,
			ExitStatus: uint8(exitStatus),
			Stderr:     stdErr,
		}, nil
	}
}

type KPBN struct {
	Processname string
	Processtype string
}

// KillProcessByName disconnect a user from an active tty from the current host
// func (s Service) KillProcessByName(arg interface{}, dependency ...OtherParams) (rpi.Exec, error) {
func (s Service) KillProcessByName(arg interface{}) (rpi.Exec, error) {
	var processname string
	var processtype string

	switch v := arg.(type) {
	case KPBN:
		processname = v.Processname
		processtype = v.Processtype
	case OtherParams:
		processname = arg.(OtherParams).Value["processname"]
		processtype = arg.(OtherParams).Value["processtype"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"processname", "processtype"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	var err error
	if processtype == "terminal" {
		_, err = exec.Command("sh", "-c", "pkill -t "+processname).Output()
	} else {
		_, err = exec.Command("sh", "-c", "pkill "+processname).Output()
	}

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       KillProcessByName,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// FileOrDirectory is the argument used when wanting to modified a file only (ex: comment)
type FileOrDirectory struct {
	Path string
}

// DeleteFile deletes a file or (empty) directory
func (s Service) DeleteFile(arg interface{}) (rpi.Exec, error) {
	var path string

	switch v := arg.(type) {
	case FileOrDirectory:
		path = v.Path
	case OtherParams:
		path = arg.(OtherParams).Value["path"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"path"}}
	}
	// execution start time
	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string
	e := os.Remove(path)
	if e != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(e)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DeleteFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// DataToFile is an argument used when wanting to add new data to a file
type DataToFile struct {
	TargetFile string
	Data       string
}

// ChangeHostnameInHostnameFile changes the hostname in /etc/hostname
// It should completely overwrite the file with the new hostname
func (s Service) ChangeHostnameInHostnameFile(arg interface{}) (rpi.Exec, error) {
	var hostname string
	var targetFile string

	switch v := arg.(type) {
	case DataToFile:
		targetFile = v.TargetFile
		hostname = v.Data
	case OtherParams:
		targetFile = arg.(OtherParams).Value["targetFile"]
		hostname = arg.(OtherParams).Value["hostname"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"hostname", "targetFile"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	err := OverwriteToFile(WriteToFileArg{
		File:      targetFile,
		Data:      []string{hostname},
		Multiline: false,
	})

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       ChangeHostnameInHostnameFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// ChangeHostnameInHostsFile changes the hostname in /etc/hosts
// It should replace the old hostname value with new hostname value
func (s Service) ChangeHostnameInHostsFile(arg interface{}) (rpi.Exec, error) {
	var hostname string
	var targetFile string

	switch v := arg.(type) {
	case DataToFile:
		targetFile = v.TargetFile
		hostname = v.Data
	case OtherParams:
		targetFile = arg.(OtherParams).Value["targetFile"]
		hostname = arg.(OtherParams).Value["hostname"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"hostname", "targetFile"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	info, err := host.Info()

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	} else {
		// copy the file if the file exists
		if _, err := os.Stat(targetFile); err == nil {
			err = ReplaceLineInFile(ReplaceLineInFileArg{
				File:  targetFile,
				Regex: HostnameChangeInHostsRegex,
				ReplaceType: ReplaceType{
					// old hostname is not passed as an argument from the application
					// indeed it can changed between the moment the user ask for a change
					// and the moment it actually changes
					&AllOccurrences{
						Occurrence: info.Hostname,
						NewData:    hostname,
					},
					nil,
				},
				ToAddIfNoMatch: []string{"127.0.1.1		" + hostname},
				HasUniqueLines: true,
			})

			if err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		} else {
			exitStatus, stdErr = CreateAssetFile(
				CreateAssetFileArg{
					AssetFile:  "../assets/hosts",
					TargetFile: targetFile,
					NewData: []string{"127.0.1.1		" + hostname},
					HasUniqueLine: true,
				},
			)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       ChangeHostnameInHostsFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// CP is the argument when changing the password
type CP struct {
	Password string
	Username string
}

// ChangePassword changes a password without a prompt
func (s Service) ChangePassword(arg interface{}) (rpi.Exec, error) {
	var password string
	var username string

	switch v := arg.(type) {
	case CP:
		password = v.Password
		username = v.Username
	case OtherParams:
		password = arg.(OtherParams).Value["password"]
		username = arg.(OtherParams).Value["username"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"password", "username"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	var err error
	// the command was found here:
	// https://askubuntu.com/questions/80444/how-to-set-user-passwords-using-passwd-without-a-prompt
	_, err = exec.Command(
		"sh",
		"-c",
		"usermod --password $(echo "+password+" | openssl passwd -1 -stdin) "+username,
	).Output()

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       ChangePassword,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// ADU argument for AddOrDeleteUser function
type ADU struct {
	Username string
	Password string
}

// AddUser add a user on the system
func (s Service) AddUser(arg interface{}) (rpi.Exec, error) {
	var password string
	var username string

	switch v := arg.(type) {
	case ADU:
		password = v.Password
		username = v.Username
	case OtherParams:
		password = arg.(OtherParams).Value["password"]
		username = arg.(OtherParams).Value["username"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"username", "password"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	if exitStatus != 1 {
		var err error
		_, err = exec.Command(
			"sh",
			"-c",
			fmt.Sprintf("useradd -m %v", username),
		).Output()

		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		} else {
			_, err = exec.Command(
				"sh",
				"-c",
				"usermod --password $(echo "+password+" | openssl passwd -1 -stdin) "+username,
			).Output()

			if err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       AddUser,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// DeleteUser delete a user on the system
func (s Service) DeleteUser(arg interface{}) (rpi.Exec, error) {
	var username string

	switch v := arg.(type) {
	case ADU:
		username = v.Username
	case OtherParams:
		username = arg.(OtherParams).Value["username"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"username"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	var err error
	_, err = exec.Command(
		"sh",
		"-c",
		fmt.Sprintf("userdel -r %v", username),
	).Output()

	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DeleteUser,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

type EBC struct {
	Command string
}

const ExecuteBashCommand = "execute_bash_command"

// DisableOrEnableBashCommand runs a bash command depending on the action value
func (s Service) ExecuteBashCommand(arg interface{}) (rpi.Exec, error) {
	var command string

	switch v := arg.(type) {
	case EBC:
		command = v.Command
	case OtherParams:
		command = arg.(OtherParams).Value["command"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"command"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	if command == "" {
		exitStatus = 1
		stdErr = "no command"
	} else {
		var err error
		_, err = exec.Command(
			"sh",
			"-c",
			command,
		).Output()

		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       ExecuteBashCommand,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

type CVPNAUTH struct {
	Filepath  string
	Timelimit string
}

// ConfirmVPNAuthentication checks if a VPN authentication works on not
func (s Service) ConfirmVPNAuthentication(arg interface{}) (rpi.Exec, error) {
	// in seconds
	var timelimit string
	var filepath string

	switch v := arg.(type) {
	case CVPNAUTH:
		filepath = v.Filepath
		timelimit = v.Timelimit
	case OtherParams:
		filepath = arg.(OtherParams).Value["filepath"]
		timelimit = arg.(OtherParams).Value["timelimit"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{
			"filepath", "keyword",
		}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	var exitStatus int
	var stdErr string

	timelimitint, errTL := strconv.Atoi(timelimit)
	if errTL != nil {
		return rpi.Exec{
			Name:       ConfirmVPNAuthentication,
			StartTime:  startTime,
			EndTime:    uint64(time.Now().Unix()),
			ExitStatus: uint8(1),
			Stderr:     "timelimit is not an int",
		}, nil
	}

	keyword, err := infos.New().IsFileContainsUntil(
		filepath,
		infos.IFCK{
			Name: "auth_failure",
			Keywords: []string{
				"AUTH_FAILED",
				"auth-failure",
			},
		}, infos.IFCK{
			Name: "auth_success",
			Keywords: []string{
				"Initialization Sequence Completed",
			},
		},
		timelimitint,
	)

	if err != nil {
		return rpi.Exec{
			Name:       ConfirmVPNAuthentication,
			StartTime:  startTime,
			EndTime:    uint64(time.Now().Unix()),
			ExitStatus: uint8(1),
			Stderr:     fmt.Sprint(err),
		}, nil
	}

	if keyword == "auth_failure" {
		exitStatus = 1
		stdErr = "auth_failure"
	} else if keyword == "not_found" {
		exitStatus = 1
		stdErr = "stalled"
	} else if keyword == "auth_success" {
		exitStatus = 0
	} else {
		exitStatus = 1
		stdErr = "auth_unknown_error"
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       ConfirmVPNAuthentication,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// EnableOrDisableConfig is the argument for enable or disable methods
type EnableOrDisableConfig struct {
	Action        string
	DirOrFilePath string
}

// EnableOrDisableConfigExtraFile is the argument for enable or disable methods
type TargetDestEnableOrDisableConfig struct {
	Action                   string
	TargetDirOrFilePath      string
	DestinationDirOrFilePath string
}

// WaitForNetworkAtBoot enable or disable wait for network at boot
func (s Service) WaitForNetworkAtBoot(arg interface{}) (rpi.Exec, error) {
	var directory string
	var action string

	switch v := arg.(type) {
	case EnableOrDisableConfig:
		directory = v.DirOrFilePath
		action = v.Action
	case OtherParams:
		directory = arg.(OtherParams).Value["directory"]
		action = arg.(OtherParams).Value["action"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"directory", "action"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	if action == Enable {
		// create the directory and the parent directories
		_ = os.MkdirAll(directory, 0755)

		// create a file wait.conf and populate it
		err := OverwriteToFile(WriteToFileArg{
			File: directory + "/wait.conf",
			Data: []string{
				"[Service]",
				"ExecStart=",
				"ExecStart=/usr/lib/dhcpcd5/dhcpcd -q -w",
			},
			Multiline: true,
		})

		// if error, it is logged here
		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	} else if action == Disable {
		// remove the file
		err := os.Remove(directory + "/wait.conf")

		// if error, it is logged here
		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	} else {
		exitStatus = 1
		stdErr = "bad action type"
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       WaitForNetworkAtBoot,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// DisableOrEnableRemoteGpio enable or disable remote gpio at boot
func (s Service) DisableOrEnableRemoteGpio(arg interface{}) (rpi.Exec, error) {
	var directory string
	var action string

	switch v := arg.(type) {
	case EnableOrDisableConfig:
		directory = v.DirOrFilePath
		action = v.Action
	case OtherParams:
		directory = arg.(OtherParams).Value["directory"]
		action = arg.(OtherParams).Value["action"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"directory", "action"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	if action == Enable {
		// create the directory and the parent directories
		_ = os.MkdirAll(directory, 0755)

		// create a file wait.conf and populate it
		err := OverwriteToFile(WriteToFileArg{
			File: directory + "/public.conf",
			Data: []string{
				"[Service]",
				"ExecStart=",
				"ExecStart=/usr/bin/pigpiod",
			},
			Multiline: true,
		})

		// if error, it is logged here
		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	} else if action == Disable {
		// remove the file
		err := os.Remove(directory + "/public.conf")

		// if error, it is logged here
		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	} else {
		exitStatus = 1
		stdErr = "bad action type"
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DisableOrEnableRemoteGpio,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// CommentOrUncommentConfig is the argument for comment or uncomment methods
type CommentOrUncommentConfig struct {
	DirOrFilePath string
	Action        string
}

// CommentOverscan comments overscan lines
func (s Service) CommentOverscan(arg interface{}) (rpi.Exec, error) {
	var path string
	var action string

	switch v := arg.(type) {
	case CommentOrUncommentConfig:
		path = v.DirOrFilePath
		action = v.Action
	case OtherParams:
		path = arg.(OtherParams).Value["path"]
		action = arg.(OtherParams).Value["action"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"path", "action"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string
	var regex string

	var defaultData []string

	if action == "comment" {
		regex = CommentOverscanRegex
		defaultData = []string{
			"#overscan_left=16",
			"#overscan_right=16",
			"#overscan_top=16",
			"#overscan_bottom=16",
		}
	} else {
		exitStatus = 1
		stdErr = "bad action type"
	}

	if exitStatus == 0 {
		if _, err := os.Stat(path); err == nil {
			err := CommentOrUncommentLineInFile(CommentLineInFileArg{
				File:           path,
				Regex:          regex,
				Action:         action,
				ToAddIfNoMatch: defaultData,
				HasUniqueLines: true,
			})

			if err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		} else {
			exitStatus, stdErr = CreateAssetFile(
				// no new data because already commented in assets
				CreateAssetFileArg{
					AssetFile:  "../assets/config.txt",
					TargetFile: path,
				},
			)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       CommentOverscan,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// COUSLINF comment or uncomment single line in file
type COUSLINF struct {
	FunctionName  string
	Action        string
	DirOrFilePath string
	Regex         string
	DefaultData   string
	AssetFile     string
}

const CommentOrUncommentInFile = "comment_or_uncomment_in_file"

// CommentInFile comments overscan lines
func (s Service) CommentOrUncommentInFile(arg interface{}) (rpi.Exec, error) {
	var functionName string
	var action string
	var path string
	var regex string
	var defaultData string
	var assetFile string

	switch v := arg.(type) {
	case COUSLINF:
		functionName = v.FunctionName
		action = v.Action
		path = v.DirOrFilePath
		regex = v.Regex
		defaultData = v.DefaultData
		assetFile = v.AssetFile
	case OtherParams:
		functionName = arg.(OtherParams).Value["functionName"]
		action = arg.(OtherParams).Value["action"]
		path = arg.(OtherParams).Value["path"]
		regex = arg.(OtherParams).Value["regex"]
		defaultData = arg.(OtherParams).Value["defaultData"]
		assetFile = arg.(OtherParams).Value["assetFile"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{
			"name", "action", "path", "regex", "defaultData", "assetFile",
		}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	if action != "comment" && action != "uncomment" {
		exitStatus = 1
		stdErr = "bad action type"
	}

	if exitStatus == 0 {
		if _, err := os.Stat(path); err == nil {
			err := CommentOrUncommentLineInFile(CommentLineInFileArg{
				File:           path,
				Regex:          regex,
				Action:         action,
				ToAddIfNoMatch: []string{defaultData},
				HasUniqueLines: true,
			})

			if err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		} else {
			exitStatus, stdErr = CreateAssetFile(
				// no new data because already commented in assets
				CreateAssetFileArg{
					AssetFile:  assetFile,
					TargetFile: path,
				},
			)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       functionName,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// DisableOrEnableBlanking disables or enables blanking
func (s Service) DisableOrEnableBlanking(arg interface{}) (rpi.Exec, error) {
	var target string
	var destination string
	var action string

	switch v := arg.(type) {
	case TargetDestEnableOrDisableConfig:
		target = v.TargetDirOrFilePath
		destination = v.DestinationDirOrFilePath
		action = v.Action
	case OtherParams:
		target = arg.(OtherParams).Value["target"]
		destination = arg.(OtherParams).Value["destination"]
		action = arg.(OtherParams).Value["action"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"target", "destination", "action"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	if action == Enable {
		// remove the file
		err := os.Remove(destination + "/10-blanking.conf")

		// if error, it is logged here
		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	} else if action == Disable {
		// create the directory and the parent directories
		if err := os.MkdirAll(destination, 0755); err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		} else {
			if _, err := os.Stat(target + "/10-blanking.conf"); err != nil {
				exitStatus, stdErr = CreateAssetFile(
					// no new data because already commented in assets
					CreateAssetFileArg{
						AssetFile:  "../assets/10-blanking.conf",
						TargetFile: target + "/10-blanking.conf",
					},
				)
			}

			if err := CopyFile(
				target+"/10-blanking.conf",
				destination+"/10-blanking.conf",
				DefaultFilePerm,
			); err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		}
	} else {
		exitStatus = 1
		stdErr = "bad action type"
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       DisableOrEnableBlanking,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

type EODC struct {
	Action        string
	DirOrFilePath string
	Data          string
	Regex         string
	AssetFile     string
	FunctionName  string
}

const DisableOrEnableConfig = "disable_or_enable_config"

// DisableOrEnableConfig disables or enables a config in a file
func (s Service) DisableOrEnableConfig(arg interface{}) (rpi.Exec, error) {
	var functionName string
	var action string
	var path string
	var regex string
	var data string
	var assetFile string

	switch v := arg.(type) {
	case EODC:
		functionName = v.FunctionName
		action = v.Action
		path = v.DirOrFilePath
		regex = v.Regex
		data = v.Data
		assetFile = v.AssetFile
	case OtherParams:
		functionName = arg.(OtherParams).Value["functionName"]
		action = arg.(OtherParams).Value["action"]
		path = arg.(OtherParams).Value["path"]
		regex = arg.(OtherParams).Value["regex"]
		data = arg.(OtherParams).Value["data"]
		assetFile = arg.(OtherParams).Value["assetFile"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{
			"name", "action", "path", "regex", "data", "assetFile",
		}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string
	var newData string

	if action == Enable {
		newData = data
	} else if action == Disable {
		newData = data
	} else {
		exitStatus = 1
		stdErr = "bad action type"
	}

	if exitStatus == 0 {
		if _, err := os.Stat(path); err == nil {
			err := ReplaceLineInFile(ReplaceLineInFileArg{
				File:  path,
				Regex: regex,
				ReplaceType: ReplaceType{
					nil,
					&EntireLine{NewData: newData},
				},
				HasUniqueLines: true,
				ToAddIfNoMatch: []string{newData},
			})

			if err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		} else {
			exitStatus, stdErr = CreateAssetFile(
				// it will add the new data at the end of the file
				// indeed all lines commented from asset
				CreateAssetFileArg{
					AssetFile:     assetFile,
					TargetFile:    path,
					NewData:       []string{newData},
					HasUniqueLine: true,
				},
			)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       functionName,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

const SetVariableInConfigFile = "set_variable_in_config_file"

type SVICF struct {
	File      string
	Regex     string
	Data      string
	AssetFile string
	Threshold string
}

// DisableOrEnableConfig disables or enables a config in a file
func (s Service) SetVariableInConfigFile(arg interface{}) (rpi.Exec, error) {
	var file string
	var regex string
	var data string
	var assetFile string
	var threshold string

	switch v := arg.(type) {
	case SVICF:
		file = v.File
		regex = v.Regex
		data = v.Data
		assetFile = v.AssetFile
		threshold = v.Threshold
	case OtherParams:
		file = arg.(OtherParams).Value["file"]
		regex = arg.(OtherParams).Value["regex"]
		data = arg.(OtherParams).Value["data"]
		assetFile = arg.(OtherParams).Value["assetFile"]
		threshold = arg.(OtherParams).Value["threshold"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{
			"file", "regex", "data", "assetFile", "threshold",
		}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	thr, err := strconv.Atoi(threshold)
	if err != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(err)
	} else {
		if _, err := os.Stat(file); err == nil {
			err := SetVariable(file, 0664, regex, data, true, thr)
			if err != nil {
				exitStatus = 1
				stdErr = fmt.Sprint(err)
			}
		} else {
			exitStatus, stdErr = CreateAssetFile(
				// it will add the new data at the end of the file
				// indeed all lines commented from asset
				CreateAssetFileArg{
					AssetFile:     assetFile,
					TargetFile:    file,
					NewData:       []string{data},
					HasUniqueLine: true,
				},
			)
		}
	}

	// execution end time
	endTime := uint64(time.Now().Unix())

	return rpi.Exec{
		Name:       SetVariableInConfigFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

// Call calls a function by its name and params
func Call(funcName interface{}, params []interface{}) (result interface{}, err error) {
	// defer wg.Done()
	f := reflect.ValueOf(funcName)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	res := f.Call(in)
	result = res[0].Interface()
	return
}

// FlattenPlan flattens out an execute plan
func FlattenPlan(execPlan map[int](map[int]Func)) map[string]rpi.Exec {
	progress := map[string]rpi.Exec{}
	for kp, parentExec := range execPlan {
		parentIndex := fmt.Sprint(kp)
		for kc := range parentExec {
			childIndex := Separator + fmt.Sprint(kc)
			index := parentIndex + childIndex
			progress[index] = rpi.Exec{}
		}
	}

	return progress
}

type CallRes struct {
	Index  string
	Result rpi.Exec
}

func handleResults(input chan CallRes, output chan map[string]rpi.Exec, wg *sync.WaitGroup) {
	var res = map[string]rpi.Exec{}
	for exec := range input {
		res[exec.Index] = exec.Result
		wg.Done()
	}
	output <- res
}

func concurrentExec(execs map[int]Func, index string, progress map[string]rpi.Exec) map[string]rpi.Exec {
	input := make(chan CallRes)
	output := make(chan map[string]rpi.Exec)
	var wg sync.WaitGroup
	defer close(output)

	go handleResults(input, output, &wg)

	for kc, childExec := range execs {
		wg.Add(1)
		i, _ := strconv.Atoi(index)

		// adding arguments here in absolutely essentials
		// it allows the params to be sorted in the right order
		// arguments = append(arguments, childExec.Argument...)
		if len(childExec.Dependency.Value) > 0 && i > 1 {
			otherParamValue := map[string]string{}
			var otherParam = OtherParams{}
			for varName, dep := range childExec.Dependency.Value {
				if strings.Contains(dep, Separator) {
					otherParamValue[varName] = progress[dep].Stdout
				} else {
					otherParamValue[varName] = dep
				}
				otherParam = OtherParams{Value: otherParamValue}
			}
			childExec.Argument = append(childExec.Argument, otherParam)
		}

		go func(childExec Func, kc int) {
			res, errC := Call(childExec.Reference, childExec.Argument)
			if errC != nil {
				input <- CallRes{
					Index: index + Separator + fmt.Sprint(kc),
					Result: rpi.Exec{
						Name:       childExec.Name,
						ExitStatus: 1,
						Stderr:     fmt.Sprint(errC),
					},
				}
			} else {
				input <- CallRes{
					Index:  index + Separator + fmt.Sprint(kc),
					Result: res.(rpi.Exec),
				}
			}
		}(childExec, kc)
	}

	wg.Wait()       // Wait until the count is back to zero
	close(input)    // Close the input channel
	res := <-output // Read the message written to the output channel
	return res
}

// ExecutePlan execute an action plan sequentially and in parallel
func ExecutePlan(execPlan map[int](map[int]Func), progress map[string]rpi.Exec) (map[string]rpi.Exec, uint8) {
	var exitStatus uint8
	var index string

	n := len(execPlan)

	for kp := 1; kp <= n; kp++ {
		index = fmt.Sprint(kp)
		res := concurrentExec(execPlan[kp], index, progress)

		for i, e := range res {
			progress[i] = e
			if e.ExitStatus != 0 {
				exitStatus = 1
			}
		}

		if exitStatus == 1 {
			break
		}
	}

	return progress, exitStatus
}

// WriteToFileArg is the argument to function OverwriteToFile
type WriteToFileArg struct {
	File        string
	Data        []string
	Multiline   bool
	Permissions uint32 // 0644, 0666, etc.
}

// BackupFile copies a file and adds suffix .bak to the copied file
// !!! "defer close" should absolutely not be used here !!!
// source: https://www.joeshaw.org/dont-defer-close-on-writable-files/
func CopyFile(target string, destination string, perm uint32) error {
	// copy the file if the file exists
	if _, err := os.Stat(target); err == nil {
		in, err := os.Open(target)
		if err != nil {
			return fmt.Errorf("opening source file failed")
		}

		fmt.Println("creating file here" + destination)
		out, err := os.Create(destination)
		if err != nil {
			out.Close()
			return fmt.Errorf("creating copied file failed")
		}

		if _, err = io.Copy(out, in); err != nil {
			return fmt.Errorf("copying to copied file failed")
		}

		err = in.Close()
		if err != nil {
			return fmt.Errorf("closing target failed")
		}

		err = out.Close()
		if err != nil {
			return fmt.Errorf("closing destination failed")
		}

		if err := ApplyPermissionsToFile(destination, perm); err != nil {
			return fmt.Errorf("applying permission failed")
		}
	}
	return nil
}

// ApplyPermissionsToFile apply permissions to a given file
func ApplyPermissionsToFile(path string, perm uint32) error {
	// !!! the permissions are octal numbers !!!
	// https://yourbasic.org/golang/gotcha-octal-decimal-hexadecimal-literal/
	// first number = 0 (true for every octal)
	// min number = 0
	// max number = 7
	// The reason for Go to use octal is that it makes it impossible
	// to have numbers higher than 7 which is the max for permissions (rwx)

	re := regexp.MustCompile(`^0[0-7]{3}$`)
	if re.MatchString("0" + strconv.FormatInt(int64(perm), 8)) {
		if err := os.Chmod(path, os.FileMode(perm)); err != nil {
			return fmt.Errorf("chmoding file failed")
		}
	} else {
		if err := os.Chmod(path, os.FileMode(DefaultFilePerm)); err != nil {
			return fmt.Errorf("chmoding default file permissions failed")
		}
	}
	return nil
}

func CreateOrTruncate(path string, perm uint32) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("creating file failed")
	}

	if err := ApplyPermissionsToFile(path, perm); err != nil {
		return nil, fmt.Errorf("applying permission failed")
	}

	return f, nil
}

func CloseAndRemoveBakFile(file *os.File, path string) error {
	// closing file
	err := file.Close()
	if err != nil {
		return fmt.Errorf("closing file failed")
	}

	// remove bak file
	pathBak := path + ".bak"
	if _, err := os.Stat(pathBak); err == nil {
		if err = os.Remove(pathBak); err != nil {
			return fmt.Errorf("removing bak file failed")
		}
	}

	return nil
}

// OverwriteToFile overwrite data in a given file
func OverwriteToFile(args WriteToFileArg) error {
	if err := CopyFile(args.File, args.File+".bak", DefaultFilePerm); err != nil {
		return fmt.Errorf("backuping file failed")
	}

	f, err := CreateOrTruncate(args.File, args.Permissions)
	if err != nil {
		return fmt.Errorf("creating and opening file failed")
	}

	// overwriting logic
	for _, v := range args.Data {
		if args.Multiline {
			fmt.Fprintln(f, v)
		} else {
			fmt.Fprint(f, v)
		}

		if err != nil {
			if err = os.Rename(args.File+".bak", args.File); err != nil {
				return fmt.Errorf("renaming bak file to regular file failed")
			}
			return fmt.Errorf("writing to file failed")
		}
	}

	// close file and remove bak file
	if err := CloseAndRemoveBakFile(f, args.File); err != nil {
		return fmt.Errorf("closing file and removing bak file failed")
	}

	return nil
}

// ReplaceLineFile is the argument to function ReplaceLineFile
type ReplaceLineInFileArg struct {
	File           string
	Permissions    uint32 // 0644, 0666, etc.
	Regex          string
	ReplaceType    ReplaceType
	ToAddIfNoMatch []string
	HasUniqueLines bool
}

// ReplaceType lists all the possible replace types
type ReplaceType struct {
	AllOccurrences *AllOccurrences
	EntireLine     *EntireLine
}

// AllOccurrences is a replace type that defines an occurrence to be replaced with new data
type AllOccurrences struct {
	Occurrence string
	NewData    string
}

// EntireLine is a replace type that replaces an entire file line with new data
type EntireLine struct {
	NewData string
}

func GetReplaceType(repType ReplaceType) (*string, error) {
	result := ""

	if repType.AllOccurrences != nil && repType.EntireLine != nil {
		return nil, fmt.Errorf("only one replace type allowed")
	} else if repType.AllOccurrences != nil && repType.EntireLine == nil {
		result = RepTypeAllOccurrences
	} else if repType.AllOccurrences == nil && repType.EntireLine != nil {
		result = RepTypeEntireLine
	} else {
		return nil, fmt.Errorf("at least one replace type required")
	}

	return &result, nil
}

func RemoveDuplicateStrings(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, v := range strSlice {
		if _, value := keys[v]; !value {
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}

// ReplaceLineInFile replace one or multiple line in file
func ReplaceLineInFile(args ReplaceLineInFileArg) error {
	repType, err := GetReplaceType(args.ReplaceType)
	if err != nil {
		return fmt.Errorf("getting replace type failed")
	}

	rawLines, err := infos.New().ReadFile(args.File)
	if err != nil {
		return fmt.Errorf("opening file failed")
	}

	if args.HasUniqueLines {
		rawLines = RemoveDuplicateStrings(rawLines)
	}

	newLines := []string{}
	matchCounter := 0

	for _, line := range rawLines {
		if line != "" {
			// apply the replace type here
			re := regexp.MustCompile(args.Regex)
			if re.MatchString(line) {
				switch *repType {
				case RepTypeAllOccurrences:
					line = strings.ReplaceAll(
						line,
						args.ReplaceType.AllOccurrences.Occurrence,
						args.ReplaceType.AllOccurrences.NewData,
					)
					line = strings.TrimSpace(line)
				case RepTypeEntireLine:
					line = args.ReplaceType.EntireLine.NewData
				default:
					line = args.ReplaceType.EntireLine.NewData
				}
				matchCounter++
			}
			newLines = append(newLines, strings.TrimSuffix(line, "\n"))
		}
	}

	if args.ToAddIfNoMatch != nil {
		// we want to make sure there is as much match
		// as the number of matched lines to modify
		// ----------------------
		// example: only 1 match instead of 2 theoretically
		// ----------------------
		// overscan_top=1
		// overscan_bottomX=2
		// in this case there is only 1 match (with overscan_top) instead of 2
		// indeed we also want to work with line overscan_bottomX=2
		// so the only choice here, to make sure 2 lines are added, is to add default data.
		// ----------------------
		// A case that is not taken into account is when
		// we have two identical matches and one non-match
		// while we would like to have 2 matches for 2 different values
		// This case can happens only is HasUniqueLine = False <- not safe
		// ----------------------
		// overscan_top=1
		// overscan_top=1
		// overscan_bottomX=2
		if matchCounter != len(args.ToAddIfNoMatch) {
			newLines = append(newLines, args.ToAddIfNoMatch...)
		}
	}

	if args.HasUniqueLines {
		newLines = RemoveDuplicateStrings(newLines)
	}

	// allLines is deduplicated
	if err = OverwriteToFile(WriteToFileArg{
		File:        args.File,
		Data:        newLines,
		Multiline:   true,
		Permissions: args.Permissions,
	}); err != nil {
		return fmt.Errorf("overwriting to file failed")
	}

	return nil
}

// SetVariable replace one or multiple line in file
func SetVariable(
	file string,
	permissions uint32,
	regex string,
	data string,
	hasUniqueLines bool,
	threshold int,
) error {
	rawLines, err := infos.New().ReadFile(file)
	if err != nil {
		return fmt.Errorf("opening file failed")
	}

	if hasUniqueLines {
		rawLines = RemoveDuplicateStrings(rawLines)
	}

	newLines := []string{}
	matchCounter := 0
	isToReplace := true

	for _, line := range rawLines {
		if line != "" {
			// apply the replace type here
			re := regexp.MustCompile(regex)
			if re.MatchString(line) {
				reNum := regexp.MustCompile(`[0-9]+`)
				valueStr := reNum.FindString(line)
				if valueStr != "" {
					valueNum, _ := strconv.Atoi(valueStr)
					if threshold != -1 {
						if valueNum > threshold {
							isToReplace = false
						}
					}
				}

				if matchCounter == 0 && isToReplace {
					line = strings.TrimSpace(
						strings.ReplaceAll(
							line,
							line,
							data,
						),
					)
				} else if matchCounter >= 1 {
					line = ""
				}

				matchCounter++
			}
			newLines = append(newLines, strings.TrimSuffix(line, "\n"))
		}
	}

	if hasUniqueLines {
		newLines = RemoveDuplicateStrings(newLines)
	}

	// allLines is deduplicated
	if err = OverwriteToFile(WriteToFileArg{
		File:        file,
		Data:        newLines,
		Multiline:   true,
		Permissions: permissions,
	}); err != nil {
		return fmt.Errorf("overwriting to file failed")
	}

	return nil
}

// GetVariable get variable value from file
func GetVariable(file string, regex string) (string, error) {
	var result string

	rawLines, err := infos.New().ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("opening file failed")
	}

	rawLines = RemoveDuplicateStrings(rawLines)

	for _, line := range rawLines {
		if line != "" {
			re := regexp.MustCompile(regex)
			reNum := regexp.MustCompile(`[0-9]+`)
			if re.MatchString(line) {
				result = reNum.FindString(line)
				break
			}
		}
	}

	return result, nil
}

// AddLinesEndOfFile adds one or multiple lines at the end of a file
func AddLinesEndOfFile(args WriteToFileArg) error {
	// read all lines of original file
	readLines, err := infos.New().ReadFile(args.File)
	if err != nil {
		return fmt.Errorf("reading file failed")
	}

	// populate the array at the end with the data
	readLines = append(readLines, args.Data...)

	// then create a file and overwrite the data from the modified array
	if err = OverwriteToFile(WriteToFileArg{
		File:        args.File,
		Data:        readLines,
		Multiline:   true,
		Permissions: args.Permissions,
	}); err != nil {
		return fmt.Errorf("overwriting to file failed")
	}

	return nil
}

// CommentLineInFileArg is the argument to function CommentLineInFile
type CommentLineInFileArg struct {
	File           string
	Permissions    uint32 // 0644, 0666, etc.
	Regex          string
	Action         string
	ToAddIfNoMatch []string
	HasUniqueLines bool
}

// CommentLineInFile comments one or multiple line in file
func CommentOrUncommentLineInFile(args CommentLineInFileArg) error {
	rawLines, err := infos.New().ReadFile(args.File)
	if err != nil {
		return fmt.Errorf("opening file failed")
	}

	if args.HasUniqueLines {
		rawLines = RemoveDuplicateStrings(rawLines)
	}

	newLines := []string{}
	matchCounter := 0

	for _, line := range rawLines {
		if line != "" {
			// apply the replace type here
			re := regexp.MustCompile(args.Regex)
			if re.MatchString(line) {
				line = strings.TrimSpace(line)
				if args.Action == Comment {
					line = "#" + line
				} else if args.Action == Uncomment {
					line = strings.TrimSpace(strings.Replace(line, "#", "", 1))
				} else {
					return fmt.Errorf("bad action: comment or uncomment")
				}
				matchCounter++
			}
			newLines = append(newLines, strings.TrimSuffix(line, "\n"))
		}
	}

	if args.ToAddIfNoMatch != nil {
		if len(args.ToAddIfNoMatch) != matchCounter {
			newLines = append(newLines, args.ToAddIfNoMatch...)
		}
	}

	if args.HasUniqueLines {
		newLines = RemoveDuplicateStrings(newLines)
	}

	if err = OverwriteToFile(WriteToFileArg{
		File:        args.File,
		Data:        newLines,
		Multiline:   true,
		Permissions: args.Permissions,
	}); err != nil {
		return fmt.Errorf("overwriting to file failed")
	}

	return nil
}

// CreateAssetFileArg is the argument for CreateAssetFile
type CreateAssetFileArg struct {
	AssetFile     string
	TargetFile    string
	HasUniqueLine bool
	NewData       []string
}

// CreateAssetFile creates a file from an asset file
func CreateAssetFile(args CreateAssetFileArg) (int, string) {
	exitStatus := 0
	var stdErr string

	// for go version 1.16
	// couldn't do that because of labstack color package issue
	// assetData, err := infos.New().ReadFile(args.AssetFile)

	// if err != nil {
	// 	exitStatus = 1
	// 	stdErr = fmt.Sprint(err)
	// } else {
	// 	if args.HasUniqueLine {
	// 		assetData = RemoveDuplicateStrings(append(
	// 			assetData,
	// 			args.NewData...,
	// 		))
	// 	}

	// 	if err = OverwriteToFile(
	// 		WriteToFileArg{
	// 			File:      args.TargetFile,
	// 			Data:      assetData,
	// 			Multiline: true,
	// 		},
	// 	); err != nil {
	// 		exitStatus = 1
	// 		stdErr = fmt.Sprint(err)
	// 	}
	// }

	// 	if args.HasUniqueLine {
	// 		assetData = RemoveDuplicateStrings(append(
	// 			assetData,
	// 			args.NewData...,
	// 		))
	// 	}

	if assetData := constants.FILEMAP[args.AssetFile]; assetData != nil {
		if err := OverwriteToFile(
			WriteToFileArg{
				File:      args.TargetFile,
				Data:      append(assetData, args.NewData...),
				Multiline: true,
			},
		); err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
		}
	} else {
		exitStatus = 1
		stdErr = "couldn't find asset file"
	}

	return exitStatus, stdErr
}
