package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{`""`, ""},
		{`"abc"`, "abc"},
		{`"aaa\"aaa"`, `aaa"aaa`},
		{`"\x27"`, `'`},
	}

	for _, tt := range tests {
		out := unescapeString(tt.in)

		if out != tt.out {
			t.Errorf("Codelen: got %q, expected %q", out, tt.out)
		}
	}

}
