package gamestate

import (
	"time"

	"github.com/digitalsquid7/tetris/pkg/gameevent"
	"github.com/digitalsquid7/tetris/pkg/gamestate/board"
	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/direction"
	"github.com/digitalsquid7/tetris/pkg/gamestate/tetromino"
	"github.com/digitalsquid7/tetris/pkg/gamestate/tetrominoqueue"
)

const (
	moveDelay                        = time.Millisecond * 50
	automaticDropDelay time.Duration = time.Millisecond * 500
)

type Publisher interface {
	Publish(name gameevent.Name)
}

var scoreByLinesCleared = map[int]int{
	1: 1,
	2: 3,
	3: 5,
	4: 8,
}

type GameState struct {
	publisher        Publisher
	board            *board.Board
	currentTetromino *tetromino.Tetromino
	tetrominoQueue   *tetrominoqueue.TetrominoQueue
	tetrominoHeld    *tetromino.Tetromino
	tetrominoSwapped bool
	gameOver         bool
	level            int
	linesCleared     int
	score            int
	moveStart        time.Time
	moveDelay        *time.Ticker
	dropDelay        *time.Ticker
	downTicker       *time.Ticker
}

func New(publisher *gameevent.Publisher) *GameState {
	tetrisBoard := board.NewBoard()
	tetrominoQueue := tetrominoqueue.New(tetrisBoard)
	tetrominoQueue.GenerateTetrominos()

	return &GameState{
		publisher:        publisher,
		board:            tetrisBoard,
		currentTetromino: tetrominoQueue.Pop(),
		tetrominoQueue:   tetrominoQueue,
		level:            1,
		moveDelay:        time.NewTicker(moveDelay),
		dropDelay:        time.NewTicker(moveDelay),
		downTicker:       time.NewTicker(getAutomaticDropDelay(1)),
	}
}

func (g *GameState) Board() *board.Board {
	return g.board
}

func (g *GameState) Level() int {
	return g.level
}

func (g *GameState) Score() int {
	return g.score
}

func (g *GameState) CurrentTetromino() *tetromino.Tetromino {
	return g.currentTetromino
}

func (g *GameState) HeldTetromino() *tetromino.Tetromino {
	return g.tetrominoHeld
}

func (g *GameState) ReplaceTetromino() {
	g.newTetromino()

	total := g.board.ClearLines()
	if total > 0 {
		g.linesCleared += total
		g.level = (g.linesCleared / 10) + 1
		g.downTicker.Reset(getAutomaticDropDelay(g.level))
		g.score += scoreByLinesCleared[total]
		g.publisher.Publish(gameevent.LineClear)
	}

	g.tetrominoSwapped = false
}

func (g *GameState) NextTetromino() *tetromino.Tetromino {
	return g.tetrominoQueue.Peek()
}

func (g *GameState) HoldTetromino() {
	if g.tetrominoSwapped {
		return
	}

	if g.tetrominoHeld == nil {
		g.tetrominoHeld = g.currentTetromino
		g.newTetromino()
	} else {
		g.tetrominoHeld, g.currentTetromino = g.currentTetromino, g.tetrominoHeld
		g.resetTetromino()
	}

	g.publisher.Publish(gameevent.HoldTetromino)
	g.tetrominoSwapped = true
}

func (g *GameState) MoveRight(startMove bool) {
	if startMove {
		g.moveStart = time.Now()
	} else if !g.movable() {
		return
	}

	if !g.collision(direction.Right) {
		g.currentTetromino.MoveRight(1)
	}
	return
}

func (g *GameState) MoveLeft(startMove bool) {
	if startMove {
		g.moveStart = time.Now()
	} else if !g.movable() {
		return
	}

	if !g.collision(direction.Left) {
		g.currentTetromino.MoveLeft(1)
	}
	return
}

func (g *GameState) MoveDown() bool {
	if g.collision(direction.Down) {
		g.board.LockInPlace(g.currentTetromino)
		g.ReplaceTetromino()
		g.publisher.Publish(gameevent.Drop)
		return true
	}

	g.currentTetromino.MoveDown(1)
	return false
}

func (g *GameState) SoftDrop() {
	select {
	case <-g.dropDelay.C:
		g.MoveDown()
	default:
	}
}

