package main

import (
	"gamedev/ecs/systems"
	"gamedev/ecs/types"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1920, 1080, "Boids")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	world := &types.World{}
	world.AddSystems(&systems.AlignmentSystem{}, &systems.SeperationSystem{}, &systems.CohesionSystem{}, &systems.BoidUpdateSystem{})

	for range 2000 {
		world.AddBoid(types.NewBoid())
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(250, 0)
		world.Update(rl.GetFrameTime())

		rl.EndDrawing()
	}
}
