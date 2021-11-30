package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
	flag.Parse()
}

type Device struct {
	PubKey   int
	LoopSize int
}

func (d Device) DeriveLoopSize(subject int) Device {
	d2 := d
	val := 1
	for {
		val *= subject
		val %= 20201227
		d2.LoopSize++

		if val == d2.PubKey {
			break
		}
	}

	return d2
}

func (d Device) EncryptionKey(d2 Device) int {

	val := 1

	for i := 0; i < d.LoopSize; i++ {
		val *= d2.PubKey
		val %= 20201227
	}

	return val
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	n, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatalf("Could not parse %s: %v", lines[0], err)
	}
	card := Device{
		PubKey: n,
	}
	card = card.DeriveLoopSize(7)

	n, err = strconv.Atoi(lines[1])
	if err != nil {
		log.Fatalf("Could not parse %s: %v", lines[0], err)
	}
	door := Device{
		PubKey: n,
	}
	door = door.DeriveLoopSize(7)

	part1 = card.EncryptionKey(door)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
