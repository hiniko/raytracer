package rt

type Transform = Matrix4

func NewTranslation(x, y, z float64) *Transform {
	return NewMatrix4([]float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	})
}

func NewScaling(x, y, z float64) *Transform {
	return NewMatrix4([]float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
}
