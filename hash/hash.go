package hash

import (
	"crypto/md5"
	"fmt"
	///"strings"
)

func Split(buf []byte) ([]string, error) {
	// Get file and block size
	fSize := len(buf)
	bSize := getBlockSize(fSize)

	var ret []string

	// iterate until our block size exceeds the file length
	idx := 0
	for ; (idx+1)*bSize < fSize; idx++ {
		ret = append(ret, calcMD5(buf, idx*bSize, (idx+1)*bSize))
	}

	// grab what's left from the file and fill the rest of the block with zeroes
	from := idx * bSize
	to := fSize - from
	ret = append(ret, calcMD5(buf, from, from+to))

	return ret, nil
}

func SlideWindow(curHash uint64, prevc, nextc rune) uint64 {
	curHash -= uint64(prevc)
	curHash += uint64(nextc)

	return curHash
}

func RabinKarp(block string) uint64 {
	var ret uint64
	b := []rune(block)

	for _, c := range b {
		ret += uint64(c)
	}

	return ret
}

func calcMD5(buf []byte, from, to int) string {
	h := md5.Sum(buf[from:to])
	asStr := string(h[:])
	hexStr := fmt.Sprintf("%x", asStr)

	return hexStr
}

func getBlockSize(size int) int {
	if size == 0 {
		fmt.Errorf("cannot get block size of an empty file")
	} else if size > 1024 {
		return 1024
	} else {
		return 512
	}
	return 512
}
