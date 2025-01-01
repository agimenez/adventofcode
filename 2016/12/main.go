package main

import (
	"flag"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"os"
	"strings"
	"time"
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
	part1 := solve1(lines, 0)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve1(lines, 1)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

func solve1(s []string, c int) int {
	res := 0

	reg := map[string]int{
		"a": 0,
		"b": 0,
		"c": c,
		"d": 0,
	}
	pc := 0

	for {
		if pc >= len(s) {
			break
		}
		instr := s[pc]
		dbg("INST %v (pc=%v)", instr, pc)
		dbg("REGS: %v", reg)
		pc++
		decoded := strings.Split(instr, " ")

		switch decoded[0] {
		case "cpy":
			dbg("CPY")
			var val int
			if _, ok := reg[decoded[1]]; !ok {
				val = ToInt(decoded[1])
			} else {
				val = reg[decoded[1]]
			}

			reg[decoded[2]] = val
		case "inc":
			dbg("INC")
			reg[decoded[1]]++
		case "dec":
			dbg("DEC")
			reg[decoded[1]]--
		case "jnz":
			dbg("JNZ")
			var val int
			if _, ok := reg[decoded[1]]; !ok {
				val = ToInt(decoded[1])
			} else {
				val = reg[decoded[1]]
			}

			offset := ToInt(decoded[2])
			if val != 0 {
				// -1 because we already incremented it at the beginning of the cycle
				pc = pc + offset - 1
			}
		}
	}

	res = reg["a"]

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
