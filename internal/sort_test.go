package internal_test

import (
	"strings"
	"testing"

	"github.com/yox5ro/dset/internal"
)

func TestIsSorted(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input *strings.Reader
		want  bool
	}{
		{
			name:  "sorted",
			input: strings.NewReader("1\n2\n3\n"),
			want:  true,
		},
		{
			name:  "unsorted",
			input: strings.NewReader("1\n3\n2\n"),
			want:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.IsSorted(tt.input); got != tt.want {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
