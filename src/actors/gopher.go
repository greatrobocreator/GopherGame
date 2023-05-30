package actors

import (
	"fmt"
	"math"
	"time"

	global "gitlab.com/slon/shad-go/wasm/flappygopher/src"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/events"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/physics"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type AGopher struct {
	AMovableActor

	jumpForce float64
	size      utils.Vector
	dead      bool

	walking bool
	inAir   bool
}

func NewAGopher(jumpForce float64, size utils.Vector) *AGopher {
	gopher := &AGopher{AMovableActor: *NewAMovableActor(), jumpForce: jumpForce, size: size}

	//gopher.gravityScale = 0.75

	return gopher
}

func (g *AGopher) Render(canvas *rendering.Canvas, deltaTime time.Duration) {
	canvas.DrawImage("gopher.png", utils.NewRectangle(utils.NewVector(-g.size.X/2, -g.size.Y/2), utils.NewVector(g.size.X, g.size.Y)))
}

func (g *AGopher) Collider() utils.Shape {
	return utils.NewRectangle(
		g.position.Add(utils.NewVector(-g.size.X/2, -g.size.Y/2)),
		utils.NewVector(g.size.X, g.size.Y),
	)
}

func (g *AGopher) EventHit(other physics.PhysicsBody) {
	if floor, ok := other.(*AFloor); ok {
		newPosY := floor.GetPosition().Y - g.size.Y/2

		//g.velocity.Y -= (g.position.Y - newPosY) / 5
		//g.velocity.Y *= -1

		g.position.Y = newPosY

		g.inAir = false
	} else if !g.dead {
		//g.position.X = other.(physics.Object).GetPosition().X - g.size.X*0.6
		g.rotation = math.Pi / 2
		g.velocity = utils.NewVector(0, 0)
		g.dead = true
		global.Game().SendEvent(events.GameOverEvent{})
	}
}

func (g *AGopher) EventSpace() {
	fmt.Println("Gopher: EventSpace", g.dead, g.walking, g.inAir)
	if !g.dead && (!g.walking || !g.inAir) {
		g.inAir = true
		g.velocity.Y = -g.jumpForce
	}
}

func (g *AGopher) BecomeWalking() {
	g.walking = true
	g.inAir = true
	//g.jumpForce *= 0.5
}

/*func (g *AGopher) EventLeftDown() {
	fmt.Println("Gopher: EventLeft", g.dead, g.walking, g.inAir)
	if !g.dead && g.walking {
		g.velocity.X = -20
	}
}

func (g *AGopher) EventLeftDown() {
	fmt.Println("Gopher: EventLeft", g.dead, g.walking, g.inAir)
	if !g.dead && g.walking {
		g.velocity.X = -20
	}
}*/

func (g *AGopher) EventMoveRight(axisValue float64) {
	if !g.dead && g.walking {
		g.velocity.X = 20 * axisValue
	}
}
