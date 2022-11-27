package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func swapInt(x, y int) (int, int) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i2, j2 int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func packagesVariablesFunctions() {
	fmt.Println(c, python, java)

	fmt.Println("Welcome to the playground!")
	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(100))
	rand.Seed(1234)
	fmt.Println("Pi = ", math.Pi)
	fmt.Println(add(42, 2342352))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	c, d := swapInt(324, 234)
	fmt.Println(c, d)
	fmt.Println(c, d)

	var i int
	fmt.Println(i, c, python, java)

	var python2, java2 = false, "no!"
	fmt.Println(i2, j2, python2, java2)

	var i3, j3 int = 1, 2
	k3 := 3
	c3, python3, java3 := true, false, "no!"
	fmt.Println(i3, j3, k3, c3, python3, java3)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x5, y5 int = 3, 4
	var f5 float64 = math.Sqrt(float64(x5*x5 + y5*y5))
	var z5 uint = uint(f5)
	fmt.Println(x5, y5, z5)

	fmt.Println(split(14))
}
