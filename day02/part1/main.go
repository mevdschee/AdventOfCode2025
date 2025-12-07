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
	var rangesText string
	for scanner.Scan() {
		rangesText += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sum := 0
	ranges := strings.Split(rangesText, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		for id := start; id <= end; id++ {
			s := strconv.Itoa(id)
			length := len(s)
			// ID must have even length to be repeated twice
			if length%2 == 0 {
				half := length / 2
				// Check if first half equals second half
				if s[:half] == s[half:] {
					sum += id
				}
			}
		}
	}

	fmt.Println("Sum:", sum)
}
