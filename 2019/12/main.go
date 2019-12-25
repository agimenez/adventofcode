package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt+"\n", v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

type Point struct {
	x, y, z int
}

type Moon struct {
	Pos Point
	Vel Point
}

func getMoons(in io.Reader) []Moon {
	scanner := bufio.NewScanner(in)
	moons := []Moon{}
	for scanner.Scan() {
		line := scanner.Text()
		dbg(1, "Line: %v", line)
		pattern := regexp.MustCompile(`<x=([-\d]+), y=([-\d]+), z=([-\d]+)`)
		nums := pattern.FindStringSubmatch(line)

		m := Moon{}
		m.Pos.x, _ = strconv.Atoi(nums[1])
		m.Pos.y, _ = strconv.Atoi(nums[2])
		m.Pos.z, _ = strconv.Atoi(nums[3])

		moons = append(moons, m)
	}

	return moons
}

func timeStep(moons []Moon) {
	for i, m1 := range moons {
		for j := i + 1; j < len(moons); j++ {
			if i == j {
				continue
			}
			m2 := moons[j]

			// Compare m1 with m2
			if m1.Pos.x < m2.Pos.x {
				m1.Vel.x++
				m2.Vel.x--
			} else if m1.Pos.x > m2.Pos.x {
				m1.Vel.x--
				m2.Vel.x++
			}

			if m1.Pos.y < m2.Pos.y {
				m1.Vel.y++
				m2.Vel.y--
			} else if m1.Pos.y > m2.Pos.y {
				m1.Vel.y--
				m2.Vel.y++
			}

			if m1.Pos.z < m2.Pos.z {
				m1.Vel.z++
				m2.Vel.z--
			} else if m1.Pos.z > m2.Pos.z {
				m1.Vel.z--
				m2.Vel.z++
			}

			moons[i] = m1
			moons[j] = m2
		}

		moons[i].Pos = m1.Pos.Add(m1.Vel)
		dbg(2, "Moon[%d] = %v", i, moons[i])
	}
}

func (p Point) Add(p2 Point) Point {
	return Point{
		x: p.x + p2.x,
		y: p.y + p2.y,
		z: p.z + p2.z,
	}
}
func main() {
	moons := getMoons(os.Stdin)

	for i := 1; i <= 10; i++ {
		dbg(2, "step %d", i)
		timeStep(moons)
	}
}
