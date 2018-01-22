package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main() {
	filename := flag.String("input", "time.txt", "a file in the format of 'begintime endtime'")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s\n", *filename))
	}
	br := bufio.NewReader(file)
	r := NewActivityReader(br, 0, 1)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided file.")
	}
	activities := parseLines(lines)

	unprodictiveHrs := 0
	for i, a := range activities {
		fmt.Printf("Activity #%d: Start: %d End: %d\n", i+1, a.startHr, a.endHr)
	}

	fmt.Printf("Unproductive Hrs: %d\n", unprodictiveHrs)
}

func parseLines(lines [][]string) []Activity {
	ret := make([]Activity, len(lines))
	for i, line := range lines {
		s, err := strconv.Atoi(line[0])
		e, err := strconv.Atoi(line[1])
		if err != nil {
			exit(fmt.Sprintf("Failed to convert to integers at line #%d.", i+1))
		}
		ret[i] = Activity{
			startHr: s,
			endHr: e,
		}
	}
	return ret
}

type Activity struct {
	startHr int
	endHr int
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}