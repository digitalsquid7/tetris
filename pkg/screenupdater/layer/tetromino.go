package layer

import (
	"image/color"

	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type Tetromino struct {
	state   *gamestate.GameState
	sprites tetrissprites.Sprites
}

func NewTetromino(state *gamestate.GameState, sprites tetrissprites.Sprites) *Tetromino {
	return &Tetromino{
		state:   state,
		sprites: sprites,
	}
}

func (b *Tetromino) Update(window *pixelgl.Window) {
	sprite := b.state.CurrentTetromino().BlockSprite()

	for _, coor := range b.state.GhostCoordinates() {
		if coor.Y() < 0 {
			continue
		}

		x := float64(blockOffset + (coor.X() * blockSize))
		y := float64(screenHeight - (blockOffset + (coor.Y() * blockSize)))
		b.sprites[sprite].DrawColorMask(window, pixel.IM.Moved(pixel.V(x, y)), color.Alpha{64})
	}

	for _, coor := range b.state.CurrentTetromino().Coordinates() {
		x := float64(blockOffset + (coor.X() * blockSize))
		y := float64(screenHeight - (blockOffset + (coor.Y() * blockSize)))
		b.sprites[sprite].Draw(window, pixel.IM.Moved(pixel.V(x, y)))
	}
}
