package vector4_test

import (
	"encoding/json"
	"image/color"
	"math"
	"testing"

	"github.com/EliCDavis/vector/vector2"
	"github.com/EliCDavis/vector/vector3"
	"github.com/EliCDavis/vector/vector4"
	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	tests := map[string]struct {
		got  vector4.Float64
		want vector4.Float64
	}{
		"zero": {got: vector4.Zero[float64](), want: vector4.New(0., 0., 0., 0.)},
		"one":  {got: vector4.One[float64](), want: vector4.New(1., 1., 1., 1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), tc.got.W(), 0.000001)
		})
	}
}

func TestVectorOperations(t *testing.T) {
	start := vector4.New(1.2, -2.4, 3.7, 4.9)

	tests := map[string]struct {
		want vector4.Float64
		got  vector4.Float64
	}{
		"x":            {want: start.SetX(4), got: vector4.New(4., -2.4, 3.7, 4.9)},
		"y":            {want: start.SetY(4), got: vector4.New(1.2, 4., 3.7, 4.9)},
		"z":            {want: start.SetZ(4), got: vector4.New(1.2, -2.4, 4., 4.9)},
		"w":            {want: start.SetW(4), got: vector4.New(1.2, -2.4, 3.7, 4.)},
		"add":          {want: start.Add(vector4.New(1., -2., 3., 4.)), got: vector4.New(2.2, -4.4, 6.7, 8.9)},
		"sub":          {want: start.Sub(vector4.New(1., -2., 3., 4.)), got: vector4.New(0.2, -0.4, 0.7, 0.9)},
		"div":          {want: start.DivByConstant(2), got: vector4.New(0.6, -1.2, 1.85, 2.45)},
		"abs":          {want: start.Abs(), got: vector4.New(1.2, 2.4, 3.7, 4.9)},
		"floor":        {want: start.Floor(), got: vector4.New(1., -3., 3., 4.)},
		"ceil":         {want: start.Ceil(), got: vector4.New(2., -2., 4., 5.)},
		"round":        {want: start.Round(), got: vector4.New(1., -2., 4., 5.)},
		"multByVector": {want: start.MultByVector(vector4.New(2., 4., 6., 7.)), got: vector4.New(2.4, -9.6, 22.2, 34.3)},
		"divByVector":  {got: start.DivByVector(vector4.New(2., 4., 6., 7.)), want: vector4.New(0.6, -0.6, 0.6166666666666, .7)},
		"sqrt":         {want: start.Sqrt(), got: vector4.New(1.0954451, math.NaN(), 1.923538, 2.213594)},
		"clamp":        {want: start.Clamp(1, 2), got: vector4.New(1.2, 1., 2., 2.)},
		"center":       {want: vector4.Midpoint(start, vector4.New(2.4, 2.4, 4.7, 4.7)), got: vector4.New(1.8, 0., 4.2, 4.8)},
		"fill":         {want: vector4.Fill(9.3), got: vector4.New(9.3, 9.3, 9.3, 9.3)},
		"color black":  {want: vector4.FromColor(color.Black), got: vector4.New(0., 0., 0., 1.)},
		"color white":  {want: vector4.FromColor(color.White), got: vector4.New(1., 1., 1., 1.)},
		"flip":         {got: start.Flip(), want: vector4.New(-1.2, 2.4, -3.7, -4.9)},
		"flipX":        {got: start.FlipX(), want: vector4.New(-1.2, -2.4, 3.7, 4.9)},
		"flipY":        {got: start.FlipY(), want: vector4.New(1.2, 2.4, 3.7, 4.9)},
		"flipZ":        {got: start.FlipZ(), want: vector4.New(1.2, -2.4, -3.7, 4.9)},
		"flipW":        {got: start.FlipW(), want: vector4.New(1.2, -2.4, 3.7, -4.9)},
		"mod(1)":       {got: start.Mod(1), want: vector4.New(.2, -.4, .7, .9)},
		"normalize":    {got: start.Normalized(), want: vector4.New(0.1790845316, -0.35816906, 0.5521773, 0.73126183)},

		// Math package functions
		"log":   {got: start.Log(), want: vector4.New(0.1823215, math.NaN(), 1.308332, 1.589235)},
		"log10": {got: start.Log10(), want: vector4.New(0.0791812, math.NaN(), 0.568201, 0.690196)},
		"log2":  {got: start.Log2(), want: vector4.New(0.263034, math.NaN(), 1.887525, 2.292781)},
		"exp":   {got: start.Exp(), want: vector4.New(3.320116, 0.090717, 40.447304, 134.289779)},
		"exp2":  {got: start.Exp2(), want: vector4.New(2.297396, 0.189464, 12.996038, 29.8570557)},
		"expm1": {got: start.Expm1(), want: vector4.New(2.320116, -0.909282, 39.447304, 133.289779)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), tc.got.W(), 0.000001)
		})
	}
}

