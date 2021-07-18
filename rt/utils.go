package rt

import "math"

const SMALL_NUMBER_F64 float64 = 0.000001

func Equal(a float64, b float64) bool {
	return math.Abs(a-b) < SMALL_NUMBER_F64
}
