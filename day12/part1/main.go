package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape [][]bool

type Region struct {
	width, height int
	presents      []int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	shapes := []Shape{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, ":") {
			shape := Shape{}
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}
				row := make([]bool, len(line))
				for i, c := range line {
					row[i] = c == '#'
				}
				shape = append(shape, row)
			}
			shapes = append(shapes, shape)
		}
	}

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			regions := []Region{}
			parts := strings.Split(line, ": ")
			dims := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(dims[0])
			height, _ := strconv.Atoi(dims[1])
			presents := []int{}
			for _, p := range strings.Fields(parts[1]) {
				count, _ := strconv.Atoi(p)
				presents = append(presents, count)
			}
			regions = append(regions, Region{width, height, presents})

			for scanner.Scan() {
				line := scanner.Text()
				if line != "" {
					parts := strings.Split(line, ": ")
					dims := strings.Split(parts[0], "x")
					width, _ := strconv.Atoi(dims[0])
					height, _ := strconv.Atoi(dims[1])
					presents := []int{}
					for _, p := range strings.Fields(parts[1]) {
						count, _ := strconv.Atoi(p)
						presents = append(presents, count)
					}
					regions = append(regions, Region{width, height, presents})
				}
			}

			validCount := 0
			for _, region := range regions {
				if fits(region, shapes) {
					validCount++
				}
			}

			fmt.Println(validCount)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func fits(region Region, shapes []Shape) bool {
	grid := make([][]bool, region.height)
	for i := range grid {
		grid[i] = make([]bool, region.width)
	}

	presents := []Shape{}
	for i, count := range region.presents {
		for range count {
			presents = append(presents, shapes[i])
		}
	}

	totalArea := 0
	for _, p := range presents {
		for _, row := range p {
			for _, cell := range row {
				if cell {
					totalArea++
				}
			}
		}
	}
	if totalArea > region.width*region.height {
		return false
	}

	return dfs(grid, presents, 0)
}

func dfs(grid [][]bool, presents []Shape, i int) bool {
	if i == len(presents) {
		return true
	}

	shape := presents[i]

	for _, rotated := range orientations(shape) {
		for r := 0; r <= len(grid)-len(rotated); r++ {
			for c := 0; c <= len(grid[0])-len(rotated[0]); c++ {
				canPlace := true
				for dr := range len(rotated) {
					if !canPlace {
						break
					}
					for dc := range len(rotated[dr]) {
						if rotated[dr][dc] && grid[r+dr][c+dc] {
							canPlace = false
							break
						}
					}
				}
				if canPlace {
					for dr := range len(rotated) {
						for dc := range len(rotated[dr]) {
							if rotated[dr][dc] {
								grid[r+dr][c+dc] = true
							}
						}
					}
					if dfs(grid, presents, i+1) {
						return true
					}
					for dr := range len(rotated) {
						for dc := range len(rotated[dr]) {
							if rotated[dr][dc] {
								grid[r+dr][c+dc] = false
							}
						}
					}
				}
			}
		}
	}

	return false
}

func orientations(shape Shape) []Shape {
	seen := make(map[string]bool)
	orientations := []Shape{}

	add(shape, seen, &orientations)

	for range 3 {
		shape = rotate(shape)
		add(shape, seen, &orientations)
	}

	shape = flip(orientations[0])
	add(shape, seen, &orientations)
	for range 3 {
		shape = rotate(shape)
		add(shape, seen, &orientations)
	}

	return orientations
}

func rotate(shape Shape) Shape {
	if len(shape) == 0 {
		return shape
	}
	rows := len(shape)
	cols := len(shape[0])
	rotated := make(Shape, cols)
	for i := range rotated {
		rotated[i] = make([]bool, rows)
	}
	for r := range rows {
		for c := range cols {
			rotated[c][rows-1-r] = shape[r][c]
		}
	}
	return rotated
}

func flip(shape Shape) Shape {
	flipped := make(Shape, len(shape))
	for i := range shape {
		flipped[i] = make([]bool, len(shape[i]))
		for j := range shape[i] {
			flipped[i][len(shape[i])-1-j] = shape[i][j]
		}
	}
	return flipped
}

func add(s Shape, seen map[string]bool, orientations *[]Shape) {
	var sb strings.Builder
	for _, row := range s {
		for _, cell := range row {
			if cell {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		sb.WriteByte('|')
	}
	key := sb.String()
	if !seen[key] {
		seen[key] = true
		*orientations = append(*orientations, s)
	}
}
