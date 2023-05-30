package physics

import (
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

/*type Object struct {
	Position Vector
}*/

//func (o *Object) Tick(deltaTime time.Duration) {}

type Object interface {
	GetPosition() utils.Vector
	SetPosition(utils.Vector)
	GetRotation() float64
	EventTick(deltaTime time.Duration)
}

// TODO: Normal collision system like in UE

// Blocks all collisions
type PhysicsBody interface {
	// returns Collider of this object
	Collider() utils.Shape
	EventHit(other PhysicsBody)
}

/*type MovableObject struct {
	Object

	velocity      Vector
	gravity_scale float64
}*/

type MovableObject interface {
	Object

	GetVelocity() utils.Vector
	ApplyForce(deltaVelocity utils.Vector)
	ApplyGravity(gravity utils.Vector) // gravity = deltaVelocitys
}
