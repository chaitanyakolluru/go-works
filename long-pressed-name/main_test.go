package main

import (
	"fmt"
	"testing"
)

func TestIsLongPressedName(t *testing.T) {
	tests := []struct {
		name  string
		typed string
		want  bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"alex", "aaleexa", false},
		{"ppyplrza", "pyypllrza", false},
		{"alex", "aaleexeex", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %s %v", tt.name, tt.typed, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := isLongPressedName(tt.name, tt.typed)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}

}
