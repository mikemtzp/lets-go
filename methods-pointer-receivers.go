package main

import (
	"fmt"
	"math"
)

type Vertez struct {
	X, Y float64
}

func (v *Vertez) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertez) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodsPointerReceivers() {
	v := &Vertez{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}