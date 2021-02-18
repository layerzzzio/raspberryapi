package infos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
