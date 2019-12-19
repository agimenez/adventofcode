package main

import (
	"os"
	"testing"
)

var tests = []struct {
	testfile     string
	numAsteroids int
	best         Asteroid
	bestSight    int
}{
	{
		"testdata/test1.txt",
		10,
		Asteroid{3, 4},
		8,
	},

	{
		"testdata/test2.txt",
		40,
		Asteroid{5, 8},
		33,
	},

	{
		"testdata/test3.txt",
		40,
		Asteroid{1, 2},
		35,
	},

	{
		"testdata/test4.txt",
		50,
		Asteroid{6, 3},
		41,
	},
	{
		"testdata/test5.txt",
		300,
		Asteroid{11, 13},
		210,
	},
}

func TestParser(t *testing.T) {

	for _, tst := range tests {
		file, err := os.Open(tst.testfile)
		if err != nil {
			t.Error(err)
		}
		m := parseInput(file)
		// Basic check: the winner asteroid should exist in the structure
		if _, ok := m[tst.best]; !ok {
			t.Errorf("Asteroid %v should exist in the map (not found)", tst.best)
		}

		if len(m) != tst.numAsteroids {
			t.Errorf("Parsed wrong number of asteroids (got %d, want %d)", len(m), tst.numAsteroids)
		}

	}
}

func TestBest(t *testing.T) {
	for _, tst := range tests {
		file, err := os.Open(tst.testfile)
		if err != nil {
			t.Error(err)
		}
		m := parseInput(file)
		m.calculateAllSights()
		a := m.getBestLocation()

		if a != tst.best {
			t.Errorf("Best sight expected %v, got %v", tst.best, a)
		}

		if len(m[a]) != tst.bestSight {
			t.Errorf("Num sights expected %d, got %d", tst.bestSight, len(m[a]))
		}

	}
}
