package helper

import "math"

// mean of axis 0
func Mean(a [][]float64) []float64 {
	n_f := len(a[0])
	mean := make([]float64, n_f)

	x := 0
	for x < n_f {
		sum := 0.0
		l := 0.0
		for id := 0; id < len(a); id++ {
			sum += a[id][x]
			l += 1
		}

		mean[x] = sum / l
		x += 1
	}

	return mean
}


func EuclideanDistance(a []float64, b []float64) float64 {
	var sum float64
	for i := range a {
		if i < len(b) {
			sum += (a[i] - b[i]) * (a[i] - b[i])
		} else {
			sum = 0
		}
	}

	return math.Sqrt(sum)
}
