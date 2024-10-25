package preprocess

import (
	"math"
	"test/helper"
)

// detect outliers using clustering result
func DetectOutliers(X, centroids [][]float64, labels []int) []int {
	// get labels and centroids from FGKA Calculations
	result := make([]int, len(labels))
	for i, cen := range centroids {
		sumDistances := 0.0
		length := 0.0
		distances := []float64{}
		idx := []int{}

		for c := 0; c < len(labels); c++ {
			if labels[c] == i {
				d := helper.EuclideanDistance(X[c], cen)
				sumDistances += d
				distances = append(distances, d)
				idx = append(idx, c)
				length++
			}
		}

		mean := sumDistances / length
		sum := 0.0
		for _, d := range distances {
			sum += (d - mean) * (d - mean)
		}
		std := math.Sqrt(sum / length)

		upperBound := mean + 2*std
		lowerBound := mean - 2*std

		for s, d := range distances {
			if d > upperBound || d < lowerBound {
				result[idx[s]] = 1
			}
		}

	}

	return result
	// example we have 3 cluster : [1, 0, 2 ]
	// calculate standart deviation of every cluster
	// calculate mean

	// return [0, 1, 0 , 1, 0, 1]
	// 0 indicate data point is not an outlires, and otherwise

}
