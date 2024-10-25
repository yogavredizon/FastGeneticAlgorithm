package preprocess_test

import (
	"fmt"
	"test/preprocess"
	"testing"
)

var X [][]float64 = [][]float64{{60, 402}, {31, 182}, {49, 259}, {50, 289}, {51, 281}, {65, 464}, {72, 387}, {162, 946}, {113, 706}, {61, 329}, {48, 290}, {59, 311}}
var labels []int = []int{0, 2, 2, 2, 2, 0, 0, 1, 1, 2, 2, 2}
var centroids [][]float64 = [][]float64{{65.66666666666667, 417.6666666666667}, {137.5, 826}, {49.857142857142854, 277.2857142857143}}

func TestDetectOutliers(t *testing.T) {
	fmt.Println(preprocess.DetectOutliers(X, centroids, labels))
	fmt.Println(X)
}
