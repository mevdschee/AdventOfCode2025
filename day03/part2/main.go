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

		// Find maximum n-digit number by greedily selecting digits
		n := 12 // number of characters to use
		var selectedDigits []byte
		remainingLine := line

		for len(selectedDigits) < n && len(remainingLine) > 0 {
			// Find the maximum digit that still leaves enough digits for the rest
			maxChar := byte('0')
			maxPos := 0
			neededDigits := n - len(selectedDigits)
			searchLimit := len(remainingLine) - neededDigits + 1

			for i := 0; i < searchLimit; i++ {
				if remainingLine[i] > maxChar {
					maxChar = remainingLine[i]
					maxPos = i
				}
			}

			selectedDigits = append(selectedDigits, maxChar)
			remainingLine = remainingLine[maxPos+1:]
		}

		// Convert selected digits to a number
		joltage := 0
		for _, digit := range selectedDigits {
			joltage = joltage*10 + int(digit-'0')
		}

		totalJoltage += joltage
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalJoltage)
}
