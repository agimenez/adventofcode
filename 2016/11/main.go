package main

import (
	"cmp"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

type Item struct {
	name      string
	generator bool
	floor     int
}

func (i Item) abbrev() string {
	b := []byte{i.name[0], i.name[len(i.name)-1]}

	return strings.ToUpper(string(b))
}

type State struct {
	elevator int
	items    []Item
}

func (s State) Clone() State {
	return State{
		elevator: s.elevator,
		items:    slices.Clone(s.items),
	}
}

func getFloor(s string) int {
	str2floor := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
		"fourth": 4,
	}

	return str2floor[strings.Fields(s)[1]]
}

func dbgState(s State) {
	if !debug {
		return
	}

	var b strings.Builder
	for f := 4; f > 0; f-- {
		b.WriteString(fmt.Sprintf("F%d ", f))
		if s.elevator == f {
			b.WriteString("E | ")
		} else {
			b.WriteString(". | ")
		}
		for _, item := range s.items {
			if item.floor == f {
				b.WriteString(item.abbrev())
			} else {
				b.WriteString(". ")
			}
			b.WriteRune(' ')
		}
		b.WriteRune('\n')
	}

	fmt.Println(b.String())
}

func parseInput(in []string) State {
	allItems := []Item{}
	re := regexp.MustCompile(`([a-z]+)(?:-compatible)? (microchip|generator)`)
	for _, l := range in {
		floor := getFloor(l)
		itemsinFloor := re.FindAllStringSubmatch(l, -1)
		for _, items := range itemsinFloor {
			suffix := "M"
			isGen := false
			if items[2] == "generator" {
				suffix = "G"
				isGen = true
			}
			item := Item{
				name:      items[1] + suffix,
				floor:     floor,
				generator: isGen,
			}

			allItems = append(allItems, item)
		}

	}

	// cosmetic: keep chip and its chip together
	slices.SortFunc(allItems, func(a, b Item) int {
		return strings.Compare(a.name, b.name)
	})

	return State{elevator: 1, items: allItems}
}

func (s State) DistToSol() int {
	res := 4 - s.elevator

	res += Reduce(s.items, 0, func(acc int, i Item) int {
		return acc + 4 - i.floor
	})
	dbg("s: %s, dist: %d", s.Serial(), res)

	return res
}

type path struct {
	state State
	cost  int
	path  []State
}

func (s State) Serial() string {
	var b strings.Builder

	b.WriteRune(rune(s.elevator + '0'))
	for _, i := range s.items {
		b.WriteRune(rune(i.floor + '0'))
	}

	return b.String()
}

func (s State) MoveItems(items []Item, dir int) State {
	s.elevator += dir

	for _, it := range items {
		idx := slices.IndexFunc(s.items, func(i Item) bool { return i.name == it.name })
		s.items[idx].floor += dir
	}

	return s
}

func (s State) IsValid() bool {
	if s.elevator < 1 || s.elevator > 4 {
		return false
	}

	// debug = false
	if s.Serial() == "22232" {
		debug = true
	}
	gensByFloor := map[int]int{}
	for i := 0; i < len(s.items); i += 2 {
		gensByFloor[s.items[i].floor]++
	}
	dbg("================= IsValid ================")
	dbgState(s)
	dbg("IsValid (%s) Gens by floor: %v", s.Serial(), gensByFloor)

	for i := 0; i <= len(s.items)-1; i += 2 {
		gen := s.items[i]
		chip := s.items[i+1]

		dbg("[%d/%d] Checking [%s/%d] [%s/%d]", i, len(s.items)-1, gen.name, gen.floor, chip.name, chip.floor)

		// NOK: Chip is not with its generator AND not by itself in a floor
		if chip.floor != gen.floor && gensByFloor[chip.floor] > 0 {
			dbg("  >> [%s] [%s] INVALID 1!!!", chip.name, gen.name)
			return false
		}
	}
	// debug = false

	return true
}

func (s State) NextMoves() []State {
	nextStates := []State{}
	currentItems := Filter(s.items, func(i Item) bool {
		return i.floor == s.elevator
	})
	dbg("Current Items on floor %d: %v", s.elevator, currentItems)

	combinations := CollectCombinations(currentItems, 1)
	combinations = append(combinations, CollectCombinations(currentItems, 2)...)

	for _, dir := range []int{-1, 1} {
		for _, comb := range combinations {
			// dbg("[%d] TESTING combination %v", dir, comb)
			next := s.Clone().MoveItems(comb, dir)
			nextStates = append(nextStates, next)
		}
	}

	return nextStates
}

func dbgQueue(q []path) {
	if !debug {
		return
	}

	for _, s := range q {
		dbg("%s -> %d", s.state.Serial(), s.cost)
	}
}

func backTrackPath(path []State) {
	olddebug := debug
	debug = true
	dbg("===== BACKTRACK SOLUTION ====")
	for _, s := range path {
		dbgState(s)
	}
	dbg("=============================")
	debug = olddebug
}

func findShortestPath(s State) int {
	queue := []path{{
		state: s,
		cost:  s.DistToSol(),
		path:  []State{s},
	}}

	distances := map[string]int{
		s.Serial(): 0,
	}

	end := strings.Repeat("4", len(s.items)+1)
	dbg("Initial: %s", s.Serial())
	dbg("End sta: %s", end)

	for len(queue) > 0 {
		// Poor man's priority queue
		slices.SortFunc(queue, func(i, j path) int {
			return cmp.Compare(i.cost, j.cost)
		})

		// dbg("Q: %v", queue)
		// dbg("============ QUEUE ==============")
		// dbgQueue(queue)

		cur := queue[0]
		queue = queue[1:]
		fmt.Println("Q len: ", len(queue))

		dbg("================ CURRENT STATE ===============")
		dbgState(cur.state)
		if cur.state.Serial() == end {
			backTrackPath(cur.path)
			return distances[end]
		}

		for _, next := range cur.state.NextMoves() {
			dbg("CUR Move candidate:  %s", cur.state.Serial())
			dbg("Next Move candidate: %s", next.Serial())
			if !next.IsValid() {
				dbg(" >> INVALID STATE")
				continue
			}

			if _, ok := distances[next.Serial()]; !ok {
				dbg("  >> Adding to queue!")
				nextDist := distances[cur.state.Serial()] + 1
				distances[next.Serial()] = nextDist

				nextPath := slices.Clone(cur.path)
				nextPath = append(nextPath, next)
				queue = append(queue, path{
					state: next,
					cost:  nextDist + next.DistToSol(),
					path:  nextPath,
				})
			}
		}
		dbg("DIST: %v", distances)
		dbg("")
	}

	return -1
}

func solve1(s []string) int {
	res := 0

	state := parseInput(s)
	// dbg("%v", state)
	// dbgState(state)
	res = findShortestPath(state)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
