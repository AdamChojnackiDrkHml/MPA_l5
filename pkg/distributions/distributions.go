package distributions

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	CYC = 1
	SEQ = 2
	SET = 3
)

func Logarithmic(x float64, min int) int {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	U := r.Float64()
	p_k := -1 / math.Log(1-x)
	S := p_k
	k := 1

	for k < min {
		p_k = p_k * x * (float64(k) / float64(k+1))
		S += p_k
		k++
	}

	for U > S {
		p_k = p_k * x * (float64(k) / float64(k+1))
		S = S + p_k
		k++
	}

	return k
}

// func Logarithmic(x float64, minSize int) int {
// 	return generateMin(x, minSize, logarithmic)
// }

func Poisson(x float64, min int) int {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	U := r.Float64()
	p_k := math.Exp(-x)
	S := 0.0
	k := 0

	for k < min {
		S += p_k
		p_k = (p_k / float64(k+1)) * x
		k++
	}

	for U > S {
		S = S + p_k
		p_k = (p_k / float64(k+1)) * x
		k++
	}

	return k
}

// func Poisson(x float64, minLen int) int {
// 	return generateMin(x, minLen, poisson)
// }

func Geometric(x float64, min int) int {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	U := r.Float64()
	p_k := 1 - x
	S := 0.0
	k := 0

	for k < min {
		S += p_k
		p_k = x * p_k
		k++
	}

	for U > S {
		S = S + p_k
		p_k = x * p_k
		k++
	}

	return k
}

// func Geometric(x float64, minLen int) int {
// 	return generateMin(x, minLen, geometric)
// }

// func generateMin(x float64, min int, gf func(float64) int) int {
// 	var k int
// 	for k = geometric(x); k < min; k = geometric(x) {
// 	}

// 	return k
// }

func NextBernoulli(x float64) bool {
	dist := distuv.Bernoulli{
		P:   x,
		Src: rand.NewSource(uint64(time.Now().UnixNano())),
	}

	return dist.Rand() == 1.0
}
