package systems

import (
	"gamedev/ecs/types"
	"gamedev/ecs/utils"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SeperationSystem struct {
	multiplyer float32
}

func (cs *SeperationSystem) Update(dt float32, bs []*types.Boid) {
	const perception float32 = 24

	cs.multiplyer = rg.Slider(rl.NewRectangle(10, 60, 100, 20), "", "Seperation", cs.multiplyer, 0, 2)

	for _, b := range bs {
		total := 0
		steering := rl.Vector2{}

		for _, other := range bs {
			if other == b {
				continue
			}

			dist := rl.Vector2Distance(b.Postion, other.Postion) - 8
			if dist < perception {
				diff := rl.Vector2Subtract(b.Postion, other.Postion)
				diff = utils.DivideVectorByScalar(diff, dist*dist)
				steering = rl.Vector2Add(steering, diff)
				total++
			}
		}

		if total > 0 {
			steering = utils.DivideVectorByScalar(steering, float32(total))
			steering = utils.SetMag(steering, b.MaxSpeed)
			steering = rl.Vector2Subtract(steering, b.Velocity)
			steering = utils.LimitVector(steering, b.MaxForce)
		}

		steering = rl.Vector2Scale(steering, cs.multiplyer)
		b.Acceleration = rl.Vector2Add(b.Acceleration, steering)
	}
}
