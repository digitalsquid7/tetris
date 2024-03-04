package command

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/gopxl/pixel/pixelgl"
)

type Executor struct {
	state  *gamestate.GameState
	window *pixelgl.Window
}

func NewExecutor(state *gamestate.GameState, window *pixelgl.Window) *Executor {
	return &Executor{
		state:  state,
		window: window,
	}
}

func (e *Executor) Execute() {
	if e.state.GameOver() {
		return
	}

	if e.state.Board().ClearLinesInProgress() {
		e.state.Board().UpdateFullRows()
		return
	}

	if e.state.TopOut() {
		return
	}

	if e.state.AutomaticDrop() {
		if e.state.MoveDown() {
			return
		}
	}

	if e.window.JustPressed(pixelgl.KeySpace) {
		e.state.HardDrop()
		return
	}

	if e.window.Pressed(pixelgl.KeyS) {
		e.state.SoftDrop()
		return
	}

	if e.window.JustPressed(pixelgl.KeyLeftShift) {
		e.state.HoldTetromino()
		return
	}

	e.move()
	e.rotate()
}

func (e *Executor) move() {
	if e.window.JustPressed(pixelgl.KeyA) {
		e.state.MoveLeft(true)
		return
	}

	if e.window.Pressed(pixelgl.KeyA) {
		e.state.MoveLeft(false)
		return
	}

	if e.window.JustPressed(pixelgl.KeyD) {
		e.state.MoveRight(true)
		return
	}

	if e.window.Pressed(pixelgl.KeyD) {
		e.state.MoveRight(false)
		return
	}
}

func (e *Executor) rotate() {
	if e.window.JustPressed(pixelgl.KeyQ) {
		e.state.RotateAntiClockwise()
		return
	}

	if e.window.JustPressed(pixelgl.KeyE) {
		e.state.RotateClockwise()
		return
	}
}
