package day12

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"slices"
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
		program: slices.Clone(program),
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
		c.Set(ops[2], c.Get(ops[1]))
	case "inc":
		if v, ok := c.regs[ops[1]]; ok {
			c.regs[ops[1]] = v + 1
		}
	case "dec":
		if v, ok := c.regs[ops[1]]; ok {
			c.regs[ops[1]] = v - 1
		}
	case "jnz":
		if c.Get(ops[1]) != 0 {
			c.pos += c.Get(ops[2]) - 1
		}
	case "tgl":
		pos := c.pos + c.Get(ops[1])
		if pos >= 0 && pos < len(c.program) {
			inst := strings.Split(c.program[pos], " ")
			switch len(inst) {
			case 2:
				if inst[0] == "inc" {
					c.program[pos] = fmt.Sprintf("dec %v", inst[1])
				} else {
					c.program[pos] = fmt.Sprintf("inc %v", inst[1])
				}
			case 3:
				if inst[0] == "jnz" {
					c.program[pos] = fmt.Sprintf("cpy %v %v", inst[1], inst[2])
				} else {
					c.program[pos] = fmt.Sprintf("jnz %v %v", inst[1], inst[2])
				}
			}
		}
	// вспомогательные для избавления от циклов
	case "add":
		c.Set(ops[1], c.Get(ops[1]) + c.Get(ops[2]))
	case "mul":
		c.Set(ops[1], c.Get(ops[2]) * c.Get(ops[3]))
	}

	c.pos++
}

func (c *Computer) Get(s string) int {
	if _, ok := c.regs[s]; ok {
		return c.regs[s]
	}
	return int(h.ToInt(s))
}

func(c *Computer) Set(s string, v int) {
	if _, ok := c.regs[s]; ok {
		c.regs[s] = v
	}
}

func PartOne(lines []string) string {
	comp := NewComputer(lines)
	comp.Run()
	return h.ToString(comp.regs["a"])
}

func PartTwo(lines []string) string {
	comp := NewComputer(lines)
	comp.regs["c"] = 12
	comp.Run()
	return h.ToString(comp.regs["a"])
}
