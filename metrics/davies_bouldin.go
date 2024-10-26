package metrics

import (
	"math"
	"test/helper"
)

// func DaviesBouldinIndex(data, centroids [][]float64, labels []int) float64 {
// 	// data = [][]float64{{1, 2, 1}, {2, 1, 2}, {1, 1, 1}, {4, 4, 5}, {5, 5, 4}, {4, 5, 4}, {8, 8, 8}, {9, 8, 9}, {8, 9, 8}}
// 	// labels = []int{0, 0, 0, 1, 1, 1, 2, 2, 2}
// 	// centroids = [][]float64{{1.33, 1.33, 1.33}, {4.33, 4.67, 4.33}, {8.33, 8.33, 8.33}}
// 	si := []float64{}
// 	for c := range centroids {
// 		distances := 0.0
// 		l := 0.0
// 		for i := range data {
// 			if labels[i] == c {
// 				d := helper.EuclideanDistance(data[i], centroids[c])
// 				distances += d
// 				l += 1.0
// 			}
// 		}
// 		distances = distances / l
// 		si = append(si, distances)
// 	}

// 	distances := []float64{}
// 	i := 1

// 	for c := 0; c < len(centroids); c++ {
// 		d := helper.EuclideanDistance(centroids[c], centroids[i])
// 		distances = append(distances, d)

// 		if i == len(centroids)-1 {
// 			i = 0
// 		} else {
// 			i++
// 		}
// 	}

// 	r := []float64{}
// 	j := 1
// 	for i := 0; i < len(distances); i++ {
// 		s := (si[i] + si[j]) / distances[i]
// 		if j == len(distances)-1 {
// 			j = 0
// 		} else {
// 			j++
// 		}
// 		r = append(r, s)
// 	}

// 	sr := 0.0

// 	for _, i := range r {
// 		sr += i
// 	}

// 	dbi := sr / float64(len(r))
// 	return dbi
// }

func DaviesBouldinIndex(data, centroids [][]float64, labels []int) float64 {
	si := []float64{}
	for c := range centroids {
		distances := 0.0
		l := 0.0
		for i := range data {
			if labels[i] == c {
				d := helper.EuclideanDistance(data[i], centroids[c])
				distances += d
				l += 1.0
			}
		}
		distances = distances / l
		si = append(si, distances)
	}

	distances := make([][]float64, len(centroids))
	for i := 0; i < len(centroids); i++ {
		distances[i] = make([]float64, len(centroids))
		for j := 0; j < len(centroids); j++ {
			if i != j {
				distances[i][j] = helper.EuclideanDistance(centroids[i], centroids[j])
			}
		}
	}
	dbi := 0.0
	for i := 0; i < len(centroids); i++ {
		max_similarity := 0.0
		for j := 0; j < len(centroids); j++ {
			if i != j {
				similarity := (si[i] + si[j]) / distances[i][j]
				max_similarity = math.Max(max_similarity, similarity)
			}
		}
		dbi += max_similarity
	}
	dbi /= float64(len(centroids))

	return dbi
}
