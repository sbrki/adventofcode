package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(lines []string) int {
	var j int = 0
	var treeCount int = 0
	for i := 0; i < len(lines); i++ {
		if lines[i][j%len(lines[i])] == '#' {
			treeCount++
		}
		j += 3
	}
	return treeCount
}

func generalisedSolver(lines []string, hStep int, vStep int) int {
	i := 0
	j := 0
	treeCount := 0
	for {

		if lines[i][j%len(lines[i])] == '#' {
			treeCount++
		}

		i += hStep
		j += vStep

		// break cond
		if i >= len(lines) {
			return treeCount
		}
	}
}

func solve2(lines []string) int {
	return generalisedSolver(lines, 1, 1) * generalisedSolver(lines, 1, 3) * generalisedSolver(lines, 1, 5) * generalisedSolver(lines, 1, 7) * generalisedSolver(lines, 2, 1)
}

func main() {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(solve(lines))
	fmt.Println("part 2")
	fmt.Println(solve2(lines))
}
