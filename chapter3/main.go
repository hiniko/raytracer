package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/hiniko/raytracer/rt"
)

// 1. What happens when you invert the identity matrix?
// 2. What do you get when you multiply a matrix by its inverse?
// 3. Is there any difference between the inverse of the transpose of a matrix, and the transpose of the inverse?
// 4. Remember how multiplying the identity matrix by a tuple gives you the tuple, unchanged?
//    Now, try changing any single element of the identity matrix to a different number, and then multiplying it by a tuple.
//    What happens to the tuple?

func main() {
	fmt.Printf("Chapter 3 - Reflections on Matricies \n\n")

	var m4i = rt.NewMatrix4([]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})

	var m4i_inverted = m4i.Invert()

	spew.Dump(m4i, m4i_inverted)

	// Expectation  - The identity matrix is unchanged
	// Result - The 1's were unchanged however the sign of the 0's were inverted.
	//  				This does makes sense, given the operation to invert the signs in the determinant op
	//					I wonder if a better thing to do here is not to invert 0's?

	m4a := rt.NewMatrix4([]float64{
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	})

	m4ami := m4a.Multi(m4a.Invert())
	spew.Dump(m4a, m4ami)

	// Expectation -  I'm expecting the result to be the same as the inversion itself
	// Result -  Huh, I have no idea what this did, Need to have some research about this.

}
