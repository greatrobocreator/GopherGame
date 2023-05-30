package actors

import (
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type AFloor struct {
	AActor

	size utils.Vector
}

func NewAFloor(size utils.Vector) *AFloor {
	return &AFloor{AActor: *NewAActor(), size: size}
}

func (g *AFloor) Render(canvas *rendering.Canvas, deltaTime time.Duration) {
	//canvas.DrawCircle(utils.NewVector(0, 0), s)
	//fmt.Println("rendering floor")
	canvas.SetFillStyle("rgba(0, 255, 255, 0)")
	canvas.FillRect(utils.NewRectangle(utils.NewVector(0, 0), g.size))
}

func (g *AFloor) Collider() utils.Shape {
	return utils.NewRectangle(
		g.position,
		g.size,
	)
}
