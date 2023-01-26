package configfile_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		cofsys     mocksys.ConfigFile
		wantedData rpi.ConfigFile
		wantedErr  error
	}{
		{
			name: "success",
			infos: mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"etcpasswd": {
							Path: "/etc/passwd",
							Name: "passwd",
						},
					}
				},
				GetEnrichedConfigFilesFn: func(map[string]rpi.ConfigFileDetails) map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"etcpasswd": {
							Path:         "/etc/passwd",
							Name:         "passwd",
							IsExist:      true,
							Size:         1,
							LastModified: 2,
							Description:  "dummy desc",
						},
					}
				},
			},
			cofsys: mocksys.ConfigFile{
				ListFn: func(map[string]rpi.ConfigFileDetails) (rpi.ConfigFile, error) {
					return rpi.ConfigFile{
						IsFilesMissing: true,
						ConfigFiles: []rpi.ConfigFileDetails{
							{
								Path:         "/etc/passwd",
								Name:         "passwd",
								IsExist:      true,
								Size:         1,
								LastModified: 2,
								Description:  "dummy desc",
							},
						},
					}, nil
				},
			},
			wantedData: rpi.ConfigFile{
				IsFilesMissing: true,
				ConfigFiles: []rpi.ConfigFileDetails{
					{
						Path:         "/etc/passwd",
						Name:         "passwd",
						IsExist:      true,
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
			s := configfile.New(&tc.cofsys, tc.infos)
			enrichedConfigFiles, err := s.List()
			assert.Equal(t, tc.wantedData, enrichedConfigFiles)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
