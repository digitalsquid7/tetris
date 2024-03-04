package tetromino

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/direction"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

func NewI() *Tetromino {
	return New(
		tetrissprites.PurpleBlock,
		tetrissprites.PurpleTetromino,
		coordinate.New(4, 0),
		FindICoordinates,
	)
}

func FindICoordinates(pivot coordinate.Coordinate, dir direction.Direction) []coordinate.Coordinate {
	var coors []coordinate.Coordinate

	switch dir {
	case direction.Up:
		coors = []coordinate.Coordinate{
			pivot.Left(1),
			pivot,
			pivot.Right(1),
			pivot.Right(2),
		}
	case direction.Right:
		coors = []coordinate.Coordinate{
			pivot.Right(1),
			pivot.Right(1).Up(1),
			pivot.Right(1).Down(1),
			pivot.Right(1).Down(2),
		}
	case direction.Down:
		coors = []coordinate.Coordinate{
			pivot.Down(1).Left(1),
			pivot.Down(1),
			pivot.Down(1).Right(1),
			pivot.Down(1).Right(2),
		}
	case direction.Left:
		coors = []coordinate.Coordinate{
			pivot.Up(1),
			pivot,
			pivot.Down(1),
			pivot.Down(2),
		}
	}

	return coors
}
