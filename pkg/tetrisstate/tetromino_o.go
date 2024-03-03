package tetrisstate

import "github.com/digitalsquid7/tetris/pkg/tetrissprites"

func findOCoordinates(pivot Coordinate, _ Direction) []Coordinate {
	return []Coordinate{pivot, pivot.Down(1), pivot.Left(1), pivot.Down(1).Left(1)}
}

func NewO(board *Board) *Tetromino {
	return NewTetromino(
		board,
		tetrissprites.RedBlock,
		tetrissprites.RedTetromino,
		NewCoordinate(5, 1),
		findOCoordinates,
	)
}
