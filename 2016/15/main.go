package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
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

type Disc struct {
	positions int
	initial   int
}

type DiscStack []Disc

func NewDiscStack() DiscStack {
	return DiscStack{}
}

func (d *DiscStack) AddDisc(s string) {
	var diskno, npos, cur int
	_, err := fmt.Sscanf(s, "Disc #%d has %d positions; at time=0, it is at position %d.", &diskno, &npos, &cur)
	if err != nil {
		panic(err)
	}

	*d = append(*d, Disc{
		positions: npos,
		initial:   cur,
	})
}

func (d Disc) PositionAt(t int) int {
	return (d.initial + t) % d.positions
}

func (d DiscStack) PressButtonAt(start int) bool {
	dbg("PRESS BUTTON: t=%d", start)
	start++ // first disc gets hit at t0+1
	for i := range d {
		dbg("  >> [t=%d] Disc %d reached", start+i, i)
		slot := d[i].PositionAt(start + i)
		dbg("    > Slot %d", slot)
		if slot != 0 {
			return false
		}
	}
	dbg("")

	return true
}

func solve1(s []string) int {
	res := 0

	ds := NewDiscStack()
	for _, line := range s {
		ds.AddDisc(line)
	}
	dbg("Disk stack: %v", ds)

	for ; ; res++ {
		out := ds.PressButtonAt(res)
		if out {
			break
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
