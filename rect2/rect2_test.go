package rect2_test

import (
	"testing"

	"github.com/EliCDavis/vector/rect2"
	"github.com/EliCDavis/vector/test"
	"github.com/EliCDavis/vector/vector2"
)

func TestOperations(t *testing.T) {
	start := rect2.New(vector2.New(1.2, -2.4), vector2.New(10.8, 12.4))

	//randSource := rand.NewSource(42)
	//r := rand.New(randSource)

	tests := map[string]struct {
		want rect2.Float64
		got  rect2.Float64
	}{
		"x":     {got: start.SetX(4), want: rect2.New(vector2.New(4, -2.4), vector2.New(10.8, 12.4))},
		"y":     {got: start.SetY(4), want: rect2.New(vector2.New(1.2, 4), vector2.New(10.8, 12.4))},
		"addx":  {got: start.AddX(4), want: rect2.New(vector2.New(5.2, -2.4), vector2.New(10.8, 12.4))},
		"addy":  {got: start.AddY(4), want: rect2.New(vector2.New(1.2, 1.6), vector2.New(10.8, 12.4))},
		"floor": {got: start.Floor(), want: rect2.New(vector2.New(1., -3.), vector2.New(10., 12.))},
		"ceil":  {got: start.Ceil(), want: rect2.New(vector2.New(2., -2.), vector2.New(11., 13.))},
		"round": {got: start.Round(), want: rect2.New(vector2.New(1., -2.), vector2.New(11., 12.))},
		//"sqrt":           {got: start.Sqrt(), want: vector2.New(1.0954451, math.NaN())},
		//"clamp":          {got: start.Clamp(1, 2), want: vector2.New(1.2, 1.)},
		//"clampv":         {got: start.ClampV(vector2.New(0., 0.8), vector2.New(1., 2)), want: vector2.New(1., 0.8)},
		//"clamp0v":        {got: start.Clamp0V(vector2.New(1., 2.)), want: vector2.New(1., 0)},
		//"perpendicular":  {got: start.Perpendicular(), want: vector2.New(-2.4, -1.2)},
		//"normalized":     {got: start.Normalized(), want: vector2.New(0.447213, -.894427)},
		"scale":          {got: start.Scale(2.), want: rect2.New(vector2.New(1.2, -2.4), vector2.New(21.6, 24.8))},
		"scale f":        {got: start.ScaleF(2.), want: rect2.New(vector2.New(1.2, -2.4), vector2.New(21.6, 24.8))},
		"scale by vec":   {got: start.ScaleByVector(vector2.New(2., 4.)), want: rect2.New(vector2.New(1.2, -2.4), vector2.New(21.6, 49.6))},
		"scale by vec f": {got: start.ScaleByVectorF(vector2.New[float32](2., 4.)), want: rect2.New(vector2.New(1.2, -2.4), vector2.New(21.6, 49.6))},
		"scale by xy":    {got: start.ScaleByXY(2., 4.), want: rect2.New(vector2.New(1.2, -2.4), vector2.New(21.6, 49.6))},
		"scale by xy f":  {got: start.ScaleByXYF(2., 4.), want: rect2.New(vector2.New(1.2, -2.4), vector2.New(21.6, 49.6))},
		"zoom":           {got: start.Zoom(2.), want: rect2.New(vector2.New(2.4, -4.8), vector2.New(21.6, 24.8))},
		"zoom f":         {got: start.ZoomF(2.), want: rect2.New(vector2.New(2.4, -4.8), vector2.New(21.6, 24.8))},
		"zoom by vec":    {got: start.ZoomByVector(vector2.New(2., 4.)), want: rect2.New(vector2.New(2.4, -9.6), vector2.New(21.6, 49.6))},
		"zoom by vec f":  {got: start.ZoomByVectorF(vector2.New[float32](2., 4.)), want: rect2.New(vector2.New(2.4, -9.6), vector2.New(21.6, 49.6))},
		"zoom by xy":     {got: start.ZoomByXY(2., 4.), want: rect2.New(vector2.New(2.4, -9.6), vector2.New(21.6, 49.6))},
		"zoom by xy f":   {got: start.ZoomByXYF(2., 4.), want: rect2.New(vector2.New(2.4, -9.6), vector2.New(21.6, 49.6))},
		//"mult by vec":    {got: start.MultByVector(vector2.New(2., 4.)), want: vector2.New(2.4, -9.6)},
		//"div by vec":     {got: start.DivByVector(vector2.New(2., 4.)), want: vector2.New(0.6, -0.6)},
		//"center":         {got: vector2.Midpoint(start, vector2.New(2.4, 2.4)), want: vector2.New(1.8, 0.)},
		//"fill":           {got: vector2.Fill(9.3), want: vector2.New(9.3, 9.3)},
		//"yx":             {got: start.YX(), want: vector2.New(-2.4, 1.2)},
		//"random":         {got: vector2.Rand(r), want: vector2.New(.373028361, 0.066000496)},
		//"flip":           {got: start.Flip(), want: vector2.New(-1.2, 2.4)},
		//"flipX":          {got: start.FlipX(), want: vector2.New(-1.2, -2.4)},
		//"flipY":          {got: start.FlipY(), want: vector2.New(1.2, 2.4)},

		// Math package functions
		//"log":   {got: start.Log(), want: vector2.New(0.1823215, math.NaN())},
		//"log10": {got: start.Log10(), want: vector2.New(0.0791812, math.NaN())},
		//"log2":  {got: start.Log2(), want: vector2.New(0.263034, math.NaN())},
		//"exp":   {got: start.Exp(), want: vector2.New(3.320116, 0.090717)},
		//"exp2":  {got: start.Exp2(), want: vector2.New(2.297396, 0.189464)},
		//"expm1": {got: start.Expm1(), want: vector2.New(2.320116, -0.909282)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			test.AssertRectangle2InDelta(t, tc.want, tc.got, 0.00001)
		})
	}
}
