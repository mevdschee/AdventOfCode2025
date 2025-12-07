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

	beams := make(map[int]bool)
	for x := range lines[0] {
		if lines[0][x] == 'S' {
			beams[x] = true
		}
	}

	total := 0
	for y := 1; y < len(lines); y++ {
		line := lines[y]
		newBeams := map[int]bool{}
		for b := range beams {
			if line[b] == '^' {
				total++
				newBeams[b-1] = true
				newBeams[b+1] = true
			} else {
				newBeams[b] = true
			}
		}
		beams = newBeams
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
