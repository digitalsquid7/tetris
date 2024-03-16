package runner

import (
	"fmt"
	_ "image/png"
	"log/slog"
	"time"

	"github.com/digitalsquid7/tetris/pkg/audioupdater"
	"github.com/digitalsquid7/tetris/pkg/command"
	"github.com/digitalsquid7/tetris/pkg/config"
	"github.com/digitalsquid7/tetris/pkg/gameevent"
	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/screenupdater"
	"github.com/digitalsquid7/tetris/pkg/tetrisaudio"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/faiface/beep/speaker"
	"github.com/gopxl/pixel/pixelgl"
)

type Runner struct {
	logger *slog.Logger
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
	tetrisConfig := config.LoadTetris()
	windowConfig := config.LoadWindow()

	window, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		return fmt.Errorf("create new window: %w", err)
	}

	sprites, err := tetrissprites.Load()
	if err != nil {
		return fmt.Errorf("load sprites: %w", err)
	}

	audio, err := tetrisaudio.Load()
	if err != nil {
		return fmt.Errorf("load audio: %w", err)
	}

	err = speaker.Init(audio.SampleRate(), audio.SampleRate().N(time.Second/10))
	if err != nil {
		return fmt.Errorf("init audio: %w", err)
	}

	publisher := gameevent.NewPublisher()
	state := gamestate.New(publisher)
	commandExecutor := command.NewExecutor(state, window)
	screenUpdater := screenupdater.New(state, window, sprites)
	publisher.Subscribe(screenUpdater)

	if tetrisConfig.EnableSounds {
		audioUpdater := audioupdater.New(audio)
		publisher.Subscribe(audioUpdater)
	}

	for !window.Closed() {
		commandExecutor.Execute()

		err = publisher.Notify()
		if err != nil {
			return fmt.Errorf("notify subscribers: %w", err)
		}
	}

	return nil
}
