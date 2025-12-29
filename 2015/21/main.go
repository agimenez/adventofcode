package main

import (
	"flag"
	"fmt"
	"io"
	"iter"
	"log"
	"math"
	"os"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}
func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type Item struct {
	cost   int
	damage int
	armor  int
}

var weapons []Item = []Item{
	{cost: 8, damage: 4, armor: 0},
	{cost: 10, damage: 5, armor: 0},
	{cost: 25, damage: 6, armor: 0},
	{cost: 40, damage: 7, armor: 0},
	{cost: 74, damage: 8, armor: 0},
}

var armor []Item = []Item{
	// "None"
	{cost: 0, damage: 0, armor: 0},

	{cost: 13, damage: 0, armor: 1},
	{cost: 31, damage: 0, armor: 2},
	{cost: 53, damage: 0, armor: 3},
	{cost: 75, damage: 0, armor: 4},
	{cost: 102, damage: 0, armor: 5},
}

var rings []Item = []Item{
	{cost: 25, damage: 1, armor: 0},
	{cost: 50, damage: 2, armor: 0},
	{cost: 100, damage: 3, armor: 0},
	{cost: 20, damage: 0, armor: 1},
	{cost: 40, damage: 0, armor: 2},
	{cost: 80, damage: 0, armor: 3},
}

type Player struct {
	hp  int
	dmg int
	def int
}

func parsePlayer(s []string) Player {
	return Player{
		hp:  ToInt(strings.Split(s[0], ": ")[1]),
		dmg: ToInt(strings.Split(s[1], ": ")[1]),
		def: ToInt(strings.Split(s[2], ": ")[1]),
	}
}

func AllEquipments() iter.Seq[[]Item] {
	return func(yield func(item []Item) bool) {
		for w := 0; w < len(weapons); w++ {
			for a := 0; a < len(armor); a++ {
				equipment := []Item{
					weapons[w],
					armor[a],
				}
				for ringCombs := range 3 {
					for r := range Combinations(rings, ringCombs) {
						fullEquipment := append(equipment, r...)
						if !yield(fullEquipment) {
							return
						}
					}
				}
			}
		}
	}
}

func (p Player) Equip(item Item) (Player, int) {
	p.dmg += item.damage
	p.def += item.armor

	return p, item.cost
}

func (p Player) EquipAll(items []Item) (Player, int) {
	cost := 0

	for _, item := range items {
		var itemCost int
		p, itemCost = p.Equip(item)
		cost += itemCost
	}

	return p, cost
}

func (p Player) Attack(foe Player) Player {
	dmg := p.dmg - foe.def
	dmg = Max(1, dmg)
	foe.hp -= dmg

	return foe
}

func battleWinner(me, boss Player) bool {
	for {
		boss = me.Attack(boss)
		if boss.hp <= 0 {
			return true
		}

		me = boss.Attack(me)
		if me.hp <= 0 {
			return false
		}

	}
}

func simulate(me Player, boss Player) int {
	minGold := math.MaxInt
	for equipment := range AllEquipments() {
		meEquipped, cost := me.EquipAll(equipment)

		if battleWinner(meEquipped, boss) {
			minGold = Min(minGold, cost)
		}
	}

	return minGold
}

func solve1(s []string) int {
	res := 0

	me := Player{hp: 100, dmg: 0, def: 0}
	boss := parsePlayer(s)

	res = simulate(me, boss)

	//TEST
	// me = Player{hp: 8, dmg: 5, def: 5}
	// boss = Player{hp: 12, dmg: 7, def: 2}
	// battleWinner(me, boss)

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
