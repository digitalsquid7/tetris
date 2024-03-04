package main

import (
	"log/slog"
	"os"

	"github.com/digitalsquid7/tetris/pkg/runner"
	"github.com/gopxl/pixel/pixelgl"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)

	tetrisRunner := runner.New(logger)

	pixelgl.Run(tetrisRunner.Run)
}
