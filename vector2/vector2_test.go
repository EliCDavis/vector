package vector2_test

import (
	"encoding/json"
	"math"
	"math/rand"
	"testing"

	"github.com/EliCDavis/vector/vector2"
	"github.com/stretchr/testify/assert"
)

func TestToIntConversions(t *testing.T) {
	start := vector2.New(1.2, -2.4)

	tests := map[string]struct {
		want vector2.Int
		got  vector2.Int
	}{
		"round to int": {want: start.RoundToInt(), got: vector2.New(1, -2)},
		"floor to int": {want: start.FloorToInt(), got: vector2.New(1, -3)},
		"ceil to int":  {want: start.CeilToInt(), got: vector2.New(2, -2)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
		})
	}
}

func TestDistances(t *testing.T) {
	tests := map[string]struct {
		a    vector2.Float64
		b    vector2.Float64
		want float64
	}{
		"(0, 0), (0, 0)":  {a: vector2.Zero[float64](), b: vector2.New(0., 0.), want: 0},
		"(0, 0), (0, 1)":  {a: vector2.Zero[float64](), b: vector2.New(0., 1.), want: 1},
		"(0, -1), (0, 1)": {a: vector2.New(0., -1.), b: vector2.New(0., 1.), want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want, tc.a.Distance(tc.b), 0.000001)
		})
	}
}

func TestOperations(t *testing.T) {
	start := vector2.New(1.2, -2.4)

	randSource := rand.NewSource(42)
	r := rand.New(randSource)

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
		"mult by vec":   {want: start.MultByVector(vector2.New(2., 4.)), got: vector2.New(2.4, -9.6)},
		"center":        {want: vector2.Midpoint(start, vector2.New(2.4, 2.4)), got: vector2.New(1.8, 0.)},
		"fill":          {want: vector2.Fill(9.3), got: vector2.New(9.3, 9.3)},
		"yx":            {want: start.YX(), got: vector2.New(-2.4, 1.2)},
		"random":        {want: vector2.Rand(r), got: vector2.New(.373028361, 0.066000496)},
		"flip":          {got: start.Flip(), want: vector2.New(-1.2, 2.4)},
		"flipX":         {got: start.FlipX(), want: vector2.New(-1.2, -2.4)},
		"flipY":         {got: start.FlipY(), want: vector2.New(1.2, 2.4)},
		"mod(1)":        {got: start.Mod(1), want: vector2.New(.2, -.4)},

		// Math package functions
		"log":   {got: start.Log(), want: vector2.New(0.1823215, math.NaN())},
		"log10": {got: start.Log10(), want: vector2.New(0.0791812, math.NaN())},
		"log2":  {got: start.Log2(), want: vector2.New(0.263034, math.NaN())},
		"exp":   {got: start.Exp(), want: vector2.New(3.320116, 0.090717)},
		"exp2":  {got: start.Exp2(), want: vector2.New(2.297396, 0.189464)},
		"expm1": {got: start.Expm1(), want: vector2.New(2.320116, -0.909282)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
		})
	}
}

