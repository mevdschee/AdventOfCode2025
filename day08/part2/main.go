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

	// Create all edges with distances
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

	// Sort edges by distance
	sort.Slice(edges, func(a, b int) bool {
		return edges[a].distance < edges[b].distance
	})

	// Use Union-Find to connect until all in one circuit
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range n {
		parent[i] = i
	}

	// Find with path compression
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	numCircuits := n
	var lastEdge Edge

	for _, edge := range edges {
		// Union by rank
		rootX := find(edge.i)
		rootY := find(edge.j)

		if rootX != rootY {
			if rank[rootX] < rank[rootY] {
				parent[rootX] = rootY
			} else if rank[rootX] > rank[rootY] {
				parent[rootY] = rootX
			} else {
				parent[rootY] = rootX
				rank[rootX]++
			}
			numCircuits--
			if numCircuits == 1 {
				lastEdge = edge
				break
			}
		}
	}

	result := points[lastEdge.i].x * points[lastEdge.j].x

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(result)
}
