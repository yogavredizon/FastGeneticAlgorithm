package helper

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// mean of axis 0
func Mean(a [][]float64) []float64 {
	n_f := len(a[0])
	mean := make([]float64, n_f)

	x := 0
	for x < n_f {
		sum := 0.0
		l := 0.0
		for id := 0; id < len(a); id++ {
			sum += a[id][x]
			l += 1
		}

		mean[x] = sum / l
		x += 1
	}

	return mean
}

func ArgMin(a []float64) int {
	min := a[0]
	iMin := 0
	for iM, m := range a[1:] {
		if min > m {
			min = m
			iMin = iM + 1
		}
	}

	return iMin
}
func ArgMax(a []float64) int {
	max := a[0]
	iMax := 0
	for iMx, m := range a[1:] {
		if max < m {
			max = m
			iMax = iMx + 1
		}
	}

	return iMax
}

func EuclideanDistance(a []float64, b []float64) float64 {
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

func LegalString(solution []int, k int) []int {
	for j := 0; j < len(solution); j++ {
		if j < k {
			solution[j] = j
		} else {
			id := rand.Intn(k)
			solution[j] = id
		}
	}
	return solution
}

func ReadCsv(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	record, _ := csv.NewReader(file).ReadAll()

	m := make(map[string][]string, len(record[0]))

	nD := len(record[0])
	for l := 1; l < len(record); l++ {
		for i := 0; i < nD; i++ {
			m[record[0][i]] = append(m[record[0][i]], record[l][i])
		}
	}

	return m
}

func SelectData(data map[string][]string, columns ...string) map[string][]string {
	newM := make(map[string][]string, len(columns))

	for _, key := range columns {
		if v, ok := data[key]; ok {
			newM[key] = v
		}
	}

	return newM
}

func stringMean(a []string) float64 {
	sum := 0.0

	for _, s := range a {
		convData, _ := strconv.ParseFloat(s, 64)
		sum += convData
	}

	return sum / float64(len(a))
}
func SelectNumericData(data map[string][]string, NumericColumns ...string) map[string][][]float64 {
	l := len(NumericColumns)

	newM := map[string][][]float64{
		"data": make([][]float64, len(data[NumericColumns[0]])),
	}
	for i := 0; i < len(data[NumericColumns[0]]); i++ {
		s := make([]float64, l)
		for j := 0; j < l; j++ {
			convData, err := strconv.ParseFloat(data[NumericColumns[j]][i], 64)
			if err != nil {
				convData = stringMean(data[NumericColumns[j]])
			}
			s[j] = convData
		}
		newM["data"][i] = s
	}
	return newM
}
