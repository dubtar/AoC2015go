package day12

import (
	h "go-aoc-template/internal/helpers"
	"strings"
)

type Computer struct {
	regs    map[string]int
	pos     int
	program []string
}

func NewComputer(program []string) Computer {
	return Computer{
		regs: map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
			"d": 0,
		},
		pos:     0,
		program: program,
	}
}

func (c *Computer) Run() {
	for c.pos < len(c.program) {
		c.Step()
	}
}

func (c *Computer) Step() {
	ops := strings.Split(c.program[c.pos], " ")
	switch ops[0] {
	case "cpy":
		c.regs[ops[2]] = c.Get(ops[1])
	case "inc":
		c.regs[ops[1]]++
	case "dec":
		c.regs[ops[1]]--
	case "jnz":
		if c.Get(ops[1]) != 0 {
			c.pos += int(h.ToInt(ops[2])) - 1
		}
	}
	c.pos++
}

func (c *Computer) Get(s string) int {
	if _, ok := c.regs[s]; ok {
		return c.regs[s]
	}
	return int(h.ToInt(s))
}

func PartOne(lines []string) string {
	comp := NewComputer(lines)
	comp.Run()
	return h.ToString(comp.regs["a"])
}

func PartTwo(lines []string) string {
	comp := NewComputer(lines)
	comp.regs["c"] = 1
	comp.Run()
	return h.ToString(comp.regs["a"])
}
