package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	// read input and split into parts
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// split the operators using regex on non-space
	re := regexp.MustCompile(`\S\s+`)
	operators := re.FindAllString(lines[len(lines)-1], -1)
	lines = lines[:len(lines)-1]

	// build numbers slice
	numbers := [][]int{}
	offset := 0
	for i, op := range operators {
		len := len(op)
		operators[i] = op[:1]
		ints := []int{}
		for j := 0; j < len; j++ {
			num := ""
			for _, line := range lines {
				num += string(line[offset+j])
			}
			if strings.TrimSpace(num) != "" {
				n, _ := strconv.Atoi(strings.TrimSpace(num))
				ints = append(ints, n)
			}
		}
		numbers = append(numbers, ints)
		offset += len
	}

	total := 0
	// for each operator
	for x, op := range operators {
		sum := 0
		if op == "*" {
			sum = 1
		}
		// apply operator to each row of numbers
		for _, num := range numbers[x] {
			switch op {
			case "+":
				sum += num
			case "*":
				sum *= num
			}
		}
		total += sum
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
