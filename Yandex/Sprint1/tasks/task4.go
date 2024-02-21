package main

import (
	"fmt"
	"math"
)

const (
	place  = 0
	target = 3
)

func task4() {
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace)
	}
}

func (g *Grasshopper) WhereAmI() int {
	return g.place
}

func (g *Grasshopper) Jump() (int, error) {
	if g.place == g.target {
		return 0, fmt.Errorf("already at target")
	}

	if math.Abs(float64(g.target-g.place)) <= 5 {
		g.place = g.target
	} else {
		return g.place, fmt.Errorf("can't jump that far")
	}

	return g.place, nil
}

type Grasshopper struct {
	place, target int
}

type Jumper interface {
	WhereAmI() int
	Jump() (int, error)
}

func PlaceJumper(place, target int) Jumper {
	var g = &Grasshopper{}

	g.place = place
	g.target = target

	return g
}
