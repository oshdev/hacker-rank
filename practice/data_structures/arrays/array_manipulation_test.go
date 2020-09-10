package arrays_test

import (
	"fmt"
	"github.com/oshdev/hacker-rank/practice/data_structures/arrays"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	testCases := []struct {
		InputN   int32
		InputQ   [][]int32
		Expected int64
	}{
		{5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}, 200},
		{10, [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}, 10},
		{10, [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}}, 31},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v => %d", testCase.InputQ, testCase.Expected), func(t *testing.T) {
			got := arrays.ArrayManipulation(testCase.InputN, testCase.InputQ)
			if testCase.Expected != got {
				t.Errorf("Got %d, want %d", got, testCase.Expected)
			}
		})
	}
}
