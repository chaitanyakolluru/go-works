package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"flower", "flower", "flower", "flower"}, "flower"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %s", tt.strs, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
		})
	}
}
