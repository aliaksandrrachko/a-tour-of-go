package main

import "fmt"

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
	printSlice(arraySliceDefaultExample)
	printSlice(sliceDefaultExmpaleOne)
	printSlice(sliceDefaultExmpaleTwo)
	printSlice(sliceDefaultExmpaleThree)

	// nil slices
	var nullSlice []int
	printSlice(nullSlice)
	if nullSlice == nil {
		fmt.Println("nill!")
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type Vertex struct {
	X, Y int
}
