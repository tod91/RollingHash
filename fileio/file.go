package fileio

import (
	"os"
)

// Reads a file and returns it's byte buffer
func Open(input string) ([]byte, error) {
	// internally handles file closing, so no need to do it here
	file, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}

	return file, nil
}
