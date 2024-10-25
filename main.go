package main

import (
	"fmt"
	"test/clustering"
	"test/metrics"
	"time"
)

func main() {
	// file := helper.ReadCsv("sample_data/iris.csv")
	// X := helper.SelectNumericData(file, "SepalLengthCm", "SepalWidthCm", "PetalLengthCm", "PetalWidthCm")
	// file := helper.ReadCsv("sample_data/nutrition.csv")
	// X := helper.SelectNumericData(file, "calories", "fat", "proteins", "carbohydrate")
	x := [][]float64{{60, 402}, {31, 182}, {49, 259}, {50, 289}, {51, 281}, {65, 464}, {72, 387}, {162, 946}, {113, 706}, {61, 329}, {48, 290}, {59, 311}}
	since := time.Now()

	fg := clustering.FastGenetic{
		X:                   x,
		N_clusters:          3,
		PopSize:             50,
		MutationProbability: 0.8,
		MaxIters:            300,
		GenSize:             100,
	}

	c, l := fg.FitWithOutliers()
	fmt.Println("Silhouette Score :", metrics.Score(fg.X, c, l), len(fg.X))
	fmt.Println(c, l)

	fmt.Println(time.Since(since))
}
