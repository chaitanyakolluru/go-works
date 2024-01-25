package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{99, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.x, tt.want)

		t.Run(testname, func(t *testing.T) {
			got := isPalindrome(tt.x)

			if got != tt.want {
				t.Errorf("got %v, want: %v", got, tt.want)

			}
		})
	}
}
