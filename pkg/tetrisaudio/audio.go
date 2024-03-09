package tetrisaudio

import (
	"fmt"
	"os"

	"github.com/digitalsquid7/tetris/pkg/gameevent"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var fileNameByEventName = map[gameevent.Name]string{
	gameevent.Rotate:        "assets/sounds/rotation.mp3",
	gameevent.Drop:          "assets/sounds/drop.mp3",
	gameevent.GameOver:      "assets/sounds/game_over.mp3",
	gameevent.HoldTetromino: "assets/sounds/hold.mp3",
	gameevent.LineClear:     "assets/sounds/line_clear.mp3",
}

type Audio struct {
	sampleRate        beep.SampleRate
	bufferByEventName map[gameevent.Name]*beep.Buffer
}

func (a *Audio) SampleRate() beep.SampleRate {
	return a.sampleRate
}

func (a *Audio) PlaySound(name gameevent.Name) error {
	buffer, ok := a.bufferByEventName[name]
	if !ok {
		return fmt.Errorf("can't find buffer for event name: %s", name)
	}

	sound := buffer.Streamer(0, buffer.Len())
	speaker.Play(sound)
	return nil
}

func Load() (*Audio, error) {
	bufferByEventName := make(map[gameevent.Name]*beep.Buffer)
	var format beep.Format
	var streamer beep.Streamer

	for eventName, fileName := range fileNameByEventName {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("open file %s: %w", fileName, err)
		}

		streamer, format, err = mp3.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("mp3 decode file %s: %w", fileName, err)
		}

		buffer := beep.NewBuffer(format)
		buffer.Append(streamer)

		bufferByEventName[eventName] = buffer
	}

	return &Audio{
		sampleRate:        format.SampleRate,
		bufferByEventName: bufferByEventName,
	}, nil
}
