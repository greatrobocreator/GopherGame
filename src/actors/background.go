package actors

import (
	"fmt"
	"math"
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type ABackground struct {
	AMovableActor

	size   utils.Vector
	frames []rendering.Renderable
}

func NewABackground(size utils.Vector, frames []rendering.Renderable) *ABackground {
	wall := &ABackground{AMovableActor: *NewAMovableActor(), size: size, frames: frames}
	wall.gravityScale = 0
	return wall
}

func (b *ABackground) Render(canvas *rendering.Canvas, deltaTime time.Duration) {
	//canvas.DrawCircle(utils.NewVector(0, 0), s)
	//fmt.Println("rendering wall")
	//canvas.FillRect(utils.NewRectangle(utils.NewVector(0, 0), w.size))
	//canvas.DrawImage("pipe.png", utils.NewRectangle(utils.NewVector(0, 0), w.size))

	//x := b.GetPosition().X % b.size.X
	framesCount := math.Floor(math.Abs(b.GetPosition().X) / (b.size.X))
	x := b.GetPosition().X - b.size.X*framesCount
	ind := int(framesCount)

	fmt.Println("Background: ", framesCount, x, ind, b.size.X)

	canvas.SetGlobalCompositeOperation("destination-over")
	canvas.Translate(utils.NewVector(b.size.X*framesCount, 0))
	b.frames[ind%len(b.frames)].Render(canvas, deltaTime)
	canvas.Translate(utils.NewVector(b.size.X-1, 0))
	b.frames[(ind+1)%len(b.frames)].Render(canvas, deltaTime)
	//canvas.DrawImage(b.frames[ind%len(b.frames)], utils.NewRectangle(utils.NewVector(b.size.X*framesCount, 0), b.size))
	//canvas.DrawImage(b.frames[(ind+1)%len(b.frames)], utils.NewRectangle(utils.NewVector(b.size.X*(framesCount+1), 0), b.size))
}

func (b *ABackground) EventTick(deltaTime time.Duration) {}
