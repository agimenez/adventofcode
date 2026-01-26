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

type CPU struct {
	r  map[string]int
	pc int

	// sound playing
	out int

	prog []string
}

func NewCPU() CPU {
	return CPU{
		r: map[string]int{},
	}
}

func (c *CPU) Load(s []string) {
	c.pc = 0

	c.prog = s
}

func (c CPU) resolve(param string) int {
	var ret int
	if param[0] >= 'a' && param[0] <= 'z' {
		ret = c.r[param]
	} else {
		ret = ToInt(param)
	}

	return ret
}

func (c *CPU) Run() int {
	for {
		if c.pc < 0 || c.pc >= len(c.prog) {
			return -1
		}

		inst := c.prog[c.pc]
		op := strings.Fields(inst)
		params := op[1:]

		switch op[0] {
		case "snd":
			v := c.resolve(params[0])
			c.out = v

		case "set":
			c.r[params[0]] = c.resolve(params[1])

		case "add":
			c.r[params[0]] += c.resolve(params[1])

		case "mul":
			c.r[params[0]] = c.r[params[0]] * c.resolve(params[1])

		case "mod":
			c.r[params[0]] = c.r[params[0]] % c.resolve(params[1])

		case "rcv":
			if c.resolve(params[0]) != 0 {
				return c.out
			}

		case "jgz":
			if c.resolve(params[0]) > 0 {
				c.pc += c.resolve(params[1])
				continue
			}

		default:
			panic("Unknown instruction:" + op[0])
		}

		c.pc++
	}
}

func solve1(s []string) int {
	res := 0

	cpu := NewCPU()
	cpu.Load(s)
	res = cpu.Run()

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
