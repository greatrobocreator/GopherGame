package global

import (
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/events"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/physics"
)

type game_interface interface {
	Spawn(o physics.Object)
	Destroy(o physics.Object)
	SendEvent(e events.Event)
}

var game game_interface

func SetGame(g game_interface) {
	game = g
}

func Game() game_interface {
	return game
}
