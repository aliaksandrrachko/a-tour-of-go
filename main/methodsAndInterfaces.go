package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func methodsAndInterfaces() {
	// Methods
	v := VertexFloat{3, 4}
	fmt.Println(v.Abs())

	// Methods and functions
	fmt.Println(Abs(v))

	// Methods continued
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer receivers
	vertexPointerReceivers := VertexFloat{3, 4}
	vertexPointerReceivers.Scale(10)
	fmt.Println(vertexPointerReceivers.Abs())

	// Pointer and functions
	vertexPointerReceiversFunction := VertexFloat{3, 4}
	ScaleFunc(&vertexPointerReceiversFunction, 10)
	fmt.Println(AbsFunc(vertexPointerReceiversFunction))

	// Choosing a value or pointer receiver
	// Two reasons to use a pointer receiver.
	// 1. you can modify the value
	// 2. to avoid copying the value on each method call
	// In general, all methods on a given type should have either value or pointer receivers
	vertexChoosing := &VertexFloat{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", vertexChoosing, vertexChoosing.Abs())
	vertexChoosing.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", vertexChoosing, vertexChoosing.Abs())

	// Interfaces

}

type VertexFloat struct {
	X, Y float64
}

// this is method for struct
// this method operate copy of the original value
func (v VertexFloat) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// this is function
func Abs(v VertexFloat) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// this is method for type in the same package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(-f)
}

// this method have a pointer receiver
// to change the VertexFloat declared in the main
func (v *VertexFloat) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v VertexFloat) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *VertexFloat, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Abser interface {
	Abs() float64
}
