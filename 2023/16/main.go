package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	. "github.com/agimenez/adventofcode/utils"
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
}

type direction int

const (
	Right = iota
	Down
	Left
	Up
)

type beam struct {
	p   Point
	dir direction
}

var dir2str = map[direction]string{
	Right: "Right",
	Down:  "Down",
	Left:  "Left",
	Up:    "Up",
}

func (b beam) String() string {
	return b.p.String() + fmt.Sprintf(" (%s)", dir2str[b.dir])
}

type contraption map[Point]rune

func (c contraption) moveBeam(b beam) beam {
	switch b.dir {
	case Right:
		b.p = b.p.Right()
	case Down:
		b.p = b.p.Down()
	case Left:
		b.p = b.p.Left()
	case Up:
		b.p = b.p.Up()
	}

	return b

}

func (c contraption) print(b beam, energized map[Point]bool) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			p := Point{x, y}
			if c, ok := c[p]; ok {
				if energized[p] {
					c = '#'
				}

				if b.p == p {
					switch b.dir {
					case Right:
						c = '>'
					case Down:
						c = 'v'
					case Left:
						c = '<'
					case Up:
						c = '^'
					}
				}

				fmt.Printf("%c", c)
			}
		}
		println()
	}

}

func (c contraption) energize() map[Point]bool {
	energized := map[Point]bool{}
	beams := []beam{{Point{0, 0}, Right}}
	seenBeams := map[beam]bool{}
	for len(beams) > 0 {
		b := beams[0]
		dbg("Processing beam: %+v", b)
		beams = beams[1:]
		var tile rune
		var ok bool
		if tile, ok = c[b.p]; !ok {
			dbg("  -> Beam outside contraption: continuing")
			continue
		}

		if seenBeams[b] {
			continue
		}

		energized[b.p] = true
		seenBeams[b] = true
		dbg(" -> Tile %c", tile)
		c.print(b, energized)
		switch tile {
		case '.':
			beams = append(beams, c.moveBeam(b))
		case '/':
			switch b.dir {
			case Up:
				b.dir = Right
			case Right:
				b.dir = Up
			case Down:
				b.dir = Left
			case Left:
				b.dir = Down
			}
			beams = append(beams, c.moveBeam(b))

		case '\\':
			switch b.dir {
			case Up:
				b.dir = Left
			case Right:
				b.dir = Down
			case Down:
				b.dir = Right
			case Left:
				b.dir = Up
			}
			beams = append(beams, c.moveBeam(b))

		case '|':
			switch b.dir {
			case Up:
				fallthrough
			case Down:
				beams = append(beams, c.moveBeam(b))
			case Right:
				fallthrough
			case Left:
				b.dir = Up
				beams = append(beams, c.moveBeam(b))
				b.dir = Down
				beams = append(beams, c.moveBeam(b))
			}

		case '-':
			switch b.dir {
			case Right:
				fallthrough
			case Left:
				beams = append(beams, c.moveBeam(b))
			case Up:
				fallthrough
			case Down:
				b.dir = Left
				beams = append(beams, c.moveBeam(b))
				b.dir = Right
				beams = append(beams, c.moveBeam(b))
			}
		}
	}

	return energized

}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	grid := contraption{}
	for y, l := range lines {
		for x, c := range l {
			grid[Point{x, y}] = c
		}
	}

	part1 = len(grid.energize())

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
