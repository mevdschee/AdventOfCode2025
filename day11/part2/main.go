package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		device := parts[0]
		outputs := strings.Fields(parts[1])
		graph[device] = outputs
	}

	// reverse graph
	reverseGraph := map[string][]string{}
	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			reverseGraph[neighbor] = append(reverseGraph[neighbor], node)
		}
	}

	memo := map[string]int{}
	visited := map[string]bool{}
	visited["out"] = true

	result := countPaths(reverseGraph, "out", false, false, visited, memo)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func countPaths(graph map[string][]string, node string, visitedDac, visitedFft bool, visited map[string]bool, memo map[string]int) int {
	if node == "dac" {
		visitedDac = true
	}
	if node == "fft" {
		visitedFft = true
	}

	if node == "svr" {
		if visitedDac && visitedFft {
			return 1
		}
		return 0
	}

	key := fmt.Sprintf("%s:%t:%t", node, visitedDac, visitedFft)
	if count, found := memo[key]; found {
		return count
	}

	totalPaths := 0
	for _, next := range graph[node] {
		if !visited[next] {
			visited[next] = true
			totalPaths += countPaths(graph, next, visitedDac, visitedFft, visited, memo)
			visited[next] = false
		}
	}
	memo[key] = totalPaths

	return totalPaths
}
