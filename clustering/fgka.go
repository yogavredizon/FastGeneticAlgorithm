package clustering

import (
	"errors"
	"math/rand"
	"reflect"
	"test/helper"
)

type FastGenetic struct {
	PopSize             int
	X                   [][]float64
	N_clusters          int
	MutationProbability float64
	GenSize             int
	MaxIters            int
}

func (f *FastGenetic) GeneratePop() ([][]int, error) {
	n_data := len(f.X)

	if f.PopSize < 2 {
		return nil, errors.New("pop size must greater than 1")
	}

	if f.N_clusters == n_data || n_data-f.N_clusters < 2 || f.N_clusters < 2 {
		return nil, errors.New("not valid n clusters")
	}

	// assign each data points to a cluster by select randomly
	population := make([][]int, f.PopSize)

	for i := 0; i < len(population); i++ {

		// each allele in each genes is representasion of cluster number.
		solution := make([]int, n_data)
		for g := 0; g < len(solution); g++ {
			id := rand.Intn(f.N_clusters)
			solution[g] = id
		}

		population[i] = solution
	}

	return population, nil
}

func (f *FastGenetic) Compute_centroids(solution []int) ([][]float64, error) {
	pop_centroids := [][]float64{}

	legal := f.CheckLegal(solution)
	if !legal {
		solution = helper.LegalString(solution, f.N_clusters)
	}

	for k := 0; k < f.N_clusters; k++ {
		Xn := [][]float64{}
		for i := 0; i < len(solution); i++ {
			if solution[i] == k {
				Xn = append(Xn, f.X[i])
			}
		}

		mean := helper.Mean(Xn)
		pop_centroids = append(pop_centroids, mean)
	}

	return pop_centroids, nil
}

func (f *FastGenetic) CheckLegal(solution []int) bool {
	n := map[int]int{}

	for _, v := range solution {
		n[v] = v
	}

	return len(n) == f.N_clusters
}

func (f *FastGenetic) ComputeSquareError(solution []int, centroids [][]float64) float64 {
	twcv := 0.0

	for c := 0; c < len(centroids); c++ {
		sum := 0.0
		for d := 0; d < len(centroids[c]); d++ {
			for i := 0; i < len(solution); i++ {
				if solution[i] == c {
					sum += (f.X[i][d] - centroids[c][d]) * (f.X[i][d] - centroids[c][d])
				}
			}
		}
		twcv += sum
	}

	return twcv
}

func (f *FastGenetic) ComputeFitness(population [][]int) []float64 {
	se := []float64{}
	for i := 0; i < len(population); i++ {
		centroids, _ := f.Compute_centroids(population[i])
		sumSquare := f.ComputeSquareError(population[i], centroids)
		se = append(se, sumSquare)
	}

	max := se[0]

	for _, s := range se[1:] {
		if s > max {
			max = s
		}
	}

	fitness := []float64{}
	for i, s := range se {
		legal := f.CheckLegal(population[i])

		if legal {
			f := 1.5*max - s
			fitness = append(fitness, f)
		} else {
			f := 0.1 * 1
			fitness = append(fitness, f)

		}
	}

	return fitness
}

func (f *FastGenetic) Selection(fitness []float64) int {
	fitnessSum := 0.0

	for _, f := range fitness {
		fitnessSum += f
	}

	p := 0.0
	id := 0
	for i, f := range fitness {
		prob := f / fitnessSum
		if p < prob {
			p = prob
			id = i
		}
	}
	return id
}

// will generate offspring that will use in Kmeans calculation
func (f *FastGenetic) Mutation(parent []int) []int {
	if rand.Float64() < f.MutationProbability {
		offspring := []int{}
		centroids, _ := f.Compute_centroids(parent)

		for i := 0; i < len(parent); i++ {
			dist := []float64{}
			max := 0.0
			for c := 0; c < len(centroids); c++ {
				d := helper.EuclideanDistance(f.X[i], centroids[c])
				dist = append(dist, d)
				if max < d {
					max = d
				}
			}

			sum := dist[0]
			for _, ds := range dist[1:] {
				sum += 1.5*max - ds + 0.5
			}
			maxD := 0.0
			id := 0
			for i, ds := range dist {
				p := (1.5*max - ds + 0.5) / sum
				if maxD < p {
					maxD = p
					id = i
				}
			}
			offspring = append(offspring, id)

		}

		return offspring
	}

	return parent
}

func (f *FastGenetic) KMeans(solution []int) ([][]float64, []int) {
	centroids := make([][]float64, f.N_clusters)
	for i := 0; i < f.MaxIters; i++ {
		centroids, _ = f.Compute_centroids(solution)
		new_solution := make([]int, len(solution))

		for s := 0; s < len(solution); s++ {
			dist := make([]float64, f.N_clusters)
			for c := 0; c < len(centroids); c++ {
				d := helper.EuclideanDistance(f.X[s], centroids[c])
				dist[c] = d
			}

			id := helper.ArgMin(dist)
			new_solution[s] = id
		}
		if reflect.DeepEqual(new_solution, solution) {
			break
		}
		solution = new_solution
	}

	return centroids, solution
}

func (f *FastGenetic) Fit() ([][]float64, []int) {
	population, _ := f.GeneratePop()

	centroids := make([][]float64, f.N_clusters)
	fitness := make([]float64, f.PopSize)

	for i := 0; i < f.GenSize; i++ {
		fitness = f.ComputeFitness(population)
		id := f.Selection(fitness)
		offspring := f.Mutation(population[id])
		newCentroids, offspring := f.KMeans(offspring)

		iMin := helper.ArgMin(fitness)
		if reflect.DeepEqual(offspring, population[iMin]) {
			break
		}

		population[iMin] = offspring
		centroids = newCentroids
	}

	iMax := helper.ArgMax(fitness)
	return centroids, population[iMax]
}
