package main

import "fmt"

func sum(nums ...int) { // seems like ... mean multiple int
	fmt.Println(nums, " ")

	var result int = 0
	for _, num := range nums {
		result += num
	}

	fmt.Println(result)
}

func main() {
	sum(1, 3)

	var nums []int = []int{1, 2, 3, 4}
	sum(nums...) // seems like it means 1, 2, 3, 4, 5
}
