package metrics

import (
	"math"
	"test/helper"
)

func Score(data, centroids [][]float64, label []int) float64 {
	o1 := [][]float64{}
	for i := range data {
		if label[i] == 0 {
			o1 = append(o1, data[i])
		}
	}
	sum := 0.0
	l := 0.0
	for n := 1; n < len(o1); n++ {
		d := helper.EuclideanDistance(o1[0], o1[n])
		sum += d
		l += 1
	}
	a1 := sum / l
	min := 999999999.9
	for j := 1; j < len(centroids); j++ {
		sum := 0.0
		l := 0.0
		for i := range data {
			if label[i] == j {
				d := helper.EuclideanDistance(o1[0], data[i])
				sum += d
				l += 1.0
			}
		}
		m := sum / l
		if m < min {
			min = m
		}
	}

	score := (min - a1) / math.Max(min, a1)
	return score
}
