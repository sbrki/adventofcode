package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run(memory []int, noun int, verb int) int {
	instructions := make([]int, len(memory))
	copy(instructions, memory)

	instructions[1] = noun
	instructions[2] = verb

	var position uint
	for {
		opcode := instructions[position]
		if opcode == 99 {
			break
		}
		switch opcode {
		case 1:
			instructions[instructions[position+3]] = instructions[instructions[position+1]] +
				instructions[instructions[position+2]]
			position += 4
		case 2:

			instructions[instructions[position+3]] = instructions[instructions[position+1]] *
				instructions[instructions[position+2]]
			position += 4
		}
	}
	return instructions[0]
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	strInst := strings.Split(input, ",")

	instructions := make([]int, len(strInst))
	for idx, val := range strInst {
		instructions[idx], _ = strconv.Atoi(val)
	}

	// fmt.Println(run(instructions, 12, 2))  <-- part one

	var goal int
	fmt.Scanf("%d\n", &goal)
	for verb := 0; verb < 100; verb++ {
		for noun := 0; noun < 100; noun++ {
			result := run(instructions, noun, verb)
			if result == goal {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}
