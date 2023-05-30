package physics

import (
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type PhysicsEngine struct {
	bodies  map[PhysicsBody]struct{}
	gravity utils.Vector
}

func NewPhysicsEngine() *PhysicsEngine {
	return &PhysicsEngine{bodies: make(map[PhysicsBody]struct{}, 0), gravity: utils.NewVector(0, 9.81)}
}

func (e *PhysicsEngine) AddPhysicBody(b PhysicsBody) {
	//e.bodies = append(e.bodies, b)
	e.bodies[b] = struct{}{}
}

func (e *PhysicsEngine) DeletePhysicBody(b PhysicsBody) {
	delete(e.bodies, b)
}

func (e *PhysicsEngine) Tick(deltaTime time.Duration) {
	// Update velocities using gravity
	// Update position
	// Collide bodies

	scaleFactor := 10.0

	for body := range e.bodies {
		if v, ok := body.(MovableObject); ok {
			v.ApplyGravity(e.gravity.Mul(deltaTime.Seconds() * scaleFactor))
			v.SetPosition(v.GetPosition().Add(v.GetVelocity().Mul(deltaTime.Seconds() * scaleFactor)))
		}
	}

	for body := range e.bodies {
		if _, ok := body.(MovableObject); !ok {
			continue
		}
		for other := range e.bodies {
			if body == other {
				continue
			}

			if CheckIntersection(body.Collider(), other.Collider()) {
				body.EventHit(other)
				other.EventHit(body)
			}
		}
	}
}

func CheckIntersection(l utils.Shape, r utils.Shape) bool {
	rect1, ok1 := l.(utils.Rectangle)
	rect2, ok2 := r.(utils.Rectangle)
	if !ok1 || !ok2 {
		return false
	}

	return utils.RectanglesIntersection(rect1, rect2)
}
