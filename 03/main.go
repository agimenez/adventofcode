package main

import (
	"bufio"
	"log"
	"os"
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
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := s.Text()

	}

	log.Printf("Part 1: \n")
	log.Printf("Part 2: \n")

}
