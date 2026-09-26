package main

import "fmt"

func main() {
	numbers := []int{2, 4}
	alias := numbers

	alias[0] = 10
	alias = append(alias, 8)

	fmt.Println(numbers)
	fmt.Println(alias)
}
