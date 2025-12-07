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
			// Check if ID is made of a pattern repeated at least twice
			isInvalid := false
			for patternLen := 1; patternLen <= length/2; patternLen++ {
				if length%patternLen == 0 {
					pattern := s[:patternLen]
					valid := true
					for i := patternLen; i < length; i += patternLen {
						if s[i:i+patternLen] != pattern {
							valid = false
							break
						}
					}
					if valid {
						isInvalid = true
						break
					}
				}
			}
			if isInvalid {
				sum += id
			}
		}
	}

	fmt.Println("Sum:", sum)
}
