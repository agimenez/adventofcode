package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt, v...)
	}
}

type Asteroid struct {
	x, y int
}

type Map map[Asteroid][]Asteroid

func parseInput(in io.Reader) Map {
	m := Map{}
	scanner := bufio.NewScanner(in)
	line := 0
	for scanner.Scan() {
		row := scanner.Text()
		for i, c := range row {
			dbg(2, "Pos (%d, %d): %c\n", i, line, c)
			if c != '#' {
				continue
			}
			dbg(1, "Asteroid @ (%d, %d): %c\n", i, line, c)

			a := Asteroid{
				y: line,
				x: i,
			}

			m[a] = []Asteroid{}

		}
		line++
	}

	return m
}

func main() {

	m := parseInput(os.Stdin)
	fmt.Printf("%v\n", m)

}
