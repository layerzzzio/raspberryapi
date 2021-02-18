package actions

import (
	"bufio"
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
	"github.com/shirou/gopsutil/host"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

var (
	// DefaultFilePerm is the default file permission
	DefaultFilePerm = uint32(0644)

	// RepTypeAllOccurrences is a flag meaning all occurrences of a word should be replaced
	RepTypeAllOccurrences = "all_occurrences"

	// RepTypeEntireLine is a flag meaning all occurrences of an entire file line should be replaced
	RepTypeEntireLine = "entire_line"

	// IpRegex is the regex used to detect ip addresses in strings
	HostnameChangeInHostsRegex = `^127.0.1.1.*`
	// IpRegex = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"

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

// DF is the argument used when deleting a file
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

	err := OverwriteToFile(OverwriteToFileArg{
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
		})

		if err != nil {
			exitStatus = 1
			stdErr = fmt.Sprint(err)
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

// OverwriteToFileArg is the argument to function OverwriteToFile
type OverwriteToFileArg struct {
	File        string
	Data        []string
	Multiline   bool
	Permissions uint32 // 0644, 0666, etc.
}

// BackupFile copies a file and adds suffix .bak to the copied file
// !!! "defer close" should absolutely not be used here !!!
// source: https://www.joeshaw.org/dont-defer-close-on-writable-files/
func BackupFile(path string, perm uint32) error {
	newPath := path + ".bak"

	// copy the file if the file exists
	if _, err := os.Stat(path); err == nil {
		in, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("opening source file failed")
		}

		out, err := os.Create(newPath)
		if err != nil {
			out.Close()
			return fmt.Errorf("creating bak file failed")
		}

		if _, err = io.Copy(out, in); err != nil {
			return fmt.Errorf("copying to bak file failed")
		}

		err = in.Close()
		if err != nil {
			return fmt.Errorf("closing file failed")
		}

		err = out.Close()
		if err != nil {
			return fmt.Errorf("closing new file failed")
		}

		if err := ApplyPermissionsToFile(newPath, perm); err != nil {
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
func OverwriteToFile(args OverwriteToFileArg) error {
	err := BackupFile(args.File, DefaultFilePerm)
	if err != nil {
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
	File        string
	Permissions uint32 // 0644, 0666, etc.
	Regex       string
	ReplaceType ReplaceType
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

// ReplaceLineInFile replace one or multiple line in file
func ReplaceLineInFile(args ReplaceLineInFileArg) error {
	repType, err := GetReplaceType(args.ReplaceType)
	if err != nil {
		return fmt.Errorf("getting replace type failed")
	}

	f, err := os.Open(args.File)
	if err != nil {
		return fmt.Errorf("opening file failed")
	}

	// replacing line logic
	// source: https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	reader := bufio.NewReader(f)
	var line string
	allLines := []string{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// apply the replace type here
		re := regexp.MustCompile(args.Regex)
		if re.MatchString(line) {
			switch *repType {
			case RepTypeAllOccurrences:
				line = strings.ReplaceAll(
					line,
					args.ReplaceType.AllOccurrences.Occurrence,
					args.ReplaceType.AllOccurrences.NewData)
			case RepTypeEntireLine:
				line = args.ReplaceType.EntireLine.NewData
			default:
				line = args.ReplaceType.EntireLine.NewData
			}
		}

		// save each line (replaced or non-replaced) in an array
		allLines = append(allLines, line)

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return fmt.Errorf("reading file failed")
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("closing file failed")
	}

	if err = OverwriteToFile(OverwriteToFileArg{
		File:        args.File,
		Data:        allLines,
		Multiline:   true,
		Permissions: args.Permissions,
	}); err != nil {
		return fmt.Errorf("overwriting to file failed")
	}

	return nil
}
