package main

import (
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

func dbgMaze(m []int, idx int) {
	if !debug {
		return
	}

	var b strings.Builder
	for i, n := range m {
		if i == idx {
			b.WriteString(fmt.Sprintf("(%d)", n))
		} else {
			b.WriteString(fmt.Sprintf("%d", n))
		}

		b.WriteRune(' ')
	}
	dbg("%s", b.String())
}

func exec(m []int) int {
	steps := 1
	idx := 0

	dbgMaze(m, idx)
	for {
		jmp := m[idx]
		m[idx]++
		if idx+jmp >= len(m) {
			break
		}
		idx += jmp

		steps++
		dbgMaze(m, idx)
	}

	return steps
}

func execv2(m []int) int {
	steps := 1
	idx := 0

	dbgMaze(m, idx)
	for {
		jmp := m[idx]
		if jmp >= 3 {
			m[idx]--
		} else {
			m[idx]++
		}

		if idx+jmp >= len(m) {
			break
		}
		idx += jmp

		steps++
		dbgMaze(m, idx)
	}

	return steps
}

func solve1(s []string) int {
	res := 0
	maze := make([]int, len(s))
	for i, l := range s {
		maze[i] = ToInt(l)
	}
	dbg("%v", maze)
	res = exec(maze)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	maze := make([]int, len(s))
	for i, l := range s {
		maze[i] = ToInt(l)
	}
	res = execv2(maze)

	return res
}
