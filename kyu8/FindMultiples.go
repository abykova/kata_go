package main

import "fmt"

//Find Multiples of a Number

func main() {
	integer := 2
	limit := 6

	fmt.Println(FindMultiples(integer, limit))
}

func FindMultiples(integer, limit int) []int {
	result := make([]int, 0)

	for i := integer; i <= limit; i++ {
		a := i % integer
		if a == 0 {
			result = append(result, i)
		}

	}

	return result
}
