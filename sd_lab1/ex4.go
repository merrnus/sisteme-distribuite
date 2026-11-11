package main 

import "fmt" 

func main() {
	nums := []float32{1.5, 2.5, 3.5, 4.5, 5.5}
	
 	var idx int
	fmt.Scan(&idx)

	if idx < 0 || idx >= len(nums) {
		fmt.Println("Index Invalid")
		return
	}

	nums = append(nums[:idx], nums[idx+1:]...)

	fmt.Println(nums)
}
