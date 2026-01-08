package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	l := slices.Clone(lines)
	part1 := solve1(l, 7)
	dur[0] = time.Since(now)

	now = time.Now()
	l = slices.Clone(lines)
	part2 := solve1(l, 12)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

func solve1(s []string, a int) int {
	res := 0

	dbg("=================================")
	reg := map[string]int{
		"a": a,
		"b": 0,
		"c": 0,
		"d": 0,
	}
	pc := 0

	tick := 0
	for {
		if pc >= len(s) {
			break
		}
		instr := s[pc]
		dbg("[%d] REGS: %v", pc, reg)
		dbg("[%d] INST %v", pc, instr)
		pc++
		decoded := strings.Split(instr, " ")

		switch decoded[0] {
		case "cpy":
			var val int
			if _, ok := reg[decoded[1]]; !ok {
				val = ToInt(decoded[1])
			} else {
				val = reg[decoded[1]]
			}

			reg[decoded[2]] = val
		case "inc":
			reg[decoded[1]]++
		case "dec":
			reg[decoded[1]]--
		case "jnz":
			var val int
			if _, ok := reg[decoded[1]]; !ok {
				val = ToInt(decoded[1])
			} else {
				val = reg[decoded[1]]
			}

			var offset int
			if _, ok := reg[decoded[2]]; !ok {
				offset = ToInt(decoded[2])
			} else {
				offset = reg[decoded[2]]
			}
			// dbg("  >> JNZ %d %d", val, offset)

			if val != 0 {
				// -1 because we already incremented it at the beginning of the cycle
				pc = pc + offset - 1
			}
			// dbg("  >> JNZ pc = %d", pc)

		case "tgl":
			offset := pc - 1 + reg[decoded[1]]
			if offset < 0 || offset >= len(s) {
				// dbg("  >> TGL Target OOB (%d)", offset)
				continue
			}
			target := s[offset]
			// dbg(" >> TGL: target (%d): %v", offset, target)

			var b strings.Builder
			parts := strings.Fields(target)
			if len(parts) == 2 {
				if parts[0] == "inc" {
					b.WriteString("dec")
				} else {
					b.WriteString("inc")
				}

			} else {
				if parts[0] == "jnz" {
					b.WriteString("cpy")
				} else {
					b.WriteString("jnz")
				}
			}
			b.WriteByte(' ')
			b.WriteString(strings.Join(parts[1:], " "))
			// dbg(" >> TGL: %s", b.String())
			s[offset] = b.String()
		case "mul":
			r1 := reg[decoded[1]]
			r2 := reg[decoded[2]]
			reg[decoded[3]] = r1 * r2

		case "nop": // nothing

		default:
			panic("Unknown instruction: " + decoded[0])
		}
		// dbg("%v", strings.Join(s, "\n"))
		// fmt.Printf("[%d] IP: %d (%s) -- A: %d B: %d C: %d D: %d\n", tick, pc, decoded[0], reg["a"], reg["b"], reg["c"], reg["d"])
		tick++
		if tick%1_000_000 == 0 {
			print(".")

			if tick%120_000_000 == 0 {
				fmt.Printf(" (%d)\n", tick)
			}
		}
	}

	dbg("=== [%d] HALT: %v", pc, reg)
	res = reg["a"]

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
