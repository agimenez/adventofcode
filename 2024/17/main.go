package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
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

	part1, part2 := "", 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

type minicode struct {
	ip  int
	r   map[string]int
	mem []int
	buf []string
}

func (m minicode) flush() string {
	return strings.Join(m.buf, ",")
}

func (m minicode) load(s []string) minicode {
	tmp := []int{}

	for _, op := range s {
		tmp = append(tmp, ToInt(op))
	}

	m.mem = tmp

	return m
}

func (m minicode) run(d bool) minicode {

	origDebug := debug
	debug = d
	for {
		dbg("%v", m)
		// halt
		if m.ip > len(m.mem)-2 {
			break
		}

		op := m.mem[m.ip]
		operand := m.mem[m.ip+1]
		comboOperand := m.comboOperand(m.mem[m.ip+1])
		dbg(" - op: %v, operand, %v, comboOperand: %v", op, operand, comboOperand)

		switch op {

		// adv: A <- A / 2^combo op
		case 0:
			m.r["A"] = m.r["A"] >> comboOperand
			m.ip += 2
		// bxl: B <- XOR B, literal op
		case 1:
			m.r["B"] ^= operand
			m.ip += 2

		// bst: B <- combo modulo 8
		case 2:
			dbg("bst %v", comboOperand%8)
			m.r["B"] = (comboOperand % 8)
			m.ip += 2

		// jnz: IP = literal if A != 0
		case 3:
			if m.r["A"] != 0 {
				m.ip = operand
				break
			}
			m.ip++

		// bxc: B <- B XOR C
		case 4:
			m.r["B"] ^= m.r["C"]
			m.ip += 2

		// out: print combo mod 8
		case 5:
			m.buf = append(m.buf, fmt.Sprintf("%d", comboOperand%8))
			m.ip += 2

		// bdv: B <- A / 2^combo op
		case 6:
			m.r["B"] = m.r["A"] >> comboOperand
			m.ip += 2

		// cdv: B <- A / 2^combo op
		case 7:
			m.r["C"] = m.r["A"] >> comboOperand
			m.ip += 2
		}

	}
	debug = origDebug

	return m
}

func (m minicode) comboOperand(idx int) int {
	ret := 0
	dbg(" -> ComboOperand(%v)", idx)
	switch idx {
	case 0, 1, 2, 3:
		ret = idx
		dbg("  -> Literal %v", ret)
	case 4:
		ret = m.r["A"]
		dbg("  -> A %v", ret)
	case 5:
		ret = m.r["B"]
		dbg("  -> B %v", ret)
	case 6:
		ret = m.r["C"]
		dbg("  -> C %v", ret)

	}

	return ret
}

func parseProgram(s []string) (minicode, []string) {
	mc := minicode{
		ip: 0,
		r: map[string]int{
			"A": 0,
			"B": 0,
			"C": 0,
		},
		mem: []int{},
		buf: []string{},
	}
	var prog []string

	reReg := regexp.MustCompile(`Register ([A-C]): (\d+)`)

	for _, l := range s {
		if l == "" {
			continue
		}

		if strings.HasPrefix(l, "Register") {
			parts := reReg.FindStringSubmatch(l)

			mc.r[parts[1]] = ToInt(parts[2])

			continue
		} else { //program
			parts := strings.Fields(l)

			prog = strings.Split(parts[1], ",")

		}
	}

	return mc, prog
}

func solve1(s []string) string {
	mc, program := parseProgram(s)
	dbg("MC: %+v", mc)
	mc = mc.load(program)
	mc = mc.run(false)

	dbg("MC: %v", mc)
	out := mc.flush()

	return out
}

func solve2(s []string) int {
	res := 0
	mc, program := parseProgram(s)

	a := 0
	for proglen := len(program) - 1; proglen >= 0; proglen-- {
		a <<= 3
		dbg("ori: %v", program)
		dbg("A: %v, n: %v", a, proglen)

		m := mc
		m.r["A"] = a
		m = m.load(program)
		m = m.run(false)
		out := m.buf
		dbg("run: %q", out)
		dbg("lst: %q", program[proglen:])

		for !slices.Equal(out, program[proglen:]) {
			a++
			dbg(" -> A == %v", a)
			m := mc
			m.r["A"] = a
			m = m.load(program)
			m = m.run(false)
			out = m.buf
			dbg(" -> run: %v", out)
			dbg(" -> lst: %v", program[proglen:])
		}

	}
	res = a

	return res
}
