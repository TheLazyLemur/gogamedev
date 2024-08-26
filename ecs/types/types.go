package types

import (
	"gamedev/ecs/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Boid struct {
	Postion      rl.Vector2
	Velocity     rl.Vector2
	Acceleration rl.Vector2
	MaxForce     float32
	MaxSpeed     float32
}

func NewBoid() *Boid {
	vel := utils.RandomVec2D()
	vel = utils.SetMag(vel, utils.RandomInRange(2, 4))

	return &Boid{
		Postion: rl.Vector2{
			X: utils.Random(float32(rl.GetScreenWidth())),
			Y: utils.Random(float32(rl.GetScreenHeight())),
		},
		Velocity:     vel,
		Acceleration: rl.Vector2{},
		MaxForce:     0.2,
		MaxSpeed:     5,
	}
}

type System interface {
	Update(dt float32, bs []*Boid)
}

type World struct {
	systems []System
	boids   []*Boid
}

func NewWorld() *World {
	return &World{
		systems: make([]System, 0),
		boids:   make([]*Boid, 0),
	}
}

func (w *World) AddSystems(systems ...System) {
	if w.systems == nil {
		w.systems = make([]System, 0)
	}

	w.systems = append(w.systems, systems...)
}

func (w *World) AddBoid(b *Boid) {
	if w.boids == nil {
		w.boids = make([]*Boid, 0)
	}

	w.boids = append(w.boids, b)
}

func (w *World) Update(dt float32) {
	for _, s := range w.systems {
		s.Update(dt, w.boids)
	}
}
