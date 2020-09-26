package sys

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"

	"github.com/stretchr/testify/assert"
)

func TestFileCategory(t *testing.T) {
	cases := []struct {
		name       string
		file       string
		wantedData struct{ category, description string }
	}{
		{
			name: "file containing no category",
			file: "/dummy/path",
			wantedData: struct{ category, description string }{
				category:    "",
				description: "",
			},
		},
		{
			name: "file containing no sub category",
			file: "/bin/badaboum/py.deb",
			wantedData: struct{ category, description string }{
				category:    "/bin/",
				description: "represents some essential user command binaries",
			},
		},
		{
			name: "file containing a sub category",
			file: "/var/cache/apt/srcpkg.bin",
			wantedData: struct{ category, description string }{
				category:    "/var/cache/",
				description: "contains application cache data",
			},
		},
		{
			name: "file containing two existing category substrings",
			file: "/usr/bin/apt/srcpkg.bin",
			wantedData: struct{ category, description string }{
				category:    "/usr/bin/",
				description: "contains most of the executable files that are not needed for booting or repairing the system",
			},
		},
		{
			name: "file containing two existing category substrings at the beginning and the end",
			file: "/usr/one/two/three/bin/srcpkg.bin",
			wantedData: struct{ category, description string }{
				category:    "/usr/",
				description: "contains shareable and read-only data",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			category := FileCategory(tc.file)
			assert.Equal(t, tc.wantedData, category)
		})
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		name        string
		top100files []metrics.PathSize
		wantedData  []rpi.LargestFile
		wantedErr   error
	}{
		{
			name: "success",
			top100files: []metrics.PathSize{
				{
					Path: "/bin/file1",
					Size: 11,
				},
				{
					Path: "/usr/include/file2",
					Size: 22,
				},
				{
					Path: "/usr/dummy/file3",
					Size: 33,
				},
			},
			wantedData: []rpi.LargestFile{
				{
					Path:                "/bin/file1",
					Name:                "file1",
					Size:                11,
					Category:            "/bin",
					CategoryDescription: "represents some essential user command binaries",
				},
				{
					Path:                "/usr/include/file2",
					Name:                "file2",
					Size:                22,
					Category:            "/usr/include",
					CategoryDescription: "contains system general-use include files for the C programming language",
				},
				{
					Path:                "/usr/dummy/file3",
					Name:                "file3",
					Size:                33,
					Category:            "/usr",
					CategoryDescription: "contains shareable and read-only data",
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := largestfile.LFSYS(LargestFile{})
			users, err := s.List(tc.top100files)
			assert.Equal(t, tc.wantedData, users)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
