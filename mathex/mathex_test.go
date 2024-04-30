package mathex_test

import (
	"testing"

	"github.com/EliCDavis/vector/mathex"
	"github.com/stretchr/testify/assert"
)

func TestLerp(t *testing.T) {
	tests := map[string]struct {
		left  float64
		right float64
		t     float64
		want  float64
	}{
		"0 = 0 => 1 = 0":     {left: 0., right: 1., t: 0, want: 0.},
		"0 = 0.5 => 1 = 0.5": {left: 0., right: 1., t: 0.5, want: 0.5},
		"0 = 1 => 1 = 1":     {left: 0., right: 1., t: 1, want: 1.},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Lerp(tc.t, tc.left, tc.right)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestLerpInt(t *testing.T) {
	tests := map[string]struct {
		left  int
		right int
		t     float64
		want  int
	}{
		"1 = 0 => 3 = 1":   {left: 1, right: 3, t: 0, want: 1},
		"1 = 0.5 => 3 = 2": {left: 1, right: 3, t: 0.5, want: 2},
		"1 = 1 => 3 = 3":   {left: 1, right: 3, t: 1, want: 3},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Lerp(tc.t, tc.left, tc.right)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := map[string]struct {
		left  float64
		right float64
		t     float64
		want  float64
	}{
		"1 = 1 => 2 = 0":     {left: 1., right: 2., t: 1, want: 0.},
		"1 = 1.5 => 2 = 0.5": {left: 1., right: 2., t: 1.5, want: 0.5},
		"1 = 2 => 2 = 1":     {left: 1., right: 2., t: 2, want: 1.},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Normalize(tc.t, tc.left, tc.right)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestNormalizeInt(t *testing.T) {
	tests := map[string]struct {
		left  int
		right int
		t     int
		want  float64
	}{
		"1 = 1 => 5 = 0":   {left: 1, right: 5, t: 1, want: 0.},
		"1 = 3 => 5 = 0.5": {left: 1, right: 5, t: 3, want: 0.5},
		"1 = 5 => 5 = 1":   {left: 1, right: 5, t: 5, want: 1.},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Normalize(tc.t, tc.left, tc.right)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestRemap(t *testing.T) {
	tests := map[string]struct {
		left   float64
		right  float64
		left2  float64
		right2 float64
		t      float64
		want   float64
	}{
		"1:4 = 1 => 2:6 = 4":   {left: 1., right: 2., left2: 4., right2: 6., t: 1, want: 4.},
		"1:4 = 1.5 => 2:6 = 5": {left: 1., right: 2., left2: 4., right2: 6., t: 1.5, want: 5.},
		"1:4 = 2 => 2:6 = 6":   {left: 1., right: 2., left2: 4., right2: 6., t: 2, want: 6.},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Remap(tc.t, tc.left, tc.right, tc.left2, tc.right2)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestRemapInt(t *testing.T) {
	tests := map[string]struct {
		left   int
		right  int
		left2  int
		right2 int
		t      int
		want   int
	}{
		"1:4 = 1 => 3:6 = 4": {left: 1, right: 3, left2: 4, right2: 6, t: 1, want: 4},
		"1:4 = 2 => 3:6 = 5": {left: 1, right: 3, left2: 4, right2: 6, t: 2, want: 5},
		"1:4 = 3 => 3:6 = 6": {left: 1, right: 3, left2: 4, right2: 6, t: 3, want: 6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Remap(tc.t, tc.left, tc.right, tc.left2, tc.right2)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestWrap(t *testing.T) {
	tests := map[string]struct {
		left  float64
		right float64
		t     float64
		want  float64
	}{
		"1 = 0.3 => 2 = 1.3": {left: 1., right: 2., t: 0.3, want: 1.3},
		"1 = 1.5 => 2 = 1.5": {left: 1., right: 2., t: 1.5, want: 1.5},
		"1 = 3.6 => 2 = 1.6": {left: 1., right: 2., t: 3.6, want: 1.6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathex.Wrap(tc.t, tc.left, tc.right)

			assert.InDelta(t, tc.want, got, 0.000001)
		})
	}
}
