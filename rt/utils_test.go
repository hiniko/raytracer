package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestF64ToStr_RGB255t(t *testing.T) {

	r1 := F64ToStr_RGB255(1)
	r2 := F64ToStr_RGB255(0.5)
	r3 := F64ToStr_RGB255(0)
	r4 := F64ToStr_RGB255(1.5)
	r5 := F64ToStr_RGB255(-1.5)

	assert.Equal(t, "255", r1, "Failed to convert float to 8 bit int string")
	assert.Equal(t, "128", r2, "Failed to convert float to 8 bit int string")
	assert.Equal(t, "0", r3, "Failed to convert float to 8 bit int string")
	assert.Equal(t, "255", r4, "Failed to convert float to 8 bit int string")
	assert.Equal(t, "0", r5, "Failed to convert float to 8 bit int string")
}

func TestWriteFile(t *testing.T) {
	t.Skip()
	content := "This is a file"
	err := WriteFile("test_file.txt", content)
	assert.NoError(t, err, "Failed to create and write test file")
}
