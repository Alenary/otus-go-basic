package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangle_Area(t *testing.T) {
	tests := []struct {
		name   string
		base   float64
		height float64
		want   float64
	}{
		{
			name:   "Case default",
			base:   8.0,
			height: 6.0,
			want:   24.0,
		},
		{
			name:   "Case 2",
			base:   10.0,
			height: 5.0,
			want:   25.0,
		},
		{
			name:   "Case 3",
			base:   7.0,
			height: 3.0,
			want:   10.5,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			triangle := Triangle{Base: tc.base, Height: tc.height}
			got := triangle.Area()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{
			name:   "Case default",
			width:  10.0,
			height: 5.0,
			want:   50.0,
		},
		{
			name:   "Case 2",
			width:  10.0,
			height: 6.0,
			want:   60.0,
		},
		{
			name:   "Case 3",
			width:  7.0,
			height: 3.0,
			want:   21.0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			rectangle := Rectangle{Width: tc.width, Height: tc.height}
			got := rectangle.Area()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "Case default",
			radius: 5.0,
			want:   78.53981633974483,
		},
		{
			name:   "Case 2",
			radius: 1.0,
			want:   3.141592653589793,
		},
		{
			name:   "Case 3",
			radius: 0.5,
			want:   0.7853981633974483,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			circle := Circle{Radius: tc.radius}
			got := circle.Area()
			assert.Equal(t, tc.want, got)
		})
	}
}
