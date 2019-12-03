package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var in string

	fmt.Scan(&in)
	program := parseProgram(in)

	initialize(program)

	log.Printf("p[0] = %d\n", program[0])
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
