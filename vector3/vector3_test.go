package vector3_test

import (
	"encoding/json"
	"image/color"
	"math"
	"math/rand"
	"testing"

	"github.com/EliCDavis/vector/vector2"
	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func TestVectorOperations(t *testing.T) {
	start := vector3.New(1.2, -2.4, 3.7)

	randSource := rand.NewSource(42)
	r := rand.New(randSource)

	tests := map[string]struct {
		got  vector3.Float64
		want vector3.Float64
	}{
		"x":            {got: start.SetX(4), want: vector3.New(4., -2.4, 3.7)},
		"y":            {got: start.SetY(4), want: vector3.New(1.2, 4., 3.7)},
		"z":            {got: start.SetZ(4), want: vector3.New(1.2, -2.4, 4.)},
		"abs":          {got: start.Abs(), want: vector3.New(1.2, 2.4, 3.7)},
		"floor":        {got: start.Floor(), want: vector3.New(1., -3., 3.)},
		"ceil":         {got: start.Ceil(), want: vector3.New(2., -2., 4.)},
		"round":        {got: start.Round(), want: vector3.New(1., -2., 4.)},
		"multByVector": {got: start.MultByVector(vector3.New(2., 4., 6.)), want: vector3.New(2.4, -9.6, 22.2)},
		"sqrt":         {got: start.Sqrt(), want: vector3.New(1.0954451, math.NaN(), 1.923538)},
		"clamp":        {got: start.Clamp(1, 2), want: vector3.New(1.2, 1., 2.)},
		"cross":        {got: start.Cross(vector3.New(2., 3., 4.)), want: vector3.New(-20.7, 2.6, 8.4)},
		"center":       {got: vector3.Midpoint(start, vector3.New(2.4, 2.4, 4.7)), want: vector3.New(1.8, 0., 4.2)},
		"fill":         {got: vector3.Fill(9.3), want: vector3.New(9.3, 9.3, 9.3)},
		"color black":  {got: vector3.FromColor(color.Black), want: vector3.New(0., 0., 0.)},
		"color white":  {got: vector3.FromColor(color.White), want: vector3.New(1., 1., 1.)},
		"xzy":          {got: start.XZY(), want: vector3.New(1.2, 3.7, -2.4)},
		"zxy":          {got: start.ZXY(), want: vector3.New(3.7, 1.2, -2.4)},
		"zyx":          {got: start.ZYX(), want: vector3.New(3.7, -2.4, 1.2)},
		"yzx":          {got: start.YZX(), want: vector3.New(-2.4, 3.7, 1.2)},
		"yxz":          {got: start.YXZ(), want: vector3.New(-2.4, 1.2, 3.7)},
		"random":       {got: vector3.Rand(r), want: vector3.New(.373028361, 0.066000496, 0.604093851)},
		"flip":         {got: start.Flip(), want: vector3.New(-1.2, 2.4, -3.7)},
		"flipX":        {got: start.FlipX(), want: vector3.New(-1.2, -2.4, 3.7)},
		"flipY":        {got: start.FlipY(), want: vector3.New(1.2, 2.4, 3.7)},
		"flipZ":        {got: start.FlipZ(), want: vector3.New(1.2, -2.4, -3.7)},
		"random range": {
			got:  vector3.RandRange(r, -2., 4.),
			want: vector3.New(-0.7470877, -1.737089, 0.299159),
		},
		"random normal": {
			got:  vector3.RandNormal(r),
			want: vector3.New(0.8852213, -0.326936, -0.330901),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
		})
	}
}

func TestToVector2(t *testing.T) {
	start := vector3.New(1.2, -2.4, 3.7)

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
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
		})
	}
}

