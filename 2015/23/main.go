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

func dbg(form string, v ...interface{}) {
	if debug {
		fmt.Printf(form+"\n", v...)
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

func runProgram(s []string, a, b int) (int, int) {

	reg := map[string]int{
		"a": a,
		"b": b,
	}
	pc := 0

	for {
		if pc >= len(s) {
			break
		}
		op := strings.Split(s[pc], " ")
		dbg("==PRE==\nPC: %v\nA: %v\nB: %v", pc, reg["a"], reg["b"])
		dbg("OP: %v", op)

		switch op[0] {
		case "hlf":
			reg[op[1]] /= 2
		case "tpl":
			reg[op[1]] *= 3
		case "inc":
			reg[op[1]]++
		case "jmp":
			pc += ToInt(op[1])
			continue

		case "jie":
			r := (op[1][:1])
			offset := ToInt(op[2])

			if reg[r]%2 == 0 {
				pc += offset
			} else {
				pc++
			}
			continue

		case "jio":
			r := (op[1][:1])
			offset := ToInt(op[2])

			if reg[r] == 1 {
				pc += offset
			} else {
				pc++
			}
			continue
		}

		pc++
		dbg("==POS==\nPC: %v\nA: %v\nB: %v\n", pc, reg["a"], reg["b"])
		// time.Sleep(1 * time.Second)
	}

	return reg["a"], reg["b"]
}

func solve1(s []string) int {
	res := 0

	_, res = runProgram(s, 0, 0)
	return res
}

func solve2(s []string) int {
	res := 0
	_, res = runProgram(s, 1, 0)

	return res
}
