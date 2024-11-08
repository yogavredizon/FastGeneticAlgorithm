package clustering_test

import (
	"fmt"
	"test/clustering"
	"testing"

	"github.com/stretchr/testify/assert"
)

var X = [][]float64{{60, 402}, {31, 182}, {49, 259}, {50, 289}, {51, 281}, {65, 464}, {72, 387}, {162, 946}, {113, 706}, {61, 329}, {48, 290}, {59, 311}}

func TestGeneratePopulationCorrect(t *testing.T) {
	fg := clustering.FastGenetic{
		X:          X,
		PopSize:    10,
		N_clusters: 3,
	}

	result, err := fg.GeneratePop()

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
func TestGeneratePopulationIncorrect(t *testing.T) {
	fg := clustering.FastGenetic{
		X:          X,
		PopSize:    10,
		N_clusters: 1,
	}

	result, err := fg.GeneratePop()

	assert.Nil(t, result)
	assert.NotNil(t, err)
}
func TestLegalSolution(t *testing.T) {
	fg := clustering.FastGenetic{
		X:          X,
		PopSize:    10,
		N_clusters: 3,
	}

	result, _ := fg.GeneratePop()

	for i, v := range result {
		r := fg.CheckLegal(v)
		fmt.Println("Index ", i, "is", r)
	}

	s := fg.CheckLegal([]int{1, 1, 1, 1, 2, 1, 1, 1, 1})
	fmt.Println(s)
}

func TestComputeCentroids(t *testing.T) {
	fg := clustering.FastGenetic{
		X:          X,
		PopSize:    10,
		N_clusters: 3,
	}

	result, _ := fg.GeneratePop()
	fmt.Println(result)
	for _, v := range result {
		x, _ := fg.ComputeCentroids(v)
		fmt.Println(x)
	}
}
func TestComputeSE(t *testing.T) {
	fg := clustering.FastGenetic{
		X:          X,
		PopSize:    10,
		N_clusters: 3,
	}

	result, _ := fg.GeneratePop()
	for _, v := range result {
		fmt.Println(v)
		x, _ := fg.ComputeCentroids(v)
		fmt.Println(x)
		se := fg.ComputeSquareError(v, x)
		fmt.Println(se)
	}
}

/*
	func TestComputeFitness(t *testing.T) {
		fg := clustering.FastGenetic{
			X:          X,
			PopSize:    10,
			N_clusters: 3,
		}

		result, _ := fg.GeneratePop()
		fmt.Println(fg.ComputeFitness(result))

}

	func TestSelection(t *testing.T) {
		fg := clustering.FastGenetic{
			X:          X,
			PopSize:    10,
			N_clusters: 3,
		}

		result, _ := fg.GeneratePop()
	    f := fg.ComputeFitness(result)
	    fg.Selection(f)
	}

	func TestMutate(t *testing.T) {
		fg := clustering.FastGenetic{
			X:          X,
			PopSize:    10,
			N_clusters: 3,
	        Thresshold: 0.5,
		}

	    for i := 0; i < 1000; i++{
	        result, _ := fg.GeneratePop()
	        f := fg.ComputeFitness(result)
	        i := fg.Selection(f)
	        m := fg.Mutation(result[i])
	        fmt.Println(m)

	    }
	}
*/
func TestKmeans(t *testing.T) {
	fg := clustering.FastGenetic{
		X:                   X,
		PopSize:             10,
		N_clusters:          3,
		MutationProbability: 0.5,
		MaxIters:            100,
	}

	for i := 0; i < 1000; i++ {

		result, _ := fg.GeneratePop()
		f := []float64{}
		f = fg.ComputeFitness(result)
		j := fg.Selection(f)
		m := fg.Mutation(result[j])

		c, ci := fg.KMeans(m)
		fmt.Println(c, ci)
	}
}

func TestFit(t *testing.T) {
	fg := clustering.FastGenetic{
		X:                   X,
		PopSize:             100,
		N_clusters:          3,
		GenSize:             50,
		MutationProbability: 0.5,
		MaxIters:            100,
	}
	fmt.Println(len(fg.X))
	for i := 0; i < 1000; i++ {
		c, l := fg.Fit()
		fmt.Println(c, l)
	}
}

func TestFitWithOutliers(t *testing.T) {
	fg := clustering.FastGenetic{
		X:                   X,
		PopSize:             100,
		N_clusters:          3,
		GenSize:             100,
		MutationProbability: 0.5,
		MaxIters:            300,
	}
	c, l := fg.FitWithOutliers()
	fmt.Println(c, len(l), len(fg.X))
	fmt.Println(fg.X)
}
