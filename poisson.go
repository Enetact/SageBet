package models

import "math"

func PoissonProbability(k int, lambda float64) float64 {
	return (math.Pow(lambda, float64(k)) * math.Exp(-lambda)) / float64(factorial(k))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
