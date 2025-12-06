package main

import (
	"bytes"
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

type Stack []int

func (s Stack) Push(d int) Stack {
	s = append(s, d)

	return s
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Pop() (Stack, int) {
	if s.IsEmpty() {
		return s, 0
	}

	return s[:len(s)-1], s[len(s)-1]
}

func solve1(s []string) int {
	res := 0
	var calcs []Stack
	for _, line := range s {
		parts := strings.Fields(line)
		dbg("Read parts: %#v", parts)
		if calcs == nil {
			calcs = make([]Stack, len(parts))
		}

		for i, num := range parts {
			switch num {
			case "+":
				s := calcs[i]
				var n int
				sum := 0
				for !s.IsEmpty() {
					s, n = s.Pop()
					sum += n
				}
				res += sum

			case "*":
				var n int
				s := calcs[i]
				mul := 1
				for !s.IsEmpty() {
					s, n = s.Pop()
					mul *= n
				}
				res += mul
			default:
				calcs[i] = calcs[i].Push(ToInt(num))

			}
		}

		dbg("Stacks: %+v", calcs)
	}
	dbg("Stacks: %+v", calcs)

	return res
}

func solve2(s []string) int {
	res := 0

	nums := Stack{}
	for x := len(s[0]) - 1; x >= 0; x-- {
		var n bytes.Buffer
		done := false
		dbg("NUMS: %+v", nums)
		for y := 0; y < len(s); y++ {
			ch := s[y][x]
			switch ch {
			case '+':
				num := ToInt(n.String())
				nums = nums.Push(num)
				var n int
				sum := 0
				for !nums.IsEmpty() {
					nums, n = nums.Pop()
					sum += n
				}
				res += sum
				dbg(" >> SUM: %v", sum)
				done = true

			case '*':
				num := ToInt(n.String())
				nums = nums.Push(num)
				var n int
				mul := 1
				for !nums.IsEmpty() {
					nums, n = nums.Pop()
					mul *= n
				}
				res += mul
				dbg(" >> MUL: %v", mul)
				done = true
			case ' ': // nothing
			default:
				n.WriteByte(ch)
				dbg("  >> Partial: %v", n.String())
			}
		}

		num := ToInt(n.String())
		if !done && num != 0 {
			nums = nums.Push(num)
			dbg("Col %v, num: %v", x, num)
		}

	}

	return res
}
