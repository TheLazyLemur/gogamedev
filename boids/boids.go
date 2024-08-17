package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func limitVector2(v rl.Vector2, max float32) rl.Vector2 {
	length := rl.Vector2Length(v)

	if length > max {
		scale := max / length
		return rl.Vector2Scale(v, scale)
	}

	return v
}

func setMagnitude(v rl.Vector2, newMagnitude float32) rl.Vector2 {
	magnitude := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))

	if magnitude == 0 {
		return rl.Vector2{X: 0, Y: 0}
	}

	return rl.Vector2{
		X: (v.X / magnitude) * newMagnitude,
		Y: (v.Y / magnitude) * newMagnitude,
	}
}

func main() {

	rl.InitWindow(640, 360, "Boids")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	flock := make([]*Boid, 100)
	for i := range 100 {
		flock[i] = NewBoid()
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		for i := range flock {
			b := flock[i]
			b.edges()
			b.flock(flock)
			b.update()
			b.show()
		}

		rl.EndDrawing()
	}
}
