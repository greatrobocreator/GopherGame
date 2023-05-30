package actors

import (
	"fmt"
	"math"
	"time"

	global "gitlab.com/slon/shad-go/wasm/flappygopher/src"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/animations"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/physics"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type ABall struct {
	AMovableActor

	radius    float64
	animation *animations.BasicAnimation
}

func NewABall(radius float64) *ABall {

	frames := make([]string, 5)
	for i := 0; i < len(frames); i += 1 {
		frames[i] = fmt.Sprintf("fireball/FB00%d.png", i+1)
	}

	ball := &ABall{
		AMovableActor: *NewAMovableActor(),
		radius:        radius,
		animation: animations.NewBasicAnimation(
			frames,
			time.Second,
			utils.NewRectangle(utils.NewVector(-5*radius, -2*radius), utils.NewVector(8*radius, 4*radius))),
	}
	ball.rotation = math.Pi / 2
	return ball
}

func (b *ABall) Render(canvas *rendering.Canvas, deltaTime time.Duration) {
	//canvas.SetFillStyle("red")
	//canvas.DrawCircle(utils.NewVector(0, 0), b.radius)
	//fmt.Println("rendering wall")
	//canvas.FillRect(utils.NewRectangle(utils.NewVector(0, 0), b.size))

	/*canvas.DrawImage(".png", utils.NewRectangle(
		b.position.Add(utils.NewVector(-b.radius, -b.radius)),
		utils.NewVector(b.radius, b.radius),
	))*/
	b.animation.Render(canvas, deltaTime)
}

func (b *ABall) Collider() utils.Shape {
	return utils.NewRectangle(
		b.position.Add(utils.NewVector(-b.radius, -b.radius)),
		utils.NewVector(b.radius, b.radius),
	)
}

func (b *ABall) EventTick(deltaTime time.Duration) {
	if b.position.X < -100 {
		global.Game().Destroy(b)
	}
}

func (b *ABall) EventHit(other physics.PhysicsBody) {
	if floor, ok := other.(*AFloor); ok {
		newPosY := floor.GetPosition().Y - b.radius

		//g.velocity.Y -= (g.position.Y - newPosY) / 5
		b.velocity.Y *= -1
		b.position.Y = newPosY
	}
}
