package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	X, Y float64
}

func dist(a, b Point) float64 {
	return math.Hypot(a.X-b.X, a.Y-b.Y)
}

func averagingRendezvous(robots []Point, radius float64, maxIter int, eps float64) []Point {
	for iter := 0; iter < maxIter && len(robots) > 1; iter++ {
		nextRobots := make([]Point, len(robots))

		// Look & Compute
		for i := 0; i < len(robots); i++ {
			sumX, sumY := robots[i].X, robots[i].Y
			count := 1.0

			for j := 0; j < len(robots); j++ {
				if i != j && dist(robots[i], robots[j]) < radius {
					sumX += robots[j].X
					sumY += robots[j].Y
					count++
				}
			}
			nextRobots[i] = Point{X: sumX / count, Y: sumY / count}
		}

		// Move & comprima pozitiile identice
		var active []Point
		for _, p := range nextRobots {
			merged := false
			for _, a := range active {
				if dist(p, a) < eps {
					merged = true
					break
				}
			}
			if !merged {
				active = append(active, p)
			}
		}
		robots = active
	}
	return robots
}

func main() {
	var robots []Point

	if f, err := os.Open("input.txt"); err == nil {
		defer f.Close()
		var n int
		fmt.Fscan(f, &n)
		for i := 0; i < n; i++ {
			var x, y float64
			fmt.Fscan(f, &x, &y)
			robots = append(robots, Point{x, y})
		}
	} else {
		robots = []Point{
			{0, 0},
			{10, 0},
			{20, 5},
			{20, 20},
			{5, 15},
			{0, 10},
		}
	}

	result := averagingRendezvous(robots, 15.0, 1000, 1e-4)
	fmt.Printf("Puncte finale (%d): %v\n", len(result), result)
}
