package main

import "fmt"

func generics() {
	// Type parameters
	intSlice := []int{10, 21, 14, -10}
	fmt.Println(Index(intSlice, 14))

	stringSlice := []string{"41", "1f", "ava", "abc"}
	fmt.Println(Index(stringSlice, "1f"))

	// Generic types
	rootElement := List[string]{nil, "A"}
	rootElement.add("B").add("C").add("D")
	fmt.Println(rootElement)
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list List[T]) String() string {
	return fmt.Sprintf("{\"val\":\"%s\",\"next\":\"%s\"}", list.val, list.next)
}

func (list *List[T]) add(o T) *List[T] {
	nextElement := &List[T]{nil, o}
	list.next = nextElement
	return nextElement
}
