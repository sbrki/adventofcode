package main

import "fmt"

func intToArray(x int) (result [6]int) {
	for i := 0; i < 6; i++ {
		r := x % 10
		x = x / 10
		result[5-i] = r
	}
	return result
}

func isValid(x int) bool {
	v := intToArray(x)
	// check if non-descending
	for i := 0; i < 5; i++ {
		if v[i] > v[i+1] {
			return false
		}
	}

	// check if containts a double digit
	var containsDoubleDigit bool = false
	for i := 0; i < 5; i++ {
		if v[i] == v[i+1] {
			containsDoubleDigit = true
			break
		}
	}
	return containsDoubleDigit
}

// used in part 2
func digitsWithMultipleOccurances(x int) []int {
	result := make([]int, 0)
	v := intToArray(x)
	for i := 0; i < 5; i++ {
		if v[i] == v[i+1] {
			// add to result if not in result
			contains := false
			for _, val := range result {
				if val == v[i] {
					contains = true
					break
				}
			}
			if !contains {
				result = append(result, v[i])
			}
		}
	}
	return result
}

// used in part 2
func specialElfRequirement(x int) bool {

	multis := digitsWithMultipleOccurances(x)
	v := intToArray(x)
	// count the digits in x
	var count [10]int
	for _, v := range v {
		count[v]++
	}

	satisfied := false
	for _, c := range multis {
		if count[c] == 2 {
			satisfied = true
		}
	}
	return satisfied

}

func main() {
	// my puzzle input: 136818-685979
	count := 0
	for i := 136818; i <= 685979; i++ {
		if isValid(i) && specialElfRequirement(i) {
			count++
		}
	}
	fmt.Println(count)
}

/*
   note: I've structured code in an unoptimal manner (intToArray gets called
   more than necesarry) in order to make it easier to swtich between
   part 1 and part 2.
   Also, search space can be reduced given the password non-descending requirements,
   but it wouldn't speed things up a lot, especially given it takes
   a couple of ms to execute as-is.
*/
