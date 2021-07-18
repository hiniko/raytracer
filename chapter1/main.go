package main

import (
	"fmt"
	"os"
	"time"

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

	//spew.Dump(proj.Position)

	vel := proj.Velocity.Add(proj.Velocity)
	vel = vel.Add(env.Gravity)
	vel = vel.Add(env.Wind)

	proj.Velocity = vel.Norm()

	//spew.Dump(proj.Velocity)

}

func main() {
	fmt.Printf("Chapter 1 - Projectile Sim \n\n")

	env := Environment{
		Gravity: rt.NewVector(0, -1.0, 0),
		Wind:    rt.NewVector(0, 0, 0),
	}

	proj := Projectile{
		Position: rt.NewVector(0, 0.1, 0),
		Velocity: rt.NewVector(1, 1, 0).Norm(),
	}

	fmt.Printf("env: \n    Gravity: %s \n    Wind: %s \n\n", env.Gravity.ToString(), env.Wind.ToString())
	fmt.Printf("proj: \n    Pos : %s \n    vel: %s \n\n", proj.Position.ToString(), proj.Velocity.ToString())

	for proj.Position.Y > 0.0 {
		tick(&env, &proj)
		fmt.Printf("proj: %s \n", proj.Position.ToString())
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("Done!\n")

	os.Exit(0)
}
