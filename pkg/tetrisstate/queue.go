package tetrisstate

import "math/rand"

type TetrominoQueue struct {
	queue []*Tetromino
	board *Board
}

func NewTetrominoQueue(board *Board) *TetrominoQueue {
	return &TetrominoQueue{
		board: board,
	}
}

func (t *TetrominoQueue) Len() int {
	return len(t.queue)
}

func (t *TetrominoQueue) Pop() *Tetromino {
	item := t.queue[0]
	t.queue = t.queue[1:]
	return item
}

func (t *TetrominoQueue) Peek() *Tetromino {
	return t.queue[0]
}

func (t *TetrominoQueue) GenerateTetrominos() []*Tetromino {
	randomTetrominos := make([]*Tetromino, 0, 7)
	tetrominos := []*Tetromino{
		NewI(t.board),
		NewJ(t.board),
		NewL(t.board),
		NewO(t.board),
		NewS(t.board),
		NewT(t.board),
		NewZ(t.board),
	}

	indexes := rand.Perm(7)

	for _, i := range indexes {
		t.queue = append(t.queue, tetrominos[i])
	}

	return randomTetrominos
}
