package main

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Speed    float64
	FuelType string
	Type     string
}

type Motorcycle struct {
	Speed float64
	Type  string
}

func (car Car) CalculateTravelTime(distance float64) float64 {
	return distance / car.Speed
}

func (motorcycle Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / motorcycle.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	results := make(map[string]float64)
	for _, v := range vehicles {
		switch v.(type) {
		case Car:
			car := v.(Car)
			results[car.Type] = car.CalculateTravelTime(distance)
		case Motorcycle:
			motorcycle := v.(Motorcycle)
			results[motorcycle.Type] = motorcycle.CalculateTravelTime(distance)
		}
	}
	return results
}
