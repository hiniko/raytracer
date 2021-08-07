package rt

type Matrix struct {
	Size int
	Data [][]float64
}

func NewMatrix(size int) *Matrix {
	m := new(Matrix)
	m.Data = make([][]float64, size)

	for i := range m.Data {
		m.Data[i] = make([]float64, size)
	}
	return m
}
