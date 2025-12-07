package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	beams := make(map[int]int)
	for x := range lines[0] {
		if lines[0][x] == 'S' {
			beams[x] = 1
		}
	}

	for y := 1; y < len(lines); y++ {
		line := lines[y]
		newBeams := map[int]int{}
		for b, v := range beams {
			if line[b] == '^' {
				newBeams[b-1] = newBeams[b-1] + v
				newBeams[b+1] = newBeams[b+1] + v
			} else {
				newBeams[b] = newBeams[b] + v
			}
		}
		beams = newBeams
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	total := 0
	for _, v := range beams {
		total += v
	}

	fmt.Println(total)
}
