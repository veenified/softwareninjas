package main

import (
	"testing"
	"bytes"
	"log"
	"github.com/stretchr/testify/assert"
)

func Test_TabDelimited(t *testing.T) {
	var input = "a1\ta2\ta3\ta4\t\n" +
		"b1\tb2\tb3\tb4\t\n" +
		"c1\tc2\tc3\tc4\t"

	var expectedOutput = [][]string{
		{"a1", "a3"},
		{"b1", "b3"},
		{"c1", "c3"},
	}

	activityReaderValidator(t, expectedOutput, input, '\t', 0, 2)
}

func Test_SpaceDelimited(t *testing.T) {
	var input = "a1 a2 a3 a4\n" +
		"b1 b2 b3 b4\n" +
		"c1 c2 c3 c4"

	var expectedOutput = [][]string{
		{"a1", "a3"},
		{"b1", "b3"},
		{"c1", "c3"},
	}

	activityReaderValidator(t, expectedOutput, input, ' ', 0, 2)
}

func activityReaderValidator(t *testing.T, expectedOutput [][]string, input string, delimiter rune, fields ...int) {
	assert := assert.New(t)
	buf := bytes.NewBuffer([]byte(input))
	r := NewActivityReader(buf, fields...)
	r.Comma = delimiter
	recs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, rec := range recs {
		for j, field := range rec {
			assert.Equal(expectedOutput[i][j], field)
		}
	}
}
