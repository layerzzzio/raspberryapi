package infos

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/constants"
)

// Service represents several system scripts.
type Service struct{}

// Infos represents multiple system functions to get infos about the current system.
type Infos interface{}

// New creates a service instance.
func New() *Service {
	return &Service{}
}

// ReadFile reads a file
func (s Service) ReadFile(filePath string) ([]string, error) {
	var result []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening file failed")
	}

	// replacing line logic
	// source: https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	reader := bufio.NewReader(file)
	var line string

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if line != "" {
			result = append(result, strings.TrimSuffix(line, "\n"))
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return nil, fmt.Errorf("reading file failed")
	}

	if err = file.Close(); err != nil {
		return nil, fmt.Errorf("closing file failed")
	}

	return result, nil
}

// IsFileExists checks if a file exists
func (s Service) IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}

// GetConfigFiles returns a map of the unix config file used in the Raspberry Pi
// source: https://qvault.io/2019/10/21/golang-constant-maps-slices/
func (s Service) GetConfigFiles() map[string]rpi.ConfigFileDetails {
	return map[string]rpi.ConfigFileDetails{
		"bootconfig": {
			Path:        "/boot/config.txt",
			IsCritical:  true,
			Description: "contains some system configuration parameters. It is read at boot time by the device.",
		},
		"etcpasswd": {
			Path:        "/etc/passwd",
			IsCritical:  true,
			Description: "is a text-based database of information about users that may log into the system or other operating system user identities that own running processes.",
		},
		"waitfornetwork": {
			Path:        "/etc/systemd/system/dhcpcd.service.d/wait.conf",
			IsCritical:  false,
			Description: "is a configuration file that forces the dhcp service to wait for the network to be configured before running.",
		},
		"hosts": {
			Path:        "/etc/hosts",
			IsCritical:  true,
			Description: "is a text file that associates IP addresses with hostnames, one line per IP address.",
		},
		"hostname": {
			Path:        "/etc/hostname",
			IsCritical:  true,
			Description: "configures the name of the local system. It contains a single newline-terminated hostname string.",
		},
		"blanking": {
			Path:        "/etc/X11/xorg.conf.d/10-blanking.conf",
			IsCritical:  false,
			Description: "configures the blanking behavior of the monitor.",
		},
		"start_x_elf": {
			Path:        "/boot/start_x.elf",
			IsCritical:  true,
			Description: "is a binary blob (firmware) that is loaded on to the VideoCore in the SoC and that includes camera drivers and codec.",
		},
		"rgpio_public_conf": {
			Path:        "/etc/systemd/system/pigpiod.service.d/public.conf",
			IsCritical:  false,
			Description: "is the daemon file for the remote GPIO service.",
		},
		"iso3166": {
			Path:        "/usr/share/zoneinfo/iso3166.tab",
			IsCritical:  false,
			Description: "is a file containing the standards published by the International Organization for Standardization (ISO) that defines codes for the names of countries, dependent territories, special areas of geographical interest, and their principal subdivisions (e.g., provinces or states).",
		},
	}
}

// GetEnrichedConfigFiles returns the list of config file with some extra fields
func (s Service) GetEnrichedConfigFiles(configFiles map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails {
	for k, v := range configFiles {
		stat, err := os.Stat(v.Path)
		if err != nil {
			configFiles[k] = rpi.ConfigFileDetails{
				Path:        v.Path,
				Description: v.Description,
				Name:        filepath.Base(v.Path),
				IsExist:     false,
			}
		} else {
			configFiles[k] = rpi.ConfigFileDetails{
				Path:         v.Path,
				Description:  v.Description,
				Name:         filepath.Base(v.Path),
				IsExist:      true,
				LastModified: uint64(stat.ModTime().Unix()),
				Size:         stat.Size(),
			}
		}
	}

	return configFiles
}

// IsXscreenSaverInstalled checks if XscreenSaver is installed
func (s Service) IsXscreenSaverInstalled() (bool, error) {
	isInstalled := false

	res, err := exec.Command(
		"sh",
		"-c",
		"dpkg -l xscreensaver | tail -n 1 | cut -d ' ' -f 1",
	).Output()

	if err != nil {
		err = fmt.Errorf("checking xscreensaver installation failed")
	} else {
		if string(res) == "ii" {
			isInstalled = true
		}
	}

	return isInstalled, err
}

// IsCommandStatus checks if a command returns 0 (found) or 1 (not found)
func (s Service) IsQuietGrep(command string, grep string, grepType string) bool {
	var grepCommand string

	if grepType == "quiet" {
		grepCommand = fmt.Sprintf("%v | grep -q %v ; echo $?", command, grep)
	} else if grepType == "word-regexp" {
		grepCommand = fmt.Sprintf("%v | grep -q -w %v ; echo $?", command, grep)
	} else {
		log.Fatal("bad grep type")
	}

	res, err := exec.Command("sh", "-c", grepCommand).Output()
	if err != nil {
		log.Fatal(err)
	}

	resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
	if err != nil {
		log.Fatal(err)
	}

	if resNum > 0 {
		return false
	} else {
		return true
	}
}

// IsSSHKeyGenerating checks if SSH keys are getting generated
func (s Service) IsSSHKeyGenerating(path string) bool {
	isFinished, err := exec.Command(
		"sh",
		"-c",
		fmt.Sprintf("grep -q \"^finished\" %v ; echo $?", path),
	).Output()

	if err != nil {
		log.Fatal(err)
	}

	isFinishedNum, err := strconv.Atoi(strings.TrimSpace(string(isFinished)))
	if err != nil {
		log.Fatal(err)
	}

	isSSHLogExist := s.IsFileExists(path)

	if isFinishedNum == 1 && isSSHLogExist {
		return true
	} else {
		return false
	}
}

// IsDPKGInstalled checks if package is installed with dpkg
func (s Service) IsDPKGInstalled(packageName string) bool {
	command := fmt.Sprintf("dpkg -l %v 2> /dev/null | tail -n 1 | cut -d ' ' -f 1", packageName)
	res, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}

	if strings.TrimSpace(string(res)) == "ii" {
		return true
	} else {
		return false
	}
}

