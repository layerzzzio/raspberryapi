package sys

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
	"github.com/stretchr/testify/assert"
)

func TestViewFL(t *testing.T) {
	cases := []struct {
		name          string
		path          string
		fileStructure *rpi.File
		flattenFiles  map[string]int64
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
			flattenFiles: map[string]int64{},
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
			flattenFiles: map[string]int64{},
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
			flattenFiles: map[string]int64{
				"/dummy/path/file1": 100,
				"/dummy/path/file2": 200,
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
			flattenFiles: map[string]int64{
				"/dummy/path/file1": 100,
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

var flattenFiles101 = map[string]int64{
	"/dummy/path/file1":   1,
	"/dummy/path/file2":   2,
	"/dummy/path/file3":   3,
	"/dummy/path/file4":   4,
	"/dummy/path/file5":   5,
	"/dummy/path/file6":   6,
	"/dummy/path/file7":   7,
	"/dummy/path/file8":   8,
	"/dummy/path/file9":   9,
	"/dummy/path/file10":  10,
	"/dummy/path/file11":  11,
	"/dummy/path/file12":  12,
	"/dummy/path/file13":  13,
	"/dummy/path/file14":  14,
	"/dummy/path/file15":  15,
	"/dummy/path/file16":  16,
	"/dummy/path/file17":  17,
	"/dummy/path/file18":  18,
	"/dummy/path/file19":  19,
	"/dummy/path/file20":  20,
	"/dummy/path/file21":  21,
	"/dummy/path/file22":  22,
	"/dummy/path/file23":  23,
	"/dummy/path/file24":  24,
	"/dummy/path/file25":  25,
	"/dummy/path/file26":  26,
	"/dummy/path/file27":  27,
	"/dummy/path/file28":  28,
	"/dummy/path/file29":  29,
	"/dummy/path/file30":  30,
	"/dummy/path/file31":  31,
	"/dummy/path/file32":  32,
	"/dummy/path/file33":  33,
	"/dummy/path/file34":  34,
	"/dummy/path/file35":  35,
	"/dummy/path/file36":  36,
	"/dummy/path/file37":  37,
	"/dummy/path/file38":  38,
	"/dummy/path/file39":  39,
	"/dummy/path/file40":  40,
	"/dummy/path/file41":  41,
	"/dummy/path/file42":  42,
	"/dummy/path/file43":  43,
	"/dummy/path/file44":  44,
	"/dummy/path/file45":  45,
	"/dummy/path/file46":  46,
	"/dummy/path/file47":  47,
	"/dummy/path/file48":  48,
	"/dummy/path/file49":  49,
	"/dummy/path/file50":  50,
	"/dummy/path/file51":  51,
	"/dummy/path/file52":  52,
	"/dummy/path/file53":  53,
	"/dummy/path/file54":  54,
	"/dummy/path/file55":  55,
	"/dummy/path/file56":  56,
	"/dummy/path/file57":  57,
	"/dummy/path/file58":  58,
	"/dummy/path/file59":  59,
	"/dummy/path/file60":  60,
	"/dummy/path/file61":  61,
	"/dummy/path/file62":  62,
	"/dummy/path/file63":  63,
	"/dummy/path/file64":  64,
	"/dummy/path/file65":  65,
	"/dummy/path/file66":  66,
	"/dummy/path/file67":  67,
	"/dummy/path/file68":  68,
	"/dummy/path/file69":  69,
	"/dummy/path/file70":  70,
	"/dummy/path/file71":  71,
	"/dummy/path/file72":  72,
	"/dummy/path/file73":  73,
	"/dummy/path/file74":  74,
	"/dummy/path/file75":  75,
	"/dummy/path/file76":  76,
	"/dummy/path/file77":  77,
	"/dummy/path/file78":  78,
	"/dummy/path/file79":  79,
	"/dummy/path/file80":  80,
	"/dummy/path/file81":  81,
	"/dummy/path/file82":  82,
	"/dummy/path/file83":  83,
	"/dummy/path/file84":  84,
	"/dummy/path/file85":  85,
	"/dummy/path/file86":  86,
	"/dummy/path/file87":  87,
	"/dummy/path/file88":  88,
	"/dummy/path/file89":  89,
	"/dummy/path/file90":  90,
	"/dummy/path/file91":  91,
	"/dummy/path/file92":  92,
	"/dummy/path/file93":  93,
	"/dummy/path/file94":  94,
	"/dummy/path/file95":  95,
	"/dummy/path/file96":  96,
	"/dummy/path/file97":  97,
	"/dummy/path/file98":  98,
	"/dummy/path/file99":  99,
	"/dummy/path/file100": 100,
	"/dummy/path/file101": 101,
}

var fileStructure100 = []*rpi.File{
	{
		Path: "/dummy/path/file2",
		Size: 2,
	},
	{
		Path: "/dummy/path/file3",
		Size: 3,
	},
	{
		Path: "/dummy/path/file4",
		Size: 4,
	},
	{
		Path: "/dummy/path/file5",
		Size: 5,
	},
	{
		Path: "/dummy/path/file6",
		Size: 6,
	},
	{
		Path: "/dummy/path/file7",
		Size: 7,
	},
	{
		Path: "/dummy/path/file8",
		Size: 8,
	},
	{
		Path: "/dummy/path/file9",
		Size: 9,
	},
	{
		Path: "/dummy/path/file10",
		Size: 10,
	},
	{
		Path: "/dummy/path/file11",
		Size: 11,
	},
	{
		Path: "/dummy/path/file12",
		Size: 12,
	},
	{
		Path: "/dummy/path/file13",
		Size: 13,
	},
	{
		Path: "/dummy/path/file14",
		Size: 14,
	},
	{
		Path: "/dummy/path/file15",
		Size: 15,
	},
	{
		Path: "/dummy/path/file16",
		Size: 16,
	},
	{
		Path: "/dummy/path/file17",
		Size: 17,
	},
	{
		Path: "/dummy/path/file18",
		Size: 18,
	},
	{
		Path: "/dummy/path/file19",
		Size: 19,
	},
	{
		Path: "/dummy/path/file20",
		Size: 20,
	},
	{
		Path: "/dummy/path/file21",
		Size: 21,
	},
	{
		Path: "/dummy/path/file22",
		Size: 22,
	},
	{
		Path: "/dummy/path/file23",
		Size: 23,
	},
	{
		Path: "/dummy/path/file24",
		Size: 24,
	},
	{
		Path: "/dummy/path/file25",
		Size: 25,
	},
	{
		Path: "/dummy/path/file26",
		Size: 26,
	},
	{
		Path: "/dummy/path/file27",
		Size: 27,
	},
	{
		Path: "/dummy/path/file28",
		Size: 28,
	},
	{
		Path: "/dummy/path/file29",
		Size: 29,
	},
	{
		Path: "/dummy/path/file30",
		Size: 30,
	},
	{
		Path: "/dummy/path/file31",
		Size: 31,
	},
	{
		Path: "/dummy/path/file32",
		Size: 32,
	},
	{
		Path: "/dummy/path/file33",
		Size: 33,
	},
	{
		Path: "/dummy/path/file34",
		Size: 34,
	},
	{
		Path: "/dummy/path/file35",
		Size: 35,
	},
	{
		Path: "/dummy/path/file36",
		Size: 36,
	},
	{
		Path: "/dummy/path/file37",
		Size: 37,
	},
	{
		Path: "/dummy/path/file38",
		Size: 38,
	},
	{
		Path: "/dummy/path/file39",
		Size: 39,
	},
	{
		Path: "/dummy/path/file40",
		Size: 40,
	},
	{
		Path: "/dummy/path/file41",
		Size: 41,
	},
	{
		Path: "/dummy/path/file42",
		Size: 42,
	},
	{
		Path: "/dummy/path/file43",
		Size: 43,
	},
	{
		Path: "/dummy/path/file44",
		Size: 44,
	},
	{
		Path: "/dummy/path/file45",
		Size: 45,
	},
	{
		Path: "/dummy/path/file46",
		Size: 46,
	},
	{
		Path: "/dummy/path/file47",
		Size: 47,
	},
	{
		Path: "/dummy/path/file48",
		Size: 48,
	},
	{
		Path: "/dummy/path/file49",
		Size: 49,
	},
	{
		Path: "/dummy/path/file50",
		Size: 50,
	},
	{
		Path: "/dummy/path/file51",
		Size: 51,
	},
	{
		Path: "/dummy/path/file52",
		Size: 52,
	},
	{
		Path: "/dummy/path/file53",
		Size: 53,
	},
	{
		Path: "/dummy/path/file54",
		Size: 54,
	},
	{
		Path: "/dummy/path/file55",
		Size: 55,
	},
	{
		Path: "/dummy/path/file56",
		Size: 56,
	},
	{
		Path: "/dummy/path/file57",
		Size: 57,
	},

	{
		Path: "/dummy/path/file58",
		Size: 58,
	},
	{
		Path: "/dummy/path/file59",
		Size: 59,
	},
	{
		Path: "/dummy/path/file60",
		Size: 60,
	},
	{
		Path: "/dummy/path/file61",
		Size: 61,
	},
	{
		Path: "/dummy/path/file62",
		Size: 62,
	},
	{
		Path: "/dummy/path/file63",
		Size: 63,
	},
	{
		Path: "/dummy/path/file64",
		Size: 64,
	},
	{
		Path: "/dummy/path/file65",
		Size: 65,
	},
	{
		Path: "/dummy/path/file66",
		Size: 66,
	},
	{
		Path: "/dummy/path/file67",
		Size: 67,
	},
	{
		Path: "/dummy/path/file68",
		Size: 68,
	},
	{
		Path: "/dummy/path/file69",
		Size: 69,
	},
	{
		Path: "/dummy/path/file70",
		Size: 70,
	},
	{
		Path: "/dummy/path/file71",
		Size: 71,
	},
	{
		Path: "/dummy/path/file72",
		Size: 72,
	},
	{
		Path: "/dummy/path/file73",
		Size: 73,
	},
	{
		Path: "/dummy/path/file74",
		Size: 74,
	},
	{
		Path: "/dummy/path/file75",
		Size: 75,
	},
	{
		Path: "/dummy/path/file76",
		Size: 76,
	},
	{
		Path: "/dummy/path/file77",
		Size: 77,
	},
	{
		Path: "/dummy/path/file78",
		Size: 78,
	},
	{
		Path: "/dummy/path/file79",
		Size: 79,
	},
	{
		Path: "/dummy/path/file80",
		Size: 80,
	},
	{
		Path: "/dummy/path/file81",
		Size: 81,
	},
	{
		Path: "/dummy/path/file82",
		Size: 82,
	},
	{
		Path: "/dummy/path/file83",
		Size: 83,
	},
	{
		Path: "/dummy/path/file84",
		Size: 84,
	},
	{
		Path: "/dummy/path/file85",
		Size: 85,
	},
	{
		Path: "/dummy/path/file86",
		Size: 86,
	},
	{
		Path: "/dummy/path/file87",
		Size: 87,
	},
	{
		Path: "/dummy/path/file88",
		Size: 88,
	},
	{
		Path: "/dummy/path/file89",
		Size: 89,
	},
	{
		Path: "/dummy/path/file90",
		Size: 90,
	},
	{
		Path: "/dummy/path/file91",
		Size: 91,
	},
	{
		Path: "/dummy/path/file92",
		Size: 92,
	},
	{
		Path: "/dummy/path/file93",
		Size: 93,
	},
	{
		Path: "/dummy/path/file94",
		Size: 94,
	},
	{
		Path: "/dummy/path/file95",
		Size: 95,
	},
	{
		Path: "/dummy/path/file96",
		Size: 96,
	},
	{
		Path: "/dummy/path/file97",
		Size: 97,
	},
	{
		Path: "/dummy/path/file98",
		Size: 98,
	},
	{
		Path: "/dummy/path/file99",
		Size: 99,
	},
	{
		Path: "/dummy/path/file100",
		Size: 100,
	},
	{
		Path: "/dummy/path/file101",
		Size: 101,
	},
}
