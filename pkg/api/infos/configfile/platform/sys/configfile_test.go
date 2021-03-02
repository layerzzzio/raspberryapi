package sys_test

import (
	"sort"
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile/platform/sys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name        string
		configFiles map[string]rpi.ConfigFileDetails
		wantedData  rpi.ConfigFile
		wantedErr   error
	}{
		{
			name: "success",
			configFiles: map[string]rpi.ConfigFileDetails{
				"/etc/passwd": {
					Path:         "/etc/passwd",
					IsCritical:   true,
					Name:         "passwd",
					IsExist:      true,
					Description:  "dummy desc",
					Size:         1,
					LastModified: 2,
				},
				"/etc/hostname": {
					Path:         "/etc/hostname",
					IsCritical:   false,
					Name:         "hostname",
					IsExist:      false,
					Description:  "dummy desc",
					LastModified: 0,
					Size:         0,
				},
				"/etc/hosts": {
					Path:         "/etc/hosts",
					IsCritical:   true,
					Name:         "hosts",
					IsExist:      false,
					Description:  "dummy desc",
					LastModified: 0,
					Size:         0,
				},
			},
			wantedData: rpi.ConfigFile{
				IsFilesMissing:         true,
				IsCriticalFilesMissing: true,
				CriticalFilesMissing: []string{
					"/etc/hosts",
				},
				FilesMissing: []string{
					"/etc/hostname",
					"/etc/hosts",
				},
				ConfigFiles: []rpi.ConfigFileDetails{
					{
						Path:         "/etc/hostname",
						Name:         "hostname",
						IsExist:      false,
						IsCritical:   false,
						Description:  "dummy desc",
						LastModified: 0,
						Size:         0,
					},
					{
						Path:         "/etc/hosts",
						Name:         "hosts",
						IsExist:      false,
						IsCritical:   true,
						Description:  "dummy desc",
						LastModified: 0,
						Size:         0,
					},
					{
						Path:         "/etc/passwd",
						Name:         "passwd",
						IsExist:      true,
						IsCritical:   true,
						Size:         1,
						LastModified: 2,
						Description:  "dummy desc",
					},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configfile.COFSYS(sys.ConfigFile{})
			enrichedConfigFiles, err := s.List(tc.configFiles)

			sort.SliceStable(enrichedConfigFiles.ConfigFiles, func(i, j int) bool {
				return enrichedConfigFiles.ConfigFiles[i].Name < enrichedConfigFiles.ConfigFiles[j].Name
			})

			sort.SliceStable(enrichedConfigFiles.FilesMissing, func(i, j int) bool {
				return enrichedConfigFiles.FilesMissing[i] < enrichedConfigFiles.FilesMissing[j]
			})

			res := rpi.ConfigFile{
				IsFilesMissing:         enrichedConfigFiles.IsFilesMissing,
				IsCriticalFilesMissing: enrichedConfigFiles.IsCriticalFilesMissing,
				CriticalFilesMissing:   enrichedConfigFiles.CriticalFilesMissing,
				FilesMissing:           enrichedConfigFiles.FilesMissing,
				ConfigFiles:            enrichedConfigFiles.ConfigFiles,
			}

			assert.Equal(t, tc.wantedData, res)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
