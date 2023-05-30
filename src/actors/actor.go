package actors

import (
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/physics"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type AActor struct {
	position utils.Vector
	rotation float64
}

func (a *AActor) GetPosition() utils.Vector         { return a.position }
func (a *AActor) SetPosition(newPos utils.Vector)   { a.position = newPos }
func (a *AActor) GetRotation() float64              { return a.rotation }
func (a *AActor) EventTick(deltaTime time.Duration) {}

func (a *AActor) Collider() utils.Shape {
	return utils.NewRectangle(a.position, utils.NewVector(1, 1))
}
func (a *AActor) EventHit(other physics.PhysicsBody) {}

func NewAActor() *AActor {
	return &AActor{}
}

type AMovableActor struct {
	AActor

	velocity     utils.Vector
	gravityScale float64
}

func (a *AMovableActor) GetVelocity() utils.Vector { return a.velocity }
func (a *AMovableActor) ApplyForce(deltaVelocity utils.Vector) {
	a.velocity = a.velocity.Add(deltaVelocity)
}
func (a *AMovableActor) ApplyGravity(gravity utils.Vector) { a.ApplyForce(gravity.Mul(a.gravityScale)) }

func NewAMovableActor() *AMovableActor {
	return &AMovableActor{gravityScale: 1}
}
