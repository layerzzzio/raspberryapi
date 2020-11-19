package sys

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
	"gopkg.in/go-playground/assert.v1"
)

func TestViewFL(t *testing.T) {
	cases := []struct {
		name          string
		path          string
		fileStructure *rpi.File
		flattenFiles  map[int64]string
		wantedData    rpi.FileStructure
		wantedErr     error
	}{
		{
			name: "success: len(fileStructure.Files) = 0",
			path: "/dummy/path",
			fileStructure: &rpi.File{
				Name:  "/dummy/path",
				Size:  10,
				IsDir: false,
				Files: []*rpi.File{},
			},
			flattenFiles: map[int64]string{},
			wantedData: rpi.FileStructure{
				DirectoryPath: "/dummy/path",
			},
			wantedErr: nil,
		},
		{
			name: "success: flattenFiles = 0",
			path: "/dummy/path",
			fileStructure: &rpi.File{
				Name:  "/dummy/path",
				Size:  10,
				IsDir: false,
				Files: []*rpi.File{
					{
						Name: "dummy1",
					},
					{
						Name: "dummy2",
					},
				},
			},
			flattenFiles: map[int64]string{},
			wantedData: rpi.FileStructure{
				DirectoryPath: "/dummy/path",
			},
			wantedErr: nil,
		},
		{
			name: "success: flattenFiles > 1",
			path: "/dummy/path",
			fileStructure: &rpi.File{
				Name:  "/dummy/path",
				Size:  10,
				IsDir: false,
				Files: []*rpi.File{
					{
						Name: "dummy1",
					},
					{
						Name: "dummy2",
					},
				},
			},
			flattenFiles: map[int64]string{
				100: "/dummy/path/file1",
				200: "/dummy/path/file2",
			},
			wantedData: rpi.FileStructure{
				DirectoryPath: "/dummy/path",
				LargestFiles: []*rpi.File{
					{
						Path: "/dummy/path/file1",
						Size: 100,
					},
					{
						Path: "/dummy/path/file2",
						Size: 200,
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: flattenFiles > 1",
			path: "/dummy/path",
			fileStructure: &rpi.File{
				Name:  "/dummy/path",
				Size:  10,
				IsDir: false,
				Files: []*rpi.File{
					{
						Name: "dummy1",
					},
					{
						Name: "dummy2",
					},
				},
			},
			flattenFiles: map[int64]string{
				100: "/dummy/path/file1",
			},
			wantedData: rpi.FileStructure{
				DirectoryPath: "/dummy/path",
				LargestFiles: []*rpi.File{
					{
						Path: "/dummy/path/file1",
						Size: 100,
					},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: flattenFiles > 100",
			path: "/dummy/path",
			fileStructure: &rpi.File{
				Name:  "/dummy/path",
				Size:  10,
				IsDir: false,
				Files: []*rpi.File{
					{
						Name: "dummy1",
					},
					{
						Name: "dummy2",
					},
				},
			},
			flattenFiles: flattenFiles101,
			wantedData: rpi.FileStructure{
				DirectoryPath: "/dummy/path",
				LargestFiles:  fileStructure100,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := filestructure.FSSYS(FileStructure{})
			largestFiles, err := s.ViewLF(tc.fileStructure, tc.flattenFiles)
			assert.Equal(t, tc.wantedData, largestFiles)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

var flattenFiles101 = map[int64]string{
	1:   "/dummy/path/file1",
	2:   "/dummy/path/file1",
	3:   "/dummy/path/file1",
	4:   "/dummy/path/file1",
	5:   "/dummy/path/file1",
	6:   "/dummy/path/file1",
	7:   "/dummy/path/file1",
	8:   "/dummy/path/file1",
	9:   "/dummy/path/file1",
	10:  "/dummy/path/file1",
	11:  "/dummy/path/file1",
	12:  "/dummy/path/file1",
	13:  "/dummy/path/file1",
	14:  "/dummy/path/file1",
	15:  "/dummy/path/file1",
	16:  "/dummy/path/file1",
	17:  "/dummy/path/file1",
	18:  "/dummy/path/file1",
	19:  "/dummy/path/file1",
	20:  "/dummy/path/file1",
	21:  "/dummy/path/file1",
	22:  "/dummy/path/file1",
	23:  "/dummy/path/file1",
	24:  "/dummy/path/file1",
	25:  "/dummy/path/file1",
	26:  "/dummy/path/file1",
	27:  "/dummy/path/file1",
	28:  "/dummy/path/file1",
	29:  "/dummy/path/file1",
	30:  "/dummy/path/file1",
	31:  "/dummy/path/file1",
	32:  "/dummy/path/file1",
	33:  "/dummy/path/file1",
	34:  "/dummy/path/file1",
	35:  "/dummy/path/file1",
	36:  "/dummy/path/file1",
	37:  "/dummy/path/file1",
	38:  "/dummy/path/file1",
	39:  "/dummy/path/file1",
	40:  "/dummy/path/file1",
	41:  "/dummy/path/file1",
	42:  "/dummy/path/file1",
	43:  "/dummy/path/file1",
	44:  "/dummy/path/file1",
	45:  "/dummy/path/file1",
	46:  "/dummy/path/file1",
	47:  "/dummy/path/file1",
	48:  "/dummy/path/file1",
	49:  "/dummy/path/file1",
	50:  "/dummy/path/file1",
	51:  "/dummy/path/file1",
	52:  "/dummy/path/file1",
	53:  "/dummy/path/file1",
	54:  "/dummy/path/file1",
	55:  "/dummy/path/file1",
	56:  "/dummy/path/file1",
	57:  "/dummy/path/file1",
	58:  "/dummy/path/file1",
	59:  "/dummy/path/file1",
	60:  "/dummy/path/file1",
	61:  "/dummy/path/file1",
	62:  "/dummy/path/file1",
	63:  "/dummy/path/file1",
	64:  "/dummy/path/file1",
	65:  "/dummy/path/file1",
	66:  "/dummy/path/file1",
	67:  "/dummy/path/file1",
	68:  "/dummy/path/file1",
	69:  "/dummy/path/file1",
	70:  "/dummy/path/file1",
	71:  "/dummy/path/file1",
	72:  "/dummy/path/file1",
	73:  "/dummy/path/file1",
	74:  "/dummy/path/file1",
	75:  "/dummy/path/file1",
	76:  "/dummy/path/file1",
	77:  "/dummy/path/file1",
	78:  "/dummy/path/file1",
	79:  "/dummy/path/file1",
	80:  "/dummy/path/file1",
	81:  "/dummy/path/file1",
	82:  "/dummy/path/file1",
	83:  "/dummy/path/file1",
	84:  "/dummy/path/file1",
	85:  "/dummy/path/file1",
	86:  "/dummy/path/file1",
	87:  "/dummy/path/file1",
	88:  "/dummy/path/file1",
	89:  "/dummy/path/file1",
	90:  "/dummy/path/file1",
	91:  "/dummy/path/file1",
	92:  "/dummy/path/file1",
	93:  "/dummy/path/file1",
	94:  "/dummy/path/file1",
	95:  "/dummy/path/file1",
	96:  "/dummy/path/file1",
	97:  "/dummy/path/file1",
	98:  "/dummy/path/file1",
	99:  "/dummy/path/file1",
	100: "/dummy/path/file1",
	101: "/dummy/path/file1",
}

var fileStructure100 = []*rpi.File{
	{
		Path: "/dummy/path/file1",
		Size: 2,
	},
	{
		Path: "/dummy/path/file1",
		Size: 3,
	},
	{
		Path: "/dummy/path/file1",
		Size: 4,
	},
	{
		Path: "/dummy/path/file1",
		Size: 5,
	},
	{
		Path: "/dummy/path/file1",
		Size: 6,
	},
	{
		Path: "/dummy/path/file1",
		Size: 7,
	},
	{
		Path: "/dummy/path/file1",
		Size: 8,
	},
	{
		Path: "/dummy/path/file1",
		Size: 9,
	},
	{
		Path: "/dummy/path/file1",
		Size: 10,
	},
	{
		Path: "/dummy/path/file1",
		Size: 11,
	},
	{
		Path: "/dummy/path/file1",
		Size: 12,
	},
	{
		Path: "/dummy/path/file1",
		Size: 13,
	},
	{
		Path: "/dummy/path/file1",
		Size: 14,
	},
	{
		Path: "/dummy/path/file1",
		Size: 15,
	},
	{
		Path: "/dummy/path/file1",
		Size: 16,
	},
	{
		Path: "/dummy/path/file1",
		Size: 17,
	},

	{
		Path: "/dummy/path/file1",
		Size: 18,
	},
	{
		Path: "/dummy/path/file1",
		Size: 19,
	},
	{
		Path: "/dummy/path/file1",
		Size: 20,
	},
	{
		Path: "/dummy/path/file1",
		Size: 21,
	},
	{
		Path: "/dummy/path/file1",
		Size: 22,
	},
	{
		Path: "/dummy/path/file1",
		Size: 23,
	},
	{
		Path: "/dummy/path/file1",
		Size: 24,
	},
	{
		Path: "/dummy/path/file1",
		Size: 25,
	},
	{
		Path: "/dummy/path/file1",
		Size: 26,
	},
	{
		Path: "/dummy/path/file1",
		Size: 27,
	},
	{
		Path: "/dummy/path/file1",
		Size: 28,
	},
	{
		Path: "/dummy/path/file1",
		Size: 29,
	},
	{
		Path: "/dummy/path/file1",
		Size: 30,
	},
	{
		Path: "/dummy/path/file1",
		Size: 31,
	},
	{
		Path: "/dummy/path/file1",
		Size: 32,
	},
	{
		Path: "/dummy/path/file1",
		Size: 33,
	},
	{
		Path: "/dummy/path/file1",
		Size: 34,
	},
	{
		Path: "/dummy/path/file1",
		Size: 35,
	},
	{
		Path: "/dummy/path/file1",
		Size: 36,
	},
	{
		Path: "/dummy/path/file1",
		Size: 37,
	},
	{
		Path: "/dummy/path/file1",
		Size: 38,
	},
	{
		Path: "/dummy/path/file1",
		Size: 39,
	},
	{
		Path: "/dummy/path/file1",
		Size: 40,
	},
	{
		Path: "/dummy/path/file1",
		Size: 41,
	},
	{
		Path: "/dummy/path/file1",
		Size: 42,
	},
	{
		Path: "/dummy/path/file1",
		Size: 43,
	},
	{
		Path: "/dummy/path/file1",
		Size: 44,
	},
	{
		Path: "/dummy/path/file1",
		Size: 45,
	},
	{
		Path: "/dummy/path/file1",
		Size: 46,
	},
	{
		Path: "/dummy/path/file1",
		Size: 47,
	},
	{
		Path: "/dummy/path/file1",
		Size: 48,
	},
	{
		Path: "/dummy/path/file1",
		Size: 49,
	},
	{
		Path: "/dummy/path/file1",
		Size: 50,
	},
	{
		Path: "/dummy/path/file1",
		Size: 51,
	},
	{
		Path: "/dummy/path/file1",
		Size: 52,
	},
	{
		Path: "/dummy/path/file1",
		Size: 53,
	},
	{
		Path: "/dummy/path/file1",
		Size: 54,
	},
	{
		Path: "/dummy/path/file1",
		Size: 55,
	},
	{
		Path: "/dummy/path/file1",
		Size: 56,
	},
	{
		Path: "/dummy/path/file1",
		Size: 57,
	},

	{
		Path: "/dummy/path/file1",
		Size: 58,
	},
	{
		Path: "/dummy/path/file1",
		Size: 59,
	},
	{
		Path: "/dummy/path/file1",
		Size: 60,
	},
	{
		Path: "/dummy/path/file1",
		Size: 61,
	},
	{
		Path: "/dummy/path/file1",
		Size: 62,
	},
	{
		Path: "/dummy/path/file1",
		Size: 63,
	},
	{
		Path: "/dummy/path/file1",
		Size: 64,
	},
	{
		Path: "/dummy/path/file1",
		Size: 65,
	},
	{
		Path: "/dummy/path/file1",
		Size: 66,
	},
	{
		Path: "/dummy/path/file1",
		Size: 67,
	},
	{
		Path: "/dummy/path/file1",
		Size: 68,
	},
	{
		Path: "/dummy/path/file1",
		Size: 69,
	},
	{
		Path: "/dummy/path/file1",
		Size: 70,
	},
	{
		Path: "/dummy/path/file1",
		Size: 71,
	},
	{
		Path: "/dummy/path/file1",
		Size: 72,
	},
	{
		Path: "/dummy/path/file1",
		Size: 73,
	},
	{
		Path: "/dummy/path/file1",
		Size: 74,
	},
	{
		Path: "/dummy/path/file1",
		Size: 75,
	},
	{
		Path: "/dummy/path/file1",
		Size: 76,
	},
	{
		Path: "/dummy/path/file1",
		Size: 77,
	},
	{
		Path: "/dummy/path/file1",
		Size: 78,
	},
	{
		Path: "/dummy/path/file1",
		Size: 79,
	},
	{
		Path: "/dummy/path/file1",
		Size: 80,
	},
	{
		Path: "/dummy/path/file1",
		Size: 81,
	},
	{
		Path: "/dummy/path/file1",
		Size: 82,
	},
	{
		Path: "/dummy/path/file1",
		Size: 83,
	},
	{
		Path: "/dummy/path/file1",
		Size: 84,
	},
	{
		Path: "/dummy/path/file1",
		Size: 85,
	},
	{
		Path: "/dummy/path/file1",
		Size: 86,
	},
	{
		Path: "/dummy/path/file1",
		Size: 87,
	},
	{
		Path: "/dummy/path/file1",
		Size: 88,
	},
	{
		Path: "/dummy/path/file1",
		Size: 89,
	},
	{
		Path: "/dummy/path/file1",
		Size: 90,
	},
	{
		Path: "/dummy/path/file1",
		Size: 91,
	},
	{
		Path: "/dummy/path/file1",
		Size: 92,
	},
	{
		Path: "/dummy/path/file1",
		Size: 93,
	},
	{
		Path: "/dummy/path/file1",
		Size: 94,
	},
	{
		Path: "/dummy/path/file1",
		Size: 95,
	},
	{
		Path: "/dummy/path/file1",
		Size: 96,
	},
	{
		Path: "/dummy/path/file1",
		Size: 97,
	},
	{
		Path: "/dummy/path/file1",
		Size: 98,
	},
	{
		Path: "/dummy/path/file1",
		Size: 99,
	},
	{
		Path: "/dummy/path/file1",
		Size: 100,
	},
	{
		Path: "/dummy/path/file1",
		Size: 101,
	},
}
