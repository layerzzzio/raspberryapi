package infos

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/raspibuddy/rpi"
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
func (s Service) IsQuietGrep(command string, grep string) bool {
	res, err := exec.Command(
		"sh",
		"-c",
		fmt.Sprintf("%v | grep -q %v ; echo $?", command, grep),
	).Output()

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
