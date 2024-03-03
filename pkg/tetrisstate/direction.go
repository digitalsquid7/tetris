package tetrisstate

var clockwise = map[Direction]Direction{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

var antiClockwise = map[Direction]Direction{
	Up:    Left,
	Left:  Down,
	Down:  Right,
	Right: Up,
}

type Direction string

const (
	Up    Direction = "up"
	Right Direction = "right"
	Down  Direction = "down"
	Left  Direction = "left"
)

func (d Direction) Clockwise() Direction {
	return clockwise[d]
}

func (d Direction) AntiClockwise() Direction {
	return antiClockwise[d]
}
