package task5

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name      string
		input     io.Reader
		sequence  []byte
		expect    bool
		expectErr error
	}{
		{
			name:     "Sequence found",
			input:    bytes.NewBufferString("abcdefg"),
			sequence: []byte("cde"),
			expect:   true,
		},
		{
			name:     "Sequence not found",
			input:    bytes.NewBufferString("abcdefg"),
			sequence: []byte("xyz"),
			expect:   false,
		},
		{
			name:      "Error reading",
			input:     failingReader{},
			sequence:  []byte("abc"),
			expectErr: errors.New("error reading"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Contains(test.input, test.sequence)
			if err != nil {
				if test.expectErr == nil || err.Error() != test.expectErr.Error() {
					t.Fatalf("unexpected error: got %v, want %v", err, test.expectErr)
				}
				return
			}
			if got != test.expect {
				t.Errorf("unexpected result: got %v, want %v", got, test.expect)
			}
		})
	}
}

type failingReader struct{}

func (failingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error reading")
}
