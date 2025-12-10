package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	lights  int
	buttons []int
	joltage []int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := []Machine{}
	for scanner.Scan() {
		line := scanner.Text()

		// parse indicator lights
		start := strings.Index(line, "[")
		end := strings.Index(line, "]")
		indicators := line[start+1 : end]
		lights := 0
		for i := 0; i < len(indicators); i++ {
			if indicators[i] == '#' {
				lights |= (1 << i)
			}
		}

		// parse buttons
		buttons := []int{}
		rest := line[end+1:]
		inParen := false
		currentButton := 0
		for i := 0; i < len(rest); i++ {
			if rest[i] == '(' {
				inParen = true
				currentButton = 0
			} else if rest[i] == ')' {
				inParen = false
				buttons = append(buttons, currentButton)
			} else if rest[i] == '{' {
				break
			} else if inParen {
				if rest[i] >= '0' && rest[i] <= '9' {
					digit := int(rest[i] - '0')
					currentButton |= (1 << digit)
				}
			}
		}

		// parse joltage
		joltage := []int{}
		joltageStart := strings.Index(line, "{")
		joltageEnd := strings.Index(line, "}")
		if joltageStart != -1 && joltageEnd != -1 {
			joltageStr := line[joltageStart+1 : joltageEnd]
			parts := strings.Split(joltageStr, ",")
			for _, part := range parts {
				val, _ := strconv.Atoi(part)
				joltage = append(joltage, val)
			}
		}

		machines = append(machines, Machine{lights: lights, buttons: buttons, joltage: joltage})
	}

	// solve each machine
	totalPresses := 0
	for _, machine := range machines {
		if machine.lights == 0 {
			continue
		}

		visited := make(map[int]bool)
		queue := []struct {
			state   int
			presses int
		}{{0, 0}}
		visited[0] = true

		found := false
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			// try pressing each button
			for _, button := range machine.buttons {
				newState := current.state ^ button

				if newState == machine.lights {
					totalPresses += current.presses + 1
					found = true
					break
				}

				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, struct {
						state   int
						presses int
					}{newState, current.presses + 1})
				}
			}

			if found {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalPresses)
}
