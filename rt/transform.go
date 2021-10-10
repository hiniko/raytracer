package rt

import (
	"math"
)

type Transform = Matrix4

func NewTransform() *Transform {
	return NewMatrix4([]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

func (t *Transform) Translate(x, y, z float64) *Transform {
	t.IMulti(NewMatrix4([]float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}))
	return t
}

func (t *Transform) Scale(x, y, z float64) *Transform {
	t.IMulti(NewMatrix4([]float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}))
	return t
}

func (t *Transform) Sheer(xy, xz, yx, yz, zx, zy float64) *Transform {
	t.IMulti(NewMatrix4([]float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	}))

	return t
}

func (t *Transform) RotateX(rads float64) *Transform {
	t.IMulti(NewMatrix4([]float64{
		1, 0, 0, 0,
		0, math.Cos(rads), -math.Sin(rads), 0,
		0, math.Sin(rads), math.Cos(rads), 0,
		0, 0, 0, 1,
	}))
	return t
}

func (t *Transform) RotateY(rads float64) *Transform {
	t.IMulti(NewMatrix4([]float64{
		math.Cos(rads), 0, math.Sin(rads), 0,
		0, 1, 0, 0,
		-math.Sin(rads), 0, math.Cos(rads), 0,
		0, 0, 0, 1,
	}))
	return t
}

func (t *Transform) RotateZ(rads float64) *Transform {
	t.IMulti(NewMatrix4([]float64{
		math.Cos(rads), -math.Sin(rads), 0, 0,
		math.Sin(rads), math.Cos(rads), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}))
	return t
}
