package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBinarySearch проводит тестирование функции binarySearch с использованием табличных тестов.
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("arr=%v target=%d", tc.arr, tc.target), func(t *testing.T) {
			result := binarySearch(tc.arr, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
