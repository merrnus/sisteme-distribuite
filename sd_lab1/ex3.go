package main

import "fmt"

func main() {
	nums := []int{3, 8, 5, 12, 7, 6}
	var evens []int

	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}

	fmt.Println(evens)
}
