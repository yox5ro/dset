package internal

import (
	"bufio"
	"io"
	"strings"
)

func SubtractWrapper(w io.Writer, minuendFilename, subtrahendFilename string) error {
	minuendReader, err := WrapIsSorted(minuendFilename)
	if err != nil {
		return err
	}

	subtrahendReader, err := WrapIsSorted(subtrahendFilename)
	if err != nil {
		return err
	}

	return Subtract(w, minuendReader, subtrahendReader)
}

func Subtract(w io.Writer, minuend io.Reader, subtrahend io.Reader) error {
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