// IsSPI checks if SPI is enabled or disabled
func (s Service) IsSPI(path string) bool {
	command := fmt.Sprintf(
		"grep -q -E \"^(device_tree_param|dtparam)=([^,]*,)*spi(=(on|true|yes|1))?(,.*)?$\" %v ; echo $?",
		path,
	)
	res, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}

	resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
	if err != nil {
		log.Fatal(err)
	}

	if resNum == 1 {
		return false
	} else {
		return true
	}
}

// IsI2C checks if I2C is enabled or disabled
func (s Service) IsI2C(path string) bool {
	command := fmt.Sprintf(
		"grep -q -E \"^(device_tree_param|dtparam)=([^,]*,)*i2c(_arm)?(=(on|true|yes|1))?(,.*)?$\" %v ; echo $?",
		path,
	)
	res, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}

	resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
	if err != nil {
		log.Fatal(err)
	}

	if resNum == 1 {
		return false
	} else {
		return true
	}
}

// IsVariableSet checks if a variable equals a certain value in a file
func (s Service) IsVariableSet(rawLines []string, key string, value string) bool {
	reg := `^\s*` + key + `\s*=\s*` + value + `\s*#?.*`

	isMatch := false

	for _, line := range rawLines {
		if line != "" {
			re := regexp.MustCompile(reg)
			if re.MatchString(line) {
				isMatch = true
			}
		}
	}

	return isMatch
}

// ListWifiInterfaces returns a list of wifi interfaces
func (s Service) ListWifiInterfaces(directoryPath string) []string {
	var wifiInterfaces []string

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal()
	}

	for _, f := range files {
		wirelessPath := fmt.Sprintf("%v/%v/wireless", directoryPath, f.Name())
		fmt.Println(wirelessPath)
		if s.IsFileExists(wirelessPath) {
			wifiInterfaces = append(wifiInterfaces, f.Name())
		}
	}

	return wifiInterfaces
}

func (s Service) IsWpaSupCom() map[string]bool {
	ifaces := s.ListWifiInterfaces(constants.NETWORKINTERFACES)
	result := map[string]bool{}

	for _, i := range ifaces {
		command := fmt.Sprintf(
			"wpa_cli -i %v status > /dev/null 2>&1 ; echo $?",
			i,
		)
		res, err := exec.Command("sh", "-c", command).Output()
		if err != nil {
			log.Fatal(err)
		}

		resNum, err := strconv.Atoi(strings.TrimSpace(string(res)))
		if err != nil {
			log.Fatal(err)
		}

		if resNum == 1 {
			result[i] = false
		} else {
			result[i] = true
		}
	}
	return result
}

func (s Service) ZoneInfo(filePath string) map[string]string {
	result := make(map[string]string)
	zi, err := s.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range zi {
		if !strings.HasPrefix(v, "#") {
			split := strings.SplitAfter(v, "\t")
			countryCode := strings.ReplaceAll(split[0], "\t", "")
			countryName := split[1]
			result[countryCode] = countryName
		}
	}

	return result
}

// ListNameFilesInDirectory lists all files in directory
func (s Service) ListNameFilesInDirectory(directoryPath string) []string {
	var result []string

	// files, err := ioutil.ReadDir(directoryPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	if !file.IsDir() {
	// 		result = append(result, file.Name())
	// 	}
	// }

	err := godirwalk.Walk(directoryPath, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if !de.IsDir() {
				result = append(result, de.Name())
			}
			return nil
		},
		// (optional) set true for faster yet non-deterministic enumeration (see godoc)
		Unsorted: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(result)

	return result
}
