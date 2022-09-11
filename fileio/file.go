package fileio

import (
	"bufio"
	"errors"
	"os"
)

func Open(input string, blockSize int64) (*bufio.Reader, error) {

	// Read file
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	fileInfo, _ := file.Stat()

	// calc chunk size
	if fileInfo.Size()/blockSize < 2 {
		return nil, errors.New("chunk size less than 2")
	}

	return bufio.NewReader(file), nil
}
