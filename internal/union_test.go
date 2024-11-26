package internal_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/yox5ro/dset/internal"
)

func TestUnion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []strings.Reader
		want  string
	}{
		{
			name: "single file",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
			},
			want: "1\n2\n3\n",
		},
		{
			name: "two files",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
				*strings.NewReader("2\n3\n4\n"),
			},
			want: "1\n2\n3\n4\n",
		},
		{
			name: "three files",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
				*strings.NewReader("2\n3\n4\n"),
				*strings.NewReader("3\n4\n5\n"),
			},
			want: "1\n2\n3\n4\n5\n",
		},
		{
			name: "no overlap",
			input: []strings.Reader{
				*strings.NewReader("1\n2\n3\n"),
				*strings.NewReader("4\n5\n6\n"),
				*strings.NewReader("7\n8\n9\n"),
			},
			want: "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
		},
		{
			name: "empty files",
			input: []strings.Reader{
				*strings.NewReader(""),
				*strings.NewReader(""),
				*strings.NewReader(""),
			},
			want: "",
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
			if err := internal.Union(&got, readers...); err != nil {
				t.Errorf("Union() error = %v", err)
				return
			}
			if !bytes.Equal([]byte(got.String()), []byte(tt.want)) {
				t.Errorf("Union() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
