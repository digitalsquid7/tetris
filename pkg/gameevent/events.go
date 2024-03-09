package gameevent

type Name string

const (
	Rotate        Name = "Rotate"
	Drop          Name = "Drop"
	GameOver      Name = "GameOver"
	HoldTetromino Name = "HoldTetromono"
	LineClear     Name = "LineClear"
)

type Events struct {
	events map[Name]bool
}

func MakeEvents() Events {
	return Events{events: make(map[Name]bool)}
}

func (e *Events) Add(event Name) {
	e.events[event] = true
}

func (e *Events) Includes(event Name) bool {
	return e.events[event]
}

func (e *Events) Reset() {
	e.events = make(map[Name]bool)
}

func (e *Events) Len() int {
	return len(e.events)
}
