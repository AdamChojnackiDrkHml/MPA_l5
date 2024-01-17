package generatingfunctions

import "math"

func CycleGeneratingFunction(x float64) float64 {
	return math.Log(1.0 / (1.0 - ZGeneratingFunction(x)))
}

func SeqGeneratingFunction(x float64) float64 {
	return 1.0 / (1.0 - x)
}

func SetGeneratingFunction(x float64) float64 {
	return math.Exp(x)
}

func WheelGeneratingFunction(x float64) float64 {
	return CycleGeneratingFunction(ZGeneratingFunction(x))
}

func PlankGeneratingFunction(x float64) float64 {
	return ZGeneratingFunction(x) * ZGeneratingFunction(x) *
		(CycleGeneratingFunction(x) + OneGeneratingFunction(x))
}

func OneGeneratingFunction(x float64) float64 {
	return 1.0
}

func ZGeneratingFunction(x float64) float64 {
	return x
}

func WagonGeneratingFunction(x float64) float64 {
	return SeqGeneratingFunction(PlankGeneratingFunction(x)) - 1
}

func PassengerGeneratingFunction(x float64) float64 {
	return HeadGeneratingFunction(x) * BodyGeneratingFunction(x)
}

func HeadGeneratingFunction(x float64) float64 {
	return CycleGeneratingFunction(ZGeneratingFunction(x))
}

func BodyGeneratingFunction(x float64) float64 {
	return CycleGeneratingFunction(ZGeneratingFunction(x))
}

func TrainGeneratingFunction(x float64) float64 {
	return WagonGeneratingFunction(x) *
		SeqGeneratingFunction(WagonWithPassengersGeneratingFunction(x))
}

func WagonWithPassengersGeneratingFunction(x float64) float64 {
	return WagonGeneratingFunction(x) *
		SetGeneratingFunction(PassengerGeneratingFunction(x))
}
