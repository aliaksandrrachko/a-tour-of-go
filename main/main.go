package main

import "fmt"

func main() {
	fmt.Println("=======================================================")
	fmt.Println("-- Basics: Packages, variables, and functions.")
	fmt.Println("=======================================================")
	packagesVariablesFunctions()
	fmt.Printf("\n\n\n\n")

	fmt.Println("=======================================================")
	fmt.Println("-- Basics: Flow control statements: for, if, else, witch and defer")
	fmt.Println("=======================================================")
	flowControlStatements()
	fmt.Printf("\n\n\n\n")

	fmt.Println("=======================================================")
	fmt.Println("-- Basics: More types: struct, slices, and maps")
	fmt.Println("=======================================================")
	moreTypes()
	fmt.Printf("\n\n\n\n")

	fmt.Println("=======================================================")
	fmt.Println("-- Methods and interfaces: Methods and interfaces")
	fmt.Println("=======================================================")
	methodsAndInterfaces()
	fmt.Printf("\n\n\n\n")

	fmt.Println("=======================================================")
	fmt.Println("-- Generics: Generics")
	fmt.Println("=======================================================")
	generics()
	fmt.Printf("\n\n\n\n")

	fmt.Println("=======================================================")
	fmt.Println("-- Concurrency: Concurrency")
	fmt.Println("=======================================================")
	concurrency()
	fmt.Printf("\n\n\n\n")
}
