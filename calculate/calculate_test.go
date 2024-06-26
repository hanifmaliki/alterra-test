package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCombinations(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		maxSum   int
		expected int
	}{
		{"Test case 1: target 4, maxSum 5", 4, 5, 4},
		// 1,1,1,1
		// 1,1,2
		// 1,3
		// 4
		{"Test case 2: target 5, maxSum 3", 5, 3, 4},
		// 1,1,1,1,1
		// 1,1,1,2
		// 1,1,3
		// 2,3
		{"Test case 3: target 3, maxSum 3", 3, 3, 3},
		// 1,1,1
		// 1,2
		// 3
		{"Test case 4: target 2, maxSum 3", 2, 3, 2},
		// 1,1
		// 2
		{"Test case 5: target 6, maxSum 4", 6, 4, 5},
		// 1,1,1,1,1,1
		// 1,1,1,1,2
		// 1,1,1,3
		// 1,1,4
		// 2,4
		{"Test case 5: target 8, maxSum 2", 8, 2, 5},
		// 1,1,1,1,1,1,1,1
		// 1,1,1,1,1,1,2
		// 1,1,1,1,2,2
		// 1,1,2,2,2
		// 2,2,2,2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalculateCombinations(tt.target, tt.maxSum)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
