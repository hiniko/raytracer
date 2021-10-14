package rt

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

const OUTPUT_DIR = "output/"
const PPM_MAX_CHARS = 70 // PPM Compat line width is 70, so we need to match 69 for the newline

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
	idx := x + (y * ca.Width)
	if idx < 0 || idx >= len(ca.Data) {
		return
	}
	ca.Data[idx] = c
}

func (ca *Canvas) Get(x, y int) *Color {
	idx := x + (y * ca.Width)
	if idx < 0 || idx > len(ca.Data) {
		return nil
	}
	return ca.Data[idx]
}

func (ca *Canvas) ToPNG(filename string) {

	f, err := os.Create(filename)

	if err != nil {
		fmt.Printf("Failed to create file %s - %s", filename, err.Error())
		return
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{ca.Width, ca.Height}})

	for p := 0; p < len(ca.Data); p++ {
		d := ca.Data[p]

		y := int(p / ca.Width)
		x := int(p % ca.Width)

		if ca.Data[p] != nil {
			img.SetRGBA(x, y, color.RGBA{uint8(d.X), uint8(d.Y), uint8(d.Z), 255})
		} else {
			img.SetRGBA(x, y, color.RGBA{0, 0, 0, 255})
		}

	}
	png.Encode(f, img)
}

// I am dumb and this could be better I'm sure.
func (ca *Canvas) ToPPM() string {

	var buf strings.Builder
	// Build PPM Header
	buf.WriteString("P3\n")
	buf.WriteString(fmt.Sprintf("%d %d\n", ca.Width, ca.Height))
	buf.WriteString("255\n")

	// Build PPM Data
	clw := 0 // current line width. For compatability each line can only be 70 chars

	e := NewColor(0, 0, 0, 0).ToRGB255String() // Empty color for missing data

	for p := 0; p < len(ca.Data); p++ {

		d := ca.Data[p]

		var s string

		if d == nil {
			s = e
		} else {
			s = d.ToRGB255String()
		}

		// Check for eol, that is are we at the end of a row
		eol := (p > 0 && (p+1)%(ca.Width) == 0)

		// Are we over PPM_MAX_CHARS
		if clw+len(s) > PPM_MAX_CHARS {
			pts := strings.Split(s, " ")
			for i, p := range pts {
				l := clw + len(p)

				var np string
				var npls = 0

				if i+1 < len(pts) {
					np = pts[i+1]
					npls = l + len(np) // next part length with space
				} else {
					np = ""
				}

				// If adding this part is equal or over max chars
				if l >= PPM_MAX_CHARS {
					buf.WriteString("\n")
					buf.WriteString(p + " ")
					clw = len(p) + 1

					// Look ahead and check if we can fix the next chunk
					// if not new line and print with a space
				} else if np != "" && npls == PPM_MAX_CHARS {
					buf.WriteString(p)
					buf.WriteString("\n")
					clw = 0
				} else if np != "" && npls > PPM_MAX_CHARS {
					buf.WriteString("\n")
					buf.WriteString(p + " ")
					clw = len(p) + 1

				} else {
					if eol {
						buf.WriteString(p)
						clw = 0
					} else {
						buf.WriteString(p + " ")
						clw += len(p) + 1
					}
				}
			}

			s = "" // set the input string to empty so we can handle eol
		}

		if eol && s == "" {
			buf.WriteString("\n")
			clw = 0
		} else if eol && s != "" {
			buf.WriteString(s)
			buf.WriteString("\n")
			clw = 0
		} else {
			if s != "" {
				buf.WriteString(s + " ")
				clw += len(s) + 1
			}
		}
	}

	return buf.String()
}
