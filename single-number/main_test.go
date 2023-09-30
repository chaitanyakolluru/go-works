package main

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.nums, tt.want)

		t.Run(testname, func(t *testing.T) {
			got := singleNumber(tt.nums)

			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
