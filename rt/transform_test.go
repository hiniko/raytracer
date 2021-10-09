package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Scenario: Multiplying by a translation matrix
//   Given transform ← translation(5, -3, 2)
//   And p ← point(-3, 4, 5)
//   Then transform * p = point(2, 1, 7)

func TestTransformTranslation(t *testing.T) {

	t1 := NewTransform().Translate(5, -3, 2)
	p1 := NewPoint(-3, 4, 5)

	pr := t1.TMulti(p1)
	pe := NewPoint(2, 1, 7)

	assert.True(t, pe.Equals(pr))
}

// Scenario: Multiplying by the inverse of a translation matrix
// Given transform ← translation(5, -3, 2)
// And inv ← inverse(transform)
// And p ← point(-3, 4, 5)
// Then inv * p = point(-8, 7, 3)

func TestTransformTranslationInverse(t *testing.T) {

	t1 := NewTransform().Translate(5, -3, 2)
	ti := t1.Invert()
	p1 := NewPoint(-3, 4, 5)

	pr := ti.TMulti(p1)
	pe := NewPoint(-8, 7, 3)

	assert.True(t, pe.Equals(pr))
}

// Scenario: Translation does not affect vectors
// Given transform ← translation(5, -3, 2)
// And v ← vector(-3, 4, 5) Then transform * v = v
func TestTranformTranslationNotAffectsVectors(t *testing.T) {
	t1 := NewTransform().Translate(2, 3, 4)
	v1 := NewVector(-3, 4, 5)

	ve := NewVector(-3, 4, 5)
	vr := t1.TMulti(ve)

	assert.True(t, v1.Equals(vr))
}

// Scenario: A scaling matrix applied to a point
// Given transform ← scaling(2, 3, 4)
// And p ← point(-4, 6, 8)
// Then transform * p = point(-8, 18, 32)
func TestTransformScaling(t *testing.T) {
	t1 := NewTransform().Scale(2, 3, 4)
	p1 := NewPoint(-4, 6, 8)

	pe := NewPoint(-8, 18, 32)
	pr := t1.TMulti(p1)

	assert.True(t, pe.Equals(pr))
}

// Scenario: A scaling matrix applied to a vector
// Given transform ← scaling(2, 3, 4)
// And v ← vector(-4, 6, 8)
// Then transform * v = vector(-8, 18, 32)
func TestTransformScalingAffectsVectors(t *testing.T) {

	t1 := NewTransform().Scale(2, 3, 4)
	v1 := NewVector(-4, 6, 8)

	ve := NewVector(-8, 18, 32)
	vr := t1.TMulti(v1)

	assert.True(t, ve.Equals(vr))
}

// Scenario: Multiplying by the inverse of a scaling matrix
// Given transform ← scaling(2, 3, 4)
// And inv ← inverse(transform)
// And v ← vector(-4, 6, 8)
// Then inv * v = vector(-2, 2, 2)
func TestTransformScalingSmaller(t *testing.T) {

	t1 := NewTransform().Scale(2, 3, 4).Invert()
	v1 := NewVector(-4, 6, 8)

	ve := NewVector(-2, 2, 2)
	vr := t1.TMulti(v1)

	assert.True(t, ve.Equals(vr))
}

// Scenario: Reflection is scaling by a negative value
// Given transform ← scaling(-1, 1, 1)
// And p ← point(2, 3, 4)
// Then transform * p = point(-2, 3, 4)
func TestTransformScalingRefelction(t *testing.T) {

	t1 := NewTransform().Scale(-1, 1, 1)
	p1 := NewPoint(2, 3, 4)

	pe := NewPoint(-2, 3, 4)
	pr := t1.TMulti(p1)

	assert.True(t, pe.Equals(pr))
}

// Scenario: Rotating a point around the x axis Given p ← point(0, 1, 0)
// And half_quarter ← rotation_x(π / 4)
// And full_quarter ← rotation_x(π / 2)
// Then half_quarter * p = point(0, √2/2, √2/2)
// And full_quarter * p = point(0, 0, 1)
func TestTransformRotationX(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	half := NewTransform().RotateX(math.Pi / 4)
	full := NewTransform().RotateX(math.Pi / 2)

	p1_half := half.TMulti(p1)
	p1_full := full.TMulti(p1)

	sq2 := math.Sqrt(2) / 2

	p1_half_e := NewPoint(0, sq2, sq2)
	p1_full_e := NewPoint(0, 0, 1)

	assert.True(t, p1_half.Equals(p1_half_e))
	assert.True(t, p1_full.Equals(p1_full_e))

}

// Scenario: Rotating a point around the y axis Given p ← point(0, 0, 1)
// And half_quarter ← rotation_y(π / 4)
// And full_quarter ← rotation_y(π / 2)
// Then half_quarter * p = point(√2/2, 0, √2/2)
// And full_quarter * p = point(1, 0, 0)
func TestTranformRotationY(t *testing.T) {
	p1 := NewPoint(0, 0, 1)
	half := NewTransform().RotateY(math.Pi / 4)
	full := NewTransform().RotateY(math.Pi / 2)

	p1_half := half.TMulti(p1)
	p1_full := full.TMulti(p1)

	sq2 := math.Sqrt(2) / 2

	p1_half_e := NewPoint(sq2, 0, sq2)
	p1_full_e := NewPoint(1, 0, 0)

	assert.True(t, p1_half.Equals(p1_half_e))
	assert.True(t, p1_full.Equals(p1_full_e))
}

// Scenario: Rotating a point around the z axis Given p ← point(0, 1, 0)
// And half_quarter ← rotation_z(π / 4)
// And full_quarter ← rotation_z(π / 2)
// Then half_quarter * p = point(-√2/2, √2/2, 0)
// And full_quarter * p = point(-1, 0, 0)
func TestTranformRotationZ(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	half := NewTransform().RotateZ(math.Pi / 4)
	full := NewTransform().RotateZ(math.Pi / 2)

	p1_half := half.TMulti(p1)
	p1_full := full.TMulti(p1)

	sq2 := math.Sqrt(2) / 2

	p1_half_e := NewPoint(-sq2, sq2, 0)
	p1_full_e := NewPoint(-1, 0, 0)

	assert.True(t, p1_half.Equals(p1_half_e))
	assert.True(t, p1_full.Equals(p1_full_e))
}

// Given p ← point(1, 0, 1)
// And A ← rotation_x(π / 2)
// And B ← scaling(5, 5, 5)
// And C ← translation(10, 5, 7)
// When T ← C * B * A
// Then T * p = point(15, 0, 7)
func TestTransformChaining(t *testing.T) {

	p1 := NewPoint(1, 0, 1)
	// Note the reverse operations here C * B * A
	t1 := NewTransform().Translate(10, 5, 7).Scale(5, 5, 5).RotateX(math.Pi / 2)

	pe := NewPoint(15, 0, 7)
	pr := t1.TMulti(p1)

	assert.True(t, pr.Equals(pe))
}

// Scenario: A shearing transformation moves x in proportion to y
// Given transform ← shearing(1, 0, 0, 0, 0, 0)
// And p ← point(2, 3, 4)
// Then transform * p = point(5, 3, 4)
func TestTransformSheer(t *testing.T) {

}
