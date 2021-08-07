package main

import (
	"fmt"
	"os"

	rt "github.com/hiniko/raytracer/rt"
)

type Projectile struct {
	Position *rt.Point
	Velocity *rt.Vector
}

type Environment struct {
	Gravity *rt.Vector
	Wind    *rt.Vector
}

func tick(env *Environment, proj *Projectile) {
	proj.Position = proj.Position.Add(proj.Velocity)

	vel := proj.Velocity.Add(proj.Velocity)
	vel = vel.Add(env.Gravity)
	vel = vel.Add(env.Wind)

	proj.Velocity = vel.Norm()
}

func main() {
	fmt.Printf("Chapter 2 - Projectile Sim Image \n\n")

	env := Environment{
		Gravity: rt.NewVector(0, -0.01, 0),
		Wind:    rt.NewVector(0.01, 0, 0),
	}

	proj := Projectile{
		Position: rt.NewVector(0, 0.01, 0),
		Velocity: rt.NewVector(50, 150, 0),
	}

	ca := rt.NewCanvas(1000, 1000)
	c1 := rt.NewColor(255, 0, 0, 1)

	r := 255.0
	for proj.Position.Y > -float64(ca.Height) {
		r -= 0.1
		if r <= 0 {
			r = 255
		}

		tick(&env, &proj)
		ca.Set(int(proj.Position.X), ca.Height-int(proj.Position.Y), c1)
	}

	rt.WriteFile("chapter2-projsim.ppm", ca.ToPPM())

	fmt.Printf("\nDone!\n")

	os.Exit(0)
}
