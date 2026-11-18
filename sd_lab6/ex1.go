package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}

		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info reprezinta o pereche (id, valoare) retinuta in memoria unui nod.
type Info struct {
	ID    int
	Value int
}

// Memory este o multime de token-uri (reuniunea se face prin inserare).
type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	Crashed   bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	s, _ := strconv.Atoi(firstLine[1])

	scanner.Scan()
	tokenFields := strings.Fields(scanner.Text())
	nodes := make([]Node, n)
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(tokenFields[i])
		mem := make(Memory)
		mem[Info{ID: i, Value: val}] = struct{}{}
		nodes[i] = Node{
			ID: i,
			M:  mem,
		}
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		for _, f := range fields {
			nb, _ := strconv.Atoi(f)
			nodes[i].Neighbors = append(nodes[i].Neighbors, nb)
			adj[i] = append(adj[i], nb)
		}
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex!")
		return
	}

	crashedCount := 0
	t := 0
	maxIter := (s + 1) * diam

	for t <= maxIter {
		// Simulare crash random (maxim s defecte)
		if crashedCount < s && rand.Float64() < 0.3 {
			victim := rand.Intn(n)
			if !nodes[victim].Crashed {
				nodes[victim].Crashed = true
				crashedCount++
			}
		}

		nextM := make([]Memory, n)
		for i := 0; i < n; i++ {
			nextM[i] = make(Memory)
			for k, v := range nodes[i].M {
				nextM[i][k] = v
			}
		}

		// Fiecare nod corect trimite memoria sa vecinilor
		for i := 0; i < n; i++ {
			if nodes[i].Crashed {
				continue
			}
			for _, nb := range nodes[i].Neighbors {
				if !nodes[nb].Crashed {
					for info := range nodes[i].M {
						nextM[nb][info] = struct{}{}
					}
				}
			}
		}

		for i := 0; i < n; i++ {
			if !nodes[i].Crashed {
				nodes[i].M = nextM[i]
			}
		}

		t++
	}

	// Gasim un nod corect pentru afisare
	var correctNode Node
	for _, node := range nodes {
		if !node.Crashed {
			correctNode = node
			break
		}
	}

	var result []Info
	for info := range correctNode.M {
		result = append(result, info)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	fmt.Printf("Diametru: %d\n", diam)
	fmt.Printf("Iteratia opririi: %d\n", t)
	for _, info := range result {
		fmt.Printf("(%d, %d) ", info.ID, info.Value)
	}
	fmt.Println()
}
