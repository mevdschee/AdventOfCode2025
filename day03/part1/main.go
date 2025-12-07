package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		maxChar1 := line[0]
		runePos := 0
		for i := 0; i < len(line)-1; i++ {
			if line[i] > maxChar1 {
				maxChar1 = line[i]
				runePos = i
			}
		}
		maxChar2 := line[runePos+1]
		for i := runePos + 1; i < len(line); i++ {
			if line[i] > maxChar2 {
				maxChar2 = line[i]
			}
		}
		firstDigit := int(maxChar1 - '0')
		secondDigit := int(maxChar2 - '0')
		joltage := firstDigit*10 + secondDigit
		totalJoltage += joltage
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalJoltage)
}
