package main

import (
	"flag"
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

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
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

type dial struct {
	current int
	length  int
}

func NewDial(length int, initial int) *dial {
	d := &dial{
		length:  length,
		current: initial,
	}

	return d
}

func (d *dial) exec(rotation string) *dial {
	dir := rotation[0]
	count := ToInt(rotation[1:])

	dbg("INIT: %+v", d)
	dbg("%s (%d)", rotation, count)
	switch dir {
	case 'R':
		d.current = Mod(d.current+count, d.length)
	case 'L':
		d.current = Mod(d.current-count, d.length)
	}
	dbg("FINAL: %+v", d)

	return d
}

func solve1(s []string) int {
	res := 0
	d := NewDial(100, 50)

	dbg(">>>>> PART 1")
	for _, rot := range s {
		d = d.exec(rot)
		if d.current == 0 {
			res++
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	dbg(">>>>> PART 2")
	d := NewDial(100, 50)

	for _, rot := range s {
		count := ToInt(rot[1:])
		dbg(">>>> ROTATION: " + rot + "<<<<<<<<<<<<<<<<<<<")
		for range count {
			r := string(rot[0]) + "1"
			d = d.exec(r)
			if d.current == 0 {
				res++
			}
		}
		dbg(">>>> ROTATION END: %+v <<<<<<<<<<<<<<<", d)
	}
	return res
}
