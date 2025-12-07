package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// make slice of int slices
	numbers := [][]int{}
	operators := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// split line on one or more spaces
		nums := strings.Fields(line)
		ints := []int{}
		for _, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				operators = append(operators, n)
			} else {
				ints = append(ints, i)
			}
		}
		if len(ints) > 0 {
			numbers = append(numbers, ints)
		}
	}

	total := 0
	// for each operator
	for x, op := range operators {
		sum := 0
		if op == "*" {
			sum = 1
		}
		// apply operator to each row of numbers
		for _, row := range numbers {
			switch op {
			case "+":
				sum += row[x]
			case "*":
				sum *= row[x]
			}
		}
		total += sum
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
