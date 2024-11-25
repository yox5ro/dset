package internal

import (
	"bufio"
	"io"
)

// read lines from r and check if they are sorted
func IsSorted(r io.Reader) bool {
	bfReader := bufio.NewReader(r)
	prev := ""
	for {
		line, err := bfReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return false
		}
		if prev > line {
			return false
		}
		prev = line
	}
	return true
}
