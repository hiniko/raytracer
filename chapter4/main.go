package main

import (
	"fmt"
	"math"

	"github.com/hiniko/raytracer/rt"
)

func main() {

	canvas := rt.NewCanvas(150, 150)
	white := rt.NewColor(255, 255, 255, 1)
	red := rt.NewColor(255, 0, 0, 1)
	blue := rt.NewColor(0, 0, 255, 1)

	points := make([]*rt.Point, 50)
	for i := 0; i < len(points); i++ {
		points[i] = rt.NewPoint(float64(1*i), 0, 0)
	}

	// Draw the base line
	for _, p := range points {
		canvas.Set(int(p.X), int(p.Y), white)
	}

	ca_half_width := float64((canvas.Width / 2.0))
	ca_half_height := float64(canvas.Height / 2.0)

	// Draw after translation to center
	t1 := rt.NewTransform().Translate(ca_half_width, ca_half_height, 0)
	for _, p := range points {
		p2 := t1.TMulti(p)
		fmt.Printf("Trans: %#v\n", p2)
		canvas.Set(int(p2.X), int(p2.Y), red)
	}

	// Draw After rotation 90 degres
	for i := 1; i < 13; i++ {
		for _, p := range points {
			t2 := rt.NewTransform().Translate(ca_half_width, ca_half_height, 0).RotateZ((math.Pi * 2) / 12 * float64(i))
			p2 := t2.TMulti(p)
			// fmt.Printf("Rot: %d, %d - %f, %f\n", int(p2.X), int(p2.Y), p2.X, p2.Y)
			canvas.Set(int(p2.X), int(p2.Y), blue)
		}
	}

	canvas.ToPNG("ClockFace.png")
}
