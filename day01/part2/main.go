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
	sum := 50
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		// replace R with + and L with -
		line = strings.ReplaceAll(line, "R", "+")
		line = strings.ReplaceAll(line, "L", "-")
		// convert to int
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if num < 0 {
			for range -num {
				sum--
				sum = (sum + 100) % 100
				if sum == 0 {
					count++
				}
			}
		} else {
			for range num {
				sum++
				sum = (sum + 100) % 100
				if sum == 0 {
					count++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Count:", count)
}
