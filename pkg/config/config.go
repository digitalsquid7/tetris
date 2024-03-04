package config

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

var WindowConfig = pixelgl.WindowConfig{
	Title:  "Tetris",
	Bounds: pixel.R(0, 0, 604, 712),
	VSync:  true,
}
