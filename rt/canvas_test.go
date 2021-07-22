package rt

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Scenario: Creating a canvas
// 	Given c ← canvas(10, 20)
// 	Then c.width = 10
// 		And c.height = 20
// 		And every pixel of c is color(0, 0, 0)

func TestCanvas(t *testing.T) {
	ca := NewCanvas(10, 20)
	assert.Equal(t, 10, ca.Width, "Canvas width is wrong size")
	assert.Equal(t, 20, ca.Height, "Canvas height s wrong size")
}

func TestCanvasReadWrite(t *testing.T) {
	ca := NewCanvas(10, 20)
	red := NewColor(1, 0, 0, 1)

	ca.Set(2, 3, red)
	r := ca.Get(2, 3)
	assert.True(t, red.Equals(r), "Failed to set and get color from canvas")

}

// Scenario: Constructing the PPM header
// 	Given c ← canvas(5, 3)
// 	When ppm ← canvas_to_ppm(c)
// 	Then lines 1-3 of ppm are
// 		"""
// 		P3
// 		5 3
// 		255
// 		"""
func TestCanvasPPMHeader(t *testing.T) {
	ca := NewCanvas(5, 3)

	ppm := ca.ToPPM()

	s := bufio.NewScanner(strings.NewReader(ppm))
	line := 0
	for s.Scan() {
		line++
		switch line {
		case 1:
			assert.Equal(t, "P3", s.Text())
		case 2:
			assert.Equal(t, "5 3", s.Text())
		case 3:
			assert.Equal(t, "255", s.Text())
		}
	}
}

// Scenario: Constructing the PPM pixel data
//	 Given c ← canvas(5, 3)
//		 And c1 ← color(1.5, 0, 0)
//		 And c2 ← color(0, 0.5, 0)
//		 And c3 ← color(-0.5, 0, 1)
//	 When write_pixel(c, 0, 0, c1)
//	   And write_pixel(c, 2, 1, c2)
//	   And write_pixel(c, 4, 2, c3)
//	   And ppm ← canvas_to_ppm(c)
//	 Then lines 4-6 of ppm are
//	    """
//	    255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
//	    0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
//	    0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
//	    """

func TestCanvasPPMNewLineAfterRow(t *testing.T) {
	ca := NewCanvas(5, 3)

	c1 := NewColor(1.5, 0, 0, 0)
	c2 := NewColor(0, 0.5, 0, 0)
	c3 := NewColor(-0.5, 0, 1, 0)

	ca.Set(0, 0, c1)
	ca.Set(2, 1, c2)
	ca.Set(4, 2, c3)

	ppm := ca.ToPPM()

	fmt.Printf("%s", ppm)

	s := bufio.NewScanner(strings.NewReader(ppm))
	line := 0
	for s.Scan() {
		line++
		switch line {

		case 4:
			assert.Equal(t, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", s.Text())
		case 5:
			assert.Equal(t, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", s.Text())
		case 6:
			assert.Equal(t, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", s.Text())

		}
	}
}

// Scenario: Splitting long lines in PPM files
// 	Given c ← canvas(10, 2)
// 	When every pixel of c is set to color(1, 0.8, 0.6)
// 		And ppm ← canvas_to_ppm(c)
// 	Then lines 4-7 of ppm are
// 		"""
// 		255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
// 		153 255 204 153 255 204 153 255 204 153 255 204 153
// 		255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
// 		153 255 20

func TestCanvasPPMLineWidth70Chars(t *testing.T) {

	ca := NewCanvas(10, 20)

	// Create an cavas with data
	colors := make([]*Color, 5)
	colors[0] = NewColor(1, 0, 0, 1)
	colors[1] = NewColor(0, 1, 0, 1)
	colors[2] = NewColor(0, 0, 1, 1)
	colors[3] = NewColor(0, 0, 0, 1)
	colors[4] = NewColor(1, 1, 1, 1)

	cur, p := 0, 0
	for y := 0; y < ca.Height; y++ {
		for x := 0; x < ca.Width; x++ {
			if p%10 == 0 {
				cur = ((cur + 1) % 5)
			}

			ca.Set(x, y, colors[cur])
			p++
		}
	}

	// Convert to PPM
	output := ca.ToPPM()

	// Read PPM String to ensure it conforms
	s := bufio.NewScanner(strings.NewReader(output))
	line := 0
	for s.Scan() {
		line++

		switch line {
		case 1:
			assert.Equal(t, "P3", s.Text())
		case 2:
			assert.Equal(t, "10 20", s.Text())
		case 3:
			assert.Equal(t, "255", s.Text())
		default:
			assert.Equal(t, 70, len(s.Text()))
		}

	}

}
