package renderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFake(t *testing.T) {
	assert := assert.New(t)

	thing := true

	assert.Equal(true, thing)
}
