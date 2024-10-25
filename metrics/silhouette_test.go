package metrics_test

import (
	"test/metrics"
	"testing"
)

var X [][]float64 = [][]float64{{60, 402}, {49, 259}, {50, 289}, {51, 281}, {65, 464}, {72, 387}, {162, 946}, {113, 706}, {61, 329}, {48, 290}, {59, 311}}
var labels []int = []int{0, 2, 2, 2, 0, 0, 1, 1, 2, 2, 2}
var centroids [][]float64 = [][]float64{{65.66666667, 417.66666667}, {137.5, 826.0}, {53.0, 293.16666667}}

func TestSilhouette(t *testing.T) {
	metrics.Score(X, centroids, labels)
}
