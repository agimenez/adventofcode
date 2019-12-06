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
	total := 0
	for orbit := range orbits {
		parent := orbit

		for parent != "COM" {
			total++
			parent = orbits[parent]

		}

	}

	log.Printf("%v", total)
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
