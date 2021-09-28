package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Scenario: Multiplying by a translation matrix
//   Given transform ← translation(5, -3, 2)
//   And p ← point(-3, 4, 5)
//   Then transform * p = point(2, 1, 7)

func TestTransformTranslation(t *testing.T) {

	t1 := NewTranslation(5, -3, 2)
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

	t1 := NewTranslation(5, -3, 2)
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
	t1 := NewTranslation(2, 3, 4)
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
	t1 := NewScaling(2, 3, 4)
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

	t1 := NewScaling(2, 3, 4)
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

	t1 := NewScaling(2, 3, 4).Invert()
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

	t1 := NewScaling(-1, 1, 1)
	p1 := NewPoint(2, 3, 4)

	pe := NewPoint(-2, 3, 4)
	pr := t1.TMulti(p1)

	assert.True(t, pe.Equals(pr))
}
