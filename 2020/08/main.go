package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

type Machine struct {
	mem []string
	pc  int
	acc int
}

func NewMachine(program []string) *Machine {
	dbg("prog: %v", program)
	return &Machine{
		pc:  0,
		acc: 0,
		mem: program,
	}

}

func parseInstruction(f string) (string, int) {
	dbg("parse: %v", f)
	instr := strings.Split(f, " ")
	op := instr[0]
	arg, err := strconv.Atoi(instr[1])
	if err != nil {
		panic("error parsing argument")
	}

	return op, arg
}

func (m *Machine) Step() {
	dbg("pc: %v", m.pc)
	op, arg := parseInstruction(m.mem[m.pc])
	dbg("  op: %v, arg: %v", op, arg)

	addr := m.pc
	switch op {
	case "jmp":
		addr += arg
	case "acc":
		m.acc += arg
		addr++
	case "nop":
		addr++
	}

	m.pc = addr
}

func runUntilLoop(m *Machine) bool {
	seen := map[int]bool{}
	for {
		if _, ok := seen[m.pc]; ok {
			return false
		}
		if m.pc >= len(m.mem) {
			return true
		}
		seen[m.pc] = true

		m.Step()
		if m.pc >= len(m.mem) {
			return true
		}
	}
}

func patchInstruction(instr string) string {
	op, _ := parseInstruction(instr)
	switch op {
	case "jmp":
		return strings.Replace(instr, "jmp", "nop", 1)
	case "nop":
		return strings.Replace(instr, "nop", "jmp", 1)
	default:
		return instr
	}
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	m := NewMachine(lines)
	runUntilLoop(m)
	part1 = m.acc
	log.Printf("Part 1: %v\n", part1)

	for i := range lines {
		prog := make([]string, len(lines))
		copy(prog, lines)
		prog[i] = patchInstruction(prog[i])

		m := NewMachine(prog)
		term := runUntilLoop(m)
		if term {
			part2 = m.acc
			break
		}
	}

	log.Printf("Part 2: %v\n", part2)

}
