package internal_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/yox5ro/dset/internal"
)

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		inputMinuend    strings.Reader
		inputSubtrahend strings.Reader
		want            string
	}{
		{
			name: "no overlap",
			inputMinuend: *strings.NewReader("1\n2\n3\n"),
			inputSubtrahend: *strings.NewReader("4\n5\n6\n"),
			want: "1\n2\n3\n",
		},
		{
			name: "overlap",
			inputMinuend: *strings.NewReader("1\n2\n3\n"),
			inputSubtrahend: *strings.NewReader("2\n3\n4\n"),
			want: "1\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Builder{}
			var minuendReader io.ReadSeeker = &tt.inputMinuend
			var subtrahendReader io.ReadSeeker = &tt.inputSubtrahend
			if err := internal.Subtract(&got, minuendReader, subtrahendReader); err != nil {
				t.Errorf("Subtract() error = %v", err)
				return
			}
			if !bytes.Equal([]byte(got.String()), []byte(tt.want)) {
				t.Errorf("Subtract() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
