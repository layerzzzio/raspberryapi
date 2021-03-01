package infos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
			Description: "contains some system configuration parameters. It is read at boot time by the device.",
		},
		"etcpasswd": {
			Path:        "/etc/passwd",
			Description: "is a text-based database of information about users that may log into the system or other operating system user identities that own running processes.",
		},
		"waitfornetwork": {
			Path:        "/etc/systemd/system/dhcpcd.service.d/wait.conf",
			Description: "is a configuration file that forces the dhcp service to wait for the network to be configured before running.",
		},
		"hosts": {
			Path:        "/etc/hosts",
			Description: "is a text file that associates IP addresses with hostnames, one line per IP address.",
		},
		"hostname": {
			Path:        "/etc/hostname",
			Description: "configures the name of the local system. It contains a single newline-terminated hostname string.",
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
