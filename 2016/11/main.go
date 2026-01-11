package main

import (
	"cmp"
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
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

const (
	CHIP = iota
	GEN
)

type Pair struct {
	name  string
	floor [2]int
}

func (p Pair) Chip() string {
	b := []byte{p.name[0], 'M'}

	return strings.ToUpper(string(b))
}

func (p Pair) Gen() string {
	b := []byte{p.name[0], 'G'}

	return strings.ToUpper(string(b))

}

type State struct {
	elevator int
	items    []Pair
}

func (s *State) AddPair(name string, chipFloor, genFloor int) *State {
	p := Pair{
		name:  name,
		floor: [2]int{CHIP: chipFloor, GEN: genFloor},
	}

	s.items = append(s.items, p)
	s.sort()

	return s
}

func (s State) Clone() State {
	return State{
		elevator: s.elevator,
		items:    slices.Clone(s.items),
	}
}

func (s State) sort() {
	slices.SortFunc(s.items, func(a Pair, b Pair) int {
		return cmp.Or(
			cmp.Compare(a.floor[CHIP], b.floor[CHIP]),
			cmp.Compare(a.floor[GEN], b.floor[GEN]),
		)

	})

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

		for _, pair := range s.items {
			if pair.floor[CHIP] == f {
				b.WriteString(pair.Chip())
				b.WriteRune(' ')
			} else {
				b.WriteString(".  ")
			}

			if pair.floor[GEN] == f {
				b.WriteString(pair.Gen())
				b.WriteRune(' ')
			} else {
				b.WriteString(".  ")
			}
		}
		b.WriteRune('\n')
	}

	fmt.Println(b.String())
}

func parseInput(in []string) State {
	pairs := map[string]Pair{}
	re := regexp.MustCompile(`([a-z]+)(?:-compatible)? (microchip|generator)`)
	for _, l := range in {
		floor := getFloor(l)
		itemsinFloor := re.FindAllStringSubmatch(l, -1)
		for _, items := range itemsinFloor {
			var p Pair
			var ok bool
			if p, ok = pairs[items[1]]; !ok {
				p = Pair{name: items[1]}
			}
			if items[2] == "generator" {
				p.floor[GEN] = floor
			} else {
				p.floor[CHIP] = floor
			}

			pairs[items[1]] = p

		}

	}

	s := State{
		elevator: 1,
		items:    slices.Collect(maps.Values(pairs)),
	}
	s.sort()

	return s
}

func (s State) DistToSol() int {
	res := 4 - s.elevator

	res += Reduce(s.items, 0, func(acc int, p Pair) int {
		return acc + 8 - p.floor[CHIP] - p.floor[GEN]
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
		b.WriteRune(rune(i.floor[CHIP] + '0'))
		b.WriteRune(rune(i.floor[GEN] + '0'))
	}

	return b.String()
}

func (s State) MoveItems(items [][2]int, dir int) State {
	s.elevator += dir

	for _, item := range items {
		dbg("   >> MoveItems: got item %v", item)
		pair := item[0]
		component := item[1]
		s.items[pair].floor[component] += dir
	}
	s.sort()

	return s
}

func (s State) IsValid() bool {
	if s.elevator < 1 || s.elevator > 4 {
		return false
	}

	gensByFloor := map[int]int{}
	for i := 0; i < len(s.items); i++ {
		gensByFloor[s.items[i].floor[GEN]]++
	}

	dbg("================= IsValid ================")
	dbgState(s)
	dbg("IsValid (%s) Gens by floor: %v", s.Serial(), gensByFloor)

	for i := 0; i < len(s.items); i++ {
		pair := s.items[i]

		dbg("[%d/%d] Checking [%s] [chip: %d/ gen: %d]", i, len(s.items)-1, pair.name, pair.floor[CHIP], pair.floor[GEN])

		// NOK: Chip is not with its generator AND not by itself in a floor
		if pair.floor[CHIP] != pair.floor[GEN] && gensByFloor[pair.floor[CHIP]] > 0 {
			dbg("  >> [%s] INVALID 1!!!", pair.name)
			return false
		}
	}

	return true
}

func (s State) NextMoves() []State {
	nextStates := []State{}

	currentPairIndices := slices.Collect(func(yield func([2]int) bool) {
		for i, pair := range s.items {
			out := [2]int{i, -1}

			if pair.floor[CHIP] == s.elevator {
				out[1] = CHIP
				if !yield(out) {
					return
				}
			}

			if pair.floor[GEN] == s.elevator {
				out[1] = GEN
				if !yield(out) {
					return
				}
			}

		}
	})
	dbg("Current pairs with items on floor %d: %v", s.elevator, currentPairIndices)

	combinations := CollectCombinations(currentPairIndices, 1)
	combinations = append(combinations, CollectCombinations(currentPairIndices, 2)...)
	// dbg("Got Combinations: %v", combinations)

	for _, dir := range []int{-1, 1} {
		for _, comb := range combinations {
			// dbg("[%d] TESTING combination %v", dir, comb)
			next := s.Clone().MoveItems(comb, dir)
			next.sort()
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

	end := strings.Repeat("4", (2*len(s.items))+1)
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
		// fmt.Println("Q len: ", len(queue))

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
	dbgState(state)
	res = findShortestPath(state)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	state := parseInput(s)
	state.AddPair("elerium", 1, 1)
	state.AddPair("dilitium", 1, 1)

	dbgState(state)
	res = findShortestPath(state)

	return res
}
