package layer

import (
	"fmt"

	"github.com/digitalsquid7/tetris/pkg/gamestate"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/gopxl/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Score struct {
	state   *gamestate.GameState
	sprites tetrissprites.Sprites
}

func NewScore(state *gamestate.GameState, sprites tetrissprites.Sprites) *Score {
	return &Score{
		state:   state,
		sprites: sprites,
	}
}

func (l *Score) Update(window *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	level := text.New(pixel.V(400, 250), basicAtlas)
	_, _ = fmt.Fprintf(level, "Level: %d", l.state.Level())
	level.Draw(window, pixel.IM.Scaled(level.Orig, 2))

	score := text.New(pixel.V(400, 220), basicAtlas)
	_, _ = fmt.Fprintf(score, "Score: %d", l.state.Score())
	score.Draw(window, pixel.IM.Scaled(score.Orig, 2))
}
