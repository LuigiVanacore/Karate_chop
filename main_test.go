package main

import (
	"fmt"
	"testing"
)

func TestChop(t *testing.T) {
	var tests = []struct {
		elem     int
		array    []int
		attended int
	}{
		{3, []int{}, -1},
		{3, []int{1}, -1},
		{1, []int{1}, 0},
		{1, []int{1, 3, 5}, 0},
		{3, []int{1, 3, 5}, 1},
		{5, []int{1, 3, 5}, 2},
		{0, []int{1, 3, 5}, -1},
		{2, []int{1, 3, 5}, -1},
		{4, []int{1, 3, 5}, -1},
		{6, []int{1, 3, 5}, -1},
		{1, []int{1, 3, 5, 7}, 0},
		{3, []int{1, 3, 5, 7}, 1},
		{5, []int{1, 3, 5, 7}, 2},
		{7, []int{1, 3, 5, 7}, 3},
		{0, []int{1, 3, 5, 7}, -1},
		{2, []int{1, 3, 5, 7}, -1},
		{4, []int{1, 3, 5, 7}, -1},
		{6, []int{1, 3, 5, 7}, -1},
		{8, []int{1, 3, 5, 7}, -1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("elem %d, array %v", test.elem, test.array)
		t.Run(testname, func(t *testing.T) {
			ans := chop(test.elem, test.array)
			if ans != test.attended {
				t.Errorf("got %d, want %d", ans, test.attended)
			}
		})
	}
}
