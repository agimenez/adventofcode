package main

import (
	"bufio"
	"log"
	"os"
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

func main() {

	orbits := readSystem()

	myOrbit, distance := orbits["YOU"], 0
	//	santaOrbit, distance := orbits["SAN"], 0

	transfers := make(map[string]int)

	for myOrbit != "COM" {
		transfers[myOrbit] = distance
		myOrbit = orbits[myOrbit]
		distance++
	}

	// Find intersection between our trasnfer path and Santa's
	santaOrbit, distance := orbits["SAN"], 0
	for santaOrbit != "COM" {
		if dist, ok := transfers[santaOrbit]; ok {
			distance += dist
			break
		}

		santaOrbit = orbits[santaOrbit]
		distance++
	}

	log.Printf("%v", distance)
}

func readSystem() map[string]string {

	orbit := make(map[string]string, 100)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		planets := strings.Split(scanner.Text(), ")")
		orbit[planets[1]] = planets[0]
	}

	return orbit
}
