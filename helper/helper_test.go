package helper

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintIntSlice(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected string
	}{
		{
			name:     "single element",
			numbers:  []int{1},
			target:   1,
			expected: "1 = 1\n",
		},
		{
			name:     "multiple elements",
			numbers:  []int{1, 2, 3, 4, 5},
			target:   15,
			expected: "1,2,3,4,5 = 15\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintIntSlice(tt.numbers, tt.target)

			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)
			os.Stdout = old

			// Verify the output using testify
			got := buf.String()
			assert.Equal(t, tt.expected, got, "they should be equal")
		})
	}
}
