package main

import (
	"fmt"
)

// Iterator interface defines the methods for iterating over elements.
type Iterator interface {
	Next() int     // Return the next element.
	HasNext() bool // Check if there are more elements.
}

// ConcreteIterator represents a concrete iterator for a collection of integers.
type ConcreteIterator struct {
	collection []int
	index      int
}

func NewConcreteIterator(collection []int) *ConcreteIterator {
	return &ConcreteIterator{
		collection: collection,
		index:      0,
	}
}

func (i *ConcreteIterator) Next() int {
	if i.HasNext() {
		nextElement := i.collection[i.index]
		i.index++
		return nextElement
	}
	return -1 // Return -1 if there are no more elements.
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.collection)
}

// Aggregate interface defines the method to create an iterator.
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteAggregate represents a concrete collection of integers.
type ConcreteAggregate struct {
	data []int
}

func NewConcreteAggregate(data []int) *ConcreteAggregate {
	return &ConcreteAggregate{
		data: data,
	}
}

func (a *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(a.data)
}

func main() {
	// Create a concrete collection of integers.
	data := []int{1, 2, 3, 4, 5}

	// Create a concrete aggregate and use it to create an iterator.
	aggregate := NewConcreteAggregate(data)
	iterator := aggregate.CreateIterator()

	// Iterate over the collection using the iterator.
	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Printf("Element: %d\n", element)
	}
}
