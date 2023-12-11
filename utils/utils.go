package utils

import (
	"flag"
	"fmt"
	"log"
)

var (
	Debug int
)

func Dbg(level int, fmt string, v ...interface{}) {
	if Debug >= level {
		log.Printf(fmt+"\n", v...)
	}
}

type Point struct {
	X, Y int
}

var P0 = Point{0, 0}

func (p Point) Min(p2 Point) Point {
	return Point{
		X: Min(p.X, p2.X),
		Y: Min(p.Y, p2.Y),
	}
}

func (p Point) Max(p2 Point) Point {
	return Point{
		X: Max(p.X, p2.X),
		Y: Max(p.Y, p2.Y),
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) Up() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) Left() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

func (p Point) ManhattanDistance(p2 Point) int {
	return Abs(p.X-p2.X) + Abs(p.Y-p2.Y)
}

func init() {
	flag.IntVar(&Debug, "debug-level", 0, "debug level")
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
