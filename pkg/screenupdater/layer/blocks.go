package layer

import (
	"image/color"

	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

const (
	blockOffset  = 52
	blockSize    = 32
	screenHeight = 712
)

type Blocks struct {
	state   *gamestate.GameState
	sprites tetrissprites.Sprites
}

func NewBlocks(state *gamestate.GameState, sprites tetrissprites.Sprites) *Blocks {
	return &Blocks{
		state:   state,
		sprites: sprites,
	}
}

func (b *Blocks) Update(window *pixelgl.Window) {
	for row := range b.state.Board().Blocks {
		for col := range b.state.Board().Blocks[row] {
			block := b.state.Board().Blocks[row][col]
			if block == nil {
				continue
			}

			x := float64(blockOffset + (col * blockSize))
			y := float64(screenHeight - (blockOffset + (row * blockSize)))

			if block.Opacity == 100 {
				b.sprites[(*block).Sprite].Draw(window, pixel.IM.Moved(pixel.V(x, y)))
				continue
			}

			percentage := float64(block.Opacity) / 100
			alpha := uint8(percentage * 256)
			b.sprites[(*block).Sprite].DrawColorMask(window, pixel.IM.Moved(pixel.V(x, y)), color.Alpha{alpha})
		}
	}
}
