package main

import "fmt"

func MajorityVotingFlooding(votes []float64, A [][]float64, eps float64) (float64, int) {
	n := len(votes)
	x := make([]float64, n)
	copy(x, votes)

	for {
		nextX := make([]float64, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				nextX[i] += A[i][j] * x[j]
			}
		}
		x = nextX

		minVal, maxVal := x[0], x[0]
		for _, v := range x {
			if v < minVal {
				minVal = v
			}
			if v > maxVal {
				maxVal = v
			}
		}

		if maxVal-minVal < eps {
			break
		}
	}

	avg := x[0]
	result := 0
	if avg > 0.5 {
		result = 1
	}
	return avg, result
}

func main() {
	// Voturi initiale (0 sau 1)
	votes := []float64{1, 0, 1, 1, 0}

	// Matrice dublu stocastica (ex: graf inel)
	A := [][]float64{
		{1.0 / 3, 1.0 / 3, 0, 0, 1.0 / 3},
		{1.0 / 3, 1.0 / 3, 1.0 / 3, 0, 0},
		{0, 1.0 / 3, 1.0 / 3, 1.0 / 3, 0},
		{0, 0, 1.0 / 3, 1.0 / 3, 1.0 / 3},
		{1.0 / 3, 0, 0, 1.0 / 3, 1.0 / 3},
	}

	avg, majority := MajorityVotingFlooding(votes, A, 1e-6)
	fmt.Printf("Media: %.4f -> Vot majoritar: %d\n", avg, majority)
}
