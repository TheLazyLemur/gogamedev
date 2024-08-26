package systems

import (
	"gamedev/ecs/types"
	"gamedev/ecs/utils"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CohesionSystem struct {
	multiplyer float32
}

func (cs *CohesionSystem) Update(dt float32, bs []*types.Boid) {
	const perception float32 = 50

	cs.multiplyer = rg.Slider(rl.NewRectangle(10, 30, 100, 20), "", "Cohesion", cs.multiplyer, 0, 2)

	for _, b := range bs {
		total := 0
		cohesion := rl.Vector2{}

		for _, other := range bs {
			if other == b {
				continue
			}

			dist := rl.Vector2Distance(b.Postion, other.Postion) - 8
			if dist < perception {
				cohesion = rl.Vector2Add(cohesion, other.Postion)
				total++
			}
		}

		if total > 0 {
			cohesion = utils.DivideVectorByScalar(cohesion, float32(total))
			cohesion = rl.Vector2Subtract(cohesion, b.Postion)
			cohesion = utils.SetMag(cohesion, b.MaxSpeed)
			cohesion = rl.Vector2Subtract(cohesion, b.Velocity)
			cohesion = utils.LimitVector(cohesion, b.MaxForce)
		}

		cohesion = rl.Vector2Scale(cohesion, cs.multiplyer)
		b.Acceleration = rl.Vector2Add(b.Acceleration, cohesion)
	}
}
