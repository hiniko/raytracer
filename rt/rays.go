package rt

type Ray struct {
	Origin    *Point
	Direction *Vector
}

func NewRay(origin *Point, Direction *Vector) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: Direction,
	}
}
