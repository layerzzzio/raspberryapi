package mock

// Infos mock
type Infos struct {
	ReadFileFn func(path string) ([]string, error)
}

// HumanUser mock
func (i Infos) ReadFile(path string) ([]string, error) {
	return i.ReadFileFn(path)
}
