package helpers

type Coords struct {
	X int64
	Y int64
}

func (c Coords) Mul(steps int64) Coords {
	return Coords{X: c.X * steps, Y: c.Y * steps}
}

func (c *Coords) Add(other Coords) {
	c.X += other.X
	c.Y += other.Y
}

var DirectionsHV = map[rune]Coords{'^': {X: -1, Y: 0}, '>': {X: 0, Y: 1}, 'v': {X: 1, Y: 0}, '<': {X: 0, Y: -1}}
var DirectionsUDLR = map[rune]Coords{
	'U': {X: 0, Y: -1},
	'D': {X: 0, Y: 1},
	'L': {X: -1, Y: 0},
	'R': {X: 1, Y: 0}}

func TurnLeft(direction rune) rune {
	switch direction {
	case '^':
		return '<'
	case '>':
		return '^'
	case 'v':
		return '>'
	case '<':
		return 'v'
	default:
		panic("Invalid direction")
	}
}

func TurnRight(direction rune) rune {
	switch direction {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	default:
		panic("Invalid direction")
	}
}
