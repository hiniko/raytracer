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

// Scenario: Transposing a matrix
// Given the following matrix A:
// |0|9|3|0|
// |9|8|0|8|
// |1|8|5|3|
// |0|0|5|8|
// Then transpose(A) is the following matrix:
// |0|9|1|0|
// |9|8|8|0|
// |3|0|5|5|
// |0|8|3|8|
func TestMatrixTranspost(t *testing.T) {
	m4a := NewMatrix4([]float64{
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	})

	m4e := NewMatrix4([]float64{
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	})

	m4r := m4a.Trans()
	assert.True(t, m4r.Equal(m4e))
}

// Scenario: Transposing the identity matrix
// Given A ← transpose(identity_matrix)
// Then A = identity_matrix
func TestMatrixIdentityTranspost(t *testing.T) {
	m4r := m4i.Trans()
	assert.True(t, m4r.Equal(m4i))
}

// Scenario: Calculating the determinant of a 2x2 matrix
// Given the following 2x2 matrix A:
// | 1|5|
// | -3 | 2 |
// Then determinant(A) = 17

func TestMatrix2Determinat(t *testing.T) {

	m2 := NewMatrix2([]float64{
		1, 5,
		-3, 2,
	})

	r := m2.Deter()

	assert.True(t, Equal(r, 17))
}

// Scenario: A submatrix of a 3x3 matrix is a 2x2 matrix
// Given the following 3x3 matrix A:
// | 1|5| 0|
// |-3|2| 7|
// | 0|6|-3|
// Then submatrix(A, 0, 2) is the following 2x2 matrix:
// | -3 | 2 |
// | 0|6|
func TestMatrix3Submatrix(t *testing.T) {

	m3 := NewMatrix3([]float64{
		1, 5, 0,
		-3, 2, 7,
		0, 6, -3,
	})

	m2e := NewMatrix2([]float64{
		-3, 2,
		0, 6,
	})

	m2r := m3.SubMat(0, 2)

	assert.True(t, m2r.Equal(m2e))
}

// Scenario: A submatrix of a 4x4 matrix is a 3x3 matrix
// Given the following 4x4 matrix A:
// |-6| 1| 1| 6|
// |-8| 5| 8| 6|
// |-1| 0| 8| 2|
// |-7| 1|-1| 1|
// Then submatrix(A, 2, 1) is the following 3x3 matrix:
// |-6| 1|6|
// |-8| 8|6|
// | -7 | -1 | 1 |
func TestMatrix4Submatrix(t *testing.T) {

	m4 := NewMatrix4([]float64{
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		-1, 0, 8, 2,
		-7, 1, -1, 1,
	})

	m3e := NewMatrix3([]float64{
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	})

	m3r := m4.SubMat(2, 1)

	assert.True(t, m3r.Equal(m3e))
}

// Scenario: Calculating a minor of a 3x3 matrix
// Given the following 3x3 matrix A:
// | 3| 5| 0|
// | 2|-1|-7|
// | 6|-1| 5|
// And B ← submatrix(A, 1, 0)
// 	Then determinant(B) = 25
// 	And minor(A, 1, 0) = 25

func TestMatrix3Minor(t *testing.T) {

	m3 := NewMatrix3([]float64{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	})

	m2d := m3.SubMat(1, 0).Deter()

	m3m := m3.Minor(1, 0)

	assert.True(t, Equal(m2d, m3m))
}

// Scenario: Calculating a cofactor of a 3x3 matrix Given the following 3x3 matrix A:
// | 3| 5| 0|
// | 2|-1|-7|
//|  6|-1| 5|
// Then minor(A, 0, 0) = -12
// And cofactor(A, 0, 0) = -12
// And minor(A, 1, 0) = 25
// And cofactor(A, 1, 0) = -25

func TestMatrix3Cofactor(t *testing.T) {
	m3 := NewMatrix3([]float64{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	})

	ma := m3.Minor(0, 0)
	ca := m3.Cofactor(0, 0)

	ma2 := m3.Minor(1, 0)
	ca2 := m3.Cofactor(1, 0)

	assert.True(t, Equal(ma, -12))
	assert.True(t, Equal(ca, -12))
	assert.True(t, Equal(ma2, 25))
	assert.True(t, Equal(ca2, -25))

}

// Scenario: Calculating the determinant of a 3x3 matrix
// Given the following 3x3 matrix A:
// | 1| 2| 6|
// |-5| 8| -4|
// | 2| 6| 4|
// Then cofactor(A, 0, 0) = 56
// And cofactor(A, 0, 1) = 12
// And cofactor(A, 0, 2) = -46
// And determinant(A) = -196
func TestMatrix3Deter(t *testing.T) {

	m3 := NewMatrix3([]float64{
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	})

	c1 := m3.Cofactor(0, 0)
	c2 := m3.Cofactor(0, 1)
	c3 := m3.Cofactor(0, 2)

	assert.True(t, Equal(56, c1))
	assert.True(t, Equal(12, c2))
	assert.True(t, Equal(-46, c3))

	d := m3.At(0, 0)*c1 +
		m3.At(0, 1)*c2 +
		m3.At(0, 2)*c3

	dr := m3.Deter()

	assert.True(t, Equal(-196, d))
	assert.True(t, Equal(dr, d))
}

