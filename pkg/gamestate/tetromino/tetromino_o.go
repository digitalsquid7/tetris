package tetromino

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/direction"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

func NewO() *Tetromino {
	return New(
		tetrissprites.RedBlock,
		tetrissprites.RedTetromino,
		coordinate.New(5, 0),
		FindOCoordinates,
	)
}

func FindOCoordinates(pivot coordinate.Coordinate, _ direction.Direction) []coordinate.Coordinate {
	return []coordinate.Coordinate{pivot, pivot.Down(1), pivot.Left(1), pivot.Down(1).Left(1)}
}