func TestToIntConversions(t *testing.T) {
	start := vector3.New(1.2, -2.4, 3.7)

	tests := map[string]struct {
		want vector3.Int
		got  vector3.Int
	}{
		"round to int": {want: start.RoundToInt(), got: vector3.New(1, -2, 4)},
		"floor to int": {want: start.FloorToInt(), got: vector3.New(1, -3, 3)},
		"ceil to int":  {want: start.CeilToInt(), got: vector3.New(2, -2, 4)},
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

func TestFromArray(t *testing.T) {
	tests := map[string]struct {
		arr  []float64
		want vector3.Float64
	}{
		"nil => (0, 0, 0)":       {arr: nil, want: vector3.Zero[float64]()},
		"[] => (0, 0, 0)":        {arr: []float64{}, want: vector3.Zero[float64]()},
		"[1] => (1, 0, 0)":       {arr: []float64{1}, want: vector3.New(1., 0., 0.)},
		"[1, 2] => (1, 2, 0)":    {arr: []float64{1, 2}, want: vector3.New(1., 2., 0.)},
		"[1, 2, 3] => (1, 2, 3)": {arr: []float64{1, 2, 3}, want: vector3.New(1., 2., 3.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector3.FromArray(tc.arr)
			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
		})
	}
}

func TestAverage(t *testing.T) {
	// ASSIGN =================================================================
	vals := []vector3.Float64{
		vector3.New(1., 2., 3.),
		vector3.New(1., 2., 3.),
		vector3.New(1., 2., 3.),
	}

	// ACT ====================================================================
	avg := vector3.Average(vals)

	// ASSERT =================================================================
	assert.InDelta(t, 1., avg.X(), 0.000001)
	assert.InDelta(t, 2., avg.Y(), 0.000001)
	assert.InDelta(t, 3., avg.Z(), 0.000001)
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

func TestToArray(t *testing.T) {
	v := vector3.New(1., 2., 3.)
	arr := v.ToArr()
	assert.Len(t, arr, 3)
	assert.Equal(t, 1., arr[0])
	assert.Equal(t, 2., arr[1])
	assert.Equal(t, 3., arr[2])
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

func TestContainsNaN(t *testing.T) {
	tests := map[string]struct {
		vec  vector3.Float64
		want bool
	}{
		"x nan":  {vec: vector3.New(math.NaN(), 0., 0.), want: true},
		"y nan":  {vec: vector3.New(0., math.NaN(), 0.), want: true},
		"z nan":  {vec: vector3.New(0., 0., math.NaN()), want: true},
		"no nan": {vec: vector3.New(0., 0., 0.), want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.vec.ContainsNaN())
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

func TestBadJSON(t *testing.T) {
	out := vector3.New(0., 0., 0.)

	unmarshallErr := out.UnmarshalJSON([]byte("bad json"))

	assert.Error(t, unmarshallErr)
	assert.Equal(t, 0., out.X())
	assert.Equal(t, 0., out.Y())
	assert.Equal(t, 0., out.Z())
}

func TestDot(t *testing.T) {
	a := vector3.New(2, 3, 4)
	b := vector3.New(6, 7, 8)

	assert.Equal(t, 65., a.Dot(b))
}

func TestToInt(t *testing.T) {
	in := vector3.New(1.2, 2.3, 3.4)
	out := in.ToInt()
	assert.Equal(t, 1, out.X())
	assert.Equal(t, 2, out.Y())
	assert.Equal(t, 3, out.Z())
}

func TestToInt64(t *testing.T) {
	in := vector3.New(1.2, 2.3, 3.4)
	out := in.ToInt64()
	assert.Equal(t, int64(1), out.X())
	assert.Equal(t, int64(2), out.Y())
	assert.Equal(t, int64(3), out.Z())
}

func TestToFloat32(t *testing.T) {
	in := vector3.New(1.2, 2.3, 3.4)
	out := in.ToFloat32()
	assert.Equal(t, float32(1.2), out.X())
	assert.Equal(t, float32(2.3), out.Y())
	assert.Equal(t, float32(3.4), out.Z())
}

func TestToFloat64(t *testing.T) {
	in := vector3.New(1, 2, 3)
	out := in.ToFloat64()
	assert.Equal(t, float64(1), out.X())
	assert.Equal(t, float64(2), out.Y())
	assert.Equal(t, float64(3), out.Z())
}

func TestAngle(t *testing.T) {
	tests := map[string]struct {
		a     vector3.Float64
		b     vector3.Float64
		angle float64
	}{
		"up => down: Pi": {
			a:     vector3.Up[float64](),
			b:     vector3.Down[float64](),
			angle: math.Pi,
		},
		"up => right: Pi": {
			a:     vector3.Up[float64](),
			b:     vector3.Right[float64](),
			angle: math.Pi / 2,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.angle, tc.a.Angle(tc.b), 0.000001)
		})
	}
}

func TestMaxComponent(t *testing.T) {
	assert.Equal(t, 4., vector3.New(-2., 3., 4.).MaxComponent())
}

func TestMinComponent(t *testing.T) {
	assert.Equal(t, -2., vector3.New(-2., 3., 4.).MinComponent())
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

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		vec       vector3.Int
		formatter string
		want      string
	}{
		"1 2 3":   {vec: vector3.New(1, 2, 3), formatter: "%d %d %d", want: "1 2 3"},
		"1, 2, 3": {vec: vector3.New(1, 2, 3), formatter: "%d, %d, %d", want: "1, 2, 3"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.vec.Format(tc.formatter)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMaxMinComponents(t *testing.T) {
	tests := map[string]struct {
		a    vector3.Float64
		b    vector3.Float64
		f    func(a, b vector3.Float64) float64
		want float64
	}{
		"maxX((0, 0, 0), (1, 0, 0))": {a: vector3.New(0., 0., 0.), b: vector3.New(1., 0., 0.), f: vector3.MaxX[float64], want: 1},
		"maxX((2, 0, 0), (0, 0, 0))": {a: vector3.New(2., 0., 0.), b: vector3.New(0., 0., 0.), f: vector3.MaxX[float64], want: 2},
		"maxY((0, 0, 0), (0, 1, 0))": {a: vector3.New(0., 0., 0.), b: vector3.New(0., 1., 0.), f: vector3.MaxY[float64], want: 1},
		"maxY((0, 2, 0), (0, 0, 0))": {a: vector3.New(0., 2., 0.), b: vector3.New(0., 0., 0.), f: vector3.MaxY[float64], want: 2},
		"maxZ((0, 0, 0), (0, 0, 1))": {a: vector3.New(0., 0., 0.), b: vector3.New(0., 0., 1.), f: vector3.MaxZ[float64], want: 1},
		"maxZ((0, 0, 2), (0, 0, 0))": {a: vector3.New(0., 0., 2.), b: vector3.New(0., 0., 0.), f: vector3.MaxZ[float64], want: 2},

		"minX((0, 0, 0), (-1, 0, 0))": {a: vector3.New(0., 0., 0.), b: vector3.New(-1., 0., 0.), f: vector3.MinX[float64], want: -1},
		"minX((-2, 0, 0), (0, 0, 0))": {a: vector3.New(-2., 0., 0.), b: vector3.New(0., 0., 0.), f: vector3.MinX[float64], want: -2},
		"minY((0, 0, 0), (0, -1, 0))": {a: vector3.New(0., 0., 0.), b: vector3.New(0., -1., 0.), f: vector3.MinY[float64], want: -1},
		"minY((0, -2, 0), (0, 0, 0))": {a: vector3.New(0., -2., 0.), b: vector3.New(0., 0., 0.), f: vector3.MinY[float64], want: -2},
		"minZ((0, 0, 0), (0, 0, -1))": {a: vector3.New(0., 0., 0.), b: vector3.New(0., 0., -1.), f: vector3.MinZ[float64], want: -1},
		"minZ((0, 0, -2), (0, 0, 0))": {a: vector3.New(0., 0., -2.), b: vector3.New(0., 0., 0.), f: vector3.MinZ[float64], want: -2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.f(tc.a, tc.b))
		})
	}
}
