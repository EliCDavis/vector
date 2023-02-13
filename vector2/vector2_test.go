package vector2_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/EliCDavis/vector/vector2"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	start := vector2.New(1.2, -2.4)

	tests := map[string]struct {
		want vector2.Float64
		got  vector2.Float64
	}{
		"x":             {want: start.SetX(4), got: vector2.New(4., -2.4)},
		"y":             {want: start.SetY(4), got: vector2.New(1.2, 4.)},
		"abs":           {want: start.Abs(), got: vector2.New(1.2, 2.4)},
		"floor":         {want: start.Floor(), got: vector2.New(1., -3.)},
		"ceil":          {want: start.Ceil(), got: vector2.New(2., -2.)},
		"round":         {want: start.Round(), got: vector2.New(1., -2.)},
		"sqrt":          {want: start.Sqrt(), got: vector2.New(1.0954451, math.NaN())},
		"clamp":         {want: start.Clamp(1, 2), got: vector2.New(1.2, 1.)},
		"perpendicular": {want: start.Perpendicular(), got: vector2.New(-2.4, -1.2)},
		"normalized":    {want: start.Normalized(), got: vector2.New(0.447213, -.894427)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		want  vector2.Float64
	}{
		"0, 0, 0 + 0, 0, 0 = 0, 0, 0": {left: vector2.New(0., 0.), right: vector2.New(0., 0.), want: vector2.New(0., 0.)},
		"1, 2, 3 + 4, 5, 6 = 5, 7, 9": {left: vector2.New(1., 2.), right: vector2.New(4., 5.), want: vector2.New(5., 7.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.left.Add(tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestSub(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		want  vector2.Float64
	}{
		"0, 0, 0 - 0, 0, 0 = 0, 0, 0": {left: vector2.New(0., 0.), right: vector2.New(0., 0.), want: vector2.New(0., 0.)},
		"4, 5, 6 - 1, 2, 3 = 3, 3, 3": {left: vector2.New(4., 5.), right: vector2.New(1., 2.), want: vector2.New(3., 3.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.left.Sub(tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestDefaults(t *testing.T) {
	tests := map[string]struct {
		got  vector2.Float64
		want vector2.Float64
	}{
		"zero":  {got: vector2.Zero[float64](), want: vector2.New(0., 0.)},
		"one":   {got: vector2.One[float64](), want: vector2.New(1., 1.)},
		"left":  {got: vector2.Left[float64](), want: vector2.New(-1., 0.)},
		"right": {got: vector2.Right[float64](), want: vector2.New(1., 0.)},
		"up":    {got: vector2.Up[float64](), want: vector2.New(0., 1.)},
		"down":  {got: vector2.Down[float64](), want: vector2.New(0., -1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
		})
	}
}

func TestMidpoint(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		want  vector2.Float64
	}{
		"0, 0 m 0, 0 = 0, 0":   {left: vector2.New(0., 0.), right: vector2.New(0., 0.), want: vector2.New(0., 0.)},
		"-1, -1 m 1, 1 = 0, 0": {left: vector2.New(-1., -1.), right: vector2.New(1., 1.), want: vector2.New(0., 0.)},
		"0, 0 m 1, 2 = 0.5, 1": {left: vector2.New(0., 0.), right: vector2.New(1., 2.), want: vector2.New(0.5, 1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.left.Midpoint(tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestLerp(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		t     float64
		want  vector2.Float64
	}{
		"(0, 0) =(0)=> (0, 0) = (0, 0)":     {left: vector2.New(0., 0.), right: vector2.New(0., 0.), t: 0, want: vector2.New(0., 0.)},
		"(0, 0) =(0.5)=> (1, 2) = (0.5, 1)": {left: vector2.New(0., 0.), right: vector2.New(1., 2.), t: 0.5, want: vector2.New(0.5, 1.)},
		"(0, 0) =(1)=> (1, 2) = (1, 2)":     {left: vector2.New(0., 0.), right: vector2.New(1., 2.), t: 1, want: vector2.New(1., 2.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector2.Lerp(tc.left, tc.right, tc.t)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestMin(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		want  vector2.Float64
	}{
		"(1, 2) m (3, 2) = (1, 2)": {left: vector2.New(1., 2.), right: vector2.New(3., 2.), want: vector2.New(1., 2.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector2.Min(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		want  vector2.Float64
	}{
		"(1, 2) m (3, 2) = (3, 2)": {left: vector2.New(1., 2.), right: vector2.New(3., 2.), want: vector2.New(3., 2.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector2.Max(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestNearZero(t *testing.T) {
	tests := map[string]struct {
		vec  vector2.Float64
		want bool
	}{
		"0, 0, 0":           {vec: vector2.New(0., 0.), want: true},
		"0, 0, 1":           {vec: vector2.New(0., 1.), want: false},
		"0, 0, .0000000001": {vec: vector2.New(0., 0.0000000001), want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.vec.NearZero())
		})
	}
}

func TestJSON(t *testing.T) {
	in := vector2.New(1.2, 2.3)
	out := vector2.New(0., 0.)

	marshalledData, marshallErr := json.Marshal(in)
	unmarshallErr := json.Unmarshal(marshalledData, &out)

	assert.NoError(t, marshallErr)
	assert.NoError(t, unmarshallErr)
	assert.Equal(t, "{\"x\":1.2,\"y\":2.3}", string(marshalledData))
	assert.Equal(t, 1.2, out.X())
	assert.Equal(t, 2.3, out.Y())
}

var result float64

func BenchmarkDistance(b *testing.B) {
	var r float64
	a := vector2.New(1., 2.)
	c := vector2.New(4., 5.)
	for i := 0; i < b.N; i++ {
		r = a.Distance(c)
	}
	result = r
}
