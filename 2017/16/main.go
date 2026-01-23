package main

import (
	"bytes"
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

func Move(progs []byte, move string) []byte {
	params := strings.Split(move[1:], "/")
	switch move[0] {
	case 's': // spin
		pos := ToInt(params[0])
		progs = append(progs[len(progs)-pos:], progs[:len(progs)-pos]...)
	case 'x': // swap by position
		from := ToInt(params[0])
		to := ToInt(params[1])

		progs[to], progs[from] = progs[from], progs[to]
	case 'p': // swap by name
		from := bytes.IndexByte(progs, byte(params[0][0]))
		to := bytes.IndexByte(progs, byte(params[1][0]))

		progs[to], progs[from] = progs[from], progs[to]
	}

	return progs
}

func solve1(s []string) int {
	res := 0

	progs := []byte("abcdefghijklmnop")
	for _, move := range strings.Split(s[0], ",") {
		progs = Move(progs, move)
		dbg("%s (%s)", progs, move)
	}
	fmt.Println(string(progs))

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
