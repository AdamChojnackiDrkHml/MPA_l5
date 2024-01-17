package train

import (
	"L5/pkg/consts"
	gf "L5/pkg/generatingFunctions"
	g "L5/pkg/generators"
)

type Plank struct {
	Wheel int
}

type Wagon struct {
	Planks []Plank
}

type Passenger struct {
	Head int
	Body int
}

type WagonWithPassengers struct {
	Wagon      Wagon
	Passengers []Passenger
}

type Train struct {
	Locomotive Wagon
	Wagons     []WagonWithPassengers
}

func CreatePlank(x float64) Plank {
	if g.NextBernoulli(gf.WheelGeneratingFunction(x) / (gf.OneGeneratingFunction(x) + gf.WheelGeneratingFunction(x))) {
		return Plank{
			Wheel: g.Logarithmic(gf.WheelGeneratingFunction(x), consts.MinWheelSize),
		}
	}

	return Plank{
		Wheel: 0,
	}
}

func CreateWagon(x float64) Wagon {
	for {
		len := g.Geometric(gf.PlankGeneratingFunction(x), consts.MinWagonLength)

		planks := make([]Plank, len)

		for i := 0; i < len; i++ {
			planks[i] = CreatePlank(x)
		}

		if enoughWheels(planks, consts.MinWagonWheels) {
			return Wagon{
				Planks: planks,
			}
		}
	}
}

func enoughWheels(planks []Plank, minWheels int) bool {
	var count int

	for _, p := range planks {
		if p.Wheel != 0 {
			count++
		}
	}

	return count >= minWheels
}

func CreatePassenger(x float64) Passenger {
	return Passenger{
		Head: g.Logarithmic(gf.HeadGeneratingFunction(x), consts.MinHeadSize),
		Body: g.Logarithmic(gf.BodyGeneratingFunction(x), consts.MinBodySize),
	}
}

func CreateWagonWithPassengers(x float64) WagonWithPassengers {
	wwp := WagonWithPassengers{
		Wagon: CreateWagon(x),
	}

	var passengers []Passenger

	for {
		len := g.Poisson(gf.PassengerGeneratingFunction(x), consts.MinPassengerCount)
		passengers = make([]Passenger, len)
		for i := range passengers {
			passengers[i] = CreatePassenger(x)
		}

		if unique(passengers) {
			wwp.Passengers = passengers
			return wwp
		}
	}
}

func unique(elems []Passenger) bool {
	set := make(map[[2]int]bool)

	for _, elem := range elems {
		if set[[2]int{elem.Body, elem.Head}] {
			return false
		}

		set[[2]int{elem.Body, elem.Head}] = true
	}

	return true
}

func CreateTrain(x float64) Train {
	len := g.Geometric(gf.WagonWithPassengersGeneratingFunction(x), consts.MinTrainLength)

	t := Train{
		Locomotive: CreateWagon(x),
		Wagons:     make([]WagonWithPassengers, len),
	}

	for i := range t.Wagons {
		t.Wagons[i] = CreateWagonWithPassengers(x)
	}

	return t
}
