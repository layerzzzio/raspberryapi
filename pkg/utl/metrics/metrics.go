package metrics

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// TODO: to test this method by simulating different OS scenarios in a Docker container (raspbian/strech)

// Service represents several system scripts.
type Service struct {
	m Metrics
}

// Metrics represents multiple system related scripts.
type Metrics interface {
	PsPID(p *process.Process, c chan (int32))
	PsName(p *process.Process, c chan (string))
	PsCPUPer(p *process.Process, c chan (float64))
	PsMemPer(p *process.Process, c chan (float32))
	PsUsername(p *process.Process, c chan (string))
	PsCmdLine(p *process.Process, c chan (string))
	PsStatus(p *process.Process, c chan (string))
	PsCreationTime(p *process.Process, c chan (int64))
	PsBackground(p *process.Process, c chan (bool))
	PsForeground(p *process.Process, c chan (bool))
	PsIsRunning(p *process.Process, c chan (bool))
	PsParent(p *process.Process, c chan (int32))
}

// New creates a service instance.
func New(m Metrics) *Service {
	return &Service{m: m}
}

// DStats represents a tuple composed of a partition and a mountpoint stats.
type DStats struct {
	Partition  *disk.PartitionStat
	Mountpoint *disk.UsageStat
}

// PathSize represents a tuple composed of a file path and a file size
type PathSize struct {
	Path string
	Size int
}

// File structure representing files and folders with their accumulated sizes
// type File struct {
// 	Name   string
// 	Parent *File
// 	Size   int64
// 	IsDir  bool
// 	Files  []*File
// }

// PInfo represents several process key attributes.
type PInfo struct {
	ID           int32
	Name         string
	Username     string
	CommandLine  string
	Status       string
	CreationTime int64
	Foreground   bool
	Background   bool
	IsRunning    bool
	CPUPercent   float64
	MemPercent   float32
	ParentP      int32
}

// CPUInfo returns several cpu key attributes.
func (s Service) CPUInfo() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return info, nil
}

// CPUPercent returns the cpu percentage usage stats.
func (s Service) CPUPercent(interval time.Duration, perVCore bool) ([]float64, error) {
	percent, err := cpu.Percent(interval, perVCore)
	if err != nil {
		return nil, err
	}
	return percent, nil
}

// CPUTimes returns some cpu times usage stats.
func (s Service) CPUTimes(perVCore bool) ([]cpu.TimesStat, error) {
	times, err := cpu.Times(perVCore)
	if err != nil {
		return nil, err
	}
	return times, nil
}

// SwapMemory returns the swap memory usage.
func (s Service) SwapMemory() (mem.SwapMemoryStat, error) {
	smem, err := mem.SwapMemory()
	if err != nil {
		return mem.SwapMemoryStat{}, err
	}
	return *smem, nil
}

// VirtualMemory returns the virtual memory usage.
func (s Service) VirtualMemory() (mem.VirtualMemoryStat, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return mem.VirtualMemoryStat{}, err
	}
	return *vmem, nil
}

// DiskStats returns some disk usage stats.
func (s Service) DiskStats(all bool) (map[string][]DStats, error) {
	dstats := make(map[string][]DStats)

	disks, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}

	for _, d := range disks {
		usage, err := disk.Usage(d.Mountpoint)
		if err != nil {
			return nil, err
		}

		dstats[d.Device] = append(
			dstats[d.Device],
			DStats{
				&disk.PartitionStat{
					Fstype:     d.Fstype,
					Opts:       d.Opts,
					Mountpoint: d.Mountpoint,
				},
				&disk.UsageStat{
					Path:              usage.Path,
					Fstype:            usage.Fstype,
					Total:             usage.Total,
					Free:              usage.Free,
					Used:              usage.Used,
					UsedPercent:       usage.UsedPercent,
					InodesTotal:       usage.InodesTotal,
					InodesFree:        usage.InodesFree,
					InodesUsed:        usage.InodesUsed,
					InodesUsedPercent: usage.InodesUsedPercent,
				},
			},
		)
	}

	return dstats, nil
}

// LoadAvg returns some host load stats.
func (s Service) LoadAvg() (load.AvgStat, error) {
	temp, err := load.Avg()
	if err != nil {
		return load.AvgStat{}, err
	}
	return *temp, nil
}

// LoadProcs returns some host procs stats.
func (s Service) LoadProcs() (load.MiscStat, error) {
	procs, err := load.Misc()
	if err != nil {
		return load.MiscStat{}, err
	}
	return *procs, nil
}

