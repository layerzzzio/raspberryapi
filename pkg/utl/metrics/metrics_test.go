package metrics_test

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/assert"
)

func TestProcesses(t *testing.T) {
	cases := []struct {
		name       string
		id         int32
		mps        mock.Metrics
		wantedData metrics.PInfo
		wantedErr  error
	}{
		{
			name:       "error: process id 0",
			id:         0,
			mps:        mock.Metrics{},
			wantedErr:  errors.New("process not found"),
			wantedData: metrics.PInfo{},
		},
		{
			name: "error: process not found",
			id:   66,
			mps: mock.Metrics{
				PsPIDFn: func(p *process.Process, c chan (int32)) {
					c <- 99
				},
				PsNameFn: func(p *process.Process, c chan (string)) {
					c <- "process_99"
				},
				PsCPUPerFn: func(p *process.Process, c chan (float64)) {
					c <- 1.1
				},
				PsMemPerFn: func(p *process.Process, c chan (float32)) {
					c <- 2.2
				},
			},
			wantedErr:  errors.New("process not found"),
			wantedData: metrics.PInfo{},
		},
		{
			name: "success: without process id",
			// id -1 simulate a non-existent process id
			id: -1,
			mps: mock.Metrics{
				PsPIDFn: func(p *process.Process, c chan (int32)) {
					c <- 99
				},
				PsNameFn: func(p *process.Process, c chan (string)) {
					c <- "process_99"
				},
				PsCPUPerFn: func(p *process.Process, c chan (float64)) {
					c <- 1.1
				},
				PsMemPerFn: func(p *process.Process, c chan (float32)) {
					c <- 2.2
				},
			},
			wantedData: metrics.PInfo{
				ID:         99,
				Name:       "process_99",
				CPUPercent: 1.1,
				MemPercent: 2.2,
			},
			wantedErr: nil,
		},
		{
			name: "success: with process id",
			// make sure to keep id 1 to make this test work
			id: 1,
			mps: mock.Metrics{
				PsPIDFn: func(p *process.Process, c chan (int32)) {
					c <- 1
				},
				PsNameFn: func(p *process.Process, c chan (string)) {
					c <- "process_1"
				},
				PsCPUPerFn: func(p *process.Process, c chan (float64)) {
					c <- 1.1
				},
				PsMemPerFn: func(p *process.Process, c chan (float32)) {
					c <- 2.2
				},
				PsUsernameFn: func(p *process.Process, c chan (string)) {
					c <- "pi"
				},
				PsCmdLineFn: func(p *process.Process, c chan (string)) {
					c <- "/cmd/test"
				},
				PsStatusFn: func(p *process.Process, c chan (string)) {
					c <- "S"
				},
				PsCreationTimeFn: func(p *process.Process, c chan (int64)) {
					c <- 1888
				},
				PsForegroundFn: func(p *process.Process, c chan (bool)) {
					c <- true
				},
				PsBackgroundFn: func(p *process.Process, c chan (bool)) {
					c <- false
				},
				PsIsRunningFn: func(p *process.Process, c chan (bool)) {
					c <- true
				},
				PsParentFn: func(p *process.Process, c chan (int32)) {
					c <- -1
				},
			},
			wantedData: metrics.PInfo{
				ID:           1,
				Name:         "process_1",
				CPUPercent:   1.1,
				MemPercent:   2.2,
				Username:     "pi",
				CommandLine:  "/cmd/test",
				Status:       "S",
				CreationTime: 1888,
				Foreground:   true,
				Background:   false,
				IsRunning:    true,
				ParentP:      -1,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := metrics.New(tc.mps)

			var ps []metrics.PInfo
			var err error

			if tc.id >= 0 {
				ps, err = s.Processes(tc.id)
			} else {
				ps, err = s.Processes()
			}

			if (tc.wantedData != metrics.PInfo{}) {
				assert.Equal(t, tc.wantedData, ps[0])
			}
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

type fakeFile struct {
	fileName  string
	fileSize  int64
	fakeFiles []fakeFile
}

func (f fakeFile) Name() string       { return f.fileName }
func (f fakeFile) Size() int64        { return f.fileSize }
func (f fakeFile) Mode() os.FileMode  { return 0 }
func (f fakeFile) ModTime() time.Time { return time.Now() }
func (f fakeFile) IsDir() bool        { return len(f.fakeFiles) > 0 }
func (f fakeFile) Sys() interface{}   { return nil }

func createReadDir(ff fakeFile) metrics.ReadDir {
	return func(path string) ([]os.FileInfo, error) {
		names := strings.Split(path, "/")
		fakeFolder := ff
		var found bool
		for _, name := range names {
			found = false
			for _, testFile := range fakeFolder.fakeFiles {
				if testFile.fileName == name {
					fakeFolder = testFile
					found = true
					break
				}
			}
			if !found {
				return []os.FileInfo{}, fmt.Errorf("file not found")
			}

		}
		result := make([]os.FileInfo, len(fakeFolder.fakeFiles))
		for i, resultFile := range fakeFolder.fakeFiles {
			result[i] = resultFile
		}
		return result, nil
	}
}

func TestFilePath(t *testing.T) {
	root := test_utl.NewTestFolder("root",
		test_utl.NewTestFolder("folder1",
			test_utl.NewTestFile("file1", 0),
		),
	)
	want := "root/folder1/file1"
	file1 := test_utl.FindTestFile(root, "file1")
	assert.Equal(t, want, metrics.Path(file1))
}

func TestWalkFolderOnSimpleDir(t *testing.T) {
	testStructure := fakeFile{"a", 0, []fakeFile{
		{"b", 0, []fakeFile{
			{"c", 100, []fakeFile{}},
			{"d", 0, []fakeFile{
				{"e", 50, []fakeFile{}},
				{"f", 30, []fakeFile{}},
				{"g", 70, []fakeFile{ //thisfolder should get ignored
					{"h", 10, []fakeFile{}},
					{"i", 20, []fakeFile{}},
				}},
			}},
		}},
	}}
	dummyIgnoreFunction := func(path string) bool { return path == "b/d/g" }
	progress := make(chan int, 3)
	s := metrics.New(metrics.Service{})
	result, _ := s.WalkFolder("b", createReadDir(testStructure), 180, 0, dummyIgnoreFunction, progress)
	buildExpected := func() *rpi.File {
		b := &rpi.File{
			Name:   "b",
			Parent: nil,
			Size:   180,
			IsDir:  true,
			Files:  []*rpi.File{},
		}
		c := &rpi.File{
			Name:   "c",
			Parent: b,
			Size:   100,
			IsDir:  false,
			Files:  []*rpi.File{},
		}
		d := &rpi.File{
			Name:   "d",
			Parent: b,
			Size:   80,
			IsDir:  true,
			Files:  []*rpi.File{},
		}
		b.Files = []*rpi.File{c, d}

		e := &rpi.File{
			Name:   "e",
			Parent: nil,
			Size:   50,
			IsDir:  false,
			Files:  []*rpi.File{},
		}
		e.Parent = d
		f := &rpi.File{
			Name:   "f",
			Parent: nil,
			Size:   30,
			IsDir:  false,
			Files:  []*rpi.File{},
		}
		g := &rpi.File{
			Name:   "g",
			Parent: nil,
			Size:   0,
			IsDir:  true,
			Files:  []*rpi.File{},
		}
		f.Parent = d
		g.Parent = d
		d.Files = []*rpi.File{e, f, g}

		return b
	}
	expected := buildExpected()
	assert.Equal(t, expected, result)
	resultProgress := 0
	resultProgress += <-progress
	resultProgress += <-progress
	_, more := <-progress
	assert.Equal(t, 2, resultProgress)
	assert.False(t, more, "the progress channel should be closed")
}

func TestWalkFolderNotInLimit(t *testing.T) {
	testStructure := fakeFile{"a", 0, []fakeFile{
		{"b", 0, []fakeFile{
			{"c", 90, []fakeFile{}},
			{"d", 0, []fakeFile{
				{"e", 50, []fakeFile{}},
				{"f", 30, []fakeFile{}},
				{"g", 70, []fakeFile{ //thisfolder should get ignored
					{"h", 10, []fakeFile{}},
					{"i", 20, []fakeFile{}},
				}},
			}},
		}},
	}}
	dummyIgnoreFunction := func(path string) bool { return path == "b/d/g" }
	progress := make(chan int, 3)
	s := metrics.New(metrics.Service{})
	result, _ := s.WalkFolder("b", createReadDir(testStructure), 180, 70, dummyIgnoreFunction, progress)
	expected := &rpi.File{
		Name:  "b",
		Size:  0,
		IsDir: true,
	}

	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Size, result.Size)
	assert.Equal(t, expected.IsDir, result.IsDir)

	resultProgress := 0
	resultProgress += <-progress
	resultProgress += <-progress
	_, more := <-progress
	assert.Equal(t, 2, resultProgress)
	assert.False(t, more, "the progress channel should be closed")
}

func TestWalkFolderHandlesError(t *testing.T) {
	failing := func(path string) ([]os.FileInfo, error) {
		return []os.FileInfo{}, errors.New("Not found")
	}
	progress := make(chan int, 2)
	s := metrics.New(metrics.Service{})
	result, _ := s.WalkFolder("xyz", failing, 0, 1, func(string) bool { return false }, progress)
	assert.Equal(t, rpi.File{}, *result, "WalkFolder didn't return empty file on ReadDir failure")
}

// TODO: test function metrics.DirSize
func TestDirSize(t *testing.T) {
	// To be analyzed with Docker
	// NewTestFolder does now work here
}
