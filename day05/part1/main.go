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

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		for _, r := range ranges {
			if num >= r.start && num <= r.end {
				total++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