func (g *GameState) HardDrop() {
	coors := g.currentTetromino.Coordinates()
	shifted := g.shiftCoordinates(coors, direction.Down)

	for g.board.FreeSpace(shifted) {
		g.currentTetromino.MoveDown(1)
		shifted = g.shiftCoordinates(shifted, direction.Down)
	}

	g.board.LockInPlace(g.currentTetromino)
	g.publisher.Publish(gameevent.Drop)
	g.ReplaceTetromino()
}

func (g *GameState) RotateClockwise() {
	g.updateDirection(g.currentTetromino.Direction().Clockwise())
}

func (g *GameState) RotateAntiClockwise() {
	g.updateDirection(g.currentTetromino.Direction().AntiClockwise())
}

func (g *GameState) AutomaticDrop() {
	select {
	case <-g.downTicker.C:
		g.MoveDown()
	default:
	}
}

func (g *GameState) GhostCoordinates() []coordinate.Coordinate {
	coors := g.currentTetromino.Coordinates()
	shifted := g.shiftCoordinates(coors, direction.Down)

	for g.board.FreeSpace(shifted) {
		coors = shifted
		shifted = g.shiftCoordinates(coors, direction.Down)
	}

	return coors
}

func (g *GameState) TopOut() bool {
	if !g.board.FreeSpace(g.currentTetromino.Coordinates()) {
		for g.board.Collision(g.currentTetromino.Coordinates()) {
			g.currentTetromino.MoveUp(1)
		}
		g.gameOver = true
		g.publisher.Publish(gameevent.GameOver)
	}

	return g.gameOver
}

func (g *GameState) GameOver() bool {
	return g.gameOver
}

func (g *GameState) resetTetromino() {
	g.currentTetromino.ResetPosition()
	g.downTicker.Reset(automaticDropDelay)
}

func (g *GameState) newTetromino() {
	g.currentTetromino = g.tetrominoQueue.Pop()

	if g.tetrominoQueue.Len() == 1 {
		g.tetrominoQueue.GenerateTetrominos()
	}
}

func (g *GameState) movable() bool {
	initialDelay := g.moveStart.Add(time.Millisecond * 250)
	if time.Now().Before(initialDelay) {
		return false
	}

	select {
	case <-g.moveDelay.C:
		return true
	default:
	}

	return false
}

func (g *GameState) updateDirection(dir direction.Direction) {
	coors := g.currentTetromino.FindCoordinates(g.currentTetromino.Pivot(), dir)

	if g.board.FreeSpace(coors) {
		g.currentTetromino.SetDirection(dir)
		g.publisher.Publish(gameevent.Rotate)
		return
	}

	if amount, ok := g.freeSpaceWithShift(coors, direction.Right); ok {
		g.currentTetromino.SetDirection(dir)
		g.currentTetromino.MoveRight(amount)
		g.publisher.Publish(gameevent.Rotate)
		return
	}

	if amount, ok := g.freeSpaceWithShift(coors, direction.Left); ok {
		g.currentTetromino.SetDirection(dir)
		g.currentTetromino.MoveLeft(amount)
		g.publisher.Publish(gameevent.Rotate)
	}
}
func (g *GameState) freeSpaceWithShift(coors []coordinate.Coordinate, direction direction.Direction) (int, bool) {
	shifted := g.shiftCoordinates(coors, direction)
	amount := 1

	if !g.board.InSideBounds(shifted) {
		shifted = g.shiftCoordinates(shifted, direction)
		amount++
	}

	if g.board.FreeSpace(shifted) {
		return amount, true
	}

	return 0, false
}

func (g *GameState) collision(direction direction.Direction) bool {
	coors := g.currentTetromino.Coordinates()
	shifted := g.shiftCoordinates(coors, direction)

	if g.board.FreeSpace(shifted) {
		return false
	}

	return true
}

func (g *GameState) shiftCoordinates(coordinates []coordinate.Coordinate, shift direction.Direction) []coordinate.Coordinate {
	shifted := make([]coordinate.Coordinate, len(coordinates))

	switch shift {
	case direction.Left:
		for i := range coordinates {
			shifted[i] = coordinates[i].Left(1)
		}
	case direction.Right:
		for i := range coordinates {
			shifted[i] = coordinates[i].Right(1)
		}
	case direction.Down:
		for i := range coordinates {
			shifted[i] = coordinates[i].Down(1)
		}
	}

	return shifted
}

func getAutomaticDropDelay(level int) time.Duration {
	multiplier := 0.8 - ((float64(level - 1)) * 0.007)
	return time.Duration(multiplier * float64(time.Second))
}
