package qc

import (
	"math"
	"sort"
)

// Mean returns the arithmetic mean (average) of a slice of float64.
func Mean(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	return SumFloat(data) / float64(len(data)), nil
}

// Median returns the median value of a slice of float64.
// The input is not modified; a copy is sorted internally.
func Median(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2], nil
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2, nil
}

// Mode returns the most frequently occurring value(s) in a slice of float64.
// Returns all values tied for highest frequency.
func Mode(data []float64) ([]float64, error) {
	if len(data) == 0 {
		return nil, ErrEmptySlice
	}
	freq := make(map[float64]int)
	for _, v := range data {
		freq[v]++
	}
	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}
	var modes []float64
	for v, f := range freq {
		if f == maxFreq {
			modes = append(modes, v)
		}
	}
	sort.Float64s(modes)
	return modes, nil
}

// Variance returns the population variance of a slice of float64.
func Variance(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	mean, _ := Mean(data)
	sumSq := 0.0
	for _, v := range data {
		diff := v - mean
		sumSq += diff * diff
	}
	return sumSq / float64(len(data)), nil
}

// SampleVariance returns the sample variance (Bessel's correction, n-1).
func SampleVariance(data []float64) (float64, error) {
	if len(data) < 2 {
		return 0, ErrEmptySlice
	}
	mean, _ := Mean(data)
	sumSq := 0.0
	for _, v := range data {
		diff := v - mean
		sumSq += diff * diff
	}
	return sumSq / float64(len(data)-1), nil
}

// StdDev returns the population standard deviation.
func StdDev(data []float64) (float64, error) {
	v, err := Variance(data)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(v), nil
}

// SampleStdDev returns the sample standard deviation (Bessel's correction).
func SampleStdDev(data []float64) (float64, error) {
	v, err := SampleVariance(data)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(v), nil
}

// Range returns the range (max - min) of a slice of float64.
func Range(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	minVal := data[0]
	maxVal := data[0]
	for _, v := range data[1:] {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal - minVal, nil
}

// Percentile returns the p-th percentile (0-100) of the data.
func Percentile(data []float64, p float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	if p < 0 || p > 100 {
		return 0, ErrInvalidInput
	}
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)
	if p == 0 {
		return sorted[0], nil
	}
	if p == 100 {
		return sorted[len(sorted)-1], nil
	}
	index := p / 100 * float64(len(sorted)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	if lower == upper {
		return sorted[lower], nil
	}
	frac := index - float64(lower)
	return sorted[lower]*(1-frac) + sorted[upper]*frac, nil
}

// Quartiles returns Q1, Q2 (median), and Q3 of the data.
func Quartiles(data []float64) (q1, q2, q3 float64, err error) {
	if len(data) == 0 {
		return 0, 0, 0, ErrEmptySlice
	}
	q1, err = Percentile(data, 25)
	if err != nil {
		return
	}
	q2, err = Percentile(data, 50)
	if err != nil {
		return
	}
	q3, err = Percentile(data, 75)
	return
}

// IQR returns the interquartile range (Q3 - Q1).
func IQR(data []float64) (float64, error) {
	q1, _, q3, err := Quartiles(data)
	if err != nil {
		return 0, err
	}
	return q3 - q1, nil
}

// Covariance returns the population covariance of two datasets.
func Covariance(x, y []float64) (float64, error) {
	if len(x) == 0 || len(y) == 0 {
		return 0, ErrEmptySlice
	}
	if len(x) != len(y) {
		return 0, ErrInvalidInput
	}
	meanX, _ := Mean(x)
	meanY, _ := Mean(y)
	sum := 0.0
	for i := range x {
		sum += (x[i] - meanX) * (y[i] - meanY)
	}
	return sum / float64(len(x)), nil
}

// Correlation returns the Pearson correlation coefficient of two datasets.
func Correlation(x, y []float64) (float64, error) {
	if len(x) == 0 || len(y) == 0 {
		return 0, ErrEmptySlice
	}
	if len(x) != len(y) {
		return 0, ErrInvalidInput
	}
	cov, err := Covariance(x, y)
	if err != nil {
		return 0, err
	}
	stdX, err := StdDev(x)
	if err != nil {
		return 0, err
	}
	stdY, err := StdDev(y)
	if err != nil {
		return 0, err
	}
	if stdX == 0 || stdY == 0 {
		return 0, ErrInvalidInput
	}
	return cov / (stdX * stdY), nil
}

// ZScore returns the z-score of a value given mean and standard deviation.
func ZScore(value, mean, stdDev float64) (float64, error) {
	if stdDev == 0 {
		return 0, ErrDivByZero
	}
	return (value - mean) / stdDev, nil
}

// MAD returns the Mean Absolute Deviation of the data.
func MAD(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	mean, _ := Mean(data)
	sum := 0.0
	for _, v := range data {
		sum += math.Abs(v - mean)
	}
	return sum / float64(len(data)), nil
}

// RMS returns the Root Mean Square of the data.
func RMS(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptySlice
	}
	sumSq := 0.0
	for _, v := range data {
		sumSq += v * v
	}
	return math.Sqrt(sumSq / float64(len(data))), nil
}

// Entropy returns the Shannon entropy of a probability distribution.
// Input should be probabilities that sum to 1.
func Entropy(probs []float64) (float64, error) {
	if len(probs) == 0 {
		return 0, ErrEmptySlice
	}
	h := 0.0
	for _, p := range probs {
		if p < 0 || p > 1 {
			return 0, ErrInvalidInput
		}
		if p > 0 {
			h -= p * math.Log2(p)
		}
	}
	return h, nil
}

// Skewness returns the skewness of the data (Fisher's definition).
func Skewness(data []float64) (float64, error) {
	if len(data) < 3 {
		return 0, ErrEmptySlice
	}
	mean, _ := Mean(data)
	std, _ := StdDev(data)
	if std == 0 {
		return 0, ErrDivByZero
	}
	n := float64(len(data))
	sum := 0.0
	for _, v := range data {
		sum += math.Pow((v-mean)/std, 3)
	}
	return (n / ((n - 1) * (n - 2))) * sum, nil
}

// Kurtosis returns the excess kurtosis of the data.
func Kurtosis(data []float64) (float64, error) {
	if len(data) < 4 {
		return 0, ErrEmptySlice
	}
	mean, _ := Mean(data)
	std, _ := StdDev(data)
	if std == 0 {
		return 0, ErrDivByZero
	}
	n := float64(len(data))
	sum := 0.0
	for _, v := range data {
		sum += math.Pow((v-mean)/std, 4)
	}
	return ((n*(n+1))/((n-1)*(n-2)*(n-3)))*sum - (3*(n-1)*(n-1))/((n-2)*(n-3)), nil
}
