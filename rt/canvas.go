package rt

import (
	"fmt"
	"strings"
)

const OUTPUT_DIR = "output/"
const PPM_MAX_CHARS = 69 // PPM Compat line width is 70, so we need to match 69 for the newline

type Canvas struct {
	Height, Width int
	Data          []*Color
}

func NewCanvas(width, height int) *Canvas {
	c := Canvas{
		Width:  width,
		Height: height,
		Data:   make([]*Color, width*height),
	}

	return &c
}

func (ca *Canvas) Set(x, y int, c *Color) {
	ca.Data[x+(y*ca.Width)] = c
}

func (ca *Canvas) Get(x, y int) *Color {
	return ca.Data[x+(y*ca.Width)]
}

func (ca *Canvas) ToPPM() string {

	var buf strings.Builder
	// Build PPM Header
	buf.WriteString("P3\n")
	buf.WriteString(fmt.Sprintf("%d %d\n", ca.Width, ca.Height))
	buf.WriteString("255\n")

	// Build PPM Data
	// clw := 0 // current line width. For compatability each line can only be 70 chars

	e := NewColor(0, 0, 0, 0).ToRGB255String() // Empty color for missing data

	for p := 0; p < len(ca.Data); p++ {

		d := ca.Data[p]

		var s string

		if d == nil {
			s = e
		} else {
			s = d.ToRGB255String()
		}

		if p > 0 && (p+1)%(ca.Width) == 0 {
			fmt.Printf("nl at %d", p)
			buf.WriteString(s)
			buf.WriteString("\n")
		} else {
			buf.WriteString(s + " ")
		}

		// nlw := clw + len(s) // New line width after we have written our colour data

		// if nlw >= PPM_MAX_CHARS { // -1 to allow for new line byte
		// 	buf.WriteString("\n")
		// 	buf.WriteString(s)
		// 	continue
		// }
	}

	return buf.String()
	// TODO: Maybe stick with the book and spit this shit out as strings first. Unsure of the agressive line breaking here after all

}
