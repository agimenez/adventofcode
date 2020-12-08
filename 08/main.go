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

func (m *Machine) Step() {
	fetch := m.mem[m.pc]
	instr := strings.Split(fetch, " ")
	op := instr[0]
	arg, err := strconv.Atoi(instr[1])
	if err != nil {
		panic("error parsing argument")
	}

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

func runUntilLoop(m *Machine) {
	seen := map[int]bool{}
	for {
		if _, ok := seen[m.pc]; ok {
			return
		}
		seen[m.pc] = true
		m.Step()
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
	dbg("%#v", m)
	runUntilLoop(m)
	part1 = m.acc

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
