// ex9: Afiseaza in O(n) indicii elementelor care insumate dau s (citit de la tastatura).

package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 8, 4, 9}

	var s int
	fmt.Scan(&s)

	seen := make(map[int]int) // değer -> index

	for j, v := range arr {
		if i, ok := seen[s-v]; ok {
			fmt.Println(i, j)
		}
		seen[v] = j
	}
}
