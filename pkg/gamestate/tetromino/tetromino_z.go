package tetromino

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/direction"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

func NewZ() *Tetromino {
	return New(
		tetrissprites.CyanBlock,
		tetrissprites.CyanTetromino,
		coordinate.New(5, 1),
		FindZCoordinates,
	)
}

func FindZCoordinates(pivot coordinate.Coordinate, dir direction.Direction) []coordinate.Coordinate {
	var coors []coordinate.Coordinate

	switch dir {
	case direction.Up:
		coors = []coordinate.Coordinate{pivot, pivot.Up(1), pivot.Up(1).Left(1), pivot.Right(1)}
	case direction.Right:
		coors = []coordinate.Coordinate{pivot, pivot.Right(1), pivot.Down(1), pivot.Up(1).Right(1)}
	case direction.Down:
		coors = []coordinate.Coordinate{pivot, pivot.Left(1), pivot.Down(1), pivot.Down(1).Right(1)}
	case direction.Left:
		coors = []coordinate.Coordinate{pivot, pivot.Up(1), pivot.Left(1), pivot.Down(1).Left(1)}
	}

	return coors
}
