package utils

import (
	"flag"
	"fmt"
	"iter"
	"log"
	"strconv"
	"strings"
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

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

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

func (p Point) Sum(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

func (p Point) Neg() Point {
	return Point{
		X: -p.X,
		Y: -p.Y,
	}
}

func (p Point) Sub(p2 Point) Point {
	return p.Sum(p2.Neg())
}

func (p Point) ManhattanDistance(p2 Point) int {
	return Abs(p.X-p2.X) + Abs(p.Y-p2.Y)
}

func (p Point) Rotate90CW() Point {
	return Point{-p.Y, p.X}
}

func (p Point) Rotate90CCW() Point {
	return Point{p.Y, -p.X}
}

func (p Point) Adjacent(diagonals bool) iter.Seq[Point] {
	return func(yield func(p Point) bool) {
		dirs := []Point{
			p.Up(),
			p.Right(),
			p.Down(),
			p.Left(),
		}

		if diagonals {
			dirs = append(dirs, []Point{
				p.Up().Right(),
				p.Down().Right(),
				p.Down().Left(),
				p.Left().Up(),
			}...)
		}

		for _, p := range dirs {
			if !yield(p) {
				return
			}
		}
	}
}

func GetChInPoint(s []string, p Point) (byte, bool) {
	if p.Y > len(s)-1 || p.Y < 0 || p.X < 0 || p.X > len(s[0])-1 {
		return ' ', false
	}

	return s[p.Y][p.X], true
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

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func SliceLCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = LCM(result, numbers[i])
	}
	return result
}

func ToInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

func LinesToIntSlice(s []string) []int {
	res := []int{}
	for _, line := range s {
		res = append(res, ToInt(line))
	}

	return res
}

func CSVToIntSlice(s string, sep string) []int {
	res := []int{}
	for _, p := range strings.Split(s, sep) {
		n, _ := strconv.Atoi(p)
		res = append(res, n)
	}

	return res
}

type ReduceFunc[T, R any] func(R, T) R

var ReduceSum = func(acc, n int) int {
	return acc + n
}
var ReduceMult = func(acc, n int) int {
	return acc * n
}

func Reduce[T, R any](s []T, init R, fn ReduceFunc[T, R]) R {
	res := init
	for i := range s {
		res = fn(res, s[i])
	}

	return res
}

func ReduceCollect[T, R any](seq iter.Seq[T], init R, fn ReduceFunc[T, R]) R {
	res := init
	for i := range seq {
		res = fn(res, i)
	}

	return res
}

type FilterFunc[T any] func(T) bool

func Filter[T any](s []T, fn FilterFunc[T]) []T {
	res := []T{}
	for _, elem := range s {
		if fn(elem) {
			res = append(res, elem)
		}
	}

	return res
}
