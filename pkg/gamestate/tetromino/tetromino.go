package tetromino

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/direction"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

type FindCoordinates func(pivot coordinate.Coordinate, direction direction.Direction) []coordinate.Coordinate

type Tetromino struct {
	FindCoordinates FindCoordinates
	blockSprite     tetrissprites.Name
	tetrominoSprite tetrissprites.Name
	defaultPivot    coordinate.Coordinate
	pivot           coordinate.Coordinate
	direction       direction.Direction
}

func New(
	blockSprite tetrissprites.Name,
	tetrominoSprite tetrissprites.Name,
	pivot coordinate.Coordinate,
	findCoordinates FindCoordinates) *Tetromino {
	return &Tetromino{
		FindCoordinates: findCoordinates,
		blockSprite:     blockSprite,
		tetrominoSprite: tetrominoSprite,
		pivot:           pivot,
		defaultPivot:    pivot,
		direction:       direction.Up,
	}
}

func (t *Tetromino) BlockSprite() tetrissprites.Name {
	return t.blockSprite
}

func (t *Tetromino) TetrominoSprite() tetrissprites.Name {
	return t.tetrominoSprite
}

func (t *Tetromino) Coordinates() []coordinate.Coordinate {
	return t.FindCoordinates(t.pivot, t.direction)
}

func (t *Tetromino) Pivot() coordinate.Coordinate {
	return t.pivot
}

func (t *Tetromino) MoveRight(amount int) {
	t.pivot = t.pivot.Right(amount)
}

func (t *Tetromino) MoveLeft(amount int) {
	t.pivot = t.pivot.Left(amount)
}

func (t *Tetromino) MoveDown(amount int) {
	t.pivot = t.pivot.Down(amount)
}

func (t *Tetromino) MoveUp(amount int) {
	t.pivot = t.pivot.Up(amount)
}

func (t *Tetromino) RotateClockwise() {
	t.direction = t.direction.Clockwise()
}

func (t *Tetromino) RotateAntiClockwise() {
	t.direction = t.direction.AntiClockwise()
}

func (t *Tetromino) Direction() direction.Direction {
	return t.direction
}

func (t *Tetromino) SetDirection(direction direction.Direction) {
	t.direction = direction
}

func (t *Tetromino) ResetPosition() {
	t.pivot = t.defaultPivot
	t.direction = direction.Up
}
