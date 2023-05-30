package rendering

import (
	"math"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
	. "gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
	"honnef.co/go/js/dom/v2"
	//"honnef.co/go/js/dom"
	//. "github.com/siongui/godom"
)

// +- equvalent to Camera
type Canvas struct {
	canvas *dom.CanvasRenderingContext2D // must be <canvas> context2d

	offset Vector

	// TODO: Add scale so velocity will be absolute, not in pixels
	scale Vector

	images map[string]*dom.HTMLImageElement
}

func NewCanvas(c *dom.CanvasRenderingContext2D) *Canvas {
	return &Canvas{canvas: c, images: make(map[string]*dom.HTMLImageElement)}
}

func (c *Canvas) WrapOffset(offset Vector) *Canvas {
	return &Canvas{canvas: c.canvas, offset: c.offset.Add(offset), images: c.images}
}

func (c *Canvas) DrawCircle(pos Vector, radius float64) {
	pos = pos.Add(c.offset)
	c.canvas.BeginPath()
	c.canvas.Arc(pos.X, pos.Y, radius, 0, math.Pi*2, false)
	c.canvas.Fill()
}

func (c *Canvas) FillRect(rect utils.Rectangle) {
	pos := rect.Pos.Add(c.offset)
	c.canvas.FillRect(pos.X, pos.Y, rect.Size.X, rect.Size.Y)
}

func (c *Canvas) LoadImage(name string) {
	//fmt.Println("Loading image: ", name)
	image := dom.GetWindow().Document().CreateElement("img").(*dom.HTMLImageElement)
	image.SetSrc("assets/" + name)
	//dom.GetWindow().Document().QuerySelector("body").AppendChild(image)

	c.images[name] = image
}

func (c *Canvas) DrawImage(name string, rect utils.Rectangle) {
	image, ok := c.images[name]
	if !ok {
		//panic("No such image")
		return
	}

	pos := rect.Pos.Add(c.offset)
	//c.canvas.DrawImageWithDst(image, pos.X, pos.Y, rect.Size.X, rect.Size.Y)
	c.canvas.Call("drawImage", image.Underlying(), pos.X, pos.Y, rect.Size.X, rect.Size.Y)
}

func (c *Canvas) ClearRect(rect utils.Rectangle) {
	pos := rect.Pos.Add(c.offset)
	c.canvas.ClearRect(pos.X, pos.Y, rect.Size.X, rect.Size.Y)
}

func (c *Canvas) FillText(s string, pos utils.Vector, maxWidth float64) {
	pos = pos.Add(c.offset)
	c.canvas.FillText(s, pos.X, pos.Y, maxWidth)
}

func (c *Canvas) SetFillStyle(s string) {
	c.canvas.SetFillStyle(s)
}

func (c *Canvas) Translate(pos utils.Vector) {
	c.canvas.Translate(pos.X, pos.Y)
}

func (c *Canvas) Rotate(angle float64) {
	c.canvas.Rotate(angle)
}

func (c *Canvas) Save() {
	c.canvas.Save()
}

func (c *Canvas) Restore() {
	c.canvas.Restore()
}

func (c *Canvas) SetGlobalCompositeOperation(s string) {
	c.canvas.SetGlobalCompositeOperation(s)
}

/*type WrapCanvas struct {
	*Canvas

	offset Vector
}

func (c *Canvas) WrapOffset(offset Vector) *WrapCanvas {
	return &WrapCanvas{Canvas: c, offset: offset}
}

func (w *WrapCanvas) WrapOffset(offset Vector) *WrapCanvas {
	return &WrapCanvas{Canvas: w.Canvas, offset: w.offset.Add(offset)}
}
*/
