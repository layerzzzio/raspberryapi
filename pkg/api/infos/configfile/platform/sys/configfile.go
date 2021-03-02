package sys

import (
	"fmt"

	"github.com/raspibuddy/rpi"
)

// ConfigFile represents a ConfigFile entity on the current system.
type ConfigFile struct{}

// List returns a list of ConfigFile info
func (co ConfigFile) List(config map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error) {
	isFilesMissing := false
	isCriticalFilesMissing := false
	criticalFilesMissing := []string{}
	filesMissing := []string{}

	var configFiles []rpi.ConfigFileDetails

	for _, v := range config {
		if !v.IsExist {
			if v.IsCritical {
				isCriticalFilesMissing = true
				criticalFilesMissing = append(criticalFilesMissing, v.Path)
			}
			isFilesMissing = true
			filesMissing = append(filesMissing, v.Path)
		}

		configFiles = append(configFiles, v)
	}

	fmt.Println(rpi.ConfigFile{
		IsFilesMissing:         isFilesMissing,
		IsCriticalFilesMissing: isCriticalFilesMissing,
		FilesMissing:           filesMissing,
		CriticalFilesMissing:   criticalFilesMissing,
		ConfigFiles:            configFiles,
	})

	return rpi.ConfigFile{
		IsFilesMissing:         isFilesMissing,
		IsCriticalFilesMissing: isCriticalFilesMissing,
		FilesMissing:           filesMissing,
		CriticalFilesMissing:   criticalFilesMissing,
		ConfigFiles:            configFiles,
	}, nil
}
