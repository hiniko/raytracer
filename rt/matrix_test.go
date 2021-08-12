package rt

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

// Scenario: Constructing and inspecting a 4x4 matrix
// 	Given the following 4x4 matrix M:
// 		| 1 | 2 | 3 | 4 |
// 		| 5.5 | 6.5 | 7.5 | 8.5 |
// 		| 9 | 10 | 11 | 12 |
// 		| 13.5 | 14.5 | 15.5 | 16.5 |
// 	Then  M[0,0] = 1
// 		And M[0,3] = 4
// 		And M[1,0] = 5.5
// 		And M[1,2] = 7.5
// 		And M[2,2] = 11
// 		And M[3,0] = 13.5
// 		And M[3,2] = 15.5

func TestMatrix(t *testing.T) {

	m4 := NewMatrix4([]float64{
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	})

	assert.True(t, Equal(m4.At(0, 0), 1))
	assert.True(t, Equal(m4.At(0, 3), 4))
	assert.True(t, Equal(m4.At(1, 0), 5.5))
	assert.True(t, Equal(m4.At(1, 2), 7.5))
	assert.True(t, Equal(m4.At(2, 2), 11))
	assert.True(t, Equal(m4.At(3, 0), 13.5))
	assert.True(t, Equal(m4.At(3, 2), 15.5))

	m3 := NewMatrix3([]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	assert.True(t, Equal(m3.At(0, 0), 1))
	assert.True(t, Equal(m3.At(1, 1), 5))
	assert.True(t, Equal(m3.At(1, 2), 6))
	assert.True(t, Equal(m3.At(2, 0), 7))

	m2 := NewMatrix2([]float64{
		9, 8,
		7, 6,
	})

	assert.True(t, Equal(m2.At(0, 1), 8))
	assert.True(t, Equal(m2.At(1, 1), 6))

}

// Scenario: Matrix equality with identical matrices Given the following matrix A:
// |1|2|3|4|
// |5|6|7|8|
// |9|8|7|6|
// |5|4|3|2|
// And the following matrix B:
// |1|2|3|4|
// |5|6|7|8|
// |9|8|7|6|
// |5|4|3|2|
// Then A = B

// Scenario: Matrix equality with different matrices Given the following matrix A:
// |1|2|3|4|
// |5|6|7|8|
// |9|8|7|6|
// |5|4|3|2|
// And the following matrix B:
// |2|3|4|5|
// |6|7|8|9|
// |8|7|6|5|
// |4|3|2|1|
// Then A != B

func TestMatrixEquality(t *testing.T) {

	// Test that a matrix can equal itself
	m4 := NewMatrix4([]float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	assert.True(t, m4.Equal(m4))

	m4b := NewMatrix4([]float64{
		2, 3, 4, 5,
		6, 7, 8, 9,
		8, 7, 6, 5,
		4, 3, 2, 1,
	})

	assert.False(t, m4b.Equal(m4))
}

// Scenario: Multiplying two matrices Given the following matrix A:
// |1|2|3|4|
// |5|6|7|8|
// |9|8|7|6|
// |5|4|3|2|
// And the following matrix B:
// |-2|1|2| 3|
// | 3|2|1|-1|
// | 4|3|6| 5|
// | 1|2|7| 8|
// Then A * B is the following 4x4 matrix:
// |20| 22| 50| 48|
// |44| 54|114|108|
// |40| 58|110|102|
// |16| 26| 46| 42|
func TestMatrixMulti(t *testing.T) {

	m4a := NewMatrix4([]float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	m4b := NewMatrix4([]float64{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	})

	m4e := NewMatrix4([]float64{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	})

	m4r := m4a.Multi(m4b)

	spew.Dump(m4r)

	assert.True(t, m4r.Equal(m4e))
}

// Scenario: A matrix multiplied by a tuple Given the following matrix A:
// |1|2|3|4|
// |2|4|4|2|
// |8|6|4|1|
// |0|0|0|1|
// And b ← tuple(1, 2, 3, 1)
// Then A * b = tuple(18, 24, 33, 1)
func TestMatrixTupleMulti(t *testing.T) {

	m4 := NewMatrix4([]float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	})

	t1 := NewTuple(1, 2, 3, 1)

	te := NewTuple(18, 24, 33, 1)

	tr := m4.TMulti(t1)

	assert.True(t, tr.Equals(te))
}

// Scenario: Multiplying a matrix by the identity matrix
// Given the following matrix A:
// |0|1| 2| 4|
// |1|2| 4| 8|
// |2|4| 8|16|
// |4|8|16|32|
// Then A * identity_matrix = A
func TestMatrixIdentity(t *testing.T) {

	m4 := NewMatrix4([]float64{
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	})

	m4r := m4.Multi(m4i)
	assert.True(t, m4r.Equal(m4))
}

// Scenario: Multiplying the identity matrix by a tuple
// Given a ← tuple(1, 2, 3, 4)
// Then identity_matrix * a = a

func TestMatrixIdentityTuple(t *testing.T) {

	t1 := NewTuple(1, 2, 3, 4)

	tr := m4i.TMulti(t1)

	assert.True(t, tr.Equals(t1))
}
