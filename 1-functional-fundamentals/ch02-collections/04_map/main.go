package main

import (
	"fmt"
	"strings"
)

type Object interface{}
type Collection []Object

func NewCollection(size int) Collection {
	return make(Collection, size)
}

type Callback func(current, currentKey, src Object) Object

func Map(c Collection, cb Callback) Collection {
	if c == nil {
		return Collection{}
	} else if cb == nil {
		return c
	}

	result := NewCollection(len(c))
	for index, val := range c {
		result[index] = cb(val, index, c)
	}
	return result
}

func main() {
	transformation10 := func(curVal, _, _ Object) Object {
		return curVal.(int) * 10 // reflection
	}

	result := Map(Collection{1, 2, 3, 4}, transformation10)
	fmt.Printf("result: %v\n", result)

	transformationUpper := func(curVal, _, _ Object) Object {
		return strings.ToUpper(curVal.(string)) // reflection
	}

	result = Map(Collection{"alice", "bob", "cindy"}, transformationUpper)
	fmt.Printf("result: %v\n", result)
}
