package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

func euclidean_distance(a []float64, b []float64) float64 {
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

func euclidean_distance_wrapper(a [][]float64, b [][]float64) [][]float64 {
	p := [][]float64{}
	for i := range a {
		c := []float64{}
		for j := range b {
			r := euclidean_distance(a[i], b[j])
			c = append(c, r)
		}
		p = append(p, c)
	}

	return p
}

func getMinIndex(a [][]float64) []int8 {
	ids := make([]int8, len(a))

	for i, v := range a {
		min := v[0]
		for j, s := range v[1:] {
			if min > s {
				ids[i] = int8(j + 1)
				min = s
			}
		}
	}

	return ids
}

func compute_centroids(a [][]float64, labels []int8, n int8) [][]float64 {
	centroids := make([][]float64, n)

	n_f := len(a[0])
	for i := int8(0); i < n; i++ {
		centroid := []float64{}
		x := 0
		for x < n_f {
			sum := 0.0
			l := 0.0
			for id := 0; id < len(a); id++ {
				if i == labels[id] {
					sum += a[id][x]
					l += 1
				}
			}

			centroid = append(centroid, sum/l)
			x += 1
		}
		centroids[i] = centroid
	}

	return centroids
}

func kmeans(X [][]float64, n_clusters int, max_iters int16) ([][]float64, []int8) {

	// initiate centroids
	var centroids = make([][]float64, n_clusters)

	// this loop used to change random number when the next number is same as previous
	var prevs = make([][]float64, n_clusters)
	n := 0

	for n < n_clusters {
		i := rand.Intn(len(X))
		elem := X[i]

		isUnique := true
		for _, v := range prevs {
			if reflect.DeepEqual(elem, v) {
				isUnique = false
				break
			}
		}

		if isUnique {
			centroids[n] = X[i]
			prevs = append(prevs, X[i])
			n++
		}
	}

	// generate labels
	indices := []int8{}
	for i := 0; i < int(max_iters); i++ {
		r := euclidean_distance_wrapper(X, centroids)
		indices = getMinIndex(r)
		centroids = compute_centroids(X, indices, int8(n_clusters))
	}

	return centroids, indices
}

func main() {
	since := time.Now()
	x := [][]float64{{60, 402}, {31, 182}, {49, 259}, {50, 289}, {51, 281}, {65, 464}, {72, 387}, {162, 946}, {113, 706}, {61, 329}, {48, 290}, {59, 311}}

	c, l := kmeans(x, 3, 20)
	fmt.Println(c, l)

	fmt.Println(time.Since(since))
}
