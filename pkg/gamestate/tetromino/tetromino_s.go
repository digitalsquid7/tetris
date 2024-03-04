package tetromino

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/direction"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

func NewS() *Tetromino {
	return New(
		tetrissprites.YellowBlock,
		tetrissprites.YellowTetromino,
		coordinate.New(4, 1),
		FindSCoordinates,
	)
}

func FindSCoordinates(pivot coordinate.Coordinate, dir direction.Direction) []coordinate.Coordinate {
	var coors []coordinate.Coordinate

	switch dir {
	case direction.Up:
		coors = []coordinate.Coordinate{pivot, pivot.Up(1), pivot.Up(1).Right(1), pivot.Left(1)}
	case direction.Right:
		coors = []coordinate.Coordinate{pivot, pivot.Up(1), pivot.Right(1), pivot.Down(1).Right(1)}
	case direction.Down:
		coors = []coordinate.Coordinate{pivot, pivot.Right(1), pivot.Down(1), pivot.Down(1).Left(1)}
	case direction.Left:
		coors = []coordinate.Coordinate{pivot, pivot.Down(1), pivot.Left(1), pivot.Left(1).Up(1)}
	}

	return coors
}
