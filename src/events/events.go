package events

type Event interface{}

type GameOverEvent struct{}

type EventSpace interface {
	EventSpace()
}

type EventMoveRight interface {
	EventMoveRight(axisValue float64) // o or +-1
}

/*type EventLeftDown interface {
	EventLeftDown()
}

type EventRightDown interface {
	EventRightDown()
}

type EventLeftUp interface {
	EventLeftUp()
}

type EventRightUp interface {
	EventRightUp()
}*/