//Processes returns some host process related stats.
func (s Service) Processes(id ...int32) ([]PInfo, error) {
	var pinfo []PInfo
	cID := make(chan (int32))
	cName := make(chan (string))
	cCPUPer := make(chan (float64))
	cMemPer := make(chan (float32))
	var cU chan (string)
	var cCL chan (string)
	var cS chan (string)
	var cCT chan (int64)
	var cFG chan (bool)
	var cBG chan (bool)
	var cIR chan (bool)
	var cP chan (int32)

	ps, err := process.Processes()
	if err != nil {
		error.Error(err)
	}

	pid := int32(-1)
	if len(id) == 1 {
		pid = id[0]
		if id[0] > 0 {
			cU = make(chan (string))
			cCL = make(chan (string))
			cS = make(chan (string))
			cCT = make(chan (int64))
			cFG = make(chan (bool))
			cBG = make(chan (bool))
			cIR = make(chan (bool))
			cP = make(chan (int32))
		}
	} else if len(id) > 1 {
		panic("only one id is authorized")
	}

	if pid > 0 {
		for i := range ps {
			if ps[i].Pid == pid {
				go s.m.PsPID(ps[i], cID)
				go s.m.PsName(ps[i], cName)
				go s.m.PsCPUPer(ps[i], cCPUPer)
				go s.m.PsMemPer(ps[i], cMemPer)
				go s.m.PsUsername(ps[i], cU)
				go s.m.PsCmdLine(ps[i], cCL)
				go s.m.PsStatus(ps[i], cS)
				go s.m.PsCreationTime(ps[i], cCT)
				go s.m.PsForeground(ps[i], cFG)
				go s.m.PsBackground(ps[i], cBG)
				go s.m.PsIsRunning(ps[i], cIR)
				go s.m.PsParent(ps[i], cP)

				pinfo = []PInfo{
					{
						ID:           <-cID,
						Name:         <-cName,
						CPUPercent:   <-cCPUPer,
						MemPercent:   <-cMemPer,
						Username:     <-cU,
						CommandLine:  <-cCL,
						Status:       <-cS,
						CreationTime: <-cCT,
						Foreground:   <-cFG,
						Background:   <-cBG,
						IsRunning:    <-cIR,
						ParentP:      <-cP,
					},
				}

				close(cID)
				close(cName)
				close(cCPUPer)
				close(cMemPer)
				close(cU)
				close(cCL)
				close(cS)
				close(cCT)
				close(cFG)
				close(cBG)
				close(cIR)
				close(cP)

				return pinfo, nil
			}
		}
		return nil, errors.New("process not found")
	} else if pid == 0 {
		return nil, errors.New("process not found")
	}

	for i := range ps {
		go s.m.PsPID(ps[i], cID)
		go s.m.PsName(ps[i], cName)
		go s.m.PsCPUPer(ps[i], cCPUPer)
		go s.m.PsMemPer(ps[i], cMemPer)

		pinfo = append(pinfo,
			PInfo{
				ID:         <-cID,
				Name:       <-cName,
				CPUPercent: <-cCPUPer,
				MemPercent: <-cMemPer,
			})
	}

	close(cID)
	close(cName)
	close(cCPUPer)
	close(cMemPer)

	return pinfo, nil
}

// PsPID feeds a channel with a process id.
func (s Service) PsPID(p *process.Process, c chan (int32)) {
	c <- p.Pid
}

// PsName feeds a channel with a process name.
func (s Service) PsName(p *process.Process, c chan (string)) {
	name, err := p.Name()
	if err != nil {
		log.Error()
	}
	c <- name
}

// PsCPUPer feeds a channel with a process cpu usage.
func (s Service) PsCPUPer(p *process.Process, c chan (float64)) {
	cpuper, err := p.CPUPercent()
	if err != nil {
		log.Error()
	}
	c <- cpuper
}

// PsMemPer feeds a channel with a process memory usage.
func (s Service) PsMemPer(p *process.Process, c chan (float32)) {
	memper, err := p.MemoryPercent()
	if err != nil {
		log.Error()
	}
	c <- memper
}

// PsUsername feeds a channel with a process username.
func (s Service) PsUsername(p *process.Process, c chan (string)) {
	u, err := p.Username()
	if err != nil {
		log.Error()
	}
	c <- u
}

// PsCmdLine feeds a channel with a process command line.
func (s Service) PsCmdLine(p *process.Process, c chan (string)) {
	cl, err := p.Cmdline()
	if err != nil {
		log.Error()
	}
	c <- cl
}

// PsStatus feeds a channel with a process status.
func (s Service) PsStatus(p *process.Process, c chan (string)) {
	st, err := p.Status()
	if err != nil {
		log.Error()
	}
	c <- st
}

// PsCreationTime feeds a channel with a process creation time.
func (s Service) PsCreationTime(p *process.Process, c chan (int64)) {
	ct, err := p.CreateTime()
	if err != nil {
		log.Error()
	}

	c <- ct
}

