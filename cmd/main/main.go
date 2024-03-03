package main

import (
	"log/slog"
	"os"

	"github.com/digitalsquid7/tetris/pkg/tetrisrun"

	"github.com/gopxl/pixel/pixelgl"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	runner := tetrisrun.NewRunner(logger)

	pixelgl.Run(runner.Run)
}
