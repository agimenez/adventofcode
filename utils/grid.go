package utils

import (
	"bytes"
	"iter"
	"maps"
)

type Grid struct {
	grid map[Point]rune

	w int
	h int
}

func NewGrid(w, h int) Grid {
	return Grid{
		h:    h,
		w:    w,
		grid: map[Point]rune{},
	}
}

func NewGridFromStr(lines []string) Grid {
	g := Grid{
		h: len(lines),
		w: len(lines[0]),

		grid: map[Point]rune{},
	}

	for y, l := range lines {
		for x, c := range l {
			p := NewPoint(x, y)
			g.grid[p] = c
		}
	}

	return g
}

func (g Grid) Height() int {
	return g.h
}

func (g Grid) Width() int {
	return g.w
}

func (g *Grid) SetRune(p Point, r rune) {
	g.grid[p] = r
}

func (g *Grid) GetRune(p Point) rune {
	return g.grid[p]
}

func (g Grid) Clone() Grid {
	return Grid{
		grid: maps.Clone(g.grid),
		w:    g.w,
		h:    g.h,
	}
}

func (g Grid) String() string {
	var b bytes.Buffer

	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			p := NewPoint(x, y)
			b.WriteRune(g.grid[p])
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')

	return b.String()

}

func (g Grid) OutOfBounds(p Point) bool {
	return p.X > g.w || p.X < 0 || p.Y > g.h || p.Y < 0
}

func (g Grid) AdjacentPoints(p Point, diagonals bool) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for p := range p.Adjacent(diagonals) {
			if g.OutOfBounds(p) {
				continue
			}

			if !yield(p) {
				return
			}
		}
	}
}

// Mapfunc applies f() to all contents of the grid
func (g Grid) MapFunc(f func(r rune)) {
	for _, v := range g.grid {
		f(v)
	}
}
