package renderer

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

// Scenario: A tuple with w=1.0 is a point
// 	Given a ← tuple(4.3, -4.2, 3.1, 1.0)
// 	Then a.x = 4.3
// 		And a.y = -4.2
// 		And a.z = 3.1
// 		And a.w = 1.0
// 		And a is a point
// 		And a is not a vector
func TestPointTuple(t *testing.T) {
	t1 := NewTuple(4.3, -4.2, 3.1, 1)

	assert.Equal(t, 4.3, t1.X, "X component is not equal")
	assert.Equal(t, -4.2, t1.Y, "Y component is not equal")
	assert.Equal(t, 3.1, t1.Z, "Z component is not equal")
	assert.Equal(t, 1.0, t1.W, "Y component is not equal")

	assert.True(t, t1.IsPoint(), "Tuple is not a Point")
	assert.False(t, t1.IsVector(), "Tuple should not be a Vector")
}

// Scenario: A Vector with w=0 is a vector
// 	Given a ← tuple(4.3, -4.2, 3.1, 0.0)
// 	Then a.x = 4.3
// 		And a.y = -4.2
// 		And a.z = 3.1
// 		And a.w = 0.0
// 		And a is not a point
// 		And a is a vector
func TestVectorTuple(t *testing.T) {

	t1 := NewTuple(4.3, -4.2, 3.1, 0.0)

	assert.Equal(t, 4.3, t1.X, "X component is not equal")
	assert.Equal(t, -4.2, t1.Y, "Y component is not equal")
	assert.Equal(t, 3.1, t1.Z, "Z component is not equal")
	assert.Equal(t, 0.0, t1.W, "W Component should be 0")

	assert.True(t, t1.IsVector(), "Tuple is not a Vector")
	assert.False(t, t1.IsPoint(), "Tuple should not be a Point")
}

// Test Type aliases
func TestPointType(t *testing.T) {
	p := NewPoint(1.2, 2.3, 3.4)

	assert.Equal(t, 1.2, p.X, "X component is not equal")
	assert.Equal(t, 2.3, p.Y, "Y component is not equal")
	assert.Equal(t, 3.4, p.Z, "Z component is not equal")
	assert.Equal(t, 1.0, p.W, "W Component should be 0")

	assert.True(t, p.IsPoint(), "Tuple is not a Point")
	assert.False(t, p.IsVector(), "Tuple should not be a Vector")
}

func TestVectorType(t *testing.T) {
	v := NewVector(1.2, 2.3, 3.4)

	assert.Equal(t, 1.2, v.X, "X component is not equal")
	assert.Equal(t, 2.3, v.Y, "Y component is not equal")
	assert.Equal(t, 3.4, v.Z, "Z comvponent is not equal")
	assert.Equal(t, 0.0, v.W, "W Component should be 0")

	assert.True(t, v.IsVector(), "Vector is not a Point")
	assert.False(t, v.IsPoint(), "Vector should not be a Vector")
}

// Test Tuple equations
func TestTupleEquals(t *testing.T) {
	t1 := NewTuple(1.1, -1.2, 1.3, 0)
	t2 := NewTuple(1.1, -1.2, 1.3, 0)

	assert.True(t, t1.Equals(t2), "Tuples should be equal up to Epslion (SMALL_NUMBER_F64)")

	t3 := NewTuple(1.0000001, -1.2, 1.3, 0)
	t4 := NewTuple(1.0000005, -1.2, 1.3, 0)

	assert.True(t, t3.Equals(t4), "Tuples should be equal up to Epslion (SMALL_NUMBER_F64)")

	t5 := NewTuple(-1.0000001, -1.2, 1.3, 0)
	t6 := NewTuple(-1.0000002, -1.2, 1.3, 0)

	assert.True(t, t5.Equals(t6), "Tuples should be equal up to Epslion (SMALL_NUMBER_F64)")

	t7 := NewTuple(1.00001, -1.2, 1.3, 0)
	t8 := NewTuple(1.00002, -1.2, 1.3, 0)

	assert.False(t, t7.Equals(t8), "Tuples should not be equal as they are bigger than Epsilion (SMALL_NUMBER_F64)")
}

// Scenario: Adding two tuples
// 	Given a1 ← tuple(3, -2, 5, 1)
// 		And a2 ← tuple(-2, 3, 1, 0)
// 	Then a1 + a2 = tuple(1, 1, 6, 1)

