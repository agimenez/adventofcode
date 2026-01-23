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
		spin := ToInt(params[0])
		newprogs := make([]byte, len(progs))
		pos := len(progs) - spin

		copy(newprogs, progs[pos:])
		copy(newprogs[len(progs)-pos:], progs[:pos])
		progs = newprogs
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
	fmt.Println("part 1:", string(progs))

	dbg("")
	return res
}

func printCache(c map[string][]byte) {
	fmt.Println("Cache contents:")
	for k, v := range c {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println("---------------")
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	progs := []byte("abcdefghijklmnop")
	moves := strings.Split(s[0], ",")
	cache := map[string][]byte{}
	positions := map[int]string{}
	cycle := 0
	hits := 0
	for range 1000000000 {
		origprog := string(progs)
		var result []byte
		var ok bool
		positions[cycle] = origprog
		if result, ok = cache[origprog]; !ok {
			for _, move := range moves {
				progs = Move(progs, move)
				// dbg("%s (%s)", progs, move)
			}

			cache[origprog] = bytes.Clone(progs)
			dbg("[%d] CACHED: %s -> %s cache[%s] = %s", cycle, origprog, string(progs), origprog, cache[origprog])
		} else {
			dbg("Position %d repeats (%s -> %s (cached result))", cycle, origprog, result)
			hits++
			break
		}
		dbg("[%d]         %s -> %s", cycle, origprog, positions[cycle])
		// printCache(cache)
		cycle++

		if cycle%100_000 == 0 {
			fmt.Print(".")
		}

		if cycle%10_000_000 == 0 {
			fmt.Printf(" (%d) - cache: %d, %d hits\n", cycle, len(cache), hits)
		}

	}
	pos := 1_000_000_000 % cycle
	dbg("Position in the array: %d -> %s", pos, positions[pos])
	fmt.Println("part 2:", positions[pos])

	return res
}
