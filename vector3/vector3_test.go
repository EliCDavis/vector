package vector3_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func TestVectorOperations(t *testing.T) {
	start := vector3.New(1.2, -2.4, 3.7)

	tests := map[string]struct {
		want vector3.Float64
		got  vector3.Float64
	}{
		"x":            {want: start.SetX(4), got: vector3.New(4., -2.4, 3.7)},
		"y":            {want: start.SetY(4), got: vector3.New(1.2, 4., 3.7)},
		"z":            {want: start.SetZ(4), got: vector3.New(1.2, -2.4, 4.)},
		"abs":          {want: start.Abs(), got: vector3.New(1.2, 2.4, 3.7)},
		"floor":        {want: start.Floor(), got: vector3.New(1., -3., 3.)},
		"ceil":         {want: start.Ceil(), got: vector3.New(2., -2., 4.)},
		"round":        {want: start.Round(), got: vector3.New(1., -2., 4.)},
		"multByVector": {want: start.MultByVector(vector3.New(2., 4., 6.)), got: vector3.New(2.4, -9.6, 22.2)},
		"sqrt":         {want: start.Sqrt(), got: vector3.New(1.0954451, math.NaN(), 1.923538)},
		"clamp":        {want: start.Clamp(1, 2), got: vector3.New(1.2, 1., 2.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
		})
	}
}

func TestDistances(t *testing.T) {
	tests := map[string]struct {
		a    vector3.Float64
		b    vector3.Float64
		want float64
	}{
		"(0, 0, 0), (0, 0, 0)":  {a: vector3.Zero[float64](), b: vector3.New(0., 0., 0.), want: 0},
		"(0, 0, 0), (0, 1, 0)":  {a: vector3.Zero[float64](), b: vector3.New(0., 1., 0.), want: 1},
		"(0, -1, 0), (0, 1, 0)": {a: vector3.New(0., -1., 0.), b: vector3.New(0., 1., 0.), want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want, tc.a.Distance(tc.b), 0.000001)
		})
	}
}

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

