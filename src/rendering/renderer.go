package rendering

import (
	"fmt"
	"math"
	"time"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/physics"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type Renderer struct {
	canvas     *Canvas
	windowSize utils.Vector

	deltaTime float64
	alpha     float64
}

type Renderable interface {
	Render(canvas *Canvas, deltaTime time.Duration)
}

func NewRenderer(c *Canvas, windowSize utils.Vector) *Renderer {
	return &Renderer{canvas: c, windowSize: windowSize, alpha: 0.03}
}

/*func IsRenderable(obj interface{}) (Renderable, bool) {
	if v, ok := obj.(Renderable); ok {
		return v, true
	} else {
		return nil, false
	}
}*/

func (r *Renderer) RenderObjects(objects map[physics.Object]struct{}, deltaTime time.Duration) {
	r.canvas.ClearRect(utils.Rectangle{Size: r.windowSize})

	r.deltaTime = r.alpha*float64(deltaTime.Milliseconds()) + (1-r.alpha)*r.deltaTime
	fps := int(math.Floor(1000 / r.deltaTime))
	r.canvas.FillText(fmt.Sprintf("FPS: %d", fps), utils.NewVector(10, 10), 200)

	for object := range objects {
		//fmt.Println("Rendering", object)
		if v, ok := object.(Renderable); ok {
			r.canvas.Save()
			r.canvas.Translate(object.GetPosition())
			r.canvas.Rotate(object.GetRotation())
			//v.Render(r.canvas.WrapOffset(object.GetPosition()), deltaTime)
			v.Render(r.canvas, deltaTime)
			r.canvas.Restore()
		}
	}
}
