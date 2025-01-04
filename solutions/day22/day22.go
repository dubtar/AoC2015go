package day22

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"strings"

	pq "github.com/idsulik/go-collections/v2/priorityqueue"
)

type EffectFunc func(*Game)

type Effect struct {
	turnsLeft int
	apply     EffectFunc
	stop      EffectFunc
	name      string
}

type Game struct {
	personHitpoints int
	personArmor     int
	personMana      int
	bossHitpoints   int
	bossDamage      int
	extraDamage     int
	effects         []Effect
	manaSpent       int
	nextSpell       Spell
	log             []string
}

func (g *Game) ApplyEffects() {
	next_effects := make([]Effect, 0, len(g.effects))
	for _, effect := range g.effects {
		if effect.apply != nil {
			effect.apply(g)
		}
		effect.turnsLeft -= 1
		g.log[len(g.log)-1] += fmt.Sprintf(". %d left", effect.turnsLeft)
		if effect.turnsLeft <= 0 {
			if effect.stop != nil {
				effect.stop(g)
			}
		} else {
			next_effects = append(next_effects, effect)
		}
	}
	g.effects = next_effects
}

func castMissile(game *Game) {
	game.bossHitpoints -= 4
	game.log = append(game.log, fmt.Sprintf("Casts Missile, boss HP=%d", game.bossHitpoints))
}
func castDrain(game *Game) {
	game.bossHitpoints -= 2
	game.personHitpoints += 2
	game.log = append(game.log, fmt.Sprintf("Casts Drain, boss HP=%d player HP=%d", game.bossHitpoints, game.personHitpoints))
}
func stopShield(game *Game) {
	game.personArmor -= 7
	game.log = append(game.log, fmt.Sprintf("Shield stops, player armor=%d", game.personArmor))
}
func castShield(game *Game) {
	game.personArmor += 7
	game.effects = append(game.effects, Effect{6, nil, stopShield, "shield"})
	game.log = append(game.log, fmt.Sprintf("Shield casts, player armor=%d", game.personArmor))
}

func poison(game *Game) {
	game.bossHitpoints -= 3
	game.log = append(game.log, fmt.Sprintf("Poison effect, boss HP=%d", game.bossHitpoints))
}
func castPoison(game *Game) {
	game.effects = append(game.effects, Effect{6, poison, nil, "poison"})
	game.log = append(game.log, fmt.Sprintf("Poison cast"))
}
func recharge(game *Game) {
	game.personMana += 101
	game.log = append(game.log, fmt.Sprintf("Recharge effect, player mana=%d", game.personMana))
}
func castRecharge(game *Game) {
	game.effects = append(game.effects, Effect{5, recharge, nil, "recharge"})
	game.log = append(game.log, fmt.Sprintf("Recharge cast"))
}

type Spell struct {
	name string
	mana int
	cast EffectFunc
}

var spells = []Spell{{"missile", 53, castMissile}, {"drain", 73, castDrain}, {"shield", 113, castShield}, {"poison", 173, castPoison}, {"recharge", 229, castRecharge}}

func parseBoss(lines []string, game *Game) {
	for _, line := range lines {
		if strings.HasPrefix(line, "Hit Points:") {
			game.bossHitpoints = int(h.ToInt(line[len("Hit Points: "):]))
		} else if strings.HasPrefix(line, "Damage: ") {
			game.bossDamage = int(h.ToInt(line[len("Damage: "):]))
		}
	}
}

func Run(init Game) int {
	queue := pq.New(func(a, b Game) bool {
		return a.manaSpent < b.manaSpent || a.manaSpent == b.manaSpent && a.nextSpell.mana < b.nextSpell.mana
	})

	for _, spell := range spells {
		newgame := init
		newgame.nextSpell = spell
		queue.Push(newgame)
	}
	for {
		game, ok := queue.Pop()
		if !ok {
			panic("No win found")
		}
		game.log = append(game.log,
			fmt.Sprintf("--- Start turn: player HP=%d, mana=%d, armor=%d; Boss HP=%d",
				game.personHitpoints,
				game.personMana,
				game.personArmor,
				game.bossHitpoints))
		if game.extraDamage > 0 {
			game.personHitpoints -= game.extraDamage
			game.log = append(game.log, fmt.Sprintf("Extra damage %d, HP=%d", game.extraDamage, game.personHitpoints))
		}
		if game.personHitpoints <= 0 {
			continue // loose
		}
		game.ApplyEffects()
		if game.bossHitpoints <= 0 {
			// fmt.Println(strings.Join(game.log, "\n"))
			return game.manaSpent // win
		}
		// person's turn
		if game.personMana < game.nextSpell.mana {
			continue // not enough mana
		}
		skip := false
		for _, effect := range game.effects {
			if effect.name == game.nextSpell.name {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		game.nextSpell.cast(&game)
		game.manaSpent += game.nextSpell.mana
		game.personMana -= game.nextSpell.mana
		if game.bossHitpoints <= 0 {
			// fmt.Println(strings.Join(game.log, "\n"))
			return game.manaSpent // win
		}
		// boss turn
		game.log = append(game.log,
			fmt.Sprintf("--- Boss turn: player HP=%d, mana=%d, armor=%d; Boss HP=%d; Cost=%d",
				game.personHitpoints,
				game.personMana,
				game.personArmor,
				game.bossHitpoints,
				game.manaSpent))
		game.ApplyEffects()
		if game.bossHitpoints <= 0 {
			// fmt.Println(strings.Join(game.log, "\n"))
			return game.manaSpent // win
		}
		bossAttack := max(1, game.bossDamage-game.personArmor)
		game.personHitpoints -= bossAttack
		game.log = append(game.log, fmt.Sprintf("Boss attacks for %d damage, player HP=%d", bossAttack, game.personHitpoints))
		if game.personHitpoints <= 0 {
			continue // loose
		}
		for _, spell := range spells {
			newgame := game
			newgame.effects = nil
			newgame.effects = append(newgame.effects, game.effects...)
			newgame.log = nil
			newgame.log = append(newgame.log, game.log...)
			newgame.nextSpell = spell
			queue.Push(newgame)
		}
	}
}

func PartOne(lines []string) string {
	game := Game{personHitpoints: 50, personMana: 500}
	parseBoss(lines, &game)
	result := Run(game)
	return fmt.Sprint(result)
}

func TestGame() Game {
	return Game{
		personHitpoints: 10, personMana: 250,
		bossHitpoints: 13, bossDamage: 8,
	}
}
func PartTwo(lines []string) string {
	game := Game{personHitpoints: 50, personMana: 500, extraDamage: 1}
	parseBoss(lines, &game)
	result := Run(game)
	return fmt.Sprint(result)
}