func TestSub(t *testing.T) {
	tests := map[string]struct {
		left  vector3.Float64
		right vector3.Float64
		want  vector3.Float64
	}{
		"0, 0, 0 - 0, 0, 0 = 0, 0, 0": {left: vector3.New(0., 0., 0.), right: vector3.New(0., 0., 0.), want: vector3.New(0., 0., 0.)},
		"4, 5, 6 - 1, 2, 3 = 3, 3, 3": {left: vector3.New(4., 5., 6.), right: vector3.New(1., 2., 3.), want: vector3.New(3., 3., 3.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.left.Sub(tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestMidpoint(t *testing.T) {
	tests := map[string]struct {
		left  vector3.Float64
		right vector3.Float64
		want  vector3.Float64
	}{
		"0, 0, 0 m 0, 0, 0 = 0, 0, 0":     {left: vector3.New(0., 0., 0.), right: vector3.New(0., 0., 0.), want: vector3.New(0., 0., 0.)},
		"-1, -1, -1 m 1, 1, 1 = 0, 0, 0":  {left: vector3.New(-1., -1., -1.), right: vector3.New(1., 1., 1.), want: vector3.New(0., 0., 0.)},
		"0, 0, 0 m 1, 2, 3 = 0.5, 1, 1.5": {left: vector3.New(0., 0., 0.), right: vector3.New(1., 2., 3.), want: vector3.New(0.5, 1., 1.5)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.left.Midpoint(tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestLerp(t *testing.T) {
	tests := map[string]struct {
		left  vector3.Float64
		right vector3.Float64
		t     float64
		want  vector3.Float64
	}{
		"(0, 0, 0) =(0)=> (0, 0, 0) = (0, 0, 0)":       {left: vector3.New(0., 0., 0.), right: vector3.New(0., 0., 0.), t: 0, want: vector3.New(0., 0., 0.)},
		"(0, 0, 0) =(0.5)=> (1, 2, 3) = (0.5, 1, 1.5)": {left: vector3.New(0., 0., 0.), right: vector3.New(1., 2., 3.), t: 0.5, want: vector3.New(0.5, 1., 1.5)},
		"(0, 0, 0) =(1)=> (1, 2, 3) = (1, 2, 3)":       {left: vector3.New(0., 0., 0.), right: vector3.New(1., 2., 3.), t: 1, want: vector3.New(1., 2., 3.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector3.Lerp(tc.left, tc.right, tc.t)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestMin(t *testing.T) {
	tests := map[string]struct {
		left  vector3.Float64
		right vector3.Float64
		want  vector3.Float64
	}{
		"(1, 2, 3) m (3, 2, 1) = (1, 2, 1)": {left: vector3.New(1., 2., 3.), right: vector3.New(3., 2., 1.), want: vector3.New(1., 2., 1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector3.Min(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		left  vector3.Float64
		right vector3.Float64
		want  vector3.Float64
	}{
		"(1, 2, 3) m (3, 2, 1) = (3, 2, 3)": {left: vector3.New(1., 2., 3.), right: vector3.New(3., 2., 1.), want: vector3.New(3., 2., 3.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector3.Max(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestScaleVecFloat(t *testing.T) {
	tests := map[string]struct {
		vec    vector3.Float64
		scalar float64
		want   vector3.Float64
	}{
		"1, 2, 3 *  2 =  2,  4,  6": {vec: vector3.New(1., 2., 3.), scalar: 2, want: vector3.New(2., 4., 6.)},
		"1, 2, 3 *  0 =  0,  0,  0": {vec: vector3.New(1., 2., 3.), scalar: 0, want: vector3.New(0., 0., 0.)},
		"1, 2, 3 * -2 = -2, -4, -6": {vec: vector3.New(1., 2., 3.), scalar: -2, want: vector3.New(-2., -4., -6.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.vec.Scale(tc.scalar)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestScaleVecInt(t *testing.T) {
	tests := map[string]struct {
		vec    vector3.Int
		scalar float64
		want   vector3.Int
	}{
		"1, 2, 3 *  2 =  2,  4,  6": {vec: vector3.New(1, 2, 3), scalar: 2, want: vector3.New(2, 4, 6)},
		"1, 2, 3 *  0 =  0,  0,  0": {vec: vector3.New(1, 2, 3), scalar: 0, want: vector3.New(0, 0, 0)},
		"1, 2, 3 * -2 = -2, -4, -6": {vec: vector3.New(1, 2, 3), scalar: -2, want: vector3.New(-2, -4, -6)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.vec.Scale(tc.scalar)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestNearZero(t *testing.T) {
	tests := map[string]struct {
		vec  vector3.Float64
		want bool
	}{
		"0, 0, 0":           {vec: vector3.New(0., 0., 0.), want: true},
		"0, 0, 1":           {vec: vector3.New(0., 0., 1.), want: false},
		"0, 0, .0000000001": {vec: vector3.New(0., 0., 0.0000000001), want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.vec.NearZero())
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

func TestJSON(t *testing.T) {
	in := vector3.New(1.2, 2.3, 3.4)
	out := vector3.New(0., 0., 0.)

	marshalledData, marshallErr := json.Marshal(in)
	unmarshallErr := json.Unmarshal(marshalledData, &out)

	assert.NoError(t, marshallErr)
	assert.NoError(t, unmarshallErr)
	assert.Equal(t, "{\"x\":1.2,\"y\":2.3,\"z\":3.4}", string(marshalledData))
	assert.Equal(t, 1.2, out.X())
	assert.Equal(t, 2.3, out.Y())
	assert.Equal(t, 3.4, out.Z())
}

var result float64

func BenchmarkDistance(b *testing.B) {
	var r float64
	a := vector3.New(1., 2., 3.)
	c := vector3.New(4., 5., 6.)
	for i := 0; i < b.N; i++ {
		r = a.Distance(c)
	}
	result = r
}

func BenchmarkDot(b *testing.B) {
	var r float64
	a := vector3.New(1., 2., 3.)
	c := vector3.New(4., 5., 6.)
	for i := 0; i < b.N; i++ {
		r = a.Dot(c)
	}
	result = r
}
