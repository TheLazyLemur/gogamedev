package systems

import (
	"fmt"
	"gamedev/ecs/types"
	"gamedev/ecs/utils"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AlignmentSystem struct {
	multiplyer float32
}

func (as *AlignmentSystem) Update(dt float32, bs []*types.Boid) {
	const perception float32 = 25

	as.multiplyer = rg.Slider(rl.NewRectangle(10, 0, 100, 20), "", "Alignment", as.multiplyer, 0, 2)

	for _, b := range bs {
		total := 0
		alignment := rl.Vector2{}

		for _, other := range bs {
			if other == b {
				continue
			}

			dist := rl.Vector2Distance(b.Postion, other.Postion) - 8
			if dist < perception {
				alignment = rl.Vector2Add(alignment, other.Velocity)
				total++
			}
		}

		if total > 0 {
			alignment = utils.DivideVectorByScalar(alignment, float32(total))
			alignment = utils.SetMag(alignment, b.MaxSpeed)
			alignment = rl.Vector2Subtract(alignment, b.Velocity)
			alignment = utils.LimitVector(alignment, b.MaxForce)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			as.multiplyer += 0.1
			fmt.Println("AlignmentSystem multiplyer")
		}

		alignment = rl.Vector2Scale(alignment, as.multiplyer)
		b.Acceleration = rl.Vector2Add(b.Acceleration, alignment)
	}
}
