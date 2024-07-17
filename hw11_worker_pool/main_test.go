package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestIncrementCounter(t *testing.T) {
	var mu sync.Mutex
	counter := 0

	incrementCounter(&mu, &counter)
	if counter != 1 {
		t.Errorf("Expected counter to be 1, but got %d", counter)
	}
}

func TestRunGoroutines(t *testing.T) {
	tests := []struct {
		numGoroutines int
		expected      int
	}{
		{1, 1},
		{5, 5},
		{10, 10},
		{0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("numGoroutines=%d", tt.numGoroutines), func(t *testing.T) {
			var counter int
			var wg sync.WaitGroup
			var mu sync.Mutex

			runGoroutines(tt.numGoroutines, &wg, &mu, &counter)

			if counter != tt.expected {
				t.Errorf("Expected counter to be %d, but got %d", tt.expected, counter)
			}
		})
	}
}
