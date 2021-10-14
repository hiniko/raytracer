package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Scenario: Creating and querying a ray
// Given origin ← point(1, 2, 3)
// And direction ← vector(4, 5, 6)
// When r ← ray(origin, direction) Then r.origin = origin
// And r.direction = direction

func TestRayNew(t *testing.T) {
	p1 := NewPoint(1, 2, 3)
	v1 := NewVector(4, 5, 6)
	r1 := NewRay(p1, v1)

	assert.True(t, p1.Equals(r1.Origin))
	assert.True(t, v1.Equals(r1.Direction))

}
