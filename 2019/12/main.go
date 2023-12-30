package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (p Point) Energy() int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func calculateEnergy(moons []Moon) int {
	total := 0
	for _, m := range moons {
		pot := m.Pos.Energy()
		kin := m.Vel.Energy()
		total += pot * kin
	}

	return total
}

func extractField(m []Moon, sub string, field string) []int {
	ret := []int{}
	for i := range m {
		switch sub + field {
		case "Px":
			ret = append(ret, m[i].Pos.x)
		case "Py":
			ret = append(ret, m[i].Pos.y)
		case "Pz":
			ret = append(ret, m[i].Pos.z)
		case "Vx":
			ret = append(ret, m[i].Vel.x)
		case "Vy":
			ret = append(ret, m[i].Vel.y)
		case "Vz":
			ret = append(ret, m[i].Vel.z)
		}
	}

	return ret
}
func axisEquals(m1, m2 []Moon, axis string) bool {
	s1 := extractField(m1, "P", axis)
	s2 := extractField(m2, "P", axis)

	s3 := extractField(m1, "V", axis)
	s4 := extractField(m2, "V", axis)

	return slices.Equal(s1, s2) && slices.Equal(s3, s4)
}

func main() {
	part1 := 0
	part2 := 0

	moons := getMoons(os.Stdin)
	origMoons := slices.Clone(moons)
	axisCycles := map[string]int{}

	for i := 1; ; i++ {
		dbg(2, "step %d", i)
		timeStep(moons)
		if i == 1000 {
			part1 = calculateEnergy(moons)
		}

		dbg(2, "Moons: %v", moons)
		dbg(2, "Orig : %v", origMoons)

		for _, axis := range []string{"x", "y", "z"} {
			if _, ok := axisCycles[axis]; !ok && axisEquals(moons, origMoons, axis) {
				axisCycles[axis] = i
				dbg(1, "Found axis cycle (%q): %d", axis, i)
				dbg(1, "Moons: %v", moons)
				dbg(1, "Orig : %v", origMoons)

			}
		}
		dbg(2, "Found cycles: %v", axisCycles)

		if len(axisCycles) == 3 {
			break
		}
	}

	part2 = sliceLCM(func() []int {
		ret := []int{}
		for _, v := range axisCycles {
			ret = append(ret, v)
		}
		return ret
	}())

	fmt.Printf("Energy: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func sliceLCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}