// PsBackground feeds a channel with a process background value.
func (s Service) PsBackground(p *process.Process, c chan (bool)) {
	bg, err := p.Background()
	if err != nil {
		log.Error()
	}
	c <- bg
}

// PsForeground feeds a channel with a process foreground value.
func (s Service) PsForeground(p *process.Process, c chan (bool)) {
	fg, err := p.Foreground()
	if err != nil {
		log.Error()
	}
	c <- fg
}

// PsIsRunning feeds a channel with a process running status value.
func (s Service) PsIsRunning(p *process.Process, c chan (bool)) {
	ir, err := p.IsRunning()
	if err != nil {
		log.Error()
	}
	c <- ir
}

// PsParent feeds a channel with a process parent id.
func (s Service) PsParent(p *process.Process, c chan (int32)) {
	var ppid int32

	ps, err := p.Parent()

	if err != nil {
		log.Error()
	}

	if ps == nil {
		ppid = -1
	} else {
		ppid = ps.Pid
	}

	c <- ppid
}

// HostInfo returns some host stats.
func (s Service) HostInfo() (host.InfoStat, error) {
	info, err := host.Info()
	if err != nil {
		return host.InfoStat{}, err
	}
	return *info, nil
}

// Users returns some host users stats.
func (s Service) Users() ([]host.UserStat, error) {
	users, err := host.Users()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Temperature returns the host temperature.
func (s Service) Temperature() (string, string, error) {
	cmd := exec.Command("sh", "-c", "vcgencmd measure_temp")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Error()
	}
	outStd, errStd := stdout.String(), stderr.String()
	return outStd, errStd, nil
}

// RaspModel returns the host Raspberry Model.
func (s Service) RaspModel() (string, string, error) {
	cmd := exec.Command("cat", "/sys/firmware/devicetree/base/model")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Error()
	}
	outStd, errStd := strings.TrimSpace(stdout.String()), stderr.String()
	return outStd, errStd, nil
}

// NetInfo returns the host net interface info.
func (s Service) NetInfo() ([]net.InterfaceStat, error) {
	netInfo, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return netInfo, nil
}

// NetStats returns the host net interface stats.
func (s Service) NetStats() ([]net.IOCountersStat, error) {
	netStats, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	return netStats, nil
}

// Path builds a file system location for given file
func Path(f *rpi.File) string {
	if f.Parent == nil {
		return f.Name
	}
	return filepath.Join(Path(f.Parent), f.Name)
}

// FlattenFile goes through subfiles and subfolders and update a map composed of file size & name
func FlattenFiles(f *rpi.File, flattenFiles map[int64]string) {
	for _, child := range f.Files {
		if child.IsDir {
			FlattenFiles(child, flattenFiles)
		} else {
			flattenFiles[child.Size] = Path(child)
		}
	}
}

// DirSize measure the size of a directory
func DirSize(path string) (float64, string) {
	pathClean := strings.ReplaceAll(path, " ", "\\ ")
	cmd := exec.Command("sh", "-c", "du -k -d0 "+pathClean+" | awk '{print $1}'")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Error()
	}

	outStdStr := strings.TrimSpace(stdout.String())

	if outStdStr == "" {
		return 0, "forbidden"
	}

	outStd, errOut := strconv.ParseFloat(outStdStr, 64)
	if errOut != nil {
		panic("errOut is not nil")
	}

	return outStd, stderr.String()
}

// UpdateSize goes through subfiles and subfolders and accumulates their size
func UpdateSize(f *rpi.File) {
	if !f.IsDir {
		return
	}
	var size int64
	for _, child := range f.Files {
		UpdateSize(child)
		size += child.Size
	}
	f.Size = size
}

// ReadDir function can return list of files for given folder path
type ReadDir func(dirname string) ([]os.FileInfo, error)

// ShouldIgnoreFolder function decides whether a folder should be ignored
type ShouldIgnoreFolder func(absolutePath string) bool

func ignoringReadDir(shouldIgnore ShouldIgnoreFolder, originalReadDir ReadDir) ReadDir {
	return func(path string) ([]os.FileInfo, error) {
		if shouldIgnore(path) {
			return []os.FileInfo{}, nil
		}
		return originalReadDir(path)
	}
}

// WalkFolder will go through a given folder and subfolders and produces file structure with aggregated file sizes
func (s Service) WalkFolder(
	path string,
	readDir ReadDir,
	pathSize uint64,
	fileLimit float32,
	ignoreFunction ShouldIgnoreFolder,
	progress chan (int),
) (*rpi.File, map[int64]string) {
	var flattenFiles = make(map[int64]string)
	var wg sync.WaitGroup
	c := make(chan bool, 2*runtime.NumCPU())
	root := walkSubFolderConcurrently(
		path,
		nil,
		ignoringReadDir(ignoreFunction, readDir),
		pathSize,
		fileLimit,
		c,
		&wg,
		progress)

	wg.Wait()
	close(progress)
	UpdateSize(root)
	FlattenFiles(root, flattenFiles)
	return root, flattenFiles
}

