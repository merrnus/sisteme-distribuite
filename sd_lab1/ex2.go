package main

import "fmt"

func main() {
	arr := [5]int{5, 2, 9, 1, 7}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr [j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	
	fmt.Println(arr)
}
