package tetrisstate

type GameState struct {
	Board                *Board
	CurrentTetromino     *Tetromino
	NextTetrominos       *TetrominoQueue
	HeldTetromino        *Tetromino
	HeldTetrominoSwapped bool
}

func New() *GameState {
	board := NewBoard()
	nextTetrominos := NewTetrominoQueue(board)
	nextTetrominos.GenerateTetrominos()

	return &GameState{
		Board:            board,
		CurrentTetromino: nextTetrominos.Pop(),
		NextTetrominos:   nextTetrominos,
	}
}

func (g *GameState) ReplaceTetromino() {
	g.newTetromino()
	g.Board.ClearLines()
	g.HeldTetrominoSwapped = false
}

func (g *GameState) NextTetromino() *Tetromino {
	return g.NextTetrominos.Peek()
}

func (g *GameState) HoldTetromino() {
	if g.HeldTetrominoSwapped {
		return
	}

	if g.HeldTetromino == nil {
		g.HeldTetromino = g.CurrentTetromino
		g.newTetromino()
	} else {
		g.HeldTetromino, g.CurrentTetromino = g.CurrentTetromino, g.HeldTetromino
		g.CurrentTetromino.pivot = g.CurrentTetromino.defaultPivot
	}

	g.HeldTetrominoSwapped = true
}

func (g *GameState) newTetromino() {
	g.CurrentTetromino = g.NextTetrominos.Pop()

	if g.NextTetrominos.Len() == 1 {
		g.NextTetrominos.GenerateTetrominos()
	}
}
