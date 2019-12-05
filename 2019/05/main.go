package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	debug = true
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

type program struct {
	mem    []int
	pc     int
	output []int
}

func newProgram(p string) *program {

	pr := &program{
		mem:    []int{},
		pc:     0,
		output: []int{},
	}

	pSlice := strings.Split(p, ",")
	for _, b := range pSlice {
		i, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}

		pr.mem = append(pr.mem, i)
	}

	return pr
}

func (p *program) run(input []int) {

	for op := p.mem[p.pc]; op != 99; {
		dbg("op = %v", op)

		opcode := op % 100
		dbg("opcode = %v", opcode)

		switch opcode {
		case 1: // Add
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			p.mem[c] = a + b
			p.pc += 4
		case 2: // Mul
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			p.mem[c] = a * b
			p.pc += 4
		case 3: // In
			var in int
			in, input = input[0], input[1:]
			p.mem[p.pc+1] = in
			p.pc += 2
		case 4: // Out
			p.output = append(p.output, p.mem[p.pc+1])
			p.pc += 2
		default:
			log.Fatalf("Bad opcode = %v", op)
		}

		op = p.mem[p.pc]
	}
}

func (p *program) fetchParameter(n int) int {
	opcode := p.mem[p.pc]
	parameter := p.mem[p.pc+n]
	mode := opcode / int(math.Pow10(n+1)) % 10

	dbg("mode: %d", mode)
	if mode == 0 {
		return p.mem[parameter]
	}

	return parameter
}

func main() {
	var in string

	fmt.Scan(&in)
	program := newProgram(in)

	program.run([]int{1})

	log.Printf("%v", program.output)
}
