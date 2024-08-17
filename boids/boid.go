package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func randomFloat() float32 {
	min := -1.0
	max := 1.0
	return float32(min + rand.Float64()*(max-min))
}

func randomScreenW() float32 {
	min := -float64(rl.GetScreenWidth())
	max := float64(rl.GetScreenWidth())

	return float32(min + rand.Float64()*(max-min))
}

func randomScreenH() float32 {
	min := -float64(rl.GetScreenHeight())
	max := float64(rl.GetScreenHeight())

	return float32(min + rand.Float64()*(max-min))
}

type Boid struct {
	postion      rl.Vector2
	velocity     rl.Vector2
	acceleration rl.Vector2
	maxSpeed     float32
}

func NewBoid() *Boid {
	return &Boid{
		postion: rl.Vector2{
			X: randomScreenW(),
			Y: randomScreenH(),
		},
		velocity: rl.Vector2{
			X: randomFloat(),
			Y: randomFloat(),
		},
		maxSpeed: 1,
	}
}

func (b *Boid) update() {
	b.postion = rl.Vector2Add(b.postion, b.velocity)
	b.velocity = rl.Vector2Add(b.velocity, b.acceleration)
	b.velocity = setMagnitude(b.velocity, 4)

	b.acceleration = rl.Vector2Multiply(b.acceleration, rl.Vector2{})
}

func (b *Boid) edges() {
	if b.postion.X > float32(rl.GetScreenWidth()) {
		b.postion.X = 0
	} else if b.postion.X < 0 {
		b.postion.X = float32(rl.GetScreenWidth())
	}

	if b.postion.Y > float32(rl.GetScreenHeight()) {
		b.postion.Y = 0
	} else if b.postion.Y < 0 {
		b.postion.Y = float32(rl.GetScreenHeight())
	}
}

func (b *Boid) flock(boids []*Boid) {
	alignment := b.align(boids)
	b.acceleration = rl.Vector2Add(b.acceleration, alignment)
}

func (b *Boid) align(boids []*Boid) rl.Vector2 {
	const perception = 50

	steeringForce := rl.Vector2{}
	total := 0

	for _, other := range boids {
		d := rl.Vector2Distance(b.postion, other.postion)
		if d < perception && other != b {
			steeringForce = rl.Vector2Add(other.velocity, steeringForce)
			total++
		}
	}

	const maxForce float32 = 1

	if total > 0 {
		steeringForce.X = steeringForce.X / float32(total)
		steeringForce.Y = steeringForce.Y / float32(total)

		steeringForce = setMagnitude(steeringForce, b.maxSpeed)
		steeringForce = rl.Vector2Subtract(steeringForce, b.velocity)
		steeringForce = limitVector2(steeringForce, maxForce)
	}

	return steeringForce
}

func (b *Boid) show() {
	rl.DrawCircleV(b.postion, 5, rl.Blue)
}
