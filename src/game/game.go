package game

import (
	"time"

	"math"
	"math/rand"

	"gitlab.com/slon/shad-go/wasm/flappygopher/src/actors"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/events"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/physics"
	"gitlab.com/slon/shad-go/wasm/flappygopher/src/utils"
)

type Game struct {
	Objects       map[physics.Object]struct{}
	physicsEngine *physics.PhysicsEngine

	wallSpeed        float64
	wallPosition     float64
	wallsDistance    float64
	wallGap          float64
	wallsCount       int
	targetWallsCount int

	gameOver    bool
	timeScale   float64
	windowsSize utils.Vector

	Player *actors.AGopher
}

func NewGame(windowsSize utils.Vector) *Game {
	return &Game{
		Objects:          make(map[physics.Object]struct{}),
		physicsEngine:    physics.NewPhysicsEngine(),
		wallSpeed:        25,
		wallsDistance:    50,
		wallGap:          windowsSize.Y * 0.4,
		targetWallsCount: 5,

		windowsSize: windowsSize,
	}
}

func (g *Game) Tick(deltaTime time.Duration) {

	g.updateObstacles(deltaTime)

	for object := range g.Objects {
		object.EventTick(deltaTime)
	}

	g.physicsEngine.Tick(deltaTime)
}

func (g *Game) updateObstacles(deltaTime time.Duration) {
	/*if g.gameOver {
		return
	}*/

	g.wallPosition += g.wallSpeed * deltaTime.Seconds()
	if g.wallsCount < g.targetWallsCount && !g.gameOver {
		g.spawnWalls()
	} else if g.wallsCount == g.targetWallsCount {
		// Stop flying
	} else {
		g.spawnBalls()
	}
	if g.wallPosition >= 5*g.wallsDistance {
		g.Player.BecomeWalking()
		g.spawnBalls()
	}

}

func (g *Game) spawnBalls() {

	distanceMultiplier := math.Min(6.6/float64(g.wallsCount-g.targetWallsCount), 1.0)
	if g.wallPosition >= g.wallsDistance*distanceMultiplier {
		g.wallPosition = 0
		g.wallsCount++

		var ball *actors.ABall
		if rand.Float64() > 0.5 {
			ball = actors.NewABall(g.windowsSize.Y / 50)
			ball.SetPosition(utils.NewVector(g.windowsSize.X, g.windowsSize.Y-200))
			ball.ApplyForce(utils.NewVector(-60, (rand.Float64()-0.25)*50))
		} else {
			ball = actors.NewAFireball(g.windowsSize.Y / 50)
			ball.SetPosition(utils.NewVector(g.windowsSize.X, g.windowsSize.Y-30))
			ball.ApplyForce(utils.NewVector(-100, 0))
		}
		g.Spawn(ball)
	}
}

func (g *Game) spawnWalls() {
	if g.wallPosition >= g.wallsDistance {

		g.wallPosition = 0
		g.wallsCount++

		gapY := rand.Float64()*(g.windowsSize.Y-2*g.wallGap) + g.wallGap/2
		wallHeight := g.windowsSize.Y

		wall1 := actors.NewAWall(utils.NewVector(wallHeight/12, wallHeight))
		wall2 := actors.NewAWall(utils.NewVector(wallHeight/12, wallHeight))

		wall1.SetPosition(utils.NewVector(g.windowsSize.X, -wallHeight+gapY))
		wall2.SetPosition(utils.NewVector(g.windowsSize.X, g.wallGap+gapY))

		wall1.ApplyForce(utils.NewVector(-25, 0))
		wall2.ApplyForce(utils.NewVector(-25, 0))

		g.Spawn(wall1)
		g.Spawn(wall2)
	}
}

func (g *Game) Spawn(obj physics.Object) {
	g.Objects[obj] = struct{}{}

	if v, ok := obj.(physics.PhysicsBody); ok {
		g.physicsEngine.AddPhysicBody(v)
	}
}

func (g *Game) Destroy(obj physics.Object) {
	delete(g.Objects, obj)
	if v, ok := obj.(physics.PhysicsBody); ok {
		g.physicsEngine.DeletePhysicBody(v)
	}
}

func (g *Game) EventSpace() {
	for object := range g.Objects {
		if v, ok := object.(events.EventSpace); ok {
			v.EventSpace()
		}
	}
}

func (g *Game) EventMoveRight(axisValue float64) {
	for object := range g.Objects {
		if v, ok := object.(events.EventMoveRight); ok {
			v.EventMoveRight(axisValue)
		}
	}
}

/*func (g *Game) EventLeftDown() {
	for object := range g.Objects {
		if v, ok := object.(events.EventLeftDown); ok {
			v.EventLeftDown()
		}
	}
}

func (g *Game) EventRightDown() {
	for object := range g.Objects {
		if v, ok := object.(events.EventRightDown); ok {
			v.EventRightDown()
		}
	}
}

func (g *Game) EventLeftUp() {
	for object := range g.Objects {
		if v, ok := object.(events.EventLeftUp); ok {
			v.EventLeftUp()
		}
	}
}

func (g *Game) EventRightUp() {
	for object := range g.Objects {
		if v, ok := object.(events.EventRightUp); ok {
			v.EventRightUp()
		}
	}
}*/

func (g *Game) SendEvent(e events.Event) {
	switch e.(type) {
	case events.GameOverEvent:
		g.GameOver()
	default:
	}
}

func (g *Game) GameOver() {
	g.gameOver = true
	for object := range g.Objects {
		if wall, ok := object.(*actors.AWall); ok {
			wall.ApplyForce(wall.GetVelocity().Mul(-1))
		}
	}
}
