package internal

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func IntersectWrapper(w io.Writer, filenames ...string) error {
	f, err := os.CreateTemp("", "intersect")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if err := UnionWrapper(f, filenames...); err != nil {
		return err
	}
	f.Seek(0, io.SeekStart)

	readers := make([]io.Reader, len(filenames))
	for i := range filenames {
		readers[i], err = OpenFile(filenames[i])
		if err != nil {
			return err
		}
	}
	return Intersect(w, f, readers...)
}

func Intersect(w io.Writer, tmpFileReader io.Reader, readers ...io.Reader) error {
	bufWriter := bufio.NewWriter(w)
	defer bufWriter.Flush()

	bufReaders := make([]*bufio.Reader, len(readers))
	for i, reader := range readers {
		bufReaders[i] = bufio.NewReader(reader)
	}
	currentStrings := make([]string, len(readers))

	tmpReader := bufio.NewReader(tmpFileReader)

	for {
		s, err := tmpReader.ReadString('\n')
		if err == io.EOF && s == "" {
			break
		} else if err != nil && err != io.EOF {
			return err
		}
		currentUnionString := strings.TrimSpace(s)

		found := true
		for i, reader := range bufReaders {
			for currentStrings[i] < currentUnionString {
				s, err = reader.ReadString('\n')
				if err == io.EOF && s == "" {
					found = false
					break
				} else if err != nil && err != io.EOF {
					return err
				}
				currentStrings[i] = strings.TrimSpace(s)
			}
			if currentStrings[i] != currentUnionString {
				found = false
				break
			}
		}

		if found {
			if _, err := bufWriter.WriteString(currentUnionString + "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}
