package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func flowControlStatements() {

	// for:
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// if:
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// if and else:
	fmt.Println(
		pow(3, 2, 10),
	)

	// loops and functions
	SqrtByScale(12341.422, 0.9999999)
	Sqrt(256)
	Sqrt(1564817)
	SqrtByScale(156481234, 0.9999999)

	// switch statements
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan0, windows...
		fmt.Printf("%s. \n", os)
	}

	// switch with no conditions
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer - deffer the execution until the surrounding function returns
	defer fmt.Println("world")

	// stacking defers
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g > = %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

/* Returns the square root calculated by Newton's method */
func Sqrt(x float64) float64 {
	return SqrtByScale(x, float64(1))
}

/* Returns the square root calculated by Newton's method with scale*/
func SqrtByScale(x, scale float64) float64 {
	result := x / 2

	for currentScale := float64(1); currentScale > 1-scale && result > 0; currentScale = (result*result - x) / x {
		result = 0.5 * (result + x/result)
	}

	fmt.Println("Sqrt '", x, "' = ", result, "; ", result, " * ", result, " = ", result*result)

	return result
}
