package rt

import "fmt"

type Matrix interface {
	Dims() (r, c int)
	At(row, col int) float64
	Set(data []float64)
	Data() []float64
	Equal(m Matrix)
}

func MatrixEqual(a Matrix, b Matrix) bool {
	ar, ac := a.Dims()
	br, bc := b.Dims()

	if ar != br || ac != bc {
		panic("Cannot compare matricies with different dimensions")
	}

	ad := a.Data()
	bd := b.Data()

	for i := 0; i < len(ad); i++ {
		if !Equal(ad[i], bd[i]) {
			return false
		}
	}

	return true
}

func MatrixCheckBounds(row, col int, m Matrix) {

	mr, mc := m.Dims()
	if row < 0 || col < 0 || row >= mr || col >= mc {
		panic(fmt.Sprintf("Invalid access to matrix data: Tried %d:%d, on %d:%d", row, col, mr, mc))
	}
}

func MatrixGet(r, c int, m Matrix) float64 {
	MatrixCheckBounds(r, c, m)
	mr, _ := m.Dims()
	return m.Data()[c+(r*mr)]
}

func MatrixCheckSet(v []float64, m Matrix) {

	c, r := m.Dims()
	l := c * r

	if v == nil {
		panic("mat data nil")
	}

	if len(v) != l || cap(v) != l {
		panic("mat data wrong shape")
	}
}

type Matrix2 struct {
	Matrix
	values []float64
}

type Matrix3 struct {
	Matrix
	values []float64
}

type Matrix4 struct {
	Matrix
	values []float64
}

func (m *Matrix2) Data() []float64 { return m.values }
func (m *Matrix3) Data() []float64 { return m.values }
func (m *Matrix4) Data() []float64 { return m.values }

func (m *Matrix2) Dims() (r, c int) { return 2, 2 }
func (m *Matrix3) Dims() (r, c int) { return 3, 3 }
func (m *Matrix4) Dims() (r, c int) { return 4, 4 }

func (m *Matrix2) At(row, col int) float64 { return MatrixGet(row, col, m) }
func (m *Matrix3) At(row, col int) float64 { return MatrixGet(row, col, m) }
func (m *Matrix4) At(row, col int) float64 { return MatrixGet(row, col, m) }

func (m *Matrix2) Set(values []float64) { MatrixCheckSet(values, m); m.values = values }
func (m *Matrix3) Set(values []float64) { MatrixCheckSet(values, m); m.values = values }
func (m *Matrix4) Set(values []float64) { MatrixCheckSet(values, m); m.values = values }

func NewMatrix2(values []float64) *Matrix2 {
	m := new(Matrix2)
	m.Set(values)
	return m
}

func NewMatrix3(values []float64) *Matrix3 {
	m := new(Matrix3)
	m.Set(values)
	return m
}

func NewMatrix4(values []float64) *Matrix4 {
	m := new(Matrix4)
	m.Set(values)
	return m
}
