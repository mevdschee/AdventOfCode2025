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
		points = append(points, Point{x, y})
	}

	// find largest rectangle with red corners using only green/red tiles
	maxArea := 0
	for i := 0; i < len(points); i++ {
		p1 := points[i]

		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			// check if they can form opposite corners (different x and y)
			if p1.x == p2.x || p1.y == p2.y {
				continue
			}

			// calculate potential area first (cheap operation)
			width := int(math.Abs(float64(p1.x-p2.x))) + 1
			height := int(math.Abs(float64(p1.y-p2.y))) + 1
			area := width * height

			// skip if this can't beat current max
			if area <= maxArea {
				continue
			}

			// check if rectangle is valid
			minX, maxX := min(p1.x, p2.x), max(p1.x, p2.x)
			minY, maxY := min(p1.y, p2.y), max(p1.y, p2.y)

			// all 4 corners must be inside or on the polygon
			corners := [4]Point{{minX, minY}, {maxX, minY}, {minX, maxY}, {maxX, maxY}}

			valid := true
			for _, corner := range corners {
				if !isInsideOrOn(corner, points) {
					valid = false
					break
				}
			}

			if valid {
				// sample edges to ensure they stay inside
				sampleDensity := max((maxX-minX)/100, 1)

				// check top and bottom edges
				for x := minX; x <= maxX && valid; x += sampleDensity {
					if !isInsideOrOn(Point{x, minY}, points) {
						valid = false
					}
					if valid && !isInsideOrOn(Point{x, maxY}, points) {
						valid = false
					}
				}

				// check left and right edges
				sampleDensity = max((maxY-minY)/100, 1)
				for y := minY; y <= maxY && valid; y += sampleDensity {
					if !isInsideOrOn(Point{minX, y}, points) {
						valid = false
					}
					if valid && !isInsideOrOn(Point{maxX, y}, points) {
						valid = false
					}
				}
			}

			if valid {
				maxArea = area
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(maxArea)
}

// utility function
func isInsideOrOn(p Point, polygon []Point) bool {
	return isOn(p, polygon) || isInside(p, polygon)
}

// check if point is inside the polygon
func isInside(p Point, polygon []Point) bool {
	// count ray intersections
	count := 0
	n := len(polygon)
	for i := range n {
		v1, v2 := polygon[i], polygon[(i+1)%len(polygon)]

		// check if the edge is crossing the ray
		if (v1.y <= p.y && p.y < v2.y) || (v2.y <= p.y && p.y < v1.y) {
			// compute x-coordinate of intersection
			x := v1.x + (p.y-v1.y)*(v2.x-v1.x)/(v2.y-v1.y)
			// check if the intersection is to the right of the point
			if p.x < x {
				count++
			}
		}
	}
	// odd count means inside
	return count%2 == 1
}

// check if point is on any edge of the polygon
func isOn(p Point, polygon []Point) bool {
	for i := range len(polygon) {
		v1, v2 := polygon[i], polygon[(i+1)%len(polygon)]
		if v1.x == v2.x && p.x == v1.x && p.y >= min(v1.y, v2.y) && p.y <= max(v1.y, v2.y) {
			return true
		}
		if v1.y == v2.y && p.y == v1.y && p.x >= min(v1.x, v2.x) && p.x <= max(v1.x, v2.x) {
			return true
		}
	}
	return false
}
