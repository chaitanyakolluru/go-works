package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.s, tt.want)

		t.Run(testname, func(t *testing.T) {
			got := isPalindrome(tt.s)

			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
