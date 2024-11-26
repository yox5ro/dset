package internal

import (
	"bufio"
	"io"
	"strings"
)

func UnionWrapper(w io.Writer, filenames ...string) error {
	readers := make([]io.Reader, len(filenames))
	for i, filename := range filenames {
		reader, err := WrapIsSorted(filename)
		if err != nil {
			return err
		}
		readers[i] = reader
	}
	return Union(w, readers...)
}

func Union(w io.Writer, readers ...io.Reader) error {
	bufWriter := bufio.NewWriter(w)
	defer bufWriter.Flush()

	bufReaders := make([]*bufio.Reader, len(readers))
	for i, reader := range readers {
		bufReaders[i] = bufio.NewReader(reader)
	}

	currentStrings := make([]string, len(readers))
	lastWriteString := ""

	for {
		for i, currentString := range currentStrings {
			if currentString == lastWriteString {
				s, err := bufReaders[i].ReadString('\n')
				if err == io.EOF {
					currentStrings[i] = ""
				} else if err != nil {
					return err
				}
				currentStrings[i] = strings.TrimSpace(s)
			}
		}

		minString := ""
		minIndex := -1
		for i, currentString := range currentStrings {
			if currentString == "" {
				continue
			}
			if minIndex == -1 || currentString < minString {
				minString = currentString
				minIndex = i
			}
		}
		if minIndex == -1 {
			break
		}

		if _, err := bufWriter.WriteString(minString + "\n"); err != nil {
			return err
		}
		lastWriteString = minString
	}
	return nil
}
