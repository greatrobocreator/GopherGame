package actors

import (
	"time"

	global "gitlab.com/slon/shad-go/wasm/flappygopher/src"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type AWall struct {
	AMovableActor

	size utils.Vector
}

func NewAWall(size utils.Vector) *AWall {
	wall := &AWall{AMovableActor: *NewAMovableActor(), size: size}
	wall.gravityScale = 0
	return wall
}

func (w *AWall) Render(canvas *rendering.Canvas, deltaTime time.Duration) {
	//canvas.DrawCircle(utils.NewVector(0, 0), s)
	//fmt.Println("rendering wall")
	//canvas.FillRect(utils.NewRectangle(utils.NewVector(0, 0), w.size))
	canvas.DrawImage("pipe.png", utils.NewRectangle(utils.NewVector(0, 0), w.size))
}

func (w *AWall) Collider() utils.Shape {
	return utils.NewRectangle(
		w.position,
		w.size,
	)
}

func (w *AWall) EventTick(deltaTime time.Duration) {
	if w.position.X < -100 {
		global.Game().Destroy(w)
	}
}
