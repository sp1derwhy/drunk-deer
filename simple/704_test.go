package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test704(t *testing.T) {
	tests := []struct {
		testName  string
		inputNums []int
		target    int
		want      int
	}{
		{"case1", []int{5}, 5, 0},
		{"case2", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"case3", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			got := search(tt.inputNums, tt.target)
			assert.Equal(t, got, tt.want)
		})
	}
}
