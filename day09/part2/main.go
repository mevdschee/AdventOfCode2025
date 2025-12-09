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
			minX, maxX := p1.x, p2.x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY, maxY := p1.y, p2.y
			if minY > maxY {
				minY, maxY = maxY, minY
			}

			// all 4 corners must be inside or on the polygon
			corners := [4]Point{
				{x: minX, y: minY},
				{x: maxX, y: minY},
				{x: minX, y: maxY},
				{x: maxX, y: maxY},
			}

			valid := true
			for _, corner := range corners {
				if !isInsideOrOn(corner, points) {
					valid = false
					break
				}
			}

			if valid {
				// sample edges more densely to ensure they stay inside
				// for large rectangles, increase sample density
				sampleDensity := max((maxX-minX)/100, 1)

				// check top and bottom edges
				for x := minX; x <= maxX && valid; x += sampleDensity {
					if !isInsideOrOn(Point{x: x, y: minY}, points) {
						valid = false
					}
					if valid && !isInsideOrOn(Point{x: x, y: maxY}, points) {
						valid = false
					}
				}
				// check final points to ensure we didn't skip
				if valid && !isInsideOrOn(Point{x: maxX, y: minY}, points) {
					valid = false
				}
				if valid && !isInsideOrOn(Point{x: maxX, y: maxY}, points) {
					valid = false
				}

				// check left and right edges
				sampleDensity = max((maxY-minY)/100, 1)
				for y := minY; y <= maxY && valid; y += sampleDensity {
					if !isInsideOrOn(Point{x: minX, y: y}, points) {
						valid = false
					}
					if valid && !isInsideOrOn(Point{x: maxX, y: y}, points) {
						valid = false
					}
				}
				// check final points
				if valid && !isInsideOrOn(Point{x: minX, y: maxY}, points) {
					valid = false
				}
				if valid && !isInsideOrOn(Point{x: maxX, y: maxY}, points) {
					valid = false
				}

				// sample interior points in a grid
				if valid && maxX-minX > 100 && maxY-minY > 100 {
					gridStep := max((maxX-minX)/10, (maxY-minY)/10)
					for sx := minX + gridStep; sx < maxX && valid; sx += gridStep {
						for sy := minY + gridStep; sy < maxY && valid; sy += gridStep {
							if !isInsideOrOn(Point{x: sx, y: sy}, points) {
								valid = false
							}
						}
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

func isInside(p Point, polygon []Point) bool {
	// ray casting algorithm - optimized version
	count := 0
	n := len(polygon)
	for i := range n {
		v1 := polygon[i]
		v2 := polygon[(i+1)%n]

		// skip horizontal edges
		if v1.y == v2.y {
			continue
		}

		if (v1.y <= p.y && p.y < v2.y) || (v2.y <= p.y && p.y < v1.y) {
			// compute x-coordinate of intersection
			x := v1.x + (p.y-v1.y)*(v2.x-v1.x)/(v2.y-v1.y)
			if p.x < x {
				count++
			}
		}
	}
	return count%2 == 1
}

// check if point is inside or on the boundary of polygon
func isInsideOrOn(p Point, polygon []Point) bool {
	// first check if on polygon boundary
	n := len(polygon)
	for i := range n {
		v1 := polygon[i]
		v2 := polygon[(i+1)%n]

		// check if point is on line segment
		if v1.x == v2.x {
			// vertical line
			if p.x == v1.x {
				minY, maxY := v1.y, v2.y
				if minY > maxY {
					minY, maxY = maxY, minY
				}
				if p.y >= minY && p.y <= maxY {
					return true
				}
			}
		} else {
			// horizontal line
			if p.y == v1.y {
				minX, maxX := v1.x, v2.x
				if minX > maxX {
					minX, maxX = maxX, minX
				}
				if p.x >= minX && p.x <= maxX {
					return true
				}
			}
		}
	}

	// then check if inside using ray casting
	return isInside(p, polygon)
}
