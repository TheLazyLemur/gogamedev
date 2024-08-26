package utils

import (
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RandomVec2D() rl.Vector2 {
	angle := rand.Float64() * 2 * math.Pi
	randomVec := rl.NewVector2(float32(math.Cos(angle)), float32(math.Sin(angle)))
	return randomVec
}

func SetMag(v rl.Vector2, m float32) rl.Vector2 {
	normalizedVelocity := rl.Vector2Normalize(v)
	v = rl.Vector2Scale(normalizedVelocity, m)
	return v
}

func Random(rv float32) float32 {
	value := rand.Float32() * rv
	return value
}

func RandomInRange(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func LimitVector(vector rl.Vector2, max float32) rl.Vector2 {
	length := rl.Vector2Length(vector)
	if length > max {
		return rl.Vector2Scale(rl.Vector2Normalize(vector), max)
	}

	return vector
}

func ZeroVector(vector rl.Vector2) rl.Vector2 {
	return rl.Vector2{X: 0, Y: 0}
}

func DivideVectorByScalar(vector rl.Vector2, scalar float32) rl.Vector2 {
	return rl.Vector2{
		X: vector.X / scalar,
		Y: vector.Y / scalar,
	}
}
