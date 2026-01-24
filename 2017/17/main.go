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

type SpinLock struct {
	lock  *ring.Ring
	len   int
	steps int
}

func NewLock(steps int) SpinLock {
	ring := ring.New(1)
	ring.Value = 0

	return SpinLock{
		lock:  ring,
		len:   1,
		steps: steps,
	}
}

func (s SpinLock) Current() int {
	return s.lock.Value.(int)
}

func (s SpinLock) Next() int {
	return s.lock.Next().Value.(int)
}

func (s SpinLock) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "len: %d, steps: %d, ring: ", s.len, s.steps)
	fmt.Fprintf(&sb, "(%d) ", s.lock.Value.(int))
	r := s.lock.Next()
	for i := 1; i < s.len; i++ {
		fmt.Fprintf(&sb, "%d ", r.Value.(int))
		r = r.Next()
	}

	return sb.String()

}

func (s *SpinLock) Exec(n int) {
	for range n {
		dbg("BEFORE: %v", s)

		s.lock = s.lock.Move(s.steps)
		dbg("MOVE:   %v", s)

		r := ring.New(1)
		r.Value = s.len
		s.len++

		s.lock.Link(r)
		dbg("LINK:   %v", s)
		s.lock = r

		dbg("AFTER:  %v", s)
		dbg("")
	}
}

func Test() {
	l := NewLock(3)
	l.Exec(2017)
	dbg("Test: %d", l.Next())
}

func solve1(s []string) int {
	res := 0

	l := NewLock(ToInt(s[0]))
	l.Exec(2017)

	res = l.Next()

	Test()

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
