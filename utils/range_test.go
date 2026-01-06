package utils_test

import (
	"fmt"
	"testing"

	"github.com/agimenez/adventofcode/utils"
)

func TestRange_MergeContiguous(t *testing.T) {
	tests := []struct {
		// Named input parameters for receiver constructor.
		in string
		// Named input parameters for target function.
		in2   string
		want  string
		want2 bool
	}{
		// Overlapping/fully contained or contiguous
		{"0-2", "3-5", "0-5", true},
		{"0-2", "2-5", "0-5", true},
		{"8-17", "12-35", "8-35", true},
		{"7-17", "9-12", "7-17", true},

		// Not overlapping or contigouous
		{"0-3", "5-10", "0-3", false},
		{"0-3", "5-5", "0-3", false},
		{"3-3", "5-5", "3-3", false},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s_%s", tt.in, tt.in2)
		t.Run(name, func(t *testing.T) {
			r := utils.NewRange(tt.in)
			r2 := utils.NewRange(tt.in2)
			m, got2 := r.MergeContiguous(r2)
			got := m.String()

			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("MergeContiguous() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("MergeContiguous() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
