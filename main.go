//go:build !solution

package main

import (
	"fmt"
	"time"

	global "gitlab.com/slon/shad-go/wasm/flappygopher/src"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/actors"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/game"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/rendering"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
	"honnef.co/go/js/dom/v2"
)

//. "gitlab.com/slon/shad-go/wasm/flappygopher/main"

func main() {

	currentGame := game.NewGame()
	global.SetGame(currentGame)

	canvas, windowSize := createCanvas()
	renderer := rendering.NewRenderer(canvas, windowSize)

	canvas.LoadImage("gopher.png")
	canvas.LoadImage("pipe.png")
	canvas.LoadImage("pipe_flipped.png")

	for i := 1; i <= 5; i += 1 {
		canvas.LoadImage(fmt.Sprintf("fireball/FB00%d.png", i))
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

	gopher := actors.NewAGopher(50, utils.NewVector(50, 44))
	gopher.SetPosition(utils.NewVector(200, 100))
	currentGame.Player = gopher
	currentGame.Spawn(gopher)

	floor := actors.NewAFloor(utils.NewVector(1500, 150))
	floor.SetPosition(utils.NewVector(0, 800))
	currentGame.Spawn(floor)

	select {}
}

func createCanvas() (*rendering.Canvas, utils.Vector) {
	domCanvas := dom.GetWindow().Document().CreateElement("canvas").(*dom.HTMLCanvasElement)
	domCanvas.SetHeight(dom.GetWindow().InnerHeight())
	domCanvas.SetWidth(dom.GetWindow().InnerWidth())
	dom.GetWindow().Document().QuerySelector("body").AppendChild(domCanvas)

	return rendering.NewCanvas(domCanvas.GetContext2d()),
		utils.Vector{
			X: float64(dom.GetWindow().InnerWidth()),
			Y: float64(dom.GetWindow().InnerHeight()),
		}
}
