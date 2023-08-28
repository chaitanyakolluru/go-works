package main

import "testing"

type argStructure struct {
	nums   []int
	target int
}

type caseStructure struct {
	args argStructure
	want int
}

func TestThreeSumClosest(t *testing.T) {
	cases := map[string]caseStructure{

		"should give 2": {
			args: argStructure{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		"should give 2 with 1,1,1,0": {
			args: argStructure{
				nums:   []int{1, 1, 1, 0},
				target: -100,
			},
			want: 2,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := threeSumClosest(tc.args.nums, tc.args.target)
			if got != tc.want {
				t.Errorf("got: %d, want: %d", got, tc.want)
			}
		})
	}

}
