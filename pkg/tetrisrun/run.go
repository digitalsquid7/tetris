package tetrisrun

import (
	"fmt"
	_ "image/png"
	"log/slog"

	"github.com/digitalsquid7/tetris/pkg/tetrisscreen"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/digitalsquid7/tetris/pkg/tetrisstate"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type Runner struct {
	logger *slog.Logger
}

func NewRunner(logger *slog.Logger) *Runner {
	return &Runner{
		logger: logger,
	}
}

func (t *Runner) Run() {
	if err := t.runGameLoop(); err != nil {
		t.logger.Error(err.Error())
	}
	return
}

func (t *Runner) runGameLoop() error {
	cfg := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, 604, 712),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return fmt.Errorf("create new window: %w", err)
	}

	state := tetrisstate.New()
	spritesLoader := tetrissprites.NewLoader()

	sprites, err := spritesLoader.Load()
	if err != nil {
		return fmt.Errorf("load sprites: %w", err)
	}

	scrn := tetrisscreen.New(win, state, sprites)

	for !win.Closed() {
		t.updateState(state, win)
		scrn.Update()
	}

	return nil
}

func (t *Runner) updateState(state *tetrisstate.GameState, window *pixelgl.Window) {
	if state.Board.ClearLinesInProgress() {
		state.Board.UpdateFullRows()
		return
	}

	if state.CurrentTetromino.AutomaticDrop() {
		if collision := state.CurrentTetromino.MoveDown(); collision {
			state.ReplaceTetromino()
		}
	}

	if window.JustPressed(pixelgl.KeySpace) {
		state.CurrentTetromino.HardDrop()
		state.ReplaceTetromino()
		return
	}

	if window.Pressed(pixelgl.KeyS) {
		if collision := state.CurrentTetromino.SoftDrop(); collision {
			state.ReplaceTetromino()
		}
		return
	}

	if window.JustPressed(pixelgl.KeyLeftShift) {
		state.HoldTetromino()
		return
	}

	t.move(state, window)
	t.rotate(state, window)
}

func (t *Runner) move(state *tetrisstate.GameState, window *pixelgl.Window) {
	if window.JustPressed(pixelgl.KeyA) {
		state.CurrentTetromino.MoveLeft(true)
		return
	}

	if window.Pressed(pixelgl.KeyA) {
		if state.CurrentTetromino.Movable() {
			state.CurrentTetromino.MoveLeft(false)
		}
		return
	}

	if window.JustPressed(pixelgl.KeyD) {
		state.CurrentTetromino.MoveRight(true)
		return
	}

	if window.Pressed(pixelgl.KeyD) {
		if state.CurrentTetromino.Movable() {
			state.CurrentTetromino.MoveRight(false)
		}
		return
	}
}

func (t *Runner) rotate(state *tetrisstate.GameState, window *pixelgl.Window) {
	if window.JustPressed(pixelgl.KeyQ) {
		state.CurrentTetromino.RotateAntiClockwise()
		return
	}

	if window.JustPressed(pixelgl.KeyE) {
		state.CurrentTetromino.RotateClockwise()
		return
	}
}
