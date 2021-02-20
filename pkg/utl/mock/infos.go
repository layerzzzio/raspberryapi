package mock

// Infos mock
type Infos struct {
	ReadFileFn     func(path string) ([]string, error)
	IsFileExistsFn func(path string) bool
}

// ReadFile mock
func (i Infos) ReadFile(path string) ([]string, error) {
	return i.ReadFileFn(path)
}

// IsFileExists mock
func (i Infos) IsFileExists(path string) bool {
	return i.IsFileExistsFn(path)
}
