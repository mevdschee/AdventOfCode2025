package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Machine struct {
	lights  int
	buttons []int
	joltage []int
}

func solveJoltage(buttons []int, target []int) int {
	m := len(buttons)
	n := len(target)

	// build Z3 SMT-LIB script
	var smtLib strings.Builder
	smtLib.WriteString("(set-logic QF_LIA)\n")

	// declare variables for button presses
	for i := 0; i < m; i++ {
		smtLib.WriteString(fmt.Sprintf("(declare-const b%d Int)\n", i))
		smtLib.WriteString(fmt.Sprintf("(assert (>= b%d 0))\n", i))
	}

	// add constraints for each counter
	for i := 0; i < n; i++ {
		var terms []string
		for j := 0; j < m; j++ {
			if buttons[j]&(1<<i) != 0 {
				terms = append(terms, fmt.Sprintf("b%d", j))
			}
		}
		if len(terms) > 0 {
			sum := strings.Join(terms, " ")
			if len(terms) > 1 {
				sum = "(+ " + sum + ")"
			}
			smtLib.WriteString(fmt.Sprintf("(assert (= %s %d))\n", sum, target[i]))
		}
	}

	// minimize the sum of all button presses
	var sumTerms []string
	for i := 0; i < m; i++ {
		sumTerms = append(sumTerms, fmt.Sprintf("b%d", i))
	}
	totalSum := strings.Join(sumTerms, " ")
	if len(sumTerms) > 1 {
		totalSum = "(+ " + totalSum + ")"
	}

	smtLib.WriteString("(declare-const total Int)\n")
	smtLib.WriteString(fmt.Sprintf("(assert (= total %s))\n", totalSum))
	smtLib.WriteString("(minimize total)\n")
	smtLib.WriteString("(check-sat)\n")
	smtLib.WriteString("(get-value (total))\n")

	// write to temp file and run z3
	tmpFile, err := os.CreateTemp("", "z3-*.smt2")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(smtLib.String()); err != nil {
		panic(err)
	}
	tmpFile.Close()

	// run z3
	cmd := exec.Command("z3", tmpFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("z3 error: %v\noutput: %s", err, output))
	}

	// parse output
	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 || lines[0] != "sat" {
		panic(fmt.Sprintf("z3 failed to find solution: %s", output))
	}

	// extract total from "((total 123))"
	valueLine := lines[1]
	valueLine = strings.TrimSpace(valueLine)
	valueLine = strings.TrimPrefix(valueLine, "((total ")
	valueLine = strings.TrimSuffix(valueLine, "))")
	total, err := strconv.Atoi(valueLine)
	if err != nil {
		panic(fmt.Sprintf("failed to parse z3 output: %s", valueLine))
	}

	return total
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
		if len(machine.joltage) == 0 {
			continue
		}

		// solve using Z3 (as suggested on Reddit)
		minPresses := solveJoltage(machine.buttons, machine.joltage)
		totalPresses += minPresses
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalPresses)
}
