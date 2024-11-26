package internal_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/yox5ro/dset/internal"
)

func TestIntersect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []strings.Reader
		tmpUnion strings.Reader
		want     string
	}{
		{
			name: "single file",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
			},
			tmpUnion: *strings.NewReader("1\n2\n3\n"),
			want:     "1\n2\n3\n",
		},
		{
			name: "two files",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
				*strings.NewReader("2\n3\n4\n"),
			},
			tmpUnion: *strings.NewReader("1\n2\n3\n4\n"),
			want:     "2\n3\n",
		},
		{
			name: "three files",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
				*strings.NewReader("2\n3\n4\n"),
				*strings.NewReader("3\n4\n5\n"),
			},
			tmpUnion: *strings.NewReader("1\n2\n3\n4\n5\n"),
			want:     "3\n",
		},
		{
			name: "no overlap",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
				*strings.NewReader("4\n5\n6\n"),
				*strings.NewReader("7\n8\n9\n"),
			},
			tmpUnion: *strings.NewReader("1\n2\n3\n4\n5\n6\n7\n8\n9\n"),
			want:     "",
		},
		{
			name: "empty files",
			input: []strings.Reader{
				*strings.NewReader(""),
				*strings.NewReader(""),
				*strings.NewReader(""),
			},
			tmpUnion: *strings.NewReader(""),
			want:     "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Builder{}
			readers := make([]io.Reader, len(tt.input))
			for i, input := range tt.input {
				readers[i] = &input
			}
			tmpUnion := &tt.tmpUnion
			if err := internal.Intersect(&got, tmpUnion, readers...); err != nil {
				t.Errorf("Intersect() error = %v", err)
				return
			}
			if !bytes.Equal([]byte(got.String()), []byte(tt.want)) {
				t.Errorf("Intersect() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
