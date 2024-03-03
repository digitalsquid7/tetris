package tetrisstate

import (
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

func findTCoordinates(pivot Coordinate, direction Direction) []Coordinate {
	var coors []Coordinate

	switch direction {
	case Up:
		coors = []Coordinate{pivot, pivot.Up(1), pivot.Right(1), pivot.Left(1)}
	case Right:
		coors = []Coordinate{pivot, pivot.Up(1), pivot.Right(1), pivot.Down(1)}
	case Down:
		coors = []Coordinate{pivot, pivot.Left(1), pivot.Right(1), pivot.Down(1)}
	case Left:
		coors = []Coordinate{pivot, pivot.Left(1), pivot.Up(1), pivot.Down(1)}
	}

	return coors
}

func NewT(board *Board) *Tetromino {
	return NewTetromino(
		board,
		tetrissprites.BlueBlock,
		tetrissprites.BlueTetromino,
		NewCoordinate(5, 1),
		findTCoordinates,
	)
}
