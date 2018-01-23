package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sort"
)

func Test_SortByTimesPrimary(t *testing.T) {
	var input = [][]int{
		{2, 4},
		{3, 6},
		{1, 3},
		{6, 8},
	}
	var expectedOutput = [][]int{
		{1, 3},
		{2, 4},
		{3, 6},
		{6, 8},
	}

	activtySortValidator(t, expectedOutput, input)
}

func Test_SortByTimesSecondary(t *testing.T) {
	var input = [][]int{
		{2, 4},
		{2, 3},
		{2, 2},
		{1, 4},
	}
	var expectedOutput = [][]int{
		{1, 4},
		{2, 2},
		{2, 3},
		{2, 4},
	}

	activtySortValidator(t, expectedOutput, input)
}

func activtySortValidator(t *testing.T, expectedOutput, input [][]int) {
	assert := assert.New(t)
	acts := make([]Activity, len(input))
	for i, in := range input {
		acts[i] = Activity{
			startHr: in[0],
			endHr:   in[1],
		}
	}
	sort.Sort(ByTimes{acts})
	for i, a := range acts {
		assert.Equal(expectedOutput[i][0], a.startHr)
		assert.Equal(expectedOutput[i][1], a.endHr)
	}
}
