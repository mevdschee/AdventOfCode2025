package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		// split by commas
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			continue
		}
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points = append(points, Point{x: x, y: y})
	}

	// calculate largest rectangle area
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := int(math.Abs(float64(points[i].x-points[j].x))) + 1
			height := int(math.Abs(float64(points[i].y-points[j].y))) + 1
			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(maxArea)
}
