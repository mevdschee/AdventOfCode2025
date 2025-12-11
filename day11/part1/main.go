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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var countPaths func(start string) int
	visited := map[string]bool{}

	countPaths = func(start string) int {
		if start == "out" {
			return 1
		}

		visited[start] = true
		defer delete(visited, start)

		count := 0
		for _, next := range graph[start] {
			if !visited[next] {
				count += countPaths(next)
			}
		}

		return count
	}

	fmt.Println(countPaths("you"))
}
