package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := map[string]int{
		"Andrei": 190, 
		"Maria": 165, 
		"Ion": 182, 
		"Elena": 172, 
		"George": 178,
		"Ana": 165,
 	}

	names := []string{}
	for name := range heights {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		return heights[names[i]] > heights[names[j]]
	})

	for _, name := range names {
		fmt.Println(name, heights[name])
	}
}
