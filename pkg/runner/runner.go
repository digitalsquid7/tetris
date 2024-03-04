package runner

import (
	"fmt"
	_ "image/png"
	"log/slog"

	"github.com/digitalsquid7/tetris/pkg/command"
	"github.com/digitalsquid7/tetris/pkg/config"
	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/screenupdater"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/gopxl/pixel/pixelgl"
)

type Runner struct {
	logger          *slog.Logger
	state           *gamestate.GameState
	screenUpdater   *screenupdater.ScreenUpdater
	window          *pixelgl.Window
	commandExecutor *command.Executor
}

func New(logger *slog.Logger) *Runner {
	return &Runner{logger: logger}
}

func (r *Runner) Run() {
	if err := r.gameLoop(); err != nil {
		r.logger.Error(err.Error())
	}
	return
}

func (r *Runner) gameLoop() error {
	window, err := pixelgl.NewWindow(config.WindowConfig)
	if err != nil {
		return fmt.Errorf("create new window: %w", err)
	}

	spritesLoader := tetrissprites.NewLoader(r.logger)

	sprites, err := spritesLoader.Load()
	if err != nil {
		return fmt.Errorf("load sprites: %w", err)
	}

	gameState := gamestate.New()
	commandExecutor := command.NewExecutor(gameState, window)
	screenUpdater := screenupdater.New(gameState, window, sprites)

	for !window.Closed() {
		commandExecutor.Execute()
		screenUpdater.Update()
	}

	return nil
}
