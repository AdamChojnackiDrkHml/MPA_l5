package samplers

import (
	"L5/pkg/consts"
	dist "L5/pkg/distributions"
	gf "L5/pkg/generatingFunctions"
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

func PlankSampler(x float64) Plank {
	if dist.NextBernoulli(gf.WheelGeneratingFunction(x) / (gf.OneGeneratingFunction(x) + gf.WheelGeneratingFunction(x))) {
		return Plank{
			Wheel: dist.Logarithmic(gf.WheelGeneratingFunction(x), consts.MinWheelSize),
		}
	}

	return Plank{
		Wheel: 0,
	}
}

func WagonSampler(x float64) Wagon {
	for {
		len := dist.Geometric(gf.PlankGeneratingFunction(x), consts.MinWagonLength)

		planks := make([]Plank, len)

		for i := 0; i < len; i++ {
			planks[i] = PlankSampler(x)
		}

		if isEnoughWheels(planks, consts.MinWagonWheels) {
			return Wagon{
				Planks: planks,
			}
		}
	}
}

func isEnoughWheels(planks []Plank, minWheels int) bool {
	var count int

	for _, p := range planks {
		if p.Wheel != 0 {
			count++
		}
	}

	return count >= minWheels
}

func PassengerSampler(x float64) Passenger {
	return Passenger{
		Head: dist.Logarithmic(gf.HeadGeneratingFunction(x), consts.MinHeadSize),
		Body: dist.Logarithmic(gf.BodyGeneratingFunction(x), consts.MinBodySize),
	}
}

func WagonWithPassengersSampler(x float64) WagonWithPassengers {
	wwp := WagonWithPassengers{
		Wagon: WagonSampler(x),
	}

	var passengers []Passenger

	for {
		len := dist.Poisson(gf.PassengerGeneratingFunction(x), consts.MinPassengerCount)
		passengers = make([]Passenger, len)
		for i := range passengers {
			passengers[i] = PassengerSampler(x)
		}

		if areUnique(passengers) {
			wwp.Passengers = passengers
			return wwp
		}
	}
}

func areUnique(elems []Passenger) bool {
	set := make(map[[2]int]bool)

	for _, elem := range elems {
		if set[[2]int{elem.Body, elem.Head}] {
			return false
		}

		set[[2]int{elem.Body, elem.Head}] = true
	}

	return true
}

func TrainSampler(x float64) Train {
	len := dist.Geometric(gf.WagonWithPassengersGeneratingFunction(x), consts.MinTrainLength)

	t := Train{
		Locomotive: WagonSampler(x),
		Wagons:     make([]WagonWithPassengers, len),
	}

	for i := range t.Wagons {
		t.Wagons[i] = WagonWithPassengersSampler(x)
	}

	return t
}
