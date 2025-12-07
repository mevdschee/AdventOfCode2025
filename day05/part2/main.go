package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const LINE = 1000

type Range struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := []Range{}

	for scanner.Scan() {
		line := scanner.Text()
		// split line on minus
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			break
		}
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start: start, end: end})
	}

	for {
		// merge overlapping ranges
		merged := []Range{}
		for _, r := range ranges {
			overlap := false
			for i, mr := range merged {
				if r.start <= mr.end && r.end >= mr.start {
					if r.start < mr.start {
						merged[i].start = r.start
					}
					if r.end > mr.end {
						merged[i].end = r.end
					}
					overlap = true
				}
			}
			if !overlap {
				merged = append(merged, r)
			}
		}
		// if no ranges were merged, we're done
		if len(merged) == len(ranges) {
			break
		}
		ranges = merged
	}

	// calculate total covered range
	total := 0
	for _, r := range ranges {
		total += r.end - r.start + 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
