package main

import (
	"L5/pkg/train"
	"fmt"
)

func printWagon(w train.Wagon) {
	fmt.Print(" Planks and wheels = ")
	for _, p := range w.Planks {
		fmt.Print(p.Wheel, " ")
	}
}

func printPassenger(p train.Passenger) {
	fmt.Printf("{%v %v}", p.Body, p.Head)
}
func printTrain(t train.Train) {
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

func main() {
	t := train.CreateTrain(0.3)

	printTrain(t)

}
