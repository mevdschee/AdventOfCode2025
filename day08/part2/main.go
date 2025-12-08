package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

type Edge struct {
	i        int
	j        int
	distance float64
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
		if len(coords) != 3 {
			continue
		}
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		points = append(points, Point{x: x, y: y, z: z})
	}

	n := len(points)

	// create all edges with distances
	edges := []Edge{}
	for i := range n {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]
			dx := float64(p1.x - p2.x)
			dy := float64(p1.y - p2.y)
			dz := float64(p1.z - p2.z)
			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)
			edges = append(edges, Edge{i: i, j: j, distance: dist})
		}
	}

	// sort edges by distance
	sort.Slice(edges, func(a, b int) bool {
		return edges[a].distance < edges[b].distance
	})

	// connect the 1000 closest pairs
	connected := make([]map[int]bool, n)
	for i := range n {
		connected[i] = make(map[int]bool)
	}

	var lastEdge Edge
	for _, edge := range edges {
		connected[edge.i][edge.j] = true
		connected[edge.j][edge.i] = true

		// Count circuits using DFS
		visited := make([]bool, n)
		circuits := 0

		var dfs func(int)
		dfs = func(node int) {
			visited[node] = true
			for neighbor := range connected[node] {
				if !visited[neighbor] {
					dfs(neighbor)
				}
			}
		}

		for i := range n {
			if !visited[i] {
				dfs(i)
				circuits++
			}
		}

		if circuits == 1 {
			lastEdge = edge
			break
		}
	}

	result := points[lastEdge.i].x * points[lastEdge.j].x

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(result)
}
