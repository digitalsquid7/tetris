package screenupdater

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/screenupdater/layer"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"

	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Layer interface {
	Update(window *pixelgl.Window)
}

type ScreenUpdater struct {
	state        *gamestate.GameState
	window       *pixelgl.Window
	sprites      tetrissprites.Sprites
	tetrisLayers []Layer
}

func New(
	state *gamestate.GameState,
	window *pixelgl.Window,
	sprites tetrissprites.Sprites) *ScreenUpdater {
	return &ScreenUpdater{
		state:   state,
		window:  window,
		sprites: sprites,
		tetrisLayers: []Layer{
			layer.NewUI(state, sprites),
			layer.NewBlocks(state, sprites),
			layer.NewTetromino(state, sprites),
		},
	}
}

func (l *ScreenUpdater) Update() {
	l.window.Clear(colornames.Black)

	for _, layer := range l.tetrisLayers {
		layer.Update(l.window)
	}

	l.window.Update()
}
