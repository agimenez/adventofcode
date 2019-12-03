package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	OpSum = 1
	OpMul = 2
)

func main() {
	var in string

	fmt.Scan(&in)
	program := parseProgram(in)

	initialize(program)
	run(program)

	log.Printf("p[0] = %d\n", program[0])
}

func run(p []int) {
	pc := 0

	for op := p[pc]; op != 99; {
		log.Printf("program = %v", p)
		log.Printf("fetch= %v", p[pc:pc+4])
		log.Printf("pc = %d, op = %d", pc, op)
		addr1 := p[pc+1]
		addr2 := p[pc+2]
		dest := p[pc+3]

		op1 := p[addr1]
		op2 := p[addr2]

		if op == OpSum {
			log.Printf("%d + %d into addr %d\n", op1, op2, dest)
			p[dest] = op1 + op2
		} else if op == OpMul {
			log.Printf("%d * %d into addr %d\n", op1, op2, dest)
			p[dest] = op1 * op2
		}

		pc += 4
		op = p[pc]
		log.Printf("NEW pc = %d, op = %d", pc, op)
	}
}

func initialize(p []int) {
	p[1] = 12
	p[2] = 2
}

func parseProgram(p string) []int {
	var bytecode []int

	pSlice := strings.Split(p, ",")
	for _, b := range pSlice {
		i, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}

		bytecode = append(bytecode, i)
	}

	return bytecode
}
