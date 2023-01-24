package vector3_test

import (
	"testing"

	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		left  vector3.Float64
		right vector3.Float64
		want  vector3.Float64
	}{
		"0, 0, 0 + 0, 0, 0 = 0, 0, 0": {left: vector3.New(0., 0., 0.), right: vector3.New(0., 0., 0.), want: vector3.New(0., 0., 0.)},
		"1, 2, 3 + 4, 5, 6 = 5, 7, 9": {left: vector3.New(1., 2., 3.), right: vector3.New(4., 5., 6.), want: vector3.New(5., 7., 9.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.left.Add(tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestDefaults(t *testing.T) {
	tests := map[string]struct {
		got  vector3.Float64
		want vector3.Float64
	}{
		"zero":    {got: vector3.Zero[float64](), want: vector3.New(0., 0., 0.)},
		"one":     {got: vector3.One[float64](), want: vector3.New(1., 1., 1.)},
		"left":    {got: vector3.Left[float64](), want: vector3.New(-1., 0., 0.)},
		"right":   {got: vector3.Right[float64](), want: vector3.New(1., 0., 0.)},
		"up":      {got: vector3.Up[float64](), want: vector3.New(0., 1., 0.)},
		"down":    {got: vector3.Down[float64](), want: vector3.New(0., -1., 0.)},
		"forward": {got: vector3.Forward[float64](), want: vector3.New(0., 0., 1.)},
		"back":    {got: vector3.Backwards[float64](), want: vector3.New(0., 0., -1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
		})
	}
}
