package animations

import (
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type Animation interface {
	rendering.Renderable
}

type BasicAnimation struct {
	state    time.Duration
	frames   []string
	duration time.Duration
	posSize  utils.Rectangle
}

func NewBasicAnimation(frames []string, duration time.Duration, posSize utils.Rectangle) *BasicAnimation {
	return &BasicAnimation{frames: frames, duration: duration, posSize: posSize}
}

func (a *BasicAnimation) Render(canvas *rendering.Canvas, deltaTime time.Duration) {
	a.state += deltaTime
	framesCount := int64(len(a.frames))
	frame := a.frames[(a.state.Milliseconds()/(a.duration.Milliseconds()/framesCount))%framesCount]
	canvas.DrawImage(frame, a.posSize)
}
