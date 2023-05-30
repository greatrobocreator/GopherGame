//go:build !solution

package main

import (
	"fmt"
	"time"

	global "gitlab.com/slon/shad-go/wasm/flappygopher/src"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/actors"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/animations"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/game"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
	"honnef.co/go/js/dom/v2"
)

//. "gitlab.com/slon/shad-go/wasm/flappygopher/main"

func main() {

	canvas, windowSize := createCanvas()
	renderer := rendering.NewRenderer(canvas, windowSize)

	currentGame := game.NewGame(windowSize)
	global.SetGame(currentGame)

	canvas.LoadImage("gopher.png")
	canvas.LoadImage("pipe.png")
	canvas.LoadImage("pipe_flipped.png")

	for i := 1; i <= 5; i += 1 {
		canvas.LoadImage(fmt.Sprintf("fireball/FB00%d.png", i))
	}

	for i := 1; i <= 3; i += 1 {
		canvas.LoadImage(fmt.Sprintf("background/%d.jpg", i))
	}

	moveRightAxisValue := 0.0

	var lastFrameTime time.Duration
	var tick func(currTime time.Duration)
	tick = func(currTime time.Duration) {

		//<-time.After(time.Millisecond * 100)
		deltaTime := currTime - lastFrameTime
		lastFrameTime = currTime
		currentGame.EventMoveRight(moveRightAxisValue)
		currentGame.Tick(deltaTime)
		renderer.RenderObjects(currentGame.Objects, deltaTime)

		dom.GetWindow().RequestAnimationFrame(tick)
	}
	dom.GetWindow().RequestAnimationFrame(tick)

	keyAxisValue := map[string]float64{
		"KeyD":       1,
		"KeyA":       -1,
		"ArrowRight": 1,
		"ArrowLeft":  -1,
	}

	dom.GetWindow().AddEventListener("keydown", false, func(e dom.Event) {
		key := e.(*dom.KeyboardEvent).Get("code").String()
		if v, ok := keyAxisValue[key]; ok {
			moveRightAxisValue = v
		}
		if key == "Space" || key == "ArrowUp" {
			currentGame.EventSpace()
		}
	})

	dom.GetWindow().AddEventListener("mousedown", false, func(e dom.Event) {
		currentGame.EventSpace()
	})

	dom.GetWindow().AddEventListener("keyup", false, func(e dom.Event) {
		key := e.(*dom.KeyboardEvent).Get("code").String()
		/*if key == "Space" || key == "ArrowUp" {
			currentGame.EventSpace()
		} else */if v, ok := keyAxisValue[key]; ok {
			if moveRightAxisValue == v {
				moveRightAxisValue = 0
			}
		}
	})

	gopher := actors.NewAGopher(50, utils.NewVector(windowSize.Y/10, windowSize.Y/10*88/100))
	gopher.SetPosition(utils.NewVector(windowSize.X*0.1, windowSize.Y*0.5))
	currentGame.Player = gopher
	currentGame.Spawn(gopher)

	floor := actors.NewAFloor(utils.NewVector(windowSize.X, windowSize.Y))
	floor.SetPosition(utils.NewVector(0, windowSize.Y))
	currentGame.Spawn(floor)

	backgroundFrames := make([]string, 3)
	for i := 0; i < len(backgroundFrames); i += 1 {
		backgroundFrames[i] = fmt.Sprintf("background/%d.jpg", i+1)
	}
	backgroundAnimation := animations.NewBasicAnimation(backgroundFrames, time.Second, utils.NewRectangle(utils.NewVector(0, 0), windowSize))

	background := actors.NewABackground(windowSize, []rendering.Renderable{backgroundAnimation})
	background.ApplyForce(utils.NewVector(-25, 0))
	currentGame.Spawn(background)

	select {}
}

func createCanvas() (*rendering.Canvas, utils.Vector) {
	domCanvas := dom.GetWindow().Document().CreateElement("canvas").(*dom.HTMLCanvasElement)
	domCanvas.SetHeight(dom.GetWindow().InnerHeight())
	domCanvas.SetWidth(dom.GetWindow().InnerWidth())
	dom.GetWindow().Document().QuerySelector("body").AppendChild(domCanvas)

	scale := float64(dom.GetWindow().InnerHeight()) / 1080.0

	ctx := domCanvas.GetContext2d()
	ctx.Scale(scale, scale)
	return rendering.NewCanvas(ctx),
		utils.Vector{
			X: float64(dom.GetWindow().InnerWidth()) / scale,
			Y: float64(dom.GetWindow().InnerHeight()) / scale,
		}
}
