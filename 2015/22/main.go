package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"math"
	"os"
	"slices"
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

type Player struct {
	hp       int
	dmg      int
	def      int
	mana     int
	usedMana int

	effects map[string]Effect
}

type EffectFunc func(p Player, b Player) (Player, Player)

type Spell struct {
	name string
	cost int

	effect Effect
}

type Effect struct {
	name      string
	duration  int
	remaining int

	onCast   EffectFunc
	onTurn   EffectFunc
	onExpire EffectFunc
}

var NilHandler = func(p Player, b Player) (Player, Player) { return p, b }

var ShieldEffect = Effect{
	name:     "Shield",
	duration: 6,

	onCast:   func(p, b Player) (Player, Player) { p.def += 7; return p, b },
	onTurn:   NilHandler,
	onExpire: func(p, b Player) (Player, Player) { p.def -= 7; return p, b },
}

var MagicMissleEffect = Effect{
	name: "Magic Missile",

	onCast:   func(p, b Player) (Player, Player) { b.hp -= 4; return p, b },
	onTurn:   NilHandler,
	onExpire: NilHandler,
}

var DrainEffect = Effect{
	name: "Drain",

	onCast:   func(p, b Player) (Player, Player) { p.hp += 2; b.hp -= 2; return p, b },
	onTurn:   NilHandler,
	onExpire: NilHandler,
}

var PoisonEffect = Effect{
	name:     "Poison",
	duration: 6,

	onCast:   NilHandler,
	onTurn:   func(p, b Player) (Player, Player) { b.hp -= 3; return p, b },
	onExpire: NilHandler,
}

var RechargeEffect = Effect{
	name:     "Recharge",
	duration: 5,

	onCast:   NilHandler,
	onTurn:   func(p, b Player) (Player, Player) { p.mana += 101; return p, b },
	onExpire: NilHandler,
}

var Spells []Spell = []Spell{
	{"Magic Missile", 53, MagicMissleEffect},
	{"Drain", 73, DrainEffect},
	{"Shield", 113, ShieldEffect},
	{"Poison", 173, PoisonEffect},
	{"Recharge", 229, RechargeEffect},
}

func (p Player) Clone() Player {
	p2 := p
	p2.effects = maps.Clone(p.effects)

	return p2
}

func ApplyEffects(p Player, b Player) (Player, Player) {
	for name, effect := range p.effects {
		p, b = effect.onTurn(p, b)
		effect.remaining--
		dbg("Applying effect: %v, timer is now %v", name, effect.remaining)
		if effect.remaining <= 0 {
			p, b = effect.onExpire(p, b)
			dbg("  >> %s expired!", name)
			delete(p.effects, name)
		} else {
			p.effects[name] = effect
		}

	}

	return p, b
}

func (p Player) AvailableSpells() []Spell {
	ret := []Spell{}
	for _, spell := range Spells {
		dbg("Player mana: %v, spell (%s) cost: %v, remaining: %v", p.mana, spell.name, spell.cost, p.effects[spell.name].remaining)
		e, ok := p.effects[spell.name]
		if p.mana >= spell.cost && (!ok || e.remaining <= 0) {
			dbg(" >>> Yelding %v", spell.name)
			ret = append(ret, spell)
		}
	}

	return ret
}

func parsePlayer(s []string) Player {
	return Player{
		hp:  ToInt(strings.Split(s[0], ": ")[1]),
		dmg: ToInt(strings.Split(s[1], ": ")[1]),
	}
}

func (s Spell) Cast() Spell {
	s.effect.remaining = s.effect.duration

	return s
}