// Scenario: Calculating the determinant of a 4x4 matrix
// Given the following 4x4 matrix A:
// |-2|-8| 3| 5|
// |-3| 1| 7| 3|
// | 1| 2|-9| 6|
// |-6| 7| 7|-9|
// Then cofactor(A, 0, 0) = 690
// And cofactor(A, 0, 1) = 447
// And cofactor(A, 0, 2) = 210
// And cofactor(A, 0, 3) = 51
// And determinant(A) = -4071
func TestMatrix4Deter(t *testing.T) {

	m4 := NewMatrix4([]float64{
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	})

	c1 := m4.Cofactor(0, 0)
	c2 := m4.Cofactor(0, 1)
	c3 := m4.Cofactor(0, 2)
	c4 := m4.Cofactor(0, 3)

	assert.True(t, Equal(690, c1))
	assert.True(t, Equal(447, c2))
	assert.True(t, Equal(210, c3))
	assert.True(t, Equal(51, c4))

	d := m4.At(0, 0)*c1 +
		m4.At(0, 1)*c2 +
		m4.At(0, 2)*c3 +
		m4.At(0, 3)*c4

	dr := m4.Deter()

	assert.True(t, Equal(-4071, d))
	assert.True(t, Equal(dr, d))
}

// Scenario: Testing an invertible matrix for invertibility
// Given the following 4x4 matrix A:
// | 6| 4| 4| 4|
// | 5| 5| 7| 6|
// | 4|-9| 3|-7|
// | 9| 1| 7|-6|
// Then determinant(A) = -2120 And A is invertible

// Scenario: Testing a noninvertible matrix for invertibility
// Given the following 4x4 matrix A:
// |-4| 2|-2|-3|
// | 9| 6| 2| 6|
// | 0|-5| 1|-5|
// | 0| 0| 0| 0|
// Then determinant(A) = 0 And A is not invertible
func TestMatrix4Invertable(t *testing.T) {

	m4a := NewMatrix4([]float64{
		6, 4, 4, 4,
		5, 5, 7, 6,
		4, -9, 3, -7,
		9, 1, 7, -6,
	})

	assert.True(t, m4a.IsInvertable())

	m4b := NewMatrix4([]float64{
		-4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0,
	})

	assert.False(t, m4b.IsInvertable())
}

// Scenario: Calculating the inverse of a matrix
// Given the following 4x4 matrix A:
// |-5| 2| 6|-8|
// | 1|-5| 1| 8|
// | 7| 7|-6|-7|
// | 1|-3| 7| 4|
// And B ← inverse(A)
// Then determinant(A) = 532
// And cofactor(A, 2, 3) = -160
// And B[3,2] = -160/532
// And cofactor(A, 3, 2) = 105
// And B[2,3] = 105/532
// And B is the following 4x4 matrix:
// |  0.21805 |  0.45113 |  0.24060 | -0.04511 |
// | -0.80827 | -1.45677 | -0.44361 |  0.52068 |
// | -0.07895 | -0.22368 | -0.05263 |  0.19737 |
// | -0.52256 | -0.81391 | -0.30075 |  0.30639 |
func TestMatrix4Inverse(t *testing.T) {

	m4 := NewMatrix4([]float64{
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	})

	m4i := m4.Invert()

	m4cod := make([]float64, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m4cod[i*4+j] = m4.Cofactor(i, j)
		}
	}

	m4cot := NewMatrix4(m4cod).Trans()

	m4d := make([]float64, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m4d[i*4+j] = m4cod[i*4+j] / 532.0
		}
	}

	// WHY ARE THESE DIFFERENT!
	spew.Dump(m4cod, m4cot, m4d, m4i)

	assert.True(t, Equal(m4.Deter(), 532), "Deter of A")
	assert.True(t, Equal(m4.Cofactor(2, 3), -160), "Cofactor of A[2,3]")
	assert.True(t, Equal(m4i.At(3, 2), -160/532), "B[3,2] is off")
	assert.True(t, Equal(m4.Cofactor(3, 2), 105), "Cofactor of B[3,2]")
	assert.True(t, Equal(m4i.At(2, 3), 105/532), "B[2,3] is off")

	m4ie := NewMatrix4([]float64{
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	})

	assert.True(t, m4i.Equal(m4ie))
}
