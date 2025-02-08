package day10

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"regexp"
)

type Bot struct {
	lowToBot     string
	highToBot    string
	lowToOutput  int64
	highToOutput int64
	inputs       []int64
}

func (bot *Bot) AddInput(input int64) {
	bot.inputs = append(bot.inputs, input)
	if len(bot.inputs) == 2 && bot.inputs[0] > bot.inputs[1] {
		bot.inputs[1], bot.inputs[0] = bot.inputs[0], bot.inputs[1]
	}
	if len(bot.inputs) > 2 {
		panic("to many inputs for bot")
	}
}

func ParseBots(lines []string) map[string]*Bot {
	bots := make(map[string]*Bot, len(lines))
	botLine := regexp.MustCompile(`^bot (\d+) gives low to (\w+) (\d+) and high to (\w+) (\d+)?$`)
	inputLine := regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	for _, line := range lines {
		botMatch := botLine.FindStringSubmatch(line)
		if len(botMatch) > 0 {
			botName := botMatch[1]
			bot, found := bots[botName]
			if !found {
				bot = new(Bot)
				bots[botName] = bot
			}
			if botMatch[2] == "bot" {
				bot.lowToBot = botMatch[3]
			} else {
				bot.lowToOutput = h.ToInt(botMatch[3])
			}
			if botMatch[4] == "bot" {
				bot.highToBot = botMatch[5]
			} else {
				bot.highToOutput = h.ToInt(botMatch[5])
			}
		} else {
			inputMatch := inputLine.FindStringSubmatch(line)
			botName := inputMatch[2]
			bot, found := bots[botName]
			if !found {
				bot = new(Bot)
				bots[botName] = bot
			}
			bot.AddInput(h.ToInt(inputMatch[1]))
		}
	}
	return bots
}
func PartOne(lines []string) string {
	bots := ParseBots(lines)
	target := []int64{17, 61}
	if len(lines) < 100 {
		target = []int64{3, 5}
	}
	for len(bots) > 0 {
		found := false
		for name, bot := range bots {
			if len(bot.inputs) == 2 {
				if bot.inputs[0] == target[0] && bot.inputs[1] == target[1] {
					return name
				}
				if bot.lowToBot != "" {
					bots[bot.lowToBot].AddInput(bot.inputs[0])
				}
				if bot.highToBot != "" {
					bots[bot.highToBot].AddInput(bot.inputs[1])
				}
				delete(bots, name)
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprint("bots remained:", bots))
		}
	}
	return "not found"
}

func PartTwo(lines []string) string {
	bots := ParseBots(lines)
	outputs := make([]int64, 0, 3)
	for len(bots) > 0 {
		found := false
		for name, bot := range bots {
			if len(bot.inputs) == 2 {
				if bot.lowToBot != "" {
					bots[bot.lowToBot].AddInput(bot.inputs[0])
				} else if bot.lowToOutput < 3 {
					outputs = append(outputs, bot.inputs[0])
				}
				if bot.highToBot != "" {
					bots[bot.highToBot].AddInput(bot.inputs[1])
				} else if bot.highToOutput < 3 {
					outputs = append(outputs, bot.inputs[1])
				}
				if len(outputs) == 3 {
					return h.ToString(outputs[0] * outputs[1] * outputs[2])
				}
				delete(bots, name)
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprint("bots remained:", bots))
		}
	}
	return "not found"
}
