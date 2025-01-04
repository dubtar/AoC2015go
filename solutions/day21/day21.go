package day21

import (
	"fmt"
	"strings"

	h "go-aoc-template/internal/helpers"
)

type Item struct {
	cost   int
	damage int
	armor  int
}
type Person struct {
	Item
	hitpoint int
}

func (p *Person) Add(item Item) {
	p.armor += item.armor
	p.cost += item.cost
	p.damage += item.damage
}

func (p *Person) Remove(item Item) {
	p.armor -= item.armor
	p.cost -= item.cost
	p.damage -= item.damage
}

var weapons = []Item{{8, 4, 0}, {10, 5, 0}, {25, 6, 0}, {40, 7, 0}, {74, 8, 0}}
var armors = []Item{{0, 0, 0}, {13, 0, 1}, {31, 0, 2}, {53, 0, 3}, {75, 0, 4}, {102, 0, 5}}
var rings = []Item{{20, 0, 1}, {25, 1, 0}, {40, 0, 2}, {50, 2, 0}, {80, 0, 3}, {100, 3, 0}}

func fight(person Person, boss Person) bool {
	for {
		boss.hitpoint -= max(1, person.damage-boss.armor)
		if boss.hitpoint <= 0 {
			return true
		}
		person.hitpoint -= max(1, boss.damage-person.armor)
		if person.hitpoint <= 0 {
			return false
		}
	}
}

var mostCost int = 0

func PartOne(lines []string) string {
	boss := Person{}
	for _, line := range lines {
		if strings.HasPrefix(line, "Hit Points:") {
			boss.hitpoint = int(h.ToInt(line[len("Hit Points: "):]))
		} else if strings.HasPrefix(line, "Damage: ") {
			boss.damage = int(h.ToInt(line[len("Damage: "):]))
		} else if strings.HasPrefix(line, "Armor:") {
			boss.armor = int(h.ToInt(line[len("Armor: "):]))
		}
	}
	bestCost := 1000
	person := Person{hitpoint: 100}
	for _, weapon := range weapons {
		person.Add(weapon)
		for _, armor := range armors {
			person.Add(armor)
			if fight(person, boss) {
				// fmt.Println("Win with weapon", weapon.cost, "and armor", armor.cost)
				bestCost = min(bestCost, person.cost)
			} else {
				// fmt.Println("Loose with weapon", weapon.cost, "and armor", armor.cost)
				mostCost = max(mostCost, person.cost)
			}
			for r1, ring := range rings {
				person.Add(ring)
				if fight(person, boss) {
					// fmt.Println("Win with weapon", weapon.cost, ", armor", armor.cost, "and ring", ring.cost)
					bestCost = min(bestCost, person.cost)
				} else {
					// fmt.Println("Loose with weapon", weapon.cost, ", armor", armor.cost, "and ring", ring.cost)
					mostCost = max(mostCost, person.cost)
				}
				for r2 := r1 + 1; r2 < len(rings); r2++ {
					ring2 := rings[r2]
					person.Add(ring2)
					if fight(person, boss) {
						// fmt.Println("Win with weapon", weapon.cost, ", armor", armor.cost, "and rings", ring.cost, ring2.cost)
						bestCost = min(bestCost, person.cost)
					} else {
						// fmt.Println("Loose with weapon", weapon.cost, ", armor", armor.cost, "and rinsg", ring.cost, ring2.cost)
						mostCost = max(mostCost, person.cost)
					}
					person.Remove(ring2)
				}
				person.Remove(ring)
			}
			person.Remove(armor)
		}
		person.Remove(weapon)

	}
	return fmt.Sprint(bestCost)
}

func PartTwo(lines []string) string {
	return fmt.Sprint(mostCost)
}
