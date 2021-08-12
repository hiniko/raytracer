package rt

import "fmt"

type Matrix interface {
	Dims() (r, c int)
	At(row, col int) float64
	Row(row int) []float64
	Col(row int) []float64
	Set(data []float64)
	Data() []float64
	Equal(m Matrix) bool
	Multi(m Matrix) Matrix
	TMulti(t *Tuple) *Tuple
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

func MatrixDimsCheck(a Matrix, b Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()

	if ar != br || ac != bc {
		panic("Cannot compare matricies with different dimensions")
	}
}

func MatrixEqual(a Matrix, b Matrix) bool {

	MatrixDimsCheck(a, b)

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

func MatrixGetRow(row int, m Matrix) []float64 {

	if row < 0 {
		panic("Cannot get a row less than 0")
	}

	r, c := m.Dims()

	if row >= r {
		panic(fmt.Sprintf("Tried to get a row greater than dims %d", r))
	}

	res := make([]float64, c)

	for ci := 0; ci < c; ci++ {
		res[ci] = m.At(row, ci)
	}

	return res
}

func MatrixGetCol(col int, m Matrix) []float64 {

	if col < 0 {
		panic("Cannot get a row less than 0")
	}

	r, c := m.Dims()

	if col >= c {
		panic(fmt.Sprintf("Tried to get a col greater than dims %d", c))
	}

	res := make([]float64, c)

	for ri := 0; ri < r; ri++ {
		res[ri] = m.At(ri, col)
	}

	return res
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

func (m *Matrix2) Row(row int) []float64 { return MatrixGetRow(row, m) }
func (m *Matrix3) Row(row int) []float64 { return MatrixGetRow(row, m) }
func (m *Matrix4) Row(row int) []float64 { return MatrixGetRow(row, m) }

func (m *Matrix2) Col(col int) []float64 { return MatrixGetCol(col, m) }
func (m *Matrix3) Col(col int) []float64 { return MatrixGetCol(col, m) }
func (m *Matrix4) Col(col int) []float64 { return MatrixGetCol(col, m) }

func (m *Matrix2) Set(values []float64) { MatrixCheckSet(values, m); m.values = values }
func (m *Matrix3) Set(values []float64) { MatrixCheckSet(values, m); m.values = values }
func (m *Matrix4) Set(values []float64) { MatrixCheckSet(values, m); m.values = values }

func (m *Matrix2) Equal(b Matrix) bool { return MatrixEqual(m, b) }
func (m *Matrix3) Equal(b Matrix) bool { return MatrixEqual(m, b) }
func (m *Matrix4) Equal(b Matrix) bool { return MatrixEqual(m, b) }

func (a *Matrix2) Multi(b Matrix) Matrix {
	MatrixDimsCheck(a, b)

	ar, ac := a.Dims()
	rd := make([]float64, ar*ac)

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			rd[ac*r+c] =
				a.At(r, 0)*b.At(0, c) +
					a.At(r, 1)*b.At(1, c)
		}
	}
	return NewMatrix2(rd)
}

func (a *Matrix3) Multi(b Matrix) Matrix {
	MatrixDimsCheck(a, b)

	ar, ac := a.Dims()
	rd := make([]float64, ar*ac)

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			rd[ac*r+c] =
				a.At(r, 0) + b.At(0, c) +
					a.At(r, 1)*b.At(1, c) +
					a.At(r, 2)*b.At(2, c)
		}
	}
	return NewMatrix3(rd)
}

func (a *Matrix4) Multi(b Matrix) Matrix {
	MatrixDimsCheck(a, b)

	ar, ac := a.Dims()
	rd := make([]float64, ar*ac)

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			rd[ar*r+c] =
				a.At(r, 0)*b.At(0, c) +
					a.At(r, 1)*b.At(1, c) +
					a.At(r, 2)*b.At(2, c) +
					a.At(r, 3)*b.At(3, c)
		}
	}
	return NewMatrix4(rd)
}

func (a *Matrix4) TMulti(b *Tuple) *Tuple {
	ar, _ := a.Dims()

	rd := make([]float64, 4)

	for t := 0; t < ar; t++ {
		rd[t] =
			a.At(t, 0)*b.X +
				a.At(t, 1)*b.Y +
				a.At(t, 2)*b.Z +
				a.At(t, 3)*b.W
	}

	return NewTuple(rd[0], rd[1], rd[2], rd[3])
}

func (a *Matrix3) TMulti(b *Tuple) *Tuple {
	return nil
}

func (a *Matrix2) TMulti(b *Tuple) *Tuple {
	return nil
}

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

// Useful matricies
var m4i = NewMatrix4([]float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
})