func TestTupleAdd(t *testing.T) {

	// Add A vector to a vector and get a vector
	v1 := NewVector(1.0, 2.0, 3.0)
	v2 := NewVector(2.0, 3.0, 4.0)

	vE := Vector{X: 3.0, Y: 5.0, Z: 7.0, W: 0}
	vR := v1.Add(v2)

	assert.True(t, Equal(vE.X, vR.X))
	assert.True(t, Equal(vE.Y, vR.Y))
	assert.True(t, Equal(vE.Z, vR.Z))
	assert.True(t, vR.IsVector())

	// Add a Point to a Point and get a point
	p1 := NewPoint(1.0, 2.0, 3.0)
	p2 := NewPoint(2.0, 3.0, 4.0)

	pE := Point{X: 3.0, Y: 5.0, Z: 7.0, W: 1}
	pR := p1.Add(p2)

	assert.True(t, Equal(pE.X, pR.X))
	assert.True(t, Equal(pE.Y, pR.Y))
	assert.True(t, Equal(pE.Z, pR.Z))
	assert.True(t, pR.IsPoint())

	// Add A point to a vector and get a point
	vpE := Point{X: 2.0, Y: 4.0, Z: 6.0, W: 1}
	vpR := v1.Add(p1)

	assert.True(t, Equal(vpE.X, vpR.X))
	assert.True(t, Equal(vpE.Y, vpR.Y))
	assert.True(t, Equal(vpE.Z, vpR.Z))
	assert.True(t, pR.IsPoint())
}

// Scenario: Subtracting two vectors
// Given v1 ← vector(3, 2, 1)
// 		And v2 ← vector(5, 6, 7)
// Then v1 - v2 = vector(-2, -4, -6)
func TestTupleVectorSub(t *testing.T) {

	// sub A vector to a vector and get a vector
	v1 := NewVector(1.0, 2.0, 3.0)
	v2 := NewVector(2.0, 3.0, 4.0)

	vE := Vector{X: -1.0, Y: -1.0, Z: -1.0, W: 0.0}
	vR := v1.Sub(v2)

	assert.True(t, Equal(vE.X, vR.X))
	assert.True(t, Equal(vE.Y, vR.Y))
	assert.True(t, Equal(vE.Z, vR.Z))
	assert.True(t, vR.IsVector(), "Should be a vector")
}

// Scenario: Subtracting two points
// 	Given p1 ← point(3, 2, 1)
// 		And p2 ← point(5, 6, 7)
// 	Then p1 - p2 = vector(-2, -4, -6)
func TestTuplePointSub(t *testing.T) {

	// Sub a Point to a Point and get a point
	p1 := NewPoint(1.0, 2.0, 3.0)
	p2 := NewPoint(2.0, 3.0, 4.0)

	pE := Point{X: -1.0, Y: -1.0, Z: -1.0, W: 1}
	pR := p1.Sub(p2)

	assert.True(t, Equal(pE.X, pR.X))
	assert.True(t, Equal(pE.Y, pR.Y))
	assert.True(t, Equal(pE.Z, pR.Z))
	assert.True(t, pR.IsPoint(), "Should be a point")

}

// Scenario: Subtracting a vector from a point
// 	Given p ← point(3, 2, 1)
// 		And v ← vector(5, 6, 7)
// 	Then p - v = point(-2, -4, -6)
func TestTupleVectorPointSub(t *testing.T) {

	v1 := NewVector(1.0, 2.0, 3.0)
	p1 := NewPoint(1.0, 2.0, 3.0)

	// Add A point to a vector and get a point
	vpE := Point{X: 0.0, Y: 0.0, Z: 0.0, W: 1}
	vpR := v1.Sub(p1)

	assert.True(t, Equal(vpE.X, vpR.X))
	assert.True(t, Equal(vpE.Y, vpR.Y))
	assert.True(t, Equal(vpE.Z, vpR.Z))
	assert.True(t, vpR.IsPoint(), "Should be a point")

}

// Scenario: Negating a tuple
// 	Given a ← tuple(1, -2, 3, -4)
func TestTupleManualNegation(t *testing.T) {
	t1 := NewTuple(1.0, 2.0, 3.0, 4.0)
	r := t1.Neg()
	e := NewTuple(-1.0, -2.0, -3.0, -4.0)

	assert.True(t, r.Equals(e), "Result is not negated")

	t2 := NewTuple(-1.0, -2.0, -3.0, -4.0)
	r2 := t2.Neg()
	e2 := NewTuple(1.0, 2.0, 3.0, 4.0)

	assert.True(t, r2.Equals(e2), "Result is not negated")
}

// Scenario: Multiplying a tuple by a scalar
// 	Given a ← tuple(1, -2, 3, -4)
// 	Then a * 3.5 = tuple(3.5, -7, 10.5, -14)
func TestTupleScalarMulti(t *testing.T) {

	t1 := NewTuple(1.0, 2.0, 3.0, 4.0)
	r1 := t1.Multi(2.0)

	e1 := NewTuple(1.0, 4.0, 6.0, 8.0)

	assert.True(t, r1.Equals(e1), "Scalar multiplication has gone wrong!")
}

// Scenario: Multiplying a tuple by a fraction
// 	Given a ← tuple(1, -2, 3, -4)
// 	Then a * 0.5 = tuple(0.5, -1, 1.5, -2)
func TestTupleFractionalMulti(t *testing.T) {

	t1 := NewTuple(1.0, 2.0, 3.0, 4.0)
	r1 := t1.Multi(0.5)

	spew.Dump(r1)
	e1 := NewTuple(0.5, 1, 1.5, 2)

	assert.True(t, r1.Equals(e1), "Frational multiplication has gone wrong!")
}
