package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/agimenez/adventofcode/utils"
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
	up = iota
	right
	down
	left
)

var dir2str = map[direction]string{
	up:    "up",
	right: "right",
	down:  "down",
	left:  "left",
}

type vector struct {
	utils.Point

	dir direction
}

func (v vector) String() string {
	return v.Point.String() + fmt.Sprintf(" [%s]", dir2str[v.dir])
}

type moveFunc func(utils.Point) utils.Point

var dir2op = map[direction]moveFunc{
	up:    utils.Point.Up,
	right: utils.Point.Right,
	down:  utils.Point.Down,
	left:  utils.Point.Left,
}

var validDirections map[direction]map[rune]direction = map[direction]map[rune]direction{
	up:    map[rune]direction{'|': up, 'F': right, '7': left},
	right: map[rune]direction{'J': up, '-': right, '7': down},
	down:  map[rune]direction{'J': left, '|': down, 'L': right},
	left:  map[rune]direction{'L': up, '-': left, 'F': down},
}

func getFirstDirection(f map[utils.Point]rune, cur utils.Point) vector {
	v := vector{
		Point: cur,
	}

	dbg("FINDING FIRST DIRECTION (%v '%c')", cur, f[cur])
	for _, dir := range []direction{up, right, down, left} {
		next := dir2op[dir](cur)
		dbg("  -> probing %s -> %v '%c'", dir2str[dir], next, f[next])
		_, ok := validDirections[dir][f[next]]
		if ok {
			v.dir = dir
			break
		}
	}

	return v
}

func nextPoint(f map[utils.Point]rune, cur vector) vector {
	dbg("NEXT FOR %v ('%c')", cur, f[cur.Point])

	np := dir2op[cur.dir](cur.Point)
	next := vector{
		Point: np,
		dir:   validDirections[cur.dir][f[np]],
	}
	dbg("  -> NEXT: %v ('%c')", next, f[next.Point])

	return next
}
func solve1(f map[utils.Point]rune, s utils.Point) map[utils.Point]direction {
	dbg("PART1: %v (%c)", s, f[s])
	start := getFirstDirection(f, s)
	path := map[utils.Point]direction{
		s: start.dir,
	}
	next := nextPoint(f, start)
	for next.Point != s {
		path[next.Point] = next.dir
		next = nextPoint(f, next)

	}

	return path
}

// https://www.reddit.com/r/adventofcode/comments/18eza5g/2023_day_10_animated_visualization/
// https://en.wikipedia.org/wiki/Nonzero-rule
func solve2(in []string, f map[utils.Point]rune, s utils.Point) int {
	steps := solve1(f, s)

	totalIn := 0
	for y, l := range in {
		windings := 0
		for x, c := range l {
			if c == 'S' { // I'm not cheating, you're cheating
				windings++
			}
			direction, ok := steps[utils.Point{x, y}]
			if ok {
				if direction == down {
					windings--
				} else {
					switch c {
					case '|':
						if direction == up {
							windings++
						}
					case 'F':
						if direction == right {
							windings++
						}
					case '7':
						if direction == left {
							windings++
						}
					}
				}
			} else {
				if windings != 0 {
					dbg("Found IN: (%v, %v)", x, y)
					totalIn++
				}
			}
		}
	}

	return totalIn
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
	field := map[utils.Point]rune{}
	var start utils.Point
	for y, l := range lines {
		for x, c := range l {
			p := utils.Point{x, y}
			field[p] = c
			if c == 'S' {
				start = p
			}
		}
	}
	part1 = len(solve1(field, start)) / 2
	part2 = solve2(lines, field, start)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
