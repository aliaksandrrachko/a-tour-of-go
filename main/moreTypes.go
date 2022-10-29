package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func moreTypes() {
	// Creating a slice with make
	sliceCreatedWithMakeOne := make([]int, 5) // len(a)=5
	printSlice("sliceCreatedWithMakeOne", sliceCreatedWithMakeOne)

	sliceCreatedWithMakeTwo := make([]int, 0, 5)
	printSlice("sliceCreatedWithMakeTwo", sliceCreatedWithMakeTwo)

	sliceCreatedThree := sliceCreatedWithMakeTwo[:2]
	printSlice("sliceCreatedThree", sliceCreatedThree)

	sliceCreatedFour := sliceCreatedThree[2:5]
	printSlice("sliceCreatedFour", sliceCreatedFour)

	// Slices of slices
	// create a tic-tac-toe ticCatToeBoard.
	ticCatToeBoard := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// the players take turns.
	ticCatToeBoard[0][0] = "X"
	ticCatToeBoard[2][2] = "O"
	ticCatToeBoard[1][2] = "X"
	ticCatToeBoard[1][0] = "O"
	ticCatToeBoard[0][2] = "X"

	printSliceOfSlice("Tic-tac-toe board", ticCatToeBoard)

	// Appending to slice
	var sliceToAppending []int
	printSlice("sliceToAppending", sliceToAppending)

	// append works on nell slices
	sliceToAppending = append(sliceToAppending, 0)
	printSlice("sliceToAppending", sliceToAppending)

	// the slice grows as needed.
	sliceToAppending = append(sliceToAppending, 1)
	printSlice("sliceToAppending", sliceToAppending)

	// we can add more than one element at a time.
	sliceToAppending = append(sliceToAppending, 2, 3, 4, 5, 2, 324, 234324)
	printSlice("sliceToAppending", sliceToAppending)

	// test array on copying after append and make
	testArray := [2]int{1, 2}
	testSliceForTestArray := testArray[:]
	printSlice("testSliceForTestArray", testSliceForTestArray)

	testSliceForTestArray[1] = 124413
	// value in testArray will be changed
	// because testSliceForTestArray links on testArray
	printSlice("testSliceForTestArrayAfterModifing", testArray[:])

	testSliceForTestArray = append(testSliceForTestArray, 1, 41, 41, 12)
	testSliceForTestArray[0] = 1234
	// value in testArray won't be changed
	// because after appen was created new array an values was coppied
	printSlice("testSliceForTestArrayAfterAppending", testArray[:])

	// Range
	sliceRangeTest := []int{1, 23, 23, 3, 24, 42, 34, 345, 45, 47, 568, 567, 46, 345, 23, 2}
	for i, v := range sliceRangeTest {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Range continued
	for _, value := range sliceRangeTest {
		fmt.Printf("%d\n", value)
	}

	// Exercise: Slice
	pic.Show(Pic)

	// Maps
	stringVertexBigMap := make(map[string]VertexBig, 0)
	stringVertexBigMap["Bell Labs"] = VertexBig{
		40.68433, -74.39967,
	}
	fmt.Println(stringVertexBigMap["Bell Labs"])
	fmt.Println(stringVertexBigMap)

	// Map literals
	stringVertexBigMap = map[string]VertexBig{
		"Bell Labs": VertexBig{
			40.68433, -74.39967,
		},
		"Google": VertexBig{
			37.42202, -122.08408,
		},
	}
	fmt.Println(stringVertexBigMap)

	// Map literals continued
	stringVertexBigMap = map[string]VertexBig{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(stringVertexBigMap)

	// Mutating Maps
	testMap := make(map[string]int)
	testMap["Answer"] = 42 // insert or update and element
	fmt.Println("The value:", testMap["Answer"])

	testMap["Answer"] = 48
	fmt.Println("The value:", testMap["Answer"])

	delete(testMap, "Answer") // delete an
	fmt.Println("The value:", testMap["Answer"])

	v, ok := testMap["Answer"] // test that a key is present
	fmt.Println("The value:", v, "Present?", ok)

	// Exercise: Maps
	wc.Test(WordCount)

	// Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function closures

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSliceOfSlice(s string, x [][]string) {
	header := s + ":\n"
	fmt.Printf(header)
	for i := 0; i < len(x); i++ {
		fmt.Printf("%s|%s|\n", getEmptyString(len(header)), strings.Join(x[i], " "))
	}
}

func getEmptyString(length int) string {
	emptyString := ""
	for i := 0; i < length; i++ {
		emptyString = emptyString + " "
	}
	return emptyString
}

// Implement Pic.
// It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
//
// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
//
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
//
// (Use uint8(intValue) to convert between types.)
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dx)

	for i := 0; i < dx; i++ {
		result[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8((dx * dy) / 2)
		}
	}

	return result
}

type VertexBig struct {
	Lat, Long float64
}

// Implement WordCount.
// It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
//
// You might find strings.Fields helpful.
func WordCount(s string) map[string]int {
	resultMap := make(map[string]int)

	for _, s := range strings.Fields(s) {
		_, ok := resultMap[s]
		if ok {
			resultMap[s] = resultMap[s] + 1
		} else {
			resultMap[s] = 1
		}
	}

	return resultMap
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
