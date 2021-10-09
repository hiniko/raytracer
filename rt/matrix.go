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
	Trans() Matrix
	Deter() float64
	SubMat(row, col int) Matrix
	Minor(row, col int) float64
	Cofactor(row, col int) float64
	IsInvertable() bool
	Invert() Matrix
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

func MatrixDeter(m Matrix) float64 {

	r, c := m.Dims()

	// 2x2 edge case
	if r == 2 && c == 2 {
		return (m.At(0, 0) * m.At(1, 1)) - (m.At(0, 1) * m.At(1, 0))
	}

	var d float64
	for ci := 0; ci < c; ci++ {
		d += m.At(0, ci) * m.Cofactor(0, ci)
	}
	return d
}

func MatrixCheckBounds(row, col int, m Matrix) {

	mr, mc := m.Dims()
	if row < 0 || col < 0 || row > mr || col > mc {
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

// IMulti is an in place multiplier
func (a *Matrix4) IMulti(b Matrix) {
	MatrixDimsCheck(a, b)

	ar, ac := a.Dims()

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			a.values[ar*r+c] =
				a.At(r, 0)*b.At(0, c) +
					a.At(r, 1)*b.At(1, c) +
					a.At(r, 2)*b.At(2, c) +
					a.At(r, 3)*b.At(3, c)
		}
	}
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
	ar, _ := a.Dims()

	rd := make([]float64, 3)

	for t := 0; t < ar; t++ {
		rd[t] =
			a.At(t, 0)*b.X +
				a.At(t, 1)*b.Y +
				a.At(t, 2)*b.Z
	}

	return NewTuple(rd[0], rd[1], rd[2], 1)
}

func (a *Matrix2) TMulti(b *Tuple) *Tuple {
	ar, _ := a.Dims()

	rd := make([]float64, 2)

	for t := 0; t < ar; t++ {
		rd[t] =
			a.At(t, 0)*b.X +
				a.At(t, 1)*b.Y
	}

	return NewTuple(rd[0], rd[1], 0, 1)
}

func (a *Matrix4) Trans() Matrix {
	ar, ac := a.Dims()

	rd := make([]float64, ar*ac)

	o := 0
	for r := 0; r < ar; r++ {
		rd[(0*ar)+o] = a.At(r, 0)
		rd[(1*ar)+o] = a.At(r, 1)
		rd[(2*ar)+o] = a.At(r, 2)
		rd[(3*ar)+o] = a.At(r, 3)
		o++
	}

	return NewMatrix4(rd)
}

func (a *Matrix3) Trans() Matrix {
	ar, ac := a.Dims()

	rd := make([]float64, ar*ac)

	o := 0
	for r := 0; r < ar; r++ {
		rd[(0*ar)+o] = a.At(r, 0)
		rd[(1*ar)+o] = a.At(r, 1)
		rd[(2*ar)+o] = a.At(r, 2)
		o++
	}

	return NewMatrix3(rd)
}

func (a *Matrix2) Trans() Matrix {
	ar, ac := a.Dims()

	rd := make([]float64, ar*ac)

	o := 0
	for r := 0; r < ar; r++ {
		rd[(0*ar)+o] = a.At(r, 0)
		rd[(1*ar)+o] = a.At(r, 1)
		o++
	}

	return NewMatrix2(rd)
}

// Row and col here refer to which row / col to delete
func (a *Matrix3) SubMat(row, col int) Matrix {
	ar, ac := a.Dims()

	if row > ar-1 || col > ac-1 {
		panic(fmt.Sprintf("submatrix out of bounds, mat dims %d, %d got %d, %d",
			ar, ac, row, col))
	}

	rd := make([]float64, 4)

	i := 0
	for r := 0; r < ar; r++ {
		if r == row {
			continue
		}
		for c := 0; c < ac; c++ {
			if c == col {
				continue
			}
			rd[i] = a.At(r, c)
			i++
		}
	}

	return NewMatrix2(rd)
}

// Row and col here refer to which row / col to delete
func (a *Matrix4) SubMat(row, col int) Matrix {
	ar, ac := a.Dims()

	if row > ar-1 || col > ac-1 {
		panic(fmt.Sprintf("submatrix out of bounds, mat dims %d, %d got %d, %d",
			ar, ac, row, col))
	}

	rd := make([]float64, 9)

	i := 0
	for r := 0; r < ar; r++ {
		if r == row {
			continue
		}
		for c := 0; c < ac; c++ {
			if c == col {
				continue
			}
			rd[i] = a.At(r, c)
			i++
		}
	}
	return NewMatrix3(rd)
}

func (a *Matrix3) Cofactor(row, col int) float64 {
	if (row+col)%2 > 0 {
		return -a.Minor(row, col)
	} else {
		return a.Minor(row, col)
	}
}

func (a *Matrix4) Cofactor(row, col int) float64 {
	if (row+col)%2 > 0 {
		return -a.Minor(row, col)
	} else {
		return a.Minor(row, col)
	}
}

func (a *Matrix3) Minor(row, col int) float64 {
	return a.SubMat(row, col).Deter()
}

func (a *Matrix4) Minor(row, col int) float64 {
	return a.SubMat(row, col).Deter()
}

func (a *Matrix2) Deter() float64 {
	return MatrixDeter(a)
}

func (a *Matrix3) Deter() float64 {
	return MatrixDeter(a)
}

func (a *Matrix4) Deter() float64 {
	return MatrixDeter(a)
}

func (a *Matrix2) IsInvertable() bool {
	return !Equal(a.Deter(), 0)
}

func (a *Matrix3) IsInvertable() bool {
	return !Equal(a.Deter(), 0)
}

func (a *Matrix4) IsInvertable() bool {
	return !Equal(a.Deter(), 0)
}

func (a *Matrix4) Invert() Matrix {

	if !a.IsInvertable() {
		panic(fmt.Sprintf("Matrix is not invertable: deter = %f", a.Deter()))
	}

	r, c := a.Dims()
	rd := make([]float64, r*c)
	d := a.Deter()

	for ri := 0; ri < r; ri++ {
		for ci := 0; ci < c; ci++ {
			co := a.Cofactor(ri, ci)
			rd[ci*r+ri] = co / d
		}
	}

	return NewMatrix4(rd)
}

// New Matrix helpers
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

// M4 Identity
var m4i = NewMatrix4([]float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
})
