package tetrisstate

import "github.com/digitalsquid7/tetris/pkg/tetrissprites"

func findICoordinates(pivot Coordinate, direction Direction) []Coordinate {
	var coors []Coordinate

	switch direction {
	case Up:
		coors = []Coordinate{
			pivot.Left(1),
			pivot,
			pivot.Right(1),
			pivot.Right(2),
		}
	case Right:
		coors = []Coordinate{
			pivot.Right(1),
			pivot.Right(1).Up(1),
			pivot.Right(1).Down(1),
			pivot.Right(1).Down(2),
		}
	case Down:
		coors = []Coordinate{
			pivot.Down(1).Left(1),
			pivot.Down(1),
			pivot.Down(1).Right(1),
			pivot.Down(1).Right(2),
		}
	case Left:
		coors = []Coordinate{
			pivot.Up(1),
			pivot,
			pivot.Down(1),
			pivot.Down(2),
		}
	}

	return coors
}

func NewI(board *Board) *Tetromino {
	return NewTetromino(
		board,
		tetrissprites.PurpleBlock,
		tetrissprites.PurpleTetromino,
		NewCoordinate(4, 0),
		findICoordinates,
	)
}
