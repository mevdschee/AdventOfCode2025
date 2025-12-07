package main

import (
	"bufio"
	"fmt"
	"os"
)

const LINE = 1000

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rolls := make(map[int]bool)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			rolls[y*LINE+x] = line[x] == '@'
		}
		y++
	}

	total := 0
	for {
		deleted := 0
		for xy := range rolls {
			x := xy % LINE
			y := xy / LINE
			if !rolls[xy] {
				continue
			}
			// loop over 8 adjacent positions
			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					if rolls[(y+dy)*LINE+(x+dx)] {
						count++
					}
				}
			}
			if count < 4 {
				deleted++
				delete(rolls, xy)
				total++
			}
		}
		if deleted == 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