func (p Player) CastSpell(s Spell, b Player) (Player, Player, bool) {
	if s.cost > p.mana {
		dbg("  >> Not enough mana (%d, need %d)", p.mana, s.cost)
		return p, b, false
	}

	effect, ok := p.effects[s.name]
	if ok && effect.remaining > 0 {
		dbg("  >> Effect %s still active (remaining %d turns)", effect.name, effect.remaining)
		return p, b, false
	}

	p.mana -= s.cost
	p.usedMana += s.cost

	s = s.Cast()
	p, b = s.effect.onCast(p, b)
	if s.effect.duration > 0 {
		p.effects[s.name] = s.effect
	}

	return p, b, true
}

const (
	CONTINUE = "continue"
	WIN      = "win"
	LOST     = "lost"
	INVALID  = "invalid cast"
)

// Single Round is 2 turns: first player, then boss.
func Round(me Player, boss Player, s Spell) (Player, Player, string) {
	// Player turn: apply effects
	me = me.Clone()
	boss = boss.Clone()

	dbg("--- Player turn ---")
	dbg(" - Player: %v", me)
	dbg(" - Boss: %v", boss)
	me, boss = ApplyEffects(me, boss)
	if boss.hp <= 0 {
		dbg("Boss is killed")
		return me, boss, WIN
	}

	// Cast player spell
	dbg(" > Player casts %s", s.name)
	var success bool
	me, boss, success = me.CastSpell(s, boss)
	if success == false {
		dbg(" > CANNOT CAST!")
		return me, boss, INVALID
	}
	if boss.hp <= 0 {
		dbg("Boss is killed")
		return me, boss, WIN
	}

	// Boss turn: apply effects + damage player
	dbg("")
	dbg("--- Boss turn ---")
	dbg(" - Player: %v", me)
	dbg(" - Boss: %v", boss)
	me, boss = ApplyEffects(me, boss)
	if boss.hp <= 0 {
		dbg("Boss is killed")
		return me, boss, WIN
	}

	dmg := Max(1, boss.dmg-me.def)
	me.hp -= dmg
	dbg("Boss attacks for %d damage (%v - %v)", dmg, boss.dmg, me.def)
	if me.hp <= 0 {
		dbg("Player is dead")
		return me, boss, LOST
	}

	return me, boss, CONTINUE
}

func (s state) Clone() state {
	s2 := s
	s2.spells = slices.Clone(s.spells)

	return s2
}

func (s state) AppendSpell(name string) state {
	s.spells = append(s.spells, name)
	return s
}

func Simulate(me Player, boss Player, currentMinMana int, s state) int {
	s.depth++
	dbg("DEPTH: %v, Spells: %+v", s.depth, s.spells)
	if me.usedMana >= currentMinMana {
		return currentMinMana
	}

	for _, spell := range Spells {
		dbg("[%d] PLAYER: %v\nSPELL: %v", s.depth, me, spell.name)
		newMe, newBoss, outcome := Round(me, boss, spell)
		if outcome == INVALID {
			continue
		}

		newState := s.Clone().AppendSpell(spell.name)

		dbg("ROUND FINISHED")
		dbg("---------------")
		dbg(" - Player: %v", newMe)
		dbg(" - Boss: %v", newBoss)
		dbg(" - Outcome: %v", outcome)

		if outcome == WIN {
			if newMe.usedMana < currentMinMana {
				fmt.Printf("New win: %d, Spells: %+v\n", Min(newMe.usedMana, currentMinMana), newState.spells)
			}
			currentMinMana = Min(newMe.usedMana, currentMinMana)
		}

		if outcome == CONTINUE {
			dbg("CONTINUING (min used: %v)", currentMinMana)
			dbg("PREVIOUS: %v || %v", me, boss)
			dbg("NEW     : %v || %v", newMe, newBoss)
			currentMinMana = Simulate(newMe, newBoss, currentMinMana, newState)
		}
		dbg("")
	}
	dbg("END SIMULATION -- No more spells available!")
	dbg("")

	return currentMinMana

}

func (p Player) String() string {
	return fmt.Sprintf("hp: %d | armor: %d | mana: %d | used mana: %d", p.hp, p.def, p.mana, p.usedMana)
}

