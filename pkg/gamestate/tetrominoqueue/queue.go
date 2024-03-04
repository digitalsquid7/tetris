package tetrominoqueue

import (
	"math/rand"

	"github.com/digitalsquid7/tetris/pkg/gamestate/board"
	"github.com/digitalsquid7/tetris/pkg/gamestate/tetromino"
)

type TetrominoQueue struct {
	queue []*tetromino.Tetromino
	board *board.Board
}

func New(board *board.Board) *TetrominoQueue {
	return &TetrominoQueue{
		board: board,
	}
}

func (t *TetrominoQueue) Len() int {
	return len(t.queue)
}

func (t *TetrominoQueue) Pop() *tetromino.Tetromino {
	item := t.queue[0]
	t.queue = t.queue[1:]
	return item
}

func (t *TetrominoQueue) Peek() *tetromino.Tetromino {
	return t.queue[0]
}

func (t *TetrominoQueue) GenerateTetrominos() []*tetromino.Tetromino {
	randomTetrominos := make([]*tetromino.Tetromino, 0, 7)

	tetrominos := []*tetromino.Tetromino{
		tetromino.NewI(),
		tetromino.NewJ(),
		tetromino.NewL(),
		tetromino.NewO(),
		tetromino.NewS(),
		tetromino.NewT(),
		tetromino.NewZ(),
	}

	indexes := rand.Perm(7)

	for _, i := range indexes {
		t.queue = append(t.queue, tetrominos[i])
	}

	return randomTetrominos
}
