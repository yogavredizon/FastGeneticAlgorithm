package main

import (
	"fmt"
	"test/clustering"
	"test/helper"
	"time"
)

func main() {
	file := helper.ReadCsv("sample_data/fileName.csv")
	X := helper.SelectNumericData(file, "feat1", "feat2")
	// x := [][]float64{{60, 402}, {31, 182}, {49, 259}, {50, 289}, {51, 281}, {65, 464}, {72, 387}, {162, 946}, {113, 706}, {61, 329}, {48, 290}, {59, 311}}

	since := time.Now()

	fg := clustering.FastGenetic{
		X:                   X["data"],
		N_clusters:          3,
		PopSize:             200,
		MutationProbability: 0.3,
		MaxIters:            300,
		GenSize:             100,
	}
	c, l := fg.Fit()
	fmt.Println(c, l)
	fmt.Println(time.Since(since))
}
