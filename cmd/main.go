package main

import (
	"L5/pkg/samplers"
	"encoding/json"
	"fmt"
	"os"
)

func printWagon(w samplers.Wagon) {
	fmt.Print(" Planks and wheels = ")
	for _, p := range w.Planks {
		fmt.Print(p.Wheel, " ")
	}
}

func printPassenger(p samplers.Passenger) {
	fmt.Printf("{%v %v}", p.Body, p.Head)
}
func printTrain(t samplers.Train) {
	fmt.Print("Lokomotywa: ")
	printWagon(t.Locomotive)
	fmt.Println()

	for _, w := range t.Wagons {
		fmt.Print("Passengers = {")
		for _, p := range w.Passengers {
			printPassenger(p)
		}
		fmt.Print("};;")

		printWagon(w.Wagon)
		fmt.Println()
	}
}

type BigResult struct {
	Results []Result
}

type Result struct {
	Trains []samplers.Train
	X      float64
}

func main() {

	numOfTrials := 100
	minX := 0.001
	maxX := 0.4851
	deltaX := 0.001

	bR := BigResult{
		Results: make([]Result, 0),
	}

	for x := minX; x <= maxX; x += deltaX {
		fmt.Println(x)

		r := Result{
			Trains: make([]samplers.Train, numOfTrials),
			X:      x,
		}

		for i := 0; i < numOfTrials; i++ {
			r.Trains[i] = samplers.TrainSampler(x)
		}

		bR.Results = append(bR.Results, r)
	}

	b, err := json.Marshal(bR)
	if err != nil {
		fmt.Println(err)
		return
	}

	// output := &bytes.Buffer{}
	// json.Indent(output, b, "", "  ")

	f, _ := os.Create("data/res.json")
	defer f.Close()
	f.WriteString(string(b))
}
