package main

import "testing"

func TestCountSafe(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		rows int
		want int
	}{
		{"..^^.", 3, 6},
		{".^^.^.^^^^", 10, 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountSafe(tt.name, tt.rows)
			if got != tt.want {
				t.Errorf("CountSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextRow(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
	}{
		{"..^^.", ".^^^^"},
		{".^^^^", "^^..^"},

		{".^^.^.^^^^", "^^^...^..^"},
		{"^.^^.^.^^.", "..^^...^^^"},
		{".^^^^.^^.^", "^^..^.^^.."},
		{"^^^^..^^^.", "^..^^^^.^^"},
		{".^^^..^.^^", "^^.^^^..^^"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextRow(tt.name)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("NextRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
