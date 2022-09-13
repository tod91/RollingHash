package hash

import (
	"crypto/md5"
	"fmt"
)

func Split(buf []byte) ([]string, error) {
	fSize := len(buf)
	bSize := getBlockSize(fSize)

	var ret []string

	for idx := 0; idx*bSize < fSize; idx++ {
		ret = append(ret, calcMD5(buf, idx*bSize, (idx+1)*bSize))
	}
	return ret, nil
}

func calcMD5(buf []byte, from, to int) string {
	h := md5.Sum(buf[from:to])
	hb := []byte(h[:])
	asStr := string(hb)
	hexStr := fmt.Sprintf("%x", asStr)

	return hexStr
}

func getBlockSize(size int) int {
	if size > 1024 {
		return 1024
	} else {
		return 512
	}
}
