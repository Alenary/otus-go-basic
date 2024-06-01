package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name string
		size int
		want string
	}{
		{
			name: "Valid size 1",
			size: 2,
			want: "# \n #\n",
		},
		{
			name: "Valid size 3",
			size: 3,
			want: "# #\n # \n# #\n",
		},
		{
			name: "Invalid size 1",
			size: 1,
			want: "Указаны некорректные данные",
		},
		{
			name: "Valid size 4",
			size: 4,
			want: "# # \n # #\n# # \n # #\n",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Board(tc.size)
			assert.Equal(t, tc.want, got)
		})
	}
}
