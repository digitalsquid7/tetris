package tetrisscreen

import (
	"image/color"

	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/digitalsquid7/tetris/pkg/tetrisstate"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	blockOffset  = 52
	blockSize    = 32
	screenHeight = 712
)

type Screen struct {
	window  *pixelgl.Window
	state   *tetrisstate.GameState
	sprites tetrissprites.Sprites
}

func New(
	window *pixelgl.Window,
	state *tetrisstate.GameState,
	sprites tetrissprites.Sprites) *Screen {
	return &Screen{
		window:  window,
		state:   state,
		sprites: sprites,
	}
}

func (l *Screen) Update() {
	l.window.Clear(colornames.Black)

	l.sprites[tetrissprites.Border].Draw(l.window, pixel.IM.Moved(pixel.V(196, 356)))
	l.sprites[tetrissprites.Next].Draw(l.window, pixel.IM.Moved(pixel.V(488, 596)))
	l.sprites[tetrissprites.Hold].Draw(l.window, pixel.IM.Moved(pixel.V(488, 384)))

	nextSprite := l.state.NextTetromino().TetrominoSprite()
	l.sprites[nextSprite].Draw(l.window, pixel.IM.Moved(pixel.V(488, 560)))

	if l.state.HeldTetromino != nil {
		holdSprite := l.state.HeldTetromino.TetrominoSprite()
		l.sprites[holdSprite].Draw(l.window, pixel.IM.Moved(pixel.V(488, 350)))
	}

	sprite := l.state.CurrentTetromino.BlockSprite()

	for _, coor := range l.state.CurrentTetromino.GhostCoordinates() {
		x := float64(blockOffset + (coor.X() * blockSize))
		y := float64(screenHeight - (blockOffset + (coor.Y() * blockSize)))
		l.sprites[sprite].DrawColorMask(l.window, pixel.IM.Moved(pixel.V(x, y)), color.Alpha{64})
	}

	for _, coor := range l.state.CurrentTetromino.Coordinates() {
		x := float64(blockOffset + (coor.X() * blockSize))
		y := float64(screenHeight - (blockOffset + (coor.Y() * blockSize)))
		l.sprites[sprite].Draw(l.window, pixel.IM.Moved(pixel.V(x, y)))
	}

	for row := range l.state.Board.Blocks {
		for col := range l.state.Board.Blocks[row] {
			block := l.state.Board.Blocks[row][col]
			if block == nil {
				continue
			}

			x := float64(blockOffset + (col * blockSize))
			y := float64(screenHeight - (blockOffset + (row * blockSize)))

			if block.Opacity == 100 {
				l.sprites[(*block).Sprite].Draw(l.window, pixel.IM.Moved(pixel.V(x, y)))
				continue
			}

			percentage := float64(block.Opacity) / 100
			alpha := uint8(percentage * 256)
			l.sprites[(*block).Sprite].DrawColorMask(l.window, pixel.IM.Moved(pixel.V(x, y)), color.Alpha{alpha})
		}
	}

	l.window.Update()
}
