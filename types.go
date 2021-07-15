package renderer

type Tuple struct {
	X, Y, Z, W float64
}

func (t *Tuple) IsVector() bool {
	return t.W == 0
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1
}

func (t *Tuple) Equals(t2 *Tuple) bool {
	return Equal(t.X, t2.X) && Equal(t.Y, t2.Y) && Equal(t.Z, t2.Z) && Equal(t.W, t2.W)
}

func (t *Tuple) Add(t2 *Tuple) Tuple {

	r := Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}

	if r.W > 1 {
		r.W = 1
	}

	return r
}

func (t *Tuple) Sub(t2 *Tuple) Tuple {

	r := Tuple{
		X: t.X - t2.X,
		Y: t.Y - t2.Y,
		Z: t.Z - t2.Z,
		W: t.W - t2.W,
	}

	if t.W == 1 && t2.W == 1 {
		r.W = 1
	} else if r.W < 0 {
		r.W = 1
	}

	return r
}

func (t *Tuple) Neg() Tuple {
	return Tuple{
		X: 0 - t.X,
		Y: 0 - t.Y,
		Z: 0 - t.Z,
		W: 0 - t.W,
	}
}

func (t *Tuple) Multi(m float64) Tuple {
	return Tuple{
		X: t.X * m,
		Y: t.Y * m,
		Z: t.Z * m,
		W: t.W * m,
	}
}

func NewTuple(x, y, z, w float64) *Tuple {
	t := new(Tuple)

	t.X = x
	t.Y = y
	t.Z = z
	t.W = w

	return t
}

type Point = Tuple
type Vector = Tuple

func NewPoint(x, y, z float64) *Point {
	return NewTuple(x, y, z, 1)
}

func NewVector(x, y, z float64) *Vector {
	return NewTuple(x, y, z, 0)
}
