package systems

import (
	"gamedev/ecs/types"
	"gamedev/ecs/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BoidUpdateSystem struct {
}

func (us *BoidUpdateSystem) Update(dt float32, bs []*types.Boid) {
	for _, b := range bs {
		b.Postion = rl.Vector2Add(b.Postion, b.Velocity)
		b.Velocity = rl.Vector2Add(b.Velocity, b.Acceleration)
		b.Velocity = utils.LimitVector(b.Velocity, b.MaxSpeed)
		b.Acceleration = utils.ZeroVector(b.Acceleration)

		// Edges
		if b.Postion.X > float32(rl.GetScreenWidth()) {
			b.Postion.X = 0
		} else if b.Postion.X < 0 {
			b.Postion.X = float32(rl.GetScreenWidth())
		}

		if b.Postion.Y > float32(rl.GetScreenWidth()) {
			b.Postion.Y = 0
		} else if b.Postion.Y < 0 {
			b.Postion.Y = float32(rl.GetScreenHeight())
		}

		// Draw
		rl.DrawCircleV(b.Postion, 5, rl.Blue)
	}
}
