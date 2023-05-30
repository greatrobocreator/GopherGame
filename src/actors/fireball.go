package actors

import "math"

type AFireball struct {
	ABall
}

func NewAFireball(radius float64) *ABall {
	fireball := NewABall(radius)
	fireball.gravityScale = 0
	fireball.rotation = math.Pi
	return fireball
}
