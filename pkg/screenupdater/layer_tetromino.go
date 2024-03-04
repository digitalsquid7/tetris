package screenupdater

import (
	"image/color"

	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type TetrominoLayer struct {
	state   *gamestate.GameState
	sprites tetrissprites.Sprites
}

func NewTetrominoLayer(state *gamestate.GameState, sprites tetrissprites.Sprites) *TetrominoLayer {
	return &TetrominoLayer{
		state:   state,
		sprites: sprites,
	}
}

func (b *TetrominoLayer) Update(window *pixelgl.Window) {
	sprite := b.state.CurrentTetromino().BlockSprite()

	for _, coor := range b.state.GhostCoordinates() {
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
