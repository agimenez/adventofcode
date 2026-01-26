package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
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
	id int
	r  map[string]int
	pc int

	// sound playing
	ch chan int

	prog []string
}

func (c CPU) Queue() chan int {
	return c.ch
}

func NewCPU(id int) CPU {
	return CPU{
		id: id,
		r:  map[string]int{"p": id},
		ch: make(chan int, 2000),
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

func (c *CPU) Run(ch chan int) int {
	sends := 0
	for {
		if c.pc < 0 || c.pc >= len(c.prog) {
			return sends
		}

		inst := c.prog[c.pc]
		op := strings.Fields(inst)
		params := op[1:]

		dbg("[%d] %v", c.id, inst)

		switch op[0] {
		case "snd":
			sends++
			fmt.Printf("[%d] SENDS: %v\n", c.id, sends)
			ch <- c.resolve(params[0])

		case "set":
			c.r[params[0]] = c.resolve(params[1])

		case "add":
			c.r[params[0]] += c.resolve(params[1])

		case "mul":
			c.r[params[0]] = c.r[params[0]] * c.resolve(params[1])

		case "mod":
			c.r[params[0]] = c.r[params[0]] % c.resolve(params[1])

		case "rcv":
			v := <-c.ch
			c.r[params[0]] = v

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

	// Part 1 is very different form part2, so here it's only part2
	// This actually deadlocks, so I guess we should be OK?
	c0 := NewCPU(0)
	c0.Load(s)

	c1 := NewCPU(1)
	c1.Load(s)

	var wg sync.WaitGroup

	wg.Go(func() {
		c0.Run(c1.Queue())
	})

	wg.Go(func() {
		res = c1.Run(c0.Queue())
	})
	wg.Wait()

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