func TestToIntConversions(t *testing.T) {
	start := vector4.New(1.2, -2.4, 3.7, 4.9)

	tests := map[string]struct {
		want vector4.Int
		got  vector4.Int
	}{
		"round to int": {want: start.RoundToInt(), got: vector4.New(1, -2, 4, 5)},
		"floor to int": {want: start.FloorToInt(), got: vector4.New(1, -3, 3, 4)},
		"ceil to int":  {want: start.CeilToInt(), got: vector4.New(2, -2, 4, 5)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), tc.got.W(), 0.000001)
		})
	}
}

func TestScaleVecFloat(t *testing.T) {
	tests := map[string]struct {
		vec    vector4.Float64
		scalar float64
		want   vector4.Float64
	}{
		"1, 2, 3, 4 *  2 =  2,  4,  6,  8": {vec: vector4.New(1., 2., 3., 4.), scalar: 2, want: vector4.New(2., 4., 6., 8.)},
		"1, 2, 3, 4 *  0 =  0,  0,  0,  0": {vec: vector4.New(1., 2., 3., 4.), scalar: 0, want: vector4.New(0., 0., 0., 0.)},
		"1, 2, 3, 4 * -2 = -2, -4, -6, -8": {vec: vector4.New(1., 2., 3., 4.), scalar: -2, want: vector4.New(-2., -4., -6., -8.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.vec.Scale(tc.scalar)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestSwizzle_Vector3(t *testing.T) {
	in := vector4.New(1., 2., 3., 4.)

	tests := map[string]struct {
		expected vector3.Float64
		got      vector3.Float64
	}{
		"XYZ": {expected: vector3.New(1., 2., 3.), got: in.XYZ()},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected.X(), tc.got.X())
			assert.Equal(t, tc.expected.Y(), tc.got.Y())
			assert.Equal(t, tc.expected.Z(), tc.got.Z())
		})
	}
}

func TestSwizzle_Vector2(t *testing.T) {
	start := vector4.New(1.2, -2.4, 3.7, 12.)

	tests := map[string]struct {
		got  vector2.Float64
		want vector2.Float64
	}{
		"xy": {got: start.XY(), want: vector2.New(1.2, -2.4)},
		"yz": {got: start.YZ(), want: vector2.New(-2.4, 3.7)},
		"xz": {got: start.XZ(), want: vector2.New(1.2, 3.7)},
		"yx": {got: start.YX(), want: vector2.New(-2.4, 1.2)},
		"zy": {got: start.ZY(), want: vector2.New(3.7, -2.4)},
		"zx": {got: start.ZX(), want: vector2.New(3.7, 1.2)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want.X(), tc.got.X())
			assert.Equal(t, tc.want.Y(), tc.got.Y())
		})
	}
}

func TestJSON(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := vector4.New(0., 0., 0., 0.)

	marshalledData, marshallErr := json.Marshal(in)
	unmarshallErr := json.Unmarshal(marshalledData, &out)

	assert.NoError(t, marshallErr)
	assert.NoError(t, unmarshallErr)
	assert.Equal(t, "{\"x\":1.2,\"y\":2.3,\"z\":3.4,\"w\":5.6}", string(marshalledData))
	assert.Equal(t, 1.2, out.X())
	assert.Equal(t, 2.3, out.Y())
	assert.Equal(t, 3.4, out.Z())
	assert.Equal(t, 5.6, out.W())
}

