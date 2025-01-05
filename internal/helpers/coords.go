package helpers

type Coords struct {
	X int
	Y int
}

func (c *Coords) Add(other Coords) {
	c.X += other.X
	c.Y += other.Y
}

var DirectionsHV = map[rune]Coords{'^': {X: -1, Y: 0}, '>': {X: 0, Y: 1}, 'v': {X: 1, Y: 0}, '<': {X: 0, Y: -1}}
