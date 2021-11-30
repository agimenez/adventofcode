package utils

import (
	"flag"
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
	x, y int
}

var P0 = Point{0, 0}

func (p Point) Min(p2 Point) Point {
	return Point{
		x: Min(p.x, p2.x),
		y: Min(p.y, p2.y),
	}
}

func (p Point) Max(p2 Point) Point {
	return Point{
		x: Max(p.x, p2.x),
		y: Max(p.y, p2.y),
	}
}

func init() {
	flag.IntVar(&Debug, "debug", 0, "debug level")
	flag.Parse()
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
