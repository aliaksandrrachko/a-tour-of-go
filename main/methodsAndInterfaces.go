package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"

	"golang.org/x/tour/reader"
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
	// In general, all methods on a given type should have either value or pointer receivers,
	// because can be problem with implementation interface's methods
	vertexChoosing := &VertexFloat{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", vertexChoosing, vertexChoosing.Abs())
	vertexChoosing.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", vertexChoosing, vertexChoosing.Abs())

	// Interfaces
	var abserInterface Abser
	floatImplementedAbser := MyFloat(-math.Sqrt2)
	vertexImplementedAbser := VertexFloat{3, 4}

	abserInterface = floatImplementedAbser
	abserInterface = vertexImplementedAbser
	abserInterface = &vertexImplementedAbser
	fmt.Println(abserInterface.Abs())

	// Interfaces are implemented implicitly
	var implementationOfI = T{"Hello"}
	implementationOfI.M()

	// Interface values
	var interfaceValues I

	interfaceValues = &T{"Hello"}
	describe(interfaceValues)
	interfaceValues.M()

	interfaceValues = F(math.Pi)
	describe(interfaceValues)
	interfaceValues.M()

	// Interface values with nil underlying values
	var interfaceNilValue I

	var tForInterfaceNilValue T
	interfaceNilValue = tForInterfaceNilValue
	describe(interfaceNilValue)
	interfaceNilValue.M()

	interfaceNilValue = &T{"Hello"}
	describe(interfaceNilValue)
	interfaceNilValue.M()

	// Nil interfaces values
	var nilInterfaceValue I
	describe(nilInterfaceValue)
	// nilInterfaceValue.M() // will produce run-time error

	// The empty interface
	var emptyInterfaceValue interface{}
	describeEmptyInterface(emptyInterfaceValue)

	emptyInterfaceValue = 42
	describeEmptyInterface(emptyInterfaceValue)

	emptyInterfaceValue = "hello"
	describeEmptyInterface(emptyInterfaceValue)

	emptyInterfaceValue = &VertexFloat{123.134, 134.14}
	describeEmptyInterface(emptyInterfaceValue)

	// Type assertions
	var testAssertionValue interface{} = "hello"

	stringAssertionValue := testAssertionValue.(string)
	fmt.Println(stringAssertionValue)

	stringAssertionValue, isString := testAssertionValue.(string)
	fmt.Println(stringAssertionValue, isString)

	floatAssertionValue, isFloat := testAssertionValue.(float64) // floatAssertionValue will be nil or default
	fmt.Println(floatAssertionValue, isFloat)

	if isFloat {
		floatAssertionValue := testAssertionValue.(float64)
		fmt.Println(floatAssertionValue)
	}

	// Type switches
	do(231)
	do("hello")
	do(true)
	do(&VertexFloat{12, 14})

	// Stringers
	aPerson := Person{"Arthur Dent", 42}
	bPerson := Person{"Zappos Beeblebrox", 90101}
	fmt.Println(aPerson, bPerson)

	// Exercise: Stringers
	// Add a "String() string" method to IPAddr.
	hosts := map[string]IPAddr{
		"loopback":      {127, 0, 0, 1},
		"googleDNS":     {8, 8, 8, 8},
		"someIpAddress": {1, 2, 3, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Exercise: Errors
	fmt.Println(SqrtToErrorExercise(2))
	fmt.Println(SqrtToErrorExercise(-2))

	// Readers
	newReaderExample := strings.NewReader("Hello, Reader!")

	arrayForReading := make([]byte, 8)
	for {
		n, err := newReaderExample.Read(arrayForReading)
		fmt.Printf("n = %v err = %v arrayForReading = %v\n", n, err, arrayForReading)
		fmt.Printf("arrayForReading[:n] = %q\n", arrayForReading[:n])
		if err == io.EOF {
			break
		}
	}

	// Exercise: Readers
	reader.Validate(MyReader{})
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

// this method have to return fmt. a pointer receiver
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

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	if &t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (t *T) MForPointer() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	stringValues := []string{}
	for _, v := range ipAddr {
		stringValues = append(stringValues, fmt.Sprintf("%d", v))
	}
	return strings.Join(stringValues, ".")
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/* Returns the square root calculated by Newton's method */
func SqrtToErrorExercise(x float64) (float64, error) {
	return SqrtByScaleToErrorExercise(x, float64(1))
}

/* Returns the square root calculated by Newton's method with scale*/
func SqrtByScaleToErrorExercise(x, scale float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(float64(x))
	}

	result := x / 2

	for currentScale := float64(1); currentScale > 1-scale && result > 0; currentScale = (result*result - x) / x {
		result = 0.5 * (result + x/result)
	}

	fmt.Println("Sqrt '", x, "' = ", result, "; ", result, " * ", result, " = ", result*result) // string to debug

	return result, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

type MyReader struct{}

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
func (reader MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 65
	}
	return len(p), nil
}