func TestLerpClamped(t *testing.T) {
	tests := map[string]struct {
		left  vector2.Float64
		right vector2.Float64
		t     float64
		want  vector2.Float64
	}{
		"(0, 0, 0) =(0)=> (0, 0, 0) = (0, 0, 0)":       {left: vector2.New(0., 0.), right: vector2.New(0., 0.), t: 0, want: vector2.New(0., 0.)},
		"(0, 0, 0) =(0.5)=> (1, 2, 3) = (0.5, 1, 1.5)": {left: vector2.New(0., 0.), right: vector2.New(1., 2.), t: 0.5, want: vector2.New(0.5, 1.)},
		"(0, 0, 0) =(1)=> (1, 2, 3) = (1, 2, 3)":       {left: vector2.New(0., 0.), right: vector2.New(1., 2.), t: 1, want: vector2.New(1., 2.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector2.LerpClamped(tc.left, tc.right, tc.t)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
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

func TestBadJSON(t *testing.T) {
	out := vector2.New(0., 0.)

	unmarshallErr := out.UnmarshalJSON([]byte("bad json"))

	assert.Error(t, unmarshallErr)
	assert.Equal(t, 0., out.X())
	assert.Equal(t, 0., out.Y())
}

func TestToArray(t *testing.T) {
	v := vector2.New(1., 2.)

	arr := v.ToArr()
	assert.Len(t, arr, 2)
	assert.Equal(t, 1., arr[0])
	assert.Equal(t, 2., arr[1])

	arrFixed := v.ToFixedArr()
	assert.Len(t, arrFixed, 2)
	assert.Equal(t, 1., arrFixed[0])
	assert.Equal(t, 2., arrFixed[1])
}

func TestDot(t *testing.T) {
	a := vector2.New(2, 3)
	b := vector2.New(6, 7)

	assert.Equal(t, 33., a.Dot(b))
}

func TestLengthSquared(t *testing.T) {
	a := vector2.New(2, 3)

	assert.Equal(t, 13., a.LengthSquared())
}

func TestFromArray(t *testing.T) {
	tests := map[string]struct {
		arr  []float64
		want vector2.Float64
	}{
		"nil => (0, 0, 0)":    {arr: nil, want: vector2.Zero[float64]()},
		"[] => (0, 0, 0)":     {arr: []float64{}, want: vector2.Zero[float64]()},
		"[1] => (1, 0, 0)":    {arr: []float64{1}, want: vector2.New(1., 0.)},
		"[1, 2] => (1, 2, 0)": {arr: []float64{1, 2}, want: vector2.New(1., 2.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector2.FromArray(tc.arr)
			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
		})
	}
}

func TestToInt(t *testing.T) {
	in := vector2.New(1.2, 2.3)
	out := in.ToInt()
	assert.Equal(t, 1, out.X())
	assert.Equal(t, 2, out.Y())
}

func TestToInt64(t *testing.T) {
	in := vector2.New(1.2, 2.3)
	out := in.ToInt64()
	assert.Equal(t, int64(1), out.X())
	assert.Equal(t, int64(2), out.Y())
}

func TestToFloat32(t *testing.T) {
	in := vector2.New(1.2, 2.3)
	out := in.ToFloat32()
	assert.Equal(t, float32(1.2), out.X())
	assert.Equal(t, float32(2.3), out.Y())
}

func TestToFloat64(t *testing.T) {
	in := vector2.New(1, 2)
	out := in.ToFloat64()
	assert.Equal(t, float64(1), out.X())
	assert.Equal(t, float64(2), out.Y())
}

func TestMaxComponent(t *testing.T) {
	assert.Equal(t, 4., vector2.New(-2., 4.).MaxComponent())
}

func TestMinComponent(t *testing.T) {
	assert.Equal(t, -2., vector2.New(-2., 4.).MinComponent())
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

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		vec       vector2.Int
		formatter string
		want      string
	}{
		"1 2":  {vec: vector2.New(1, 2), formatter: "%d %d", want: "1 2"},
		"1, 2": {vec: vector2.New(1, 2), formatter: "%d, %d", want: "1, 2"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.vec.Format(tc.formatter)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestContainsNaN(t *testing.T) {
	tests := map[string]struct {
		vec  vector2.Float64
		want bool
	}{
		"x nan":  {vec: vector2.New(math.NaN(), 0.), want: true},
		"y nan":  {vec: vector2.New(0., math.NaN()), want: true},
		"no nan": {vec: vector2.New(0., 0.), want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.vec.ContainsNaN())
		})
	}
}

func TestAngle(t *testing.T) {
	tests := map[string]struct {
		a     vector2.Float64
		b     vector2.Float64
		angle float64
	}{
		"up => down: Pi": {
			a:     vector2.Up[float64](),
			b:     vector2.Down[float64](),
			angle: math.Pi,
		},
		"up => right: Pi": {
			a:     vector2.Up[float64](),
			b:     vector2.Right[float64](),
			angle: math.Pi / 2,
		},
		"up => up: 0": {
			a:     vector2.Up[float64](),
			b:     vector2.Up[float64](),
			angle: 0,
		},
		"0 => 0: 0": {
			a:     vector2.Zero[float64](),
			b:     vector2.Zero[float64](),
			angle: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.angle, tc.a.Angle(tc.b), 0.000001)
		})
	}
}

func TestMaxMinComponents(t *testing.T) {
	tests := map[string]struct {
		a    vector2.Float64
		b    vector2.Float64
		f    func(a, b vector2.Float64) float64
		want float64
	}{
		"maxX((0, 0), (1, 0))": {a: vector2.New(0., 0.), b: vector2.New(1., 0.), f: vector2.MaxX[float64], want: 1},
		"maxX((2, 0), (0, 0))": {a: vector2.New(2., 0.), b: vector2.New(0., 0.), f: vector2.MaxX[float64], want: 2},
		"maxY((0, 0), (0, 1))": {a: vector2.New(0., 0.), b: vector2.New(0., 1.), f: vector2.MaxY[float64], want: 1},
		"maxY((0, 2), (0, 0))": {a: vector2.New(0., 2.), b: vector2.New(0., 0.), f: vector2.MaxY[float64], want: 2},

		"minX((0, 0, 0), (-1, 0))": {a: vector2.New(0., 0.), b: vector2.New(-1., 0.), f: vector2.MinX[float64], want: -1},
		"minX((-2, 0, 0), (0, 0))": {a: vector2.New(-2., 0.), b: vector2.New(0., 0.), f: vector2.MinX[float64], want: -2},
		"minY((0, 0, 0), (0, -1))": {a: vector2.New(0., 0.), b: vector2.New(0., -1.), f: vector2.MinY[float64], want: -1},
		"minY((0, -2, 0), (0, 0))": {a: vector2.New(0., -2.), b: vector2.New(0., 0.), f: vector2.MinY[float64], want: -2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.f(tc.a, tc.b))
		})
	}
}

func TestValues(t *testing.T) {
	x, y := vector2.New(1, 2).Values()
	assert.Equal(t, x, 1)
	assert.Equal(t, y, 2)
}

func TestComponent(t *testing.T) {

	v := vector2.New(1., 2.)
	tests := map[string]struct {
		component int
		want      float64
	}{
		"0": {component: 0, want: 1.},
		"1": {component: 1, want: 2.},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, v.Component(tc.component))
		})
	}

	assert.PanicsWithError(t, "invalid index: -1", func() {
		v.Component(-1)
	})
}
