package main

import (
	"log"
	"os"

	rt "github.com/hiniko/raytracer/rt"
)

type Projectile struct {
	Position rt.Point
	Velocity rt.Vector
}

type Environment struct {
	Gravity rt.Vector
	Wind    rt.Vector
}

func main() {
	log.Print("Chapter 1 - Projectile Sim")

	os.Exit(0)
}
