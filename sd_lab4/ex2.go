package main

import "fmt"

func GeneralizedFlooding(ids []int, adj [][]int) int {
	n := len(ids)
	maxID := make([]int, n)
	copy(maxID, ids)

	for round := 0; round < n; round++ {
		nextMax := make([]int, n)
		copy(nextMax, maxID)

		for u := 0; u < n; u++ {
			for _, v := range adj[u] {
				if maxID[u] > nextMax[v] {
					nextMax[v] = maxID[u]
				}
			}
		}
		maxID = nextMax
	}

	for i := 0; i < n; i++ {
		if maxID[i] == ids[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{15, 42, 8, 99, 27}
	// Graf tare conex (lista de adiacenta)
	adj := [][]int{
		{1},
		{2, 3},
		{0},
		{4},
		{1, 2},
	}

	fmt.Println("Lider:", GeneralizedFlooding(ids, adj))
}
