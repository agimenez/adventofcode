package main

import "testing"

func TestEQ(t *testing.T) {

	tests := []struct {
		code string
		in   int
		out  int
	}{
		// position mode EQ 8
		{
			code: "3,9,8,9,10,9,4,9,99,-1,8",
			in:   8,
			out:  1,
		},
		{
			code: "3,9,8,9,10,9,4,9,99,-1,8",
			in:   7,
			out:  0,
		},
		{
			code: "3,9,8,9,10,9,4,9,99,-1,8",
			in:   9,
			out:  0,
		},

		// immediate mode EQ 8
		{
			code: "3,3,1108,-1,8,3,4,3,99",
			in:   7,
			out:  0,
		},
		{
			code: "3,3,1108,-1,8,3,4,3,99",
			in:   8,
			out:  1,
		},
		{
			code: "3,3,1108,-1,8,3,4,3,99",
			in:   8,
			out:  0,
		},
	}

	for _, test := range tests {
		p := newProgram(test.code)
		p.run([]int{test.in})
		t.Logf("output array: %v", p.output)
	}
}

func TestLT(t *testing.T) {

	tests := []struct {
		code string
		in   int
		out  int
	}{
		// position mode LT 8
		{
			code: "3,9,7,9,10,9,4,9,99,-1,8",
			in:   7,
			out:  1,
		},
		{
			code: "3,9,7,9,10,9,4,9,99,-1,8",
			in:   8,
			out:  0,
		},
		{
			code: "3,9,7,9,10,9,4,9,99,-1,8",
			in:   9,
			out:  0,
		},

		// immediate mode LT 8
		{
			code: "3,3,1107,-1,8,3,4,3,99",
			in:   7,
			out:  1,
		},
		{
			code: "3,3,1107,-1,8,3,4,3,99",
			in:   8,
			out:  0,
		},
		{
			code: "3,3,1107,-1,8,3,4,3,99",
			in:   9,
			out:  0,
		},
	}

	for _, test := range tests {
		p := newProgram(test.code)
		p.run([]int{test.in})
		t.Logf("output array: %v", p.output)
	}
}

func testJMP(t *testing.T) {
	tests := []struct {
		code string
		in   int
		out  int
	}{
		// position mode, 0 IFF input 0
		{
			code: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:   0,
			out:  0,
		},
		{
			code: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:   1,
			out:  1,
		},

		// immediate mode, 0 IFF input 0
		{
			code: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:   0,
			out:  0,
		},
		{
			code: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:   1,
			out:  1,
		},
	}
	for _, test := range tests {
		p := newProgram(test.code)
		p.run([]int{test.in})
		t.Logf("output array: %v", p.output)
	}
}

func testIntegration(t *testing.T) {
	tests := []struct {
		code string
		in   int
		out  int
	}{
		// if input < 8: 999;  if input == 0: 1000; if input > 8: 1001
		{
			code: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:   7,
			out:  999,
		},
		{
			code: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:   8,
			out:  1000,
		},
		{
			code: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:   9,
			out:  1001,
		},
	}
	for _, test := range tests {
		p := newProgram(test.code)
		p.run([]int{test.in})
		t.Logf("output array: %v", p.output)
	}
}