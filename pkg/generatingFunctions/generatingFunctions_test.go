package generatingfunctions_test

import (
	generatingfunctions "L5/pkg/generatingFunctions"
	"fmt"
	"math"
	"testing"
)

func TestTrain(t *testing.T) {

	samples := make([]float64, 100)
	results := make([]float64, 100)

	for k := 0; k < 100; k++ {
		samples[k] = 0.01 * float64(k)
		results[k] = generatingfunctions.WagonWithPassengersGeneratingFunction(samples[k])
		if math.Abs(results[k]) >= 1.0 {
			fmt.Println(samples[k])
		}
	}

	fmt.Println(samples)
	fmt.Println(results)

	fmt.Println("dupa")
}
