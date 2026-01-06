package main

import (
	"container/ring"
	"flag"
	"fmt"
	"io"
	"log"
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

type Elf struct {
	id       int
	presents int
}

func NewTable(n int) *ring.Ring {
	r := ring.New(n)

	for i := 0; i < n; i++ {
		r.Value = Elf{id: i + 1, presents: 1}
		r = r.Next()
	}

	return r
}

func Simulate(t *ring.Ring) int {
	res := 0

	for ; t != t.Next(); t = t.Next() {
		cur := t.Value.(Elf)
		next := t.Next().Value.(Elf)

		cur.presents += next.presents
		next.presents = 0

		t.Value = cur
		t.Next().Value = next
		dbg("Elf %d takes Elf %d's presents", cur.id, next.id)
		if next.presents == 0 {
			dbg("Elf %d has no presents and is skipped", next.id)
			// fmt.Println(next.id)
			t.Unlink(1)
			continue
		}
	}
	res = t.Value.(Elf).id

	return res
}

func SimulateElephantJoseph(n int) int {
	res := 0

	t := NewTable(n)
	res = Simulate(t)

	return res
}

// After looking in Reddit, this looks like the Josephus Problem, and
// there is a mathematical way to solve it, but let's simulate, because
// simulations are cool!
func solve1(s []string) int {
	res := 0

	res = SimulateElephantJoseph(ToInt(s[0]))

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
