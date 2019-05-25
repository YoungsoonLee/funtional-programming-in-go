package main

import "fmt"

type numberIs func(int) bool

func lessThanTwo(i int) bool { return i < 2 }

// no curried parameters
func lessThan(x int, y int) bool {
	return x < y
}

func (f numberIs) apply(s ...int) (ret []bool) {
	for _, i := range s {
		ret = append(ret, f(i))
	}
	return ret
}

func main() {
	fmt.Println("NonCurried - lessThan(1,2): ", lessThan(1, 2))
	fmt.Println("Curried - LessThanTwo(1): ", lessThanTwo(1))

	// uase anonymous function
	isLessThanOne := numberIs(func(i int) bool {
		return i < 1
	}).apply

	isLessThanTwo := numberIs(lessThanTwo).apply // use named function.

	s := []int{0, 1, 2}
	fmt.Println("Curried, given:", s, "...")
	fmt.Println("isLessThanOne:", isLessThanOne(s...))
	fmt.Println("isLessThanTwo:", isLessThanTwo(s...))
}
