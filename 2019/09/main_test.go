package main

import (
	"fmt"
	"testing"
)

type progTest struct {
	op   string
	code string
	in   int
	out  int
}

func testProgram(t *testing.T, pt *progTest) {

	t.Helper()

	p := newProgram(pt.code)
	input := make(chan int)
	output := make(chan int)
	go p.run(input, output)
	input <- pt.in
	out := <-output
	if out != pt.out {
		t.Errorf("Got %d, expected %d", out, pt.out)
	}
}

func TestEQ(t *testing.T) {

	tests := []progTest{
		// position mode EQ 8
		{
			op:   "EQ P8",
			code: "3,9,8,9,10,9,4,9,99,-1,8",
			in:   8,
			out:  1,
		},
		{
			op:   "EQ P8",
			code: "3,9,8,9,10,9,4,9,99,-1,8",
			in:   7,
			out:  0,
		},
		{
			op:   "EQ P8",
			code: "3,9,8,9,10,9,4,9,99,-1,8",
			in:   9,
			out:  0,
		},

		// immediate mode EQ 8
		{
			op:   "EQ I8",
			code: "3,3,1108,-1,8,3,4,3,99",
			in:   7,
			out:  0,
		},
		{
			op:   "EQ I8",
			code: "3,3,1108,-1,8,3,4,3,99",
			in:   8,
			out:  1,
		},
		{
			op:   "EQ I8",
			code: "3,3,1108,-1,8,3,4,3,99",
			in:   9,
			out:  0,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %d", test.op, test.in), func(t *testing.T) {
			testProgram(t, &test)
		})
	}
}

func TestLT(t *testing.T) {

	tests := []progTest{
		// position mode LT 8
		{
			op:   "LT P8",
			code: "3,9,7,9,10,9,4,9,99,-1,8",
			in:   7,
			out:  1,
		},
		{
			op:   "LT P8",
			code: "3,9,7,9,10,9,4,9,99,-1,8",
			in:   8,
			out:  0,
		},
		{
			op:   "LT P8",
			code: "3,9,7,9,10,9,4,9,99,-1,8",
			in:   9,
			out:  0,
		},

		// immediate mode LT 8
		{
			op:   "LT I8",
			code: "3,3,1107,-1,8,3,4,3,99",
			in:   7,
			out:  1,
		},
		{
			op:   "LT I8",
			code: "3,3,1107,-1,8,3,4,3,99",
			in:   8,
			out:  0,
		},
		{
			op:   "LT I8",
			code: "3,3,1107,-1,8,3,4,3,99",
			in:   9,
			out:  0,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %d", test.op, test.in), func(t *testing.T) {
			testProgram(t, &test)
		})
	}
}

func TestJMP(t *testing.T) {
	tests := []progTest{
		// position mode, 0 IFF input 0
		{
			op:   "JMP 0 P",
			code: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:   0,
			out:  0,
		},
		{
			op:   "JMP 0 P",
			code: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:   1,
			out:  1,
		},

		// immediate mode, 0 IFF input 0
		{
			op:   "JMP 0 I",
			code: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:   0,
			out:  0,
		},
		{
			op:   "JMP 0 I",
			code: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:   1,
			out:  1,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %d", test.op, test.in), func(t *testing.T) {
			testProgram(t, &test)
		})
	}
}

func TestIntegration(t *testing.T) {
	tests := []progTest{
		// if input < 8: 999;  if input == 0: 1000; if input > 8: 1001
		{
			op:   "CMP < 8",
			code: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:   7,
			out:  999,
		},
		{
			op:   "CMP == 8",
			code: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:   8,
			out:  1000,
		},
		{
			op:   "CMP > 8",
			code: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:   9,
			out:  1001,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %d", test.op, test.in), func(t *testing.T) {
			testProgram(t, &test)
		})
	}
}

func TestAmps(t *testing.T) {
	tests := []struct {
		settings []int
		program  string
		max      int
	}{
		{
			settings: []int{9, 8, 7, 6, 5},
			program:  "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
			max:      139629729,
		},
		{
			settings: []int{9, 7, 8, 5, 6},
			program:  "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
			max:      18216,
		},
	}

	for _, test := range tests {
		max := runPermutation(test.program, test.settings)
		if max != test.max {
			t.Errorf("Got %d, expected %d", max, test.max)
		}
	}
}
