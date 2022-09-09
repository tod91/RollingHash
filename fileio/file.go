package fileio

import (
	"bufio"
	"os"
)

func Open(input string) (*bufio.Reader, error) {
	// Open file to split
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil

}
