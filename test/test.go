package test

import (
	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/rect2"
	"github.com/EliCDavis/vector/vector2"
	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func AssertVector2InDelta[T vector.Number](t assert.TestingT, expected, actual vector2.Vector[T], delta float64) {
	assert.InDelta(t, expected.X(), actual.X(), delta)
	assert.InDelta(t, expected.Y(), actual.Y(), delta)
}

func AssertVector3InDelta[T vector.Number](t assert.TestingT, expected, actual vector3.Vector[T], delta float64) {
	assert.InDelta(t, expected.X(), actual.X(), delta)
	assert.InDelta(t, expected.Y(), actual.Y(), delta)
	assert.InDelta(t, expected.Z(), actual.Z(), delta)
}

func AssertRectangle2InDelta[T vector.Number](t assert.TestingT, expected, actual rect2.Rectangle[T], delta float64) {
	AssertVector2InDelta(t, expected.Position(), actual.Position(), delta)
	AssertVector2InDelta(t, expected.Size(), actual.Size(), delta)
}
