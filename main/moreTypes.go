package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func moreTypes() {
  // pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pinter
	fmt.Println(j) // see the new value of j

	// structs - is a collection of fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)

	// structs pinter
	vertexPointer := &v
	vertexPointer.X = 1e9
	fmt.Println(v)

	// strict literals
	v1 := Vertex{1, 2}                // has type Vertex
	v2 := Vertex{X: 1}                /// Y:0 is implict
	v3 := Vertex{}                    // X:0 and Y:0
	vertexPointerTwo := &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, v2, v3, vertexPointerTwo)

	// arrays
	var arrayFirst [2]string
	arrayFirst[0] = "Hello"
	arrayFirst[1] = "World"
	fmt.Println(arrayFirst[1], arrayFirst[0])
	fmt.Println(arrayFirst)

	primes := [6]int{2, 3, 4, 5, 7, 11}
	fmt.Println(primes)

	// slices
	var sliceFirst []int = primes[1:4]
	fmt.Println(sliceFirst)

	namesArray := [4]string{
		"John",
		"Payl",
		"George",
		"Ringo",
	}
	fmt.Println(namesArray)

	sliseOfNamesArray := namesArray[0:2]
	sliseOfNamesArrayTwo := namesArray[1:3]
	fmt.Println(sliseOfNamesArray, sliseOfNamesArrayTwo)

	sliseOfNamesArrayTwo[0] = "XXX"
	fmt.Println(sliseOfNamesArray, sliseOfNamesArrayTwo)
	fmt.Println(sliseOfNamesArray)

	// sliece literals
	numbersArray := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(numbersArray)

	boleanArray := []bool{true, false, true, true, false, true}
	fmt.Println(boleanArray)

	arrayOfStruct := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{113, true},
	}
	fmt.Println(arrayOfStruct)

	// slice defaults
	arraySliceDefaultExample := []int{2, 3, 4, 5, 214, 4, 24}

	sliceDefaultExmpaleOne := arraySliceDefaultExample[1:4]
	fmt.Println(sliceDefaultExmpaleOne)

	sliceDefaultExmpaleTwo := arraySliceDefaultExample[:4]
	fmt.Println(sliceDefaultExmpaleTwo)

	sliceDefaultExmpaleThree := arraySliceDefaultExample[1:]
	fmt.Println(sliceDefaultExmpaleThree)

	// slice length and capacity
	printSliceSimple(arraySliceDefaultExample)
	printSliceSimple(sliceDefaultExmpaleOne)
	printSliceSimple(sliceDefaultExmpaleTwo)
	printSliceSimple(sliceDefaultExmpaleThree)

	// nil slices
	var nullSlice []int
	printSliceSimple(nullSlice)
	if nullSlice == nil {
		fmt.Println("nill!")
	}

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

func printSliceSimple(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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

type Vertex struct {
	X, Y int
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
