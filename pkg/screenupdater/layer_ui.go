package screenupdater

import (
	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type UILayer struct {
	state   *gamestate.GameState
	sprites tetrissprites.Sprites
}

func NewUILayer(state *gamestate.GameState, sprites tetrissprites.Sprites) *UILayer {
	return &UILayer{
		state:   state,
		sprites: sprites,
	}
}

func (b *UILayer) Update(window *pixelgl.Window) {
	b.sprites[tetrissprites.Border].Draw(window, pixel.IM.Moved(pixel.V(196, 356)))
	b.sprites[tetrissprites.Next].Draw(window, pixel.IM.Moved(pixel.V(488, 596)))
	b.sprites[tetrissprites.Hold].Draw(window, pixel.IM.Moved(pixel.V(488, 384)))

	b.sprites[b.state.NextTetromino().TetrominoSprite()].Draw(window, pixel.IM.Moved(pixel.V(488, 560)))

	if b.state.HeldTetromino() != nil {
		b.sprites[b.state.HeldTetromino().TetrominoSprite()].Draw(window, pixel.IM.Moved(pixel.V(488, 350)))
	}
}
