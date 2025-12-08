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

	attempts := 0
	for _, edge := range edges {
		connected[edge.i][edge.j] = true
		connected[edge.j][edge.i] = true
		attempts++
		if attempts == 1000 {
			break
		}
	}

	// find all circuits using DFS
	visited := make([]bool, n)
	circuits := [][]int{}

	var dfs func(int, []int) []int
	dfs = func(node int, circuit []int) []int {
		visited[node] = true
		circuit = append(circuit, node)
		for neighbor := range connected[node] {
			if !visited[neighbor] {
				circuit = dfs(neighbor, circuit)
			}
		}
		return circuit
	}

	for i := range n {
		if !visited[i] {
			circuit := dfs(i, []int{})
			circuits = append(circuits, circuit)
		}
	}

	// get the three largest circuit sizes
	sizes := []int{}
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// multiply the three largest
	total := 1
	for i := range sizes[:3] {
		total *= sizes[i]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
