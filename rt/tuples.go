package rt

import (
	"fmt"
	"math"
)

type Tuple struct {
	X, Y, Z, W float64
}

type Point = Tuple
type Vector = Tuple
type Color = Tuple // A color is basically a tuple right? Sure it doesn't say rgba but it's basically the same...

func (t *Tuple) IsVector() bool {
	return t.W == 0
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1
}

func (t *Tuple) Equals(t2 *Tuple) bool {
	return Equal(t.X, t2.X) && Equal(t.Y, t2.Y) && Equal(t.Z, t2.Z) && Equal(t.W, t2.W)
}

func (t *Tuple) Add(t2 *Tuple) *Tuple {

	r := Tuple{
		X: t.X + t2.X,
		Y: t.Y + t2.Y,
		Z: t.Z + t2.Z,
		W: t.W + t2.W,
	}

	if r.W > 1 {
		r.W = 1
	}

	return &r
}

func (t *Tuple) Sub(t2 *Tuple) *Tuple {

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

	return &r
}

func (t *Tuple) Neg() Tuple {
	return Tuple{
		X: 0 - t.X,
		Y: 0 - t.Y,
		Z: 0 - t.Z,
		W: 0 - t.W,
	}
}

func (t *Tuple) Multi(m float64) *Tuple {
	return &Tuple{
		X: t.X * m,
		Y: t.Y * m,
		Z: t.Z * m,
		W: t.W * m,
	}
}

func (t *Vector) Mag() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t *Vector) Norm() *Vector {
	mag := t.Mag()

	return &Vector{
		X: t.X / mag,
		Y: t.Y / mag,
		Z: t.Z / mag,
		W: t.W / mag,
	}
}

func (t *Vector) Dot(t2 *Vector) float64 {
	return t.X*t2.X + t.Y*t2.Y + t.Z*t2.Z + t.W*t2.W
}

func (t *Vector) Cross(t2 *Vector) *Vector {
	return &Vector{
		X: t.Y*t2.Z - t.Z*t2.Y,
		Y: t.Z*t2.X - t.X*t2.Z,
		Z: t.X*t2.Y - t.Y*t2.X,
		W: 0,
	}
}

// hadamard product for colors
func (c *Color) Prod(c2 *Color) Color {
	return Color{
		X: c.X * c2.X,
		Y: c.Y * c2.Y,
		Z: c.Z * c2.Z,
		W: c.W * c2.W,
	}
}

func (t *Tuple) ToRGB255String() string {
	return fmt.Sprintf("%s %s %s",
		F64ToStr_RGB255(t.X),
		F64ToStr_RGB255(t.Y),
		F64ToStr_RGB255(t.Z),
	)
}

func (t *Tuple) ToString() string {
	return fmt.Sprintf("X: %.4f, Y: %.4f, Z: %.4f, W: %.4f", t.X, t.Y, t.Z, t.W)
}

func NewTuple(x, y, z, w float64) *Tuple {
	t := new(Tuple)

	t.X = x
	t.Y = y
	t.Z = z
	t.W = w

	return t
}

func NewPoint(x, y, z float64) *Point {
	return NewTuple(x, y, z, 1)
}

func NewVector(x, y, z float64) *Vector {
	return NewTuple(x, y, z, 0)
}

func NewColor(r, g, b, a float64) *Color {
	return NewTuple(r, g, b, a)
}
