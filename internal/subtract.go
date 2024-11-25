package internal

import (
	"bufio"
	"io"
	"strings"
)

func Subtract(w io.Writer, minuend io.ReadSeeker, subtrahend io.ReadSeeker) error {
	if !IsSorted(minuend) || !IsSorted(subtrahend) {
		return ErrInputNotSorted
	}

	minuend.Seek(0, io.SeekStart)
	subtrahend.Seek(0, io.SeekStart)
	bufWriter := bufio.NewWriter(w)
	defer bufWriter.Flush()

	bufMinuend := bufio.NewReader(minuend)
	bufSubtrahend := bufio.NewReader(subtrahend)
	currentMinuendString := ""
	currentSubtrahendString := ""

	for {
		s, err := bufMinuend.ReadString('\n')
		if err == io.EOF && s == "" {
			break
		} else if err != nil && err != io.EOF {
			return err
		}
		currentMinuendString = strings.TrimSpace(s)

		for currentMinuendString > currentSubtrahendString {
			s, err = bufSubtrahend.ReadString('\n')
			if err == io.EOF && s == "" {
				break
			} else if err != nil && err != io.EOF {
				return err
			}
			currentSubtrahendString = strings.TrimSpace(s)
		}
		
		if currentMinuendString != currentSubtrahendString {
			if _, err := bufWriter.WriteString(currentMinuendString + "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}
