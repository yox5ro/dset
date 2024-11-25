package internal

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Intersect(w io.Writer, readers ...io.ReadSeeker) error {
	f, err := os.CreateTemp("", "intersect")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if err := Union(f, readers...); err != nil {
		return err
	}
	f.Seek(0, io.SeekStart)

	for _, reader := range readers {
		reader.Seek(0, io.SeekStart)
	}
	bufWriter := bufio.NewWriter(w)
	defer bufWriter.Flush()

	bufReaders := make([]*bufio.Reader, len(readers))
	for i, reader := range readers {
		bufReaders[i] = bufio.NewReader(reader)
	}
	currentStrings := make([]string, len(readers))

	tmpReader := bufio.NewReader(f)

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
