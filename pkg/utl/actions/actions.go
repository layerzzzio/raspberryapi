package actions

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/raspibuddy/rpi"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

var (
	// Separator separates parent and child execution
	Separator = "<|>"

	// DeleteFile is the name of the delete file exec
	DeleteFile = "delete_file"

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

type DF struct {
	Path string
}

// DeleteFile deletes a file or (empty) directory
func (s Service) DeleteFile(arg interface{}) (rpi.Exec, error) {
	var path string

	switch v := arg.(type) {
	case DF:
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

type CH struct {
	Hostname string
}

// ChangeHostnameInHostnameFile changes the hostname in /etc/hostname
func (s Service) ChangeHostnameInHostnameFile(arg interface{}) (rpi.Exec, error) {
	var hostname string

	switch v := arg.(type) {
	case CH:
		hostname = v.Hostname
	case OtherParams:
		hostname = arg.(OtherParams).Value["hostname"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"hostname"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	err := OverwriteToFile(OverwriteToFileArg{
		File:      "/etc/hostname",
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
func (s Service) ChangeHostnameInHostsFile(arg interface{}) (rpi.Exec, error) {
	var hostname string

	switch v := arg.(type) {
	case CH:
		hostname = v.Hostname
	case OtherParams:
		hostname = arg.(OtherParams).Value["hostname"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"hostname"}}
	}

	// execution start time
	startTime := uint64(time.Now().Unix())
	exitStatus := 0
	var stdErr string

	err := OverwriteToFile(OverwriteToFileArg{
		File:      "/etc/hosts",
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
		Name:       ChangeHostnameInHostsFile,
		StartTime:  startTime,
		EndTime:    endTime,
		ExitStatus: uint8(exitStatus),
		Stderr:     stdErr,
	}, nil
}

type CP struct {
	Password string
}

// ChangePassword changes a password
func (s Service) ChangePassword(arg interface{}) (rpi.Exec, error) {
	var password string

	switch v := arg.(type) {
	case CP:
		password = v.Password
	case OtherParams:
		password = arg.(OtherParams).Value["password"]
	default:
		return rpi.Exec{ExitStatus: 1}, &Error{[]string{"password"}}
	}
	// execution start time
	startTime := uint64(time.Now().Unix())

	exitStatus := 0
	var stdErr string
	e := os.Remove(password)
	if e != nil {
		exitStatus = 1
		stdErr = fmt.Sprint(e)
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

// OverwriteToFileArg is the argument to function OverwriteToFile
type OverwriteToFileArg struct {
	File        string
	Data        []string
	Multiline   bool
	Permissions int // 0644, 0666, etc.
}

// IsDirectory check is a file is a directory or not
// func IsDirectory(path string) (bool, error) {
// 	fileInfo, err := os.Stat(path)
// 	if err != nil {
// 		return false, err
// 	}
// 	return fileInfo.IsDir(), err
// }

// OverwriteToFile write data to file
func OverwriteToFile(args OverwriteToFileArg) error {
	perm := 0644

	f, err := os.Create(args.File)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return fmt.Errorf("creating file failed")
	}

	if args.Permissions != 0 {
		perm = args.Permissions
	}

	if err := os.Chmod(args.File, os.FileMode(perm)); err != nil {
		return fmt.Errorf("chmoding file failed")
	}

	for _, v := range args.Data {
		if args.Multiline {
			fmt.Fprintln(f, v)
		} else {
			fmt.Fprint(f, v)
		}

		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("writing to file failed")
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("closing file failed")
	}

	return nil
}

// ReplaceLineFile replace one or multiple line in file
func ReplaceLineFile(args OverwriteToFileArg) error {
	perm := 0644

	f, err := os.Create(args.File)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return fmt.Errorf("creating file failed")
	}

	if args.Permissions != 0 {
		perm = args.Permissions
	}

	if err := os.Chmod(args.File, os.FileMode(perm)); err != nil {
		return fmt.Errorf("chmoding file failed")
	}

	for _, v := range args.Data {
		if args.Multiline {
			fmt.Fprintln(f, v)
		} else {
			fmt.Fprint(f, v)
		}

		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("writing to file failed")
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("closing file failed")
	}

	return nil
}
