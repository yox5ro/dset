package internal

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
)

func isGzipFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	magicNumber := make([]byte, 2)
	if _, err = file.Read(magicNumber); err != nil {
		return false, err
	}

	// 0x1F 0x8B is the magic number for gzip files
	if bytes.Equal(magicNumber, []byte{0x1F, 0x8B}) {
		return true, nil
	}
	return false, nil
}

func OpenFile(filename string) (io.ReadCloser, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	isGzip, err := isGzipFile(filename)
	if err != nil {
		return nil, err
	}
	if isGzip {
		return gzip.NewReader(f)
	}
	return f, nil
}

func WrapIsSorted(filename string) (io.Reader, error) {
	f, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}

	if !IsSorted(f) {
		return nil, ErrInputNotSorted
	}
	f.Close()

	f, err = OpenFile(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}