func updateProgress(progress chan<- int, count *int) {
	if *count > 0 {
		progress <- *count
	}
}

func walkSubFolderConcurrently(
	path string,
	parent *rpi.File,
	readDir ReadDir,
	pathSize uint64,
	fileLimit float32,
	c chan bool,
	wg *sync.WaitGroup,
	progress chan int,
) *rpi.File {
	result := &rpi.File{}
	entries, err := readDir(path)
	if err != nil {
		log.Print(err)
		return result
	}
	dirName, name := filepath.Split(path)
	result.Files = make([]*rpi.File, 0, len(entries))
	numSubFolders := 0

	defer updateProgress(progress, &numSubFolders)
	var mutex sync.Mutex
	for _, entry := range entries {
		fileRatio := float64(entry.Size()) / float64(pathSize)
		fmt.Printf("file %v - fileRatio %v > fileLimit %v\n", entry.Name(), fileRatio, float64(fileLimit)/100)
		if entry.IsDir() {
			numSubFolders++
			subFolderPath := filepath.Join(path, entry.Name())
			wg.Add(1)
			go func() {
				c <- true
				subFolder := walkSubFolderConcurrently(subFolderPath, result, readDir, pathSize, fileLimit, c, wg, progress)
				mutex.Lock()
				result.Files = append(result.Files, subFolder)
				mutex.Unlock()
				<-c
				wg.Done()
			}()
			// check if file size > x% of path size
		} else if fileRatio > float64(fileLimit)/100 && fileRatio <= 1 {
			size := entry.Size()
			file := &rpi.File{
				Name:   entry.Name(),
				Parent: result,
				Size:   size,
				IsDir:  false,
				Files:  []*rpi.File{},
			}
			mutex.Lock()
			result.Files = append(result.Files, file)
			mutex.Unlock()
		} else {
			continue
		}
	}

	if parent != nil {
		result.Name = name
		result.Parent = parent
	} else {
		// Root dir
		// TODO unit test this Join
		result.Name = filepath.Join(dirName, name)
	}
	result.IsDir = true
	return result
}

func IgnoreBasedOnIgnoreFile(ignoreFile []string) ShouldIgnoreFolder {
	ignoredFolders := map[string]struct{}{}
	for _, line := range ignoreFile {
		ignoredFolders[line] = struct{}{}
	}
	return func(absolutePath string) bool {
		_, name := filepath.Split(absolutePath)
		// _, name := filepath.Split(absolutePath)
		_, ignored := ignoredFolders[name]
		return ignored
	}
}

func ReadIgnoreFile() []string {
	usr, err := user.Current()
	if err != nil {
		log.Print("Wasn't able to retrieve current user at runtime")
		return []string{}
	}
	ignoreFileName := filepath.Join(usr.HomeDir, ".goduignore")
	if _, err := os.Stat(ignoreFileName); os.IsNotExist(err) {
		return []string{}
	}
	ignoreFile, err := os.Open(ignoreFileName)
	if err != nil {
		log.Printf("Failed to read ingorefile because %s\n", err.Error())
		return []string{}
	}
	defer ignoreFile.Close()
	scanner := bufio.NewScanner(ignoreFile)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Top100Files returns the top 100 largest files in path.
// func (s Service) Top100Files(path string) ([]PathSize, string, error) {
// 	cmd := exec.Command("sh", "-c", "find "+path+" -type f -printf '%s<sep>%p<end>\n' | sort -n -r | head -100")
// 	// on mac with zsh and gfind installed
// 	// cmd := exec.Command("zsh", "-c", "gfind "+path+" -type f -printf '%s<sep>%p<end>\n' | sort -n -r | head -100")
// 	var stdout, stderr bytes.Buffer
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()

// 	if err != nil {
// 		log.Error()
// 	}

// 	allFiles := make([]PathSize, 100)
// 	allFilesWithSep := strings.Split(strings.TrimSpace(stdout.String()), "<end>\n")

// 	for index, file := range allFilesWithSep {
// 		sizePath := strings.Split(file, "<sep>")
// 		allFiles[index].Path = strings.ReplaceAll(sizePath[1], "<end>", "")

// 		size, err := strconv.Atoi(sizePath[0])
// 		if err != nil {
// 			size = -1
// 		}

// 		allFiles[index].Size = size
// 	}

// 	outStd, errStd := allFiles, stderr.String()
// 	return outStd, errStd, nil
// }
