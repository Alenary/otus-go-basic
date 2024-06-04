package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "case 1",
			input: "truffle krevetka felicia felicia",
			expected: map[string]int{
				"truffle":  1,
				"krevetka": 1,
				"felicia":  2,
			},
		},
		{
			name:     "case 2 empty",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:  "case 3 single word",
			input: "cat",
			expected: map[string]int{
				"cat": 1,
			},
		},
		{
			name:  "case 4 insensitive",
			input: "Cat cat CAT",
			expected: map[string]int{
				"cat": 3,
			},
		},
		{
			name:  "case 5 punctuation",
			input: "Cats...cat!!!",
			expected: map[string]int{
				"cats": 1,
				"cat":  1,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := countWords(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
