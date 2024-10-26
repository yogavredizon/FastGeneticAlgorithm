package metrics

import (
	"math"
	"test/helper"
)

// func SilhoutteScore(data, centroids [][]float64, label []int) float64 {
// 	o1 := [][]float64{}
// 	for i := range data {
// 		if label[i] == 0 {
// 			o1 = append(o1, data[i])
// 		}
// 	}
// 	sum := 0.0
// 	l := 0.0
// 	for n := 1; n < len(o1); n++ {
// 		d := helper.EuclideanDistance(o1[0], o1[n])
// 		sum += d
// 		l += 1
// 	}
// 	a1 := sum / l
// 	min := math.Inf(100)
// 	for j := 1; j < len(centroids); j++ {
// 		sum := 0.0
// 		l := 0.0
// 		for i := range data {
// 			if label[i] == j {
// 				d := helper.EuclideanDistance(o1[0], data[i])
// 				sum += d
// 				l += 1.0
// 			}
// 		}
// 		m := sum / l
// 		if m < min {
// 			min = m
// 		}
// 	}

//		score := (min - a1) / math.Max(min, a1)
//		return score
//	}
func SilhouetteScore(data [][]float64, centroids [][]float64, label []int) float64 {
	numPoints := float64(len(data))
	if numPoints == 0 {
		return 0.0
	}

	totalScore := 0.0

	// Create a map to hold cluster members
	clusterMembers := make(map[int][]int)
	for i, lbl := range label {
		clusterMembers[lbl] = append(clusterMembers[lbl], i)
	}

	for pointIndex := range data {
		clusterLabel := label[pointIndex]

		// Calculate a(i): average dissimilarity to all other points in the same cluster
		aSum := 0.0
		sameClusterIndices := clusterMembers[clusterLabel]

		for _, idx := range sameClusterIndices {
			if idx != pointIndex { // Avoid distance to itself
				aSum += helper.EuclideanDistance(data[pointIndex], data[idx])
			}
		}
		a := aSum / float64(len(sameClusterIndices)-1) // Exclude the point itself

		// Calculate b(i): average dissimilarity to points in the nearest different cluster
		minDistance := math.Inf(1)
		nearestClusterIndex := -1

		// Find the closest cluster centroid
		for otherClusterIndex, centroid := range centroids {
			if otherClusterIndex != clusterLabel {
				d := helper.EuclideanDistance(data[pointIndex], centroid) // Distance to the centroid
				if d < minDistance {
					minDistance = d
					nearestClusterIndex = otherClusterIndex
				}
			}
		}

		// Calculate the average dissimilarity to all points in the nearest cluster
		if nearestClusterIndex != -1 {
			bSum := 0.0
			count := 0
			for _, idx := range clusterMembers[nearestClusterIndex] {
				bSum += helper.EuclideanDistance(data[pointIndex], data[idx])
				count++
			}
			b := bSum / float64(count) // Average distance to the nearest cluster

			// Calculate silhouette score for the current point
			score := (b - a) / math.Max(b, a)
			totalScore += score
		}
	}

	// Return the average silhouette score for all points
	return totalScore / numPoints
}