func TestBadJSON(t *testing.T) {
	out := vector4.New(0., 0., 0., 0.)

	unmarshallErr := out.UnmarshalJSON([]byte("bad json"))

	assert.Error(t, unmarshallErr)
	assert.Equal(t, 0., out.X())
	assert.Equal(t, 0., out.Y())
	assert.Equal(t, 0., out.Z())
	assert.Equal(t, 0., out.W())
}

func TestDot(t *testing.T) {
	a := vector4.New(2, 3, 4, 5)
	b := vector4.New(6, 7, 8, 9)

	assert.Equal(t, 110., a.Dot(b))
}

func TestFromArray(t *testing.T) {
	tests := map[string]struct {
		arr  []float64
		want vector4.Float64
	}{
		"nil => (0, 0, 0, 0)":          {arr: nil, want: vector4.Zero[float64]()},
		"[] => (0, 0, 0, 0)":           {arr: []float64{}, want: vector4.Zero[float64]()},
		"[1] => (1, 0, 0, 0)":          {arr: []float64{1}, want: vector4.New(1., 0., 0., 0.)},
		"[1, 2] => (1, 2, 0, 0)":       {arr: []float64{1, 2}, want: vector4.New(1., 2., 0., 0.)},
		"[1, 2, 3] => (1, 2, 3, 0)":    {arr: []float64{1, 2, 3}, want: vector4.New(1., 2., 3., 0.)},
		"[1, 2, 3, 4] => (1, 2, 3, 4)": {arr: []float64{1, 2, 3, 4}, want: vector4.New(1., 2., 3., 4.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.FromArray(tc.arr)
			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestAverage(t *testing.T) {
	// ASSIGN =================================================================
	vals := []vector4.Float64{
		vector4.New(1., 2., 3., 4.),
		vector4.New(1., 2., 3., 4.),
		vector4.New(1., 2., 3., 4.),
	}

	// ACT ====================================================================
	avg := vector4.Average(vals)

	// ASSERT =================================================================
	assert.InDelta(t, 1., avg.X(), 0.000001)
	assert.InDelta(t, 2., avg.Y(), 0.000001)
	assert.InDelta(t, 3., avg.Z(), 0.000001)
	assert.InDelta(t, 4., avg.W(), 0.000001)
}

func TestLerp(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		t     float64
		want  vector4.Float64
	}{
		"(0, 0, 0, 0) =(0)=> (0, 0, 0, 0) = (0, 0, 0, 0)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(0., 0., 0., 0.),
			t:     0,
			want:  vector4.New(0., 0., 0., 0.),
		},
		"(0, 0, 0, 0) =(0.5)=> (1, 2, 3, 4) = (0.5, 1, 1.5, 2.0)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(1., 2., 3., 4.),
			t:     0.5,
			want:  vector4.New(0.5, 1., 1.5, 2.),
		},
		"(0, 0, 0, 0) =(1)=> (1, 2, 3, 4) = (1, 2, 3, 4)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(1., 2., 3., 4.),
			t:     1,
			want:  vector4.New(1., 2., 3., 4.),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.Lerp(tc.left, tc.right, tc.t)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestLerpClamped(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		t     float64
		want  vector4.Float64
	}{
		"(0, 0, 0, 0) =(0)=> (0, 0, 0, 0) = (0, 0, 0, 0)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(0., 0., 0., 0.),
			t:     0,
			want:  vector4.New(0., 0., 0., 0.),
		},
		"(0, 0, 0, 0) =(0.5)=> (1, 2, 3, 4) = (0.5, 1, 1.5, 2.0)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(1., 2., 3., 4.),
			t:     0.5,
			want:  vector4.New(0.5, 1., 1.5, 2.),
		},
		"(0, 0, 0, 0) =(1)=> (1, 2, 3, 4) = (1, 2, 3, 4)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(1., 2., 3., 4.),
			t:     1,
			want:  vector4.New(1., 2., 3., 4.),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.LerpClamped(tc.left, tc.right, tc.t)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestToArray(t *testing.T) {
	v := vector4.New(1., 2., 3., 4.)

	arr := v.ToArr()
	assert.Len(t, arr, 4)
	assert.Equal(t, 1., arr[0])
	assert.Equal(t, 2., arr[1])
	assert.Equal(t, 3., arr[2])
	assert.Equal(t, 4., arr[3])

	arrFixed := v.ToFixedArr()
	assert.Len(t, arrFixed, 4)
	assert.Equal(t, 1., arrFixed[0])
	assert.Equal(t, 2., arrFixed[1])
	assert.Equal(t, 3., arrFixed[2])
	assert.Equal(t, 4., arrFixed[3])
}

func TestMin(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		want  vector4.Float64
	}{
		"(1, 2, 3, 4) m (4, 3, 2, 1) = (1, 2, 2, 1)": {left: vector4.New(1., 2., 3., 4.), right: vector4.New(4., 3., 2., 1.), want: vector4.New(1., 2., 2., 1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.Min(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		want  vector4.Float64
	}{
		"(1, 2, 3, 4) m (4, 3, 2, 1) = (1, 2, 2, 1)": {left: vector4.New(1., 2., 3., 4.), right: vector4.New(4., 3., 2., 1.), want: vector4.New(4., 3., 3., 4.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.Max(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestToInt(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := in.ToInt()
	assert.Equal(t, 1, out.X())
	assert.Equal(t, 2, out.Y())
	assert.Equal(t, 3, out.Z())
	assert.Equal(t, 5, out.W())
}

func TestToInt64(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := in.ToInt64()
	assert.Equal(t, int64(1), out.X())
	assert.Equal(t, int64(2), out.Y())
	assert.Equal(t, int64(3), out.Z())
	assert.Equal(t, int64(5), out.W())
}

func TestToFloat32(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := in.ToFloat32()
	assert.Equal(t, float32(1.2), out.X())
	assert.Equal(t, float32(2.3), out.Y())
	assert.Equal(t, float32(3.4), out.Z())
	assert.Equal(t, float32(5.6), out.W())
}

func TestToFloat64(t *testing.T) {
	in := vector4.New(1, 2, 3, 5)
	out := in.ToFloat64()
	assert.Equal(t, float64(1), out.X())
	assert.Equal(t, float64(2), out.Y())
	assert.Equal(t, float64(3), out.Z())
	assert.Equal(t, float64(5), out.W())
}

func TestMaxComponent(t *testing.T) {
	assert.Equal(t, 4., vector4.New(-2., 3., 4., -1.).MaxComponent())
}

func TestMinComponent(t *testing.T) {
	assert.Equal(t, -2., vector4.New(-2., 3., 4., -1.).MinComponent())
}

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		vec       vector4.Int
		formatter string
		want      string
	}{
		"1 2 3 4":    {vec: vector4.New(1, 2, 3, 4), formatter: "%d %d %d %d", want: "1 2 3 4"},
		"1, 2, 3, 4": {vec: vector4.New(1, 2, 3, 4), formatter: "%d, %d, %d, %d", want: "1, 2, 3, 4"},
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
		vec  vector4.Float64
		want bool
	}{
		"x nan":  {vec: vector4.New(math.NaN(), 0., 0., 0.), want: true},
		"y nan":  {vec: vector4.New(0., math.NaN(), 0., 0.), want: true},
		"z nan":  {vec: vector4.New(0., 0., math.NaN(), 0.), want: true},
		"w nan":  {vec: vector4.New(0., 0., 0., math.NaN()), want: true},
		"no nan": {vec: vector4.New(0., 0., 0., 0.), want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.vec.ContainsNaN())
		})
	}
}

func TestNearZero(t *testing.T) {
	tests := map[string]struct {
		vec  vector4.Float64
		want bool
	}{
		"0, 0, 0":           {vec: vector4.New(0., 0., 0., 0.), want: true},
		"0, 0, 1":           {vec: vector4.New(0., 1., 0., 0.), want: false},
		"0, 0, .0000000001": {vec: vector4.New(0., 0.0000000001, 0., 0.), want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.vec.NearZero())
		})
	}
}

func TestMaxMinComponents(t *testing.T) {
	tests := map[string]struct {
		a    vector4.Float64
		b    vector4.Float64
		f    func(a, b vector4.Float64) float64
		want float64
	}{
		"maxX((0, 0, 0, 0), (1, 0, 0, 0))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(1., 0., 0., 0.), f: vector4.MaxX[float64], want: 1},
		"maxX((2, 0, 0, 0), (0, 0, 0, 0))": {a: vector4.New(2., 0., 0., 0.), b: vector4.New(0., 0., 0., 0.), f: vector4.MaxX[float64], want: 2},
		"maxY((0, 0, 0, 0), (0, 1, 0, 0))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(0., 1., 0., 0.), f: vector4.MaxY[float64], want: 1},
		"maxY((0, 2, 0, 0), (0, 0, 0, 0))": {a: vector4.New(0., 2., 0., 0.), b: vector4.New(0., 0., 0., 0.), f: vector4.MaxY[float64], want: 2},
		"maxZ((0, 0, 0, 0), (0, 0, 1, 0))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(0., 0., 1., 0.), f: vector4.MaxZ[float64], want: 1},
		"maxZ((0, 0, 2, 0), (0, 0, 0, 0))": {a: vector4.New(0., 0., 2., 0.), b: vector4.New(0., 0., 0., 0.), f: vector4.MaxZ[float64], want: 2},
		"maxW((0, 0, 0, 0), (0, 0, 0, 1))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(0., 0., 0., 1.), f: vector4.MaxW[float64], want: 1},
		"maxW((0, 0, 0, 2), (0, 0, 0, 0))": {a: vector4.New(0., 0., 0., 2.), b: vector4.New(0., 0., 0., 0.), f: vector4.MaxW[float64], want: 2},

		"minX((0, 0, 0, 0), (-1, 0, 0, 0))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(-1., 0., 0., 0.), f: vector4.MinX[float64], want: -1},
		"minX((-2, 0, 0, 0), (0, 0, 0, 0))": {a: vector4.New(-2., 0., 0., 0.), b: vector4.New(0., 0., 0., 0.), f: vector4.MinX[float64], want: -2},
		"minY((0, 0, 0, 0), (0, -1, 0, 0))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(0., -1., 0., 0.), f: vector4.MinY[float64], want: -1},
		"minY((0, -2, 0, 0), (0, 0, 0, 0))": {a: vector4.New(0., -2., 0., 0.), b: vector4.New(0., 0., 0., 0.), f: vector4.MinY[float64], want: -2},
		"minZ((0, 0, 0, 0), (0, 0, -1, 0))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(0., 0., -1., 0.), f: vector4.MinZ[float64], want: -1},
		"minZ((0, 0, -2, 0), (0, 0, 0, 0))": {a: vector4.New(0., 0., -2., 0.), b: vector4.New(0., 0., 0., 0.), f: vector4.MinZ[float64], want: -2},
		"minW((0, 0, 0, 0), (0, 0, 0, -1))": {a: vector4.New(0., 0., 0., 0.), b: vector4.New(0., 0., 0., -1.), f: vector4.MinW[float64], want: -1},
		"minW((0, 0, 0, -2), (0, 0, 0, 0))": {a: vector4.New(0., 0., 0., -2.), b: vector4.New(0., 0., 0., 0.), f: vector4.MinW[float64], want: -2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.f(tc.a, tc.b))
		})
	}
}

func TestValues(t *testing.T) {
	x, y, z, w := vector4.New(1, 2, 3, 4).Values()
	assert.Equal(t, x, 1)
	assert.Equal(t, y, 2)
	assert.Equal(t, z, 3)
	assert.Equal(t, w, 4)
}

func TestComponent(t *testing.T) {

	v := vector4.New(1., 2., 3., 4.)
	tests := map[string]struct {
		component int
		want      float64
	}{
		"0": {component: 0, want: 1.},
		"1": {component: 1, want: 2.},
		"2": {component: 2, want: 3.},
		"3": {component: 3, want: 4.},
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

var result vector4.Float64

func BenchmarkLerp(b *testing.B) {
	var r vector4.Float64
	a := vector4.New(1., 2., 3., 7.)
	c := vector4.New(4., 5., 6., 8.)
	for i := 0; i < b.N; i++ {
		r = vector4.Lerp(a, c, float64(i)/float64(b.N))
	}
	result = r
}
