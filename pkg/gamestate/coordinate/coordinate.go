package coordinate

type Coordinate [2]int

func New(x, y int) Coordinate {
	return Coordinate{x, y}
}

func (c Coordinate) X() int {
	return c[0]
}

func (c Coordinate) Y() int {
	return c[1]
}

func (c Coordinate) Up(amount int) Coordinate {
	return Coordinate{c.X(), c.Y() - amount}
}

func (c Coordinate) Right(amount int) Coordinate {
	return Coordinate{c.X() + amount, c.Y()}
}

func (c Coordinate) Down(amount int) Coordinate {
	return Coordinate{c.X(), c.Y() + amount}
}

func (c Coordinate) Left(amount int) Coordinate {
	return Coordinate{c.X() - amount, c.Y()}
}