func TestRun() {
	me := Player{hp: 10, mana: 250, effects: map[string]Effect{}}
	boss := Player{hp: 13, dmg: 8}

	dbg("====== TEST RUN ======")
	var TestSpells []Spell = []Spell{
		{"Poison", 173, PoisonEffect},
		{"Magic Missile", 53, MagicMissleEffect},
	}

	for _, spell := range TestSpells {
		dbg("PRE ROUND: %v", me)
		me, boss, _ = Round(me, boss, spell)
		dbg("POST ROUND: %v", me)
		dbg("")
	}
	dbg("====== END TEST RUN ======")
}

func TestRun2() {
	me := Player{hp: 10, mana: 250, effects: map[string]Effect{}}
	boss := Player{hp: 14, dmg: 8}

	dbg("====== TEST RUN ======")
	var TestSpells []Spell = []Spell{
		{"Recharge", 229, RechargeEffect},
		{"Shield", 113, ShieldEffect},
		{"Drain", 73, DrainEffect},
		{"Poison", 173, PoisonEffect},
		{"Magic Missile", 53, MagicMissleEffect},
	}

	for _, spell := range TestSpells {
		me, boss, _ = Round(me, boss, spell)
		dbg("")
	}
	dbg("====== END TEST RUN ======")
}

func TestRun3() {
	me := Player{hp: 50, mana: 500, effects: map[string]Effect{}}
	boss := Player{hp: 71, dmg: 10}

	dbg("====== TEST RUN ======")
	// 1824 mana spent: 173 (), 229 (Recharge), 113 (Shield), 173 (), 229 (Recharge), 113 (Shield), 173 (), 229 (Recharge), 113 (Shield), 53 (), 173 (), 53 (),
	// 1824 mana spent:
	//
	// 173 (Poison)
	// 229 (Recharge)
	// 113 (Shield)
	// 173 (Poison)
	// 229 (Recharge)
	// 113 (Shield)
	// 173 (Poison)
	// 229 (Recharge)
	// 113 (Shield)
	// 53 (MM)
	// 173 (Poison)
	// 53 (MM)

	var TestSpells []Spell = []Spell{
		{"Poison", 173, PoisonEffect},
		{"Recharge", 229, RechargeEffect},
		{"Shield", 113, ShieldEffect},
		{"Poison", 173, PoisonEffect},
		{"Recharge", 229, RechargeEffect},
		{"Shield", 113, ShieldEffect},
		{"Poison", 173, PoisonEffect},
		{"Recharge", 229, RechargeEffect},
		{"Shield", 113, ShieldEffect},

		{"Magic Missile", 53, MagicMissleEffect},
		{"Poison", 173, PoisonEffect},
		{"Magic Missile", 53, MagicMissleEffect},
		// HERE should end

		// {"Recharge", 229, RechargeEffect},
		// {"Shield", 113, ShieldEffect},
		// {"Drain", 73, DrainEffect},
		// {"Poison", 173, PoisonEffect},
		// {"Magic Missile", 53, MagicMissleEffect},
	}

	var outcome string
	for _, spell := range TestSpells {
		me, boss, outcome = Round(me, boss, spell)
		dbg("OUTCOME: %v", outcome)
		dbg("")
	}
	dbg("========= RESULTS ========")
	dbg(" - Player: %v", me)
	dbg(" - Boss: %v", boss)
	dbg("====== END TEST RUN ======")
}

type state struct {
	me, boss Player
	depth    int
	spells   []string
}

func solve1(s []string) int {
	res := 0
	me := Player{hp: 50, mana: 500, effects: map[string]Effect{}}
	boss := parsePlayer(s)

	if debug {
		// TestRun()
		// TestRun2()
		// TestRun3()
	}

	state := state{
		me:     me,
		boss:   boss,
		depth:  0,
		spells: []string{},
	}
	res = Simulate(me, boss, math.MaxInt, state)

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
