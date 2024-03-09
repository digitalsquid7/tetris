package audioupdater

import (
	"fmt"

	"github.com/digitalsquid7/tetris/pkg/gameevent"
	"github.com/digitalsquid7/tetris/pkg/tetrisaudio"
)

type AudioUpdater struct {
	audio      *tetrisaudio.Audio
	eventNames []gameevent.Name
}

func New(audio *tetrisaudio.Audio) *AudioUpdater {
	return &AudioUpdater{
		audio: audio,
		eventNames: []gameevent.Name{
			gameevent.LineClear,
			gameevent.HoldTetromino,
			gameevent.Rotate,
			gameevent.GameOver,
			gameevent.Drop,
		},
	}
}

func (a *AudioUpdater) Update(events gameevent.Events) error {
	for _, eventName := range a.eventNames {
		if !events.Includes(eventName) {
			continue
		}

		if err := a.audio.PlaySound(eventName); err != nil {
			return fmt.Errorf("play %s: %w", eventName, err)
		}
	}

	return nil
}
