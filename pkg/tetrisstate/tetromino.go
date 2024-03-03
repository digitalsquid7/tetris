package tetrisstate

import (
	"time"

	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

type FindCoordinates func(pivot Coordinate, direction Direction) []Coordinate

type Tetromino struct {
	board           *Board
	blockSprite     tetrissprites.Name
	tetrominoSprite tetrissprites.Name
	defaultPivot    Coordinate
	pivot           Coordinate
	direction       Direction
	findCoordinates FindCoordinates
	moveStart       time.Time
	moveDelay       <-chan time.Time
	dropDelay       <-chan time.Time
	downTicker      <-chan time.Time
}

func NewTetromino(
	board *Board,
	blockSprite tetrissprites.Name,
	tetrominoSprite tetrissprites.Name,
	pivot Coordinate,
	findCoordinates FindCoordinates) *Tetromino {
	return &Tetromino{
		board:           board,
		blockSprite:     blockSprite,
		tetrominoSprite: tetrominoSprite,
		pivot:           pivot,
		defaultPivot:    pivot,
		findCoordinates: findCoordinates,
		direction:       Up,
		moveDelay:       time.Tick(time.Millisecond * 50),
		dropDelay:       time.Tick(time.Millisecond * 50),
		downTicker:      time.Tick(time.Millisecond * 250),
	}
}

func (t *Tetromino) BlockSprite() tetrissprites.Name {
	return t.blockSprite
}

func (t *Tetromino) TetrominoSprite() tetrissprites.Name {
	return t.tetrominoSprite
}

func (t *Tetromino) Coordinates() []Coordinate {
	return t.findCoordinates(t.pivot, t.direction)
}

func (t *Tetromino) MoveRight(startMove bool) {
	if startMove {
		t.moveStart = time.Now()
	}

	if !t.collision(Right) {
		t.pivot = t.pivot.Right(1)
	}
	return
}

func (t *Tetromino) MoveLeft(startMove bool) {
	if startMove {
		t.moveStart = time.Now()
	}

	if !t.collision(Left) {
		t.pivot = t.pivot.Left(1)
	}
	return
}

func (t *Tetromino) Movable() bool {
	initialDelay := t.moveStart.Add(time.Millisecond * 250)
	if time.Now().Before(initialDelay) {
		return false
	}

	select {
	case <-t.moveDelay:
		return true
	default:
	}

	return false
}

func (t *Tetromino) MoveDown() bool {
	if t.collision(Down) {
		t.board.LockInPlace(t)
		return true
	}

	t.pivot = t.pivot.Down(1)
	return false
}

func (t *Tetromino) SoftDrop() bool {
	select {
	case <-t.dropDelay:
		return t.MoveDown()
	default:
	}

	return false
}

func (t *Tetromino) HardDrop() {
	coors := t.findCoordinates(t.pivot, t.direction)
	shifted := t.shiftCoordinates(coors, Down)

	for t.board.FreeSpace(shifted) {
		t.pivot = t.pivot.Down(1)
		shifted = t.shiftCoordinates(shifted, Down)
	}

	t.board.LockInPlace(t)
}

func (t *Tetromino) RotateClockwise() {
	t.updateDirection(t.direction.Clockwise())
}

func (t *Tetromino) RotateAntiClockwise() {
	t.updateDirection(t.direction.AntiClockwise())
}

func (t *Tetromino) AutomaticDrop() bool {
	select {
	case <-t.downTicker:
		return true
	default:
	}

	return false
}

func (t *Tetromino) GhostCoordinates() []Coordinate {
	coors := t.Coordinates()
	shifted := t.shiftCoordinates(coors, Down)

	for t.board.FreeSpace(shifted) {
		coors = shifted
		shifted = t.shiftCoordinates(coors, Down)
	}

	return coors
}

func (t *Tetromino) updateDirection(direction Direction) {
	coors := t.findCoordinates(t.pivot, direction)

	if t.board.FreeSpace(coors) {
		t.direction = direction
		return
	}

	if amount, ok := t.freeSpaceWithShift(coors, Right); ok {
		t.direction = direction
		t.pivot = t.pivot.Right(amount)
		return
	}

	if amount, ok := t.freeSpaceWithShift(coors, Left); ok {
		t.direction = direction
		t.pivot = t.pivot.Left(amount)
	}
}
func (t *Tetromino) freeSpaceWithShift(coors []Coordinate, direction Direction) (int, bool) {
	shifted := t.shiftCoordinates(coors, direction)
	amount := 1

	// There may be 2 blocks extending from the pivot in the case of "I", so need to try shifting up to 2 blocks away
	if !t.board.InSideBounds(shifted) {
		shifted = t.shiftCoordinates(shifted, direction)
		amount++
	}

	if t.board.FreeSpace(shifted) {
		return amount, true
	}

	return 0, false
}

func (t *Tetromino) collision(direction Direction) bool {
	coors := t.findCoordinates(t.pivot, t.direction)
	shifted := t.shiftCoordinates(coors, direction)

	if t.board.FreeSpace(shifted) {
		return false
	}

	return true
}

func (t *Tetromino) shiftCoordinates(coordinates []Coordinate, shift Direction) []Coordinate {
	shifted := make([]Coordinate, len(coordinates))

	switch shift {
	case Left:
		for i := range coordinates {
			shifted[i] = coordinates[i].Left(1)
		}
	case Right:
		for i := range coordinates {
			shifted[i] = coordinates[i].Right(1)
		}
	case Down:
		for i := range coordinates {
			shifted[i] = coordinates[i].Down(1)
		}
	}

	return shifted
}
