package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{2, 5, 5, 11}, 10, []int{1, 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			res := twoSum(tt.nums, tt.target)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("got: %v, want: %v", res, tt.want)
			}
		})
	}

}
